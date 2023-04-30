import { defineStore } from "pinia";
import { parameters } from "../../wailsjs/go/models";
import { Heater } from "../types/Heater";
import { HeaterListener } from "../types/HeaterListener";
import { HeatersGet } from "../../wailsjs/go/backend/Backend";

export const useHeatersStore = defineStore('heaters', {
    state: () => {
        return {
            heaters: [] as Heater[],
        }
    },
    actions: {
        init() {
            this.reload()
            HeaterListener.subscribe(this.update)
        },

        reload() {
            HeatersGet().then(
                (got) => {
                    let newHeaters: Heater[] = []
                    got.forEach((heater: parameters.Heater) => {
                        let newHeater = new Heater(heater.ID, heater.enabled)
                        newHeaters.push(newHeater)
                    })

                    this.heaters = newHeaters.sort((a: Heater, b: Heater) => {
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

        update(h: parameters.Heater) {
            let idx = this.heaters.findIndex(i => i.heater.ID == h.ID)
            if (idx != -1) {
                let heater = new Heater(h.ID, h.enabled)
                this.heaters[idx] = heater
            }
        }
    }
})