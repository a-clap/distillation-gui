import { defineStore } from "pinia";
import { PhasesDisable, PhasesEnable, PhasesMoveToNext, PhasesValidateConfig } from "../../wailsjs/go/backend/Backend";
import { ProcessListener } from "../types/ProcessListener";
import { distillation } from "../../wailsjs/go/models";

export interface Button {
    is_enabled: boolean;
    enable: Function;
}

export interface Heater {
    id: string;
    pwr: string;
}


function padTo2Digits(num: number) {
    return num.toString().padStart(2, '0');
}

function formatDate(date: Date) {
    return (
        [
            padTo2Digits(date.getHours()),
            padTo2Digits(date.getMinutes()),
            padTo2Digits(date.getSeconds()),
        ].join(':') 
    );
}

export const useProcessStore = defineStore('process', {
    state: () => {
        return {
            enable: {} as Button,
            moveToNext: {} as Button,
            disable: {} as Button,
            heaters: [] as Heater[],
            is_valid: false,
            running: false,
            show_status: false,
            start_time: "",
            end_time: "",
            current_phase: "",
            current_type_time: false,
            phase_timeleft: "",
            phase_sensor: "",
            phase_sensor_threshold: "",
        }
    },
    actions: {
        init() {
            this.enable.enable = this.processEnable
            this.moveToNext.enable = this.processMoveToNext
            this.disable.enable = this.processDisable

            this.updateButtons()
            ProcessListener.subscribeValidate(this.onValidate)
            ProcessListener.subscribeStatus(this.onStatus)
        },

        reload() {
            PhasesValidateConfig()
        },

        toggle() {
            this.is_valid = !this.is_valid
            let v = new distillation.ProcessConfigValidation()
            v.valid = this.is_valid
            this.onValidate(v)

            this.start_time = formatDate(new Date())
            this.end_time = formatDate(new Date())
            this.show_status = this.is_valid
        },

        onStatus(v: distillation.ProcessStatus) {
            this.running = v.running
            if(v.running || v.done) {
                this.current_phase = v.phase_number.toString()
                // Time
                this.current_type_time = v.next.type == 0
                this.phase_timeleft = v.next.time_left.toString()
                if (v.next.temperature) {
                    this.phase_sensor_threshold = v.next.temperature?.sensor_threshold.toString()
                    this.phase_sensor = v.next.temperature?.sensor_id.toString()
                }
                this.start_time = v.start_time
                if (v.done) {
                    this.end_time = v.end_time
                } else {
                    this.end_time = ""
                }

                let heaters: Heater[] = []
                v.heaters.forEach((v) => {
                    let h: Heater = {id: v.ID, pwr: v.power.toString()}
                })
                
                this.heaters = heaters.sort((a: Heater, b: Heater) => {
                    if (a.id > b.id) {
                        return 1
                    }
                    if (a.id < b.id) {
                        return -1
                    }
                    return 0
                })

            }
            this.updateButtons()
            this.show_status = v.running || v.done

            console.log(v)
        },
        
        updateButtons() {
            this.moveToNext.is_enabled = this.running
            this.disable.is_enabled = this.running
            
            this.enable.is_enabled = !this.running && this.is_valid
        },

        onValidate(v: distillation.ProcessConfigValidation) {
            this.is_valid = v.valid
            this.updateButtons()
        },

        processEnable() {
            PhasesEnable()
        },
        processMoveToNext() {
            PhasesMoveToNext()
        },
        processDisable() {
            PhasesDisable()
        },
    }
})
