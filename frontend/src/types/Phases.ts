import { configProviderContextKey } from "element-plus";
import { PhasesSetConfig, PhasesSetPhaseCount } from "../../wailsjs/go/backend/Backend";
import { distillation, process } from "../../wailsjs/go/models";
import Parameter, { writeCallbackType } from "./Parameter";

declare type Notify = (args: ProcessPhaseConfig) => void;
class MoveToNextConfig {
    type: number;
    sensorID: string;
    sensorThreshold: Parameter;
    temperatureHoldSeconds: Parameter;
    secondsToMove: Parameter;

    constructor(n: process.MoveToNextConfig, cb: Notify, args: any) {
        let callback = function (_: any) {
            cb(args)
        }
        this.type = n.type
        this.sensorID = n.sensor_id
        this.sensorThreshold = new Parameter(n.sensor_threshold, true, callback)
        this.temperatureHoldSeconds = new Parameter(n.temperature_hold_seconds, false, callback)
        this.secondsToMove = new Parameter(n.seconds_to_move, false, callback)
    }
}

class GPIOPhaseConfig {
    id: string;
    t_low: Parameter;
    t_high: Parameter;
    hysteresis: Parameter;
    private sensorID: string;
    private inverted_: boolean;
    private callback: writeCallbackType;

    constructor(gpio: process.GPIOPhaseConfig, callback: Notify, args: any) {
        this.callback = function (_: any = 0) {
            callback(args)
        }
        this.id = gpio.id
        this.sensorID = gpio.sensor_id
        this.t_low = new Parameter(gpio.t_low, true, this.callback)
        this.t_high = new Parameter(gpio.t_high, true, this.callback)
        this.hysteresis = new Parameter(gpio.hysteresis, true, this.callback)
        this.inverted_ = gpio.inverted
    }


    get sensor_id(): string {
        return this.sensorID
    }

    set sensor_id(v: string) {
        this.sensorID = v
        this.callback(0)
    }

    get inverted(): boolean {
        return this.inverted_
    }

    set inverted(v: boolean) {
        this.inverted_ = v
        this.callback(0)
    }
}

export class HeaterPhaseConfig {
    id: string;
    power: Parameter;
    constructor(heater: process.HeaterPhaseConfig, writeCallback: Notify, args: any) {
        let callback = function (_: any) {
            writeCallback(args)
        }
        this.id = heater.ID
        this.power = new Parameter(heater.power, false, callback)
    }

}


export class ProcessPhaseConfig {
    private id: number;
    private next: MoveToNextConfig;
    private heaters_: HeaterPhaseConfig[];
    private gpios_: GPIOPhaseConfig[];
    private avail_sensors: string[];
    private component: process.Components;

    constructor(id: number, next: process.MoveToNextConfig, heaters: process.HeaterPhaseConfig[], gpios: process.GPIOPhaseConfig[], component: process.Components) {
        this.id = id
        this.next = new MoveToNextConfig(next, this.update, this)
        this.heaters_ = []
        this.gpios_ = []
        this.component = component
        this.avail_sensors = []

        if (heaters != null) {
            heaters.forEach((v: process.HeaterPhaseConfig) => {
                this.heaters_.push(new HeaterPhaseConfig(v, this.update, this))
            })
        }
        if (gpios != null) {
            gpios.forEach((v: process.GPIOPhaseConfig) => {
                this.gpios_.push(new GPIOPhaseConfig(v, this.update, this))
            })
        }

        // Build list of sensors
        this.component.sensors.forEach(element => {
            this.avail_sensors.push(element)
        });

        // Add needed Configs for Heater
        this.component.heaters.forEach(elem => {
            let item = this.heaters_.findIndex(i => i.id == elem)
            if (item == -1) {
                let conf = new process.HeaterPhaseConfig()
                conf.ID = elem
                conf.power = 0
                this.heaters_.push(new HeaterPhaseConfig(conf, this.update, this))
            }
        })

        // Add needed Configs for Heater
        this.component.outputs.forEach(elem => {
            let item = this.gpios_.findIndex(i => i.id == elem)
            if (item == -1) {
                let conf = new process.GPIOPhaseConfig()
                conf.id = elem
                conf.t_high = 0
                conf.t_low = 0
                conf.hysteresis = 0
                conf.inverted = false
                this.gpios_.push(new GPIOPhaseConfig(conf, this.update, this))
            }

        })
    }

    update(p: ProcessPhaseConfig) {
        let cfg = new distillation.ProcessPhaseConfig()
        cfg.heaters = []
        cfg.gpio = []

        // Next
        let next = new process.MoveToNextConfig()
        next.type = p.next.type
        next.sensor_id = p.next.sensorID
        next.sensor_threshold = Number(p.next.sensorThreshold.value)
        next.temperature_hold_seconds = Number(p.next.temperatureHoldSeconds.value)
        next.seconds_to_move = Number(p.next.secondsToMove.value)
        cfg.next = next

        // Heaters
        p.heaters.forEach((value: HeaterPhaseConfig) => {
            let heater = new process.HeaterPhaseConfig()
            heater.ID = value.id
            heater.power = Number(value.power.value)
            cfg.heaters.push(heater)
        })
        // GPIO
        p.gpios.forEach((value: GPIOPhaseConfig) => {
            let gpio = new process.GPIOPhaseConfig()
            gpio.id = value.id
            gpio.sensor_id = value.sensor_id
            gpio.t_low = Number(value.t_low.value)
            gpio.t_high = Number(value.t_high.value)
            gpio.hysteresis = Number(value.hysteresis.value)
            gpio.inverted = value.inverted

            cfg.gpio.push(gpio)
        })

        PhasesSetConfig(p.id, cfg)
    }


    get next_type(): boolean {
        return this.next.type == 1
    }

    set next_type(v: boolean) {
        this.next.type = v ? 1 : 0
        this.update(this)
    }

    get next_timeleft(): Parameter {
        if (this.next.type) {
            return this.next.secondsToMove
        } else {
            return this.next.temperatureHoldSeconds
        }
    }

    get next_sensor_threshold(): Parameter {
        return this.next.sensorThreshold
    }

    get next_avail_sensors(): string[] {
        return this.avail_sensors
    }

    get next_sensor(): string {
        return this.next.sensorID
    }

    set next_sensor(v: string) {
        this.next.sensorID = v
        this.update(this)
    }

    get heaters(): HeaterPhaseConfig[] {
        return this.heaters_
    }

    get gpios(): GPIOPhaseConfig[] {
        return this.gpios_
    }
}

export class Phases {
    phases: ProcessPhaseConfig[];
    phaseCount: Parameter;

    constructor(phases: ProcessPhaseConfig[] = [], phaseCount: number = 0) {
        this.phases = phases
        this.phaseCount = new Parameter(phaseCount, false, this.setPhaseCount)
    }

    private setPhaseCount(cnt: number) {
        PhasesSetPhaseCount(cnt)
    }
}