import { defineStore } from "pinia";
import { PhasesValidateConfig } from "../../wailsjs/go/backend/Backend";
import { ProcessListener } from "../types/ProcessListener";
import { distillation } from "../../wailsjs/go/models";

export interface Button {
    is_enabled: boolean;
    enable: Function;
}

export const useProcessStore = defineStore('process', {
    state: () => {
        return {
            enable: {} as Button,
            moveToNext: {} as Button,
            disable: {} as Button,
            is_valid: false,
        }
    },
    actions: {
        init() {
            this.enable.is_enabled = false
            this.enable.enable = () => { this.echo("enable") }

            this.moveToNext.is_enabled = false
            this.moveToNext.enable = () => { this.echo("moveToNext") }

            this.disable.is_enabled = false
            this.disable.enable = () => { this.echo("stop") }

            ProcessListener.subscribeValidate(this.onValidate)
        },

        reload() {
            PhasesValidateConfig()
        },

        onValidate(v: distillation.ProcessConfigValidation) {
            this.is_valid = v.valid
            if (this.is_valid) {
                this.enable.is_enabled = true
                this.moveToNext.is_enabled = true
                this.disable.is_enabled = true
            } else {
                this.enable.is_enabled = false
                this.moveToNext.is_enabled = false
                this.disable.is_enabled = false
            }
        },

        echo(msg: string) {
            console.log(msg)
        }

    }
})
