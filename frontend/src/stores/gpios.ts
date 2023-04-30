import { defineStore } from "pinia";
import { GPIOGet } from "../../wailsjs/go/backend/Backend";
import { GPIO } from "../types/GPIO";
import { parameters } from "../../wailsjs/go/models";
import { GPIOListener } from "../types/GPIOListener";

export const useGpioStore = defineStore('gpio', {
    state: () => {
        return {
            gpios: [] as GPIO[],
        }
    },
    actions: {
        init() {
            GPIOListener.subscribe(this.update)
            this.reload()
        },

        reload() {
            GPIOGet().then(
                (gpio: parameters.GPIO[]) => {
                    let newGPIO: GPIO[] = []
                    gpio.forEach((p: parameters.GPIO) => {
                        let gp = new GPIO(p.id, p.active_level, p.value)
                        newGPIO.push(gp)
                    })

                    this.gpios = newGPIO.sort((a: GPIO, b: GPIO) => {
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

        update(p: parameters.GPIO) {
            let idx = this.gpios.findIndex(i => i.name == p.id)
            if (idx != -1) {
                this.gpios[idx] = new GPIO(p.id, p.active_level, p.value)
            }
        }
    }
})