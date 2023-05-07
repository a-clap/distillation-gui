import { defineStore } from "pinia";
import { PhasesGetGlobalConfig } from "../../wailsjs/go/backend/Backend";
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
            ProcessListener.subscribeGlobalConfig(this.updatePhases)
            ProcessListener.subscribePhaseConfig(this.phaseConfigUpdate)
            ProcessListener.subscribePhaseCount(this.phaseCountUpdate)
            this.reload()
        },

        reload() {
            PhasesGetGlobalConfig().then(
                result => { this.updatePhases(result) },
                error => {
                    console.debug(error)
                    setTimeout(() => { this.init() }, 200)
                },
            )
        },

        phaseCountUpdate(_: distillation.ProcessPhaseCount) {
            this.reload()
        },

        updatePhases(value: process.Config) {
            let configs: ProcessPhaseConfig[] = []

            value.phases.forEach((v: distillation.ProcessPhaseConfig, i: number) => {
                configs.push(new ProcessPhaseConfig(i, v.next, v.heaters, v.gpio))
            })
            this.phases = new Phases(configs, value.global_gpio, value.sensors)
        },

        phaseConfigUpdate(n: number, v: distillation.ProcessPhaseConfig) {
            this.phases.phases[n] = new ProcessPhaseConfig(n, v.next, v.heaters, v.gpio)
        },
    }
})