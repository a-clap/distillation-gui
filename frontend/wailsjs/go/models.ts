export namespace distillation {
	
	export class ProcessConfigValidation {
	    valid: boolean;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessConfigValidation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.valid = source["valid"];
	        this.error = source["error"];
	    }
	}
	export class ProcessPhaseConfig {
	    next: process.MoveToNextConfig;
	    heaters: process.HeaterPhaseConfig[];
	    gpio: process.GPIOConfig[];
	
	    static createFrom(source: any = {}) {
	        return new ProcessPhaseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.next = this.convertValues(source["next"], process.MoveToNextConfig);
	        this.heaters = this.convertValues(source["heaters"], process.HeaterPhaseConfig);
	        this.gpio = this.convertValues(source["gpio"], process.GPIOConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProcessPhaseCount {
	    phase_number: number;
	
	    static createFrom(source: any = {}) {
	        return new ProcessPhaseCount(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.phase_number = source["phase_number"];
	    }
	}
	export class ProcessStatus {
	    running: boolean;
	    done: boolean;
	    phase_number: number;
	    // Go type: time
	    start_time: any;
	    // Go type: time
	    end_time: any;
	    // Go type: process
	    next: any;
	    heaters: process.HeaterPhaseStatus[];
	    temperature: process.TemperaturePhaseStatus[];
	    gpio: process.GPIOPhaseStatus[];
	    errors: string[];
	
	    static createFrom(source: any = {}) {
	        return new ProcessStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.running = source["running"];
	        this.done = source["done"];
	        this.phase_number = source["phase_number"];
	        this.start_time = this.convertValues(source["start_time"], null);
	        this.end_time = this.convertValues(source["end_time"], null);
	        this.next = this.convertValues(source["next"], null);
	        this.heaters = this.convertValues(source["heaters"], process.HeaterPhaseStatus);
	        this.temperature = this.convertValues(source["temperature"], process.TemperaturePhaseStatus);
	        this.gpio = this.convertValues(source["gpio"], process.GPIOPhaseStatus);
	        this.errors = source["errors"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace parameters {
	
	export class DS {
	    enabled: boolean;
	    name: string;
	    id: string;
	    correction: number;
	    resolution: number;
	    poll_interval: number;
	    samples: number;
	
	    static createFrom(source: any = {}) {
	        return new DS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.name = source["name"];
	        this.id = source["id"];
	        this.correction = source["correction"];
	        this.resolution = source["resolution"];
	        this.poll_interval = source["poll_interval"];
	        this.samples = source["samples"];
	    }
	}
	export class GPIO {
	    id: string;
	    direction: number;
	    active_level: number;
	    value: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GPIO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.direction = source["direction"];
	        this.active_level = source["active_level"];
	        this.value = source["value"];
	    }
	}
	export class Heater {
	    ID: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Heater(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.enabled = source["enabled"];
	    }
	}
	export class PT {
	    enabled: boolean;
	    name: string;
	    id: string;
	    correction: number;
	    a_sync_poll: boolean;
	    poll_interval: number;
	    samples: number;
	
	    static createFrom(source: any = {}) {
	        return new PT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.name = source["name"];
	        this.id = source["id"];
	        this.correction = source["correction"];
	        this.a_sync_poll = source["a_sync_poll"];
	        this.poll_interval = source["poll_interval"];
	        this.samples = source["samples"];
	    }
	}
	export class Temperature {
	    ID: string;
	    temperature: number;
	
	    static createFrom(source: any = {}) {
	        return new Temperature(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.temperature = source["temperature"];
	    }
	}

}

export namespace process {
	
	export class GPIOConfig {
	    enabled: boolean;
	    id: string;
	    sensor_id: string;
	    t_low: number;
	    t_high: number;
	    hysteresis: number;
	    inverted: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GPIOConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.id = source["id"];
	        this.sensor_id = source["sensor_id"];
	        this.t_low = source["t_low"];
	        this.t_high = source["t_high"];
	        this.hysteresis = source["hysteresis"];
	        this.inverted = source["inverted"];
	    }
	}
	export class HeaterPhaseConfig {
	    ID: string;
	    power: number;
	
	    static createFrom(source: any = {}) {
	        return new HeaterPhaseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.power = source["power"];
	    }
	}
	export class MoveToNextConfig {
	    type: number;
	    sensors: string[];
	    sensor_id: string;
	    sensor_threshold: number;
	    time_left: number;
	
	    static createFrom(source: any = {}) {
	        return new MoveToNextConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.sensors = source["sensors"];
	        this.sensor_id = source["sensor_id"];
	        this.sensor_threshold = source["sensor_threshold"];
	        this.time_left = source["time_left"];
	    }
	}
	export class PhaseConfig {
	    next: MoveToNextConfig;
	    heaters: HeaterPhaseConfig[];
	    gpio: GPIOConfig[];
	
	    static createFrom(source: any = {}) {
	        return new PhaseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.next = this.convertValues(source["next"], MoveToNextConfig);
	        this.heaters = this.convertValues(source["heaters"], HeaterPhaseConfig);
	        this.gpio = this.convertValues(source["gpio"], GPIOConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Config {
	    phase_number: number;
	    phases: PhaseConfig[];
	    global_gpio: GPIOConfig[];
	    sensors: string[];
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.phase_number = source["phase_number"];
	        this.phases = this.convertValues(source["phases"], PhaseConfig);
	        this.global_gpio = this.convertValues(source["global_gpio"], GPIOConfig);
	        this.sensors = source["sensors"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class GPIOPhaseStatus {
	    id: string;
	    state: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GPIOPhaseStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.state = source["state"];
	    }
	}
	
	export class HeaterPhaseStatus {
	    ID: string;
	    power: number;
	
	    static createFrom(source: any = {}) {
	        return new HeaterPhaseStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.power = source["power"];
	    }
	}
	
	
	export class TemperaturePhaseStatus {
	    ID: string;
	    temperature: number;
	
	    static createFrom(source: any = {}) {
	        return new TemperaturePhaseStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.temperature = source["temperature"];
	    }
	}

}

