import { defineStore } from "pinia";
import { parameters } from "../../wailsjs/go/models";
import { DS } from "../types/DS";
import { DSListener } from "../types/DSListener";
import { DSGet } from "../../wailsjs/go/backend/Backend";
import { TemperatureErrorCodeEmptyBuffer, TemperatureErrorCodeInternal, TemperatureErrorCodeWrongID } from "../../wailsjs/go/backend/Models";
import { ErrorListener } from "../types/ErrorListener";
import { AppErrorCodes } from "./error_codes";

export const useDSStore = defineStore('ds', {
    state: () => {
        return {
            ds: [] as DS[],
            errCodeInternal: 0 as number,
            errCodeEmptyBuffer: 0 as number,
            errCodeWrongID: 0 as number,
        }
    },
    actions: {
        init() {
            TemperatureErrorCodeEmptyBuffer().then(v => {
                this.errCodeEmptyBuffer = v
            })
            TemperatureErrorCodeInternal().then(v => {
                this.errCodeInternal = v
            })
            TemperatureErrorCodeWrongID().then(v => {
                this.errCodeWrongID = v
            })

            DSListener.subscribeConfig(this.updateConfig)
            DSListener.subscribeTemperature(this.updateTemperature)
            this.reload()
        },

        reload() {
            DSGet().then(
                (ds) => {
                    let newDses: DS[] = []
                    ds.forEach((d: parameters.DS) => {
                        let ds = new DS(d.name, d.id, d.enabled, d.correction, d.samples, d.resolution)
                        newDses.push(ds)
                    })

                    this.ds = newDses.sort((a: DS, b: DS) => {
                        if (a.name > b.name) {
                            return 1
                        }
                        if (a.name < b.name) {
                            return -1
                        }
                        return 0
                    })
                },
                error => {
                    console.debug(error)
                    setTimeout(() => { this.reload() }, 200);
                }
            )
        },
        updateConfig(d: parameters.DS) {
            let idx = this.ds.findIndex(i => i.id == d.id)
            if (idx != -1) {
                let ds = new DS(d.name, d.id, d.enabled, d.correction, d.samples, d.resolution)
                ds.temperature = this.ds[idx].temperature
                this.ds[idx] = ds
            }
        },
        updateTemperature(t: parameters.Temperature) {
            let idx = this.ds.findIndex(i => i.id == t.ID)
            if (idx != -1) {
                // Discard, if it is not enabled
                if (!this.ds[idx].enable) {
                    return
                }
                switch (t.error_code) {
                    case this.errCodeEmptyBuffer:
                    case this.errCodeWrongID:
                        // Nothing to do
                        return
                    case this.errCodeInternal:
                        ErrorListener.sendError(AppErrorCodes.DSInternalError)
                        // Notify about error
                        return
                }
                this.ds[idx].temperature = t.temperature
            }
        },

    }
})