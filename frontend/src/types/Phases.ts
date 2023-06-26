import { PhasesSetConfig, PhasesSetGlobalGPIO, PhasesSetPhaseCount } from "../../wailsjs/go/backend/Backend";
import { distillation, process } from "../../wailsjs/go/models";
import Parameter, { writeCallbackType } from "./Parameter";
import { useNameStore } from "../stores/names";
import { ErrorListener } from "./ErrorListener";
import { AppErrorCodes } from "../stores/error_codes";

declare type Notify = (args: any) => void;

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
    private enabled: boolean;
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

    set enable(v: boolean) {
        this.enabled = v
        this.callback(0)
    }

    get enable(): boolean {
        return this.enabled
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

    constructor(id: number, next: process.MoveToNextConfig, heaters: process.HeaterPhaseConfig[], gpios: process.GPIOConfig[]) {
        this.id = id
        // Change ids of sensor to names
        let [newName, got] = useNameStore().id_to_name(next.sensor_id)
        if (got) {
            next.sensor_id = newName
        } else {
            ErrorListener.sendError(AppErrorCodes.SensorIDNotFound)
        }


        this.next = new MoveToNextConfig(next, this.update, this)
        this.heaters_ = []
        this.gpios_ = []

        if (heaters != null) {
            heaters.forEach((v: process.HeaterPhaseConfig) => {
                this.heaters_.push(new HeaterPhaseConfig(v, this.update, this))
            })
        }
        if (gpios != null) {
            gpios.forEach((v: process.GPIOConfig) => {
                let [newName, got] = useNameStore().id_to_name(v.sensor_id)
                if (got) {
                    v.sensor_id = newName
                } else {
                    ErrorListener.sendError(AppErrorCodes.SensorIDNotFound)
                }

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
        let [id, got] = useNameStore().name_to_id(p.next.sensorID)
        if (got) {
            next.sensor_id = id
        } else {
            next.sensor_id = p.next.sensorID
            ErrorListener.sendError(AppErrorCodes.SensorIDNotFound)
        }

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
            let [id, got] = useNameStore().name_to_id(value.sensor_id)
            if (got) {
                gpio.sensor_id = id
            } else {
                gpio.sensor_id = value.sensor_id
                ErrorListener.sendError(AppErrorCodes.SensorIDNotFound)
            }

            gpio.enabled = value.enable
            gpio.id = value.id
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
    sensors: string[]

    constructor(phases: ProcessPhaseConfig[] = [], gpios: process.GPIOConfig[] = [], sensors: string[] = []) {
        let self = this

        this.phases = phases
        this.gpios = []
        this.sensors = []
        let sensorNames: string[] = []
        if(sensors) {
            sensors.forEach((v) => {
                let [name, got] = useNameStore().id_to_name(v)
                if (got) {
                    sensorNames.push(name)
                } else {
                    ErrorListener.sendError(AppErrorCodes.SensorIDNotFound)
                    sensorNames.push(v)
                }

            })
        }
        this.sensors = sensorNames

        if (gpios != null) {
            gpios.forEach((v: process.GPIOConfig) => {
                this.gpios.push(new GPIOConfig(v, self.update, self))
            })
        }
        this.phaseCount = new Parameter(phases.length, false, this.setPhaseCount)
    }

    update(p: Phases) {
        let gpios: process.GPIOConfig[] = []
        p.gpios.forEach((v) => {
            let g = new process.GPIOConfig()
            g.enabled = v.enable
            g.id = v.id
            g.sensor_id = v.sensor_id
            g.t_low = Number(v.t_low.value)
            g.t_high = Number(v.t_high.value)
            g.hysteresis = Number(v.hysteresis.value)
            g.inverted = v.inverted

            gpios.push(g)
        })
        console.log(gpios)
        PhasesSetGlobalGPIO(gpios)
    }

    private setPhaseCount(cnt: number) {
        PhasesSetPhaseCount(cnt)
    }
}