import { defineStore } from "pinia";
import { Components, PhasesGetPhaseConfigs } from "../../wailsjs/go/backend/Backend";
import { distillation, process } from "../../wailsjs/go/models";
import { Phases, ProcessPhaseConfig } from "../types/Phases";
import { ProcessListener } from "../types/ProcessListener";

export const usePhasesStore = defineStore('phases', {
    state: () => {
        return {
            phases: new Phases() as Phases,
        }
    },
    actions: {
        init() {
            ProcessListener.subscribePhaseConfig(this.phaseConfigUpdate)
            ProcessListener.subscribePhaseCount(this.phaseCountUpdate)
            this.reload()
        },

        reload() {
            Components().then(
                components => {
                    PhasesGetPhaseConfigs().then(
                        result => { this.updatePhases(components, result) },
                        error => {
                            console.debug(error)
                            setTimeout(() => { this.init() }, 200)
                        },
                    )
                },
                error => {
                    console.debug(error)
                    setTimeout(() => { this.init() }, 200)
                },
            )
        },

        phaseCountUpdate(_: distillation.ProcessPhaseCount) {
            this.reload()
        },

        updatePhases(comp: process.Components, value: distillation.ProcessPhaseConfig[]) {
            let size = value.length
            let configs: ProcessPhaseConfig[] = []

            value.forEach((v: distillation.ProcessPhaseConfig, i: number) => {
                configs.push(new ProcessPhaseConfig(i, v.next, v.heaters, v.gpio, comp))
            })
            this.phases = new Phases(configs, size)
        },

        phaseConfigUpdate(n: number, v: distillation.ProcessPhaseConfig) {
            Components().then((components: process.Components) => {
                this.phases.phases[n] = new ProcessPhaseConfig(n, v.next, v.heaters, v.gpio, components)
            },
                error => {
                    console.debug(error)
                    setTimeout(() => {
                        this.phaseConfigUpdate(n, v)
                    }, 200);
                })
        },
        updateComponents() {
            Components().then(
                (comp: process.Components) => {
                    this.phases.phases.forEach((elem) => {
                        elem.updateComponents(comp)
                    })
                },
                error => {
                    console.debug(error)
                    setTimeout(() => { this.updateComponents() }, 200);
                }
            )
        }
    }
})