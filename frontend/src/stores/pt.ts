import { defineStore } from "pinia";
import { parameters } from "../../wailsjs/go/models";
import { PT100 } from "../types/PT100";
import { PTListener } from "../types/PTListener";
import { PTGet } from "../../wailsjs/go/backend/Backend";

export const usePTStore = defineStore('pt', {
    state: () => {
        return {
            pt: [] as PT100[],
        }
    },
    actions: {
        init() {
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
                this.pt[idx].temperature = t.temperature.toFixed(2)
            }
        },

    }
})
