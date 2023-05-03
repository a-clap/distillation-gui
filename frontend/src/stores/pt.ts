import { defineStore } from "pinia";
import { parameters } from "../../wailsjs/go/models";
import { PT100 } from "../types/PT100";
import { PTListener } from "../types/PTListener";
import { ErrorListener } from "../types/ErrorListener";
import { PTGet } from "../../wailsjs/go/backend/Backend";
import { AppErrorCodes } from "./error_codes";
import {
    TemperatureErrorCodeEmptyBuffer,
    TemperatureErrorCodeInternal,
    TemperatureErrorCodeWrongID
} from "../../wailsjs/go/backend/Models";

export const usePTStore = defineStore('pt', {
    state: () => {
        return {
            pt: [] as PT100[],
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

            PTListener.subscribeConfig(this.updateConfig)
            PTListener.subscribeTemperature(this.updateTemperature)
            this.reload()
        },
        reload() {
            PTGet().then(
                (got: parameters.PT[]) => {
                    let newPT: PT100[] = []
                    got.forEach((p: parameters.PT) => {
                        let ds = new PT100(p.name, p.id, p.enabled, p.correction, p.samples)
                        newPT.push(ds)
                    })

                    this.pt = newPT.sort((a: PT100, b: PT100) => {
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
                })
        },
        updateConfig(p: parameters.PT) {
            let idx = this.pt.findIndex(i => i.id == p.id)
            if (idx != -1) {
                let pt = new PT100(p.name, p.id, p.enabled, p.correction, p.samples)
                pt.temperature = this.pt[idx].temperature
                this.pt[idx] = pt
            }
        },
        updateTemperature(t: parameters.Temperature) {
            let idx = this.pt.findIndex(i => i.id == t.ID)
            if (idx != -1) {
                // Discard, if it is not enabled
                if (!this.pt[idx].enable) {
                    return
                }
                switch (t.error_code) {
                    case this.errCodeEmptyBuffer:
                    case this.errCodeWrongID:
                        // Nothing to do
                        return
                    case this.errCodeInternal:
                        ErrorListener.sendError(AppErrorCodes.PTInternalError)
                        // Notify about error
                        return
                }
                this.pt[idx].temperature = t.temperature.toFixed(2)
            }
        },

    }
})
