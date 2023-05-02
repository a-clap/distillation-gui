import { PhasesSetConfig, PhasesSetPhaseCount } from "../../wailsjs/go/backend/Backend";
import { distillation, process } from "../../wailsjs/go/models";
import Parameter, { writeCallbackType } from "./Parameter";

declare type Notify = (args: ProcessPhaseConfig) => void;
class MoveToNextConfig {
    type: number;
    sensorID: string;
    sensorThreshold: Parameter;
    timeleft: Parameter;

    constructor(n: process.MoveToNextConfig, cb: Notify, args: any) {
        let callback = function (_: any) {
            cb(args)
        }
        this.type = n.type
        this.sensorID = n.sensor_id
        this.sensorThreshold = new Parameter(n.sensor_threshold, true, callback)
        this.timeleft = new Parameter(n.time_left, false, callback)
    }
}

class GPIOConfig {
    enabled: boolean;
    id: string;
    t_low: Parameter;
    t_high: Parameter;
    hysteresis: Parameter;
    private sensorID: string;
    private inverted_: boolean;
    private callback: writeCallbackType;

    constructor(gpio: process.GPIOConfig, callback: Notify, args: any) {
        this.callback = function (_: any = 0) {
            callback(args)
        }
        this.enabled = gpio.enabled
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
    private gpios_: GPIOConfig[];
    private sensors: string[];

    constructor(id: number, next: process.MoveToNextConfig, heaters: process.HeaterPhaseConfig[], gpios: process.GPIOConfig[], sensors: string[]) {
        this.id = id
        this.next = new MoveToNextConfig(next, this.update, this)
        this.heaters_ = []
        this.gpios_ = []
        this.sensors = sensors

        if (heaters != null) {
            heaters.forEach((v: process.HeaterPhaseConfig) => {
                this.heaters_.push(new HeaterPhaseConfig(v, this.update, this))
            })
        }
        if (gpios != null) {
            gpios.forEach((v: process.GPIOConfig) => {
                this.gpios_.push(new GPIOConfig(v, this.update, this))
            })
        }

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
        next.time_left = Number(p.next.timeleft.value)
        cfg.next = next

        // Heaters
        p.heaters.forEach((value: HeaterPhaseConfig) => {
            let heater = new process.HeaterPhaseConfig()
            heater.ID = value.id
            heater.power = Number(value.power.value)
            cfg.heaters.push(heater)
        })
        // GPIO
        p.gpios.forEach((value: GPIOConfig) => {
            let gpio = new process.GPIOConfig()
            gpio.enabled = value.enabled
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
        return this.next.timeleft
    }

    get next_sensor_threshold(): Parameter {
        return this.next.sensorThreshold
    }

    get next_avail_sensors(): string[] {
        return this.sensors
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

    get gpios(): GPIOConfig[] {
        return this.gpios_
    }
}

export class Phases {
    phases: ProcessPhaseConfig[];
    gpios: GPIOConfig[];
    phaseCount: Parameter;

    constructor(phases: ProcessPhaseConfig[] = [], gpios: process.GPIOConfig[] = []) {
        this.phases = phases
        this.gpios = []
        gpios.forEach((v: process.GPIOConfig) => {
            // this.gpios.push(new GPIOConfig())
        })
        this.phaseCount = new Parameter(phases.length, false, this.setPhaseCount)
    }

    private setPhaseCount(cnt: number) {
        PhasesSetPhaseCount(cnt)
    }
}