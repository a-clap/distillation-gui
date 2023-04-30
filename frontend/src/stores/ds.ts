import { defineStore } from "pinia";
import { parameters } from "../../wailsjs/go/models";
import { DS } from "../types/DS";
import { DSListener } from "../types/DSListener";
import { DSGet } from "../../wailsjs/go/backend/Backend";

export const useDSStore = defineStore('ds', {
    state: () => {
        return {
            ds: [] as DS[],
        }
    },
    actions: {
        init() {
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
                this.ds[idx].temperature = t.temperature.toFixed(2)
            }
        },

    }
})