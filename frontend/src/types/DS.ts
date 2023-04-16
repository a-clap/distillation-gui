import Parameter from './Parameter'

export class DS {

    name: string;
    correction: Parameter;
    samples: Parameter;
    resolution_: number;
    temperature: number;
    enable_: boolean;

    constructor(name: string, correction: number, samples: number, resolution: number, temperature: number) {
        this.name = name
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, false, this.writeSamples)
        this.resolution_ = resolution
        this.temperature = temperature;
        this.enable_ = false
    }

    writeCorrection(value: number) {
        console.log("write correction " + value)
    }

    writeResolution(value: number) {
        console.log("writeResolution " + value)
    }

    writeSamples(value: number) {
        console.log("write samples " + value)
    }

    set enable(value: boolean) {
        this.enable_ = value
        console.log("enable " + this.enable_)
    }

    get enable(): boolean {
        return this.enable_
    }

    set resolution(value: number) {
        this.resolution_ = value
        this.writeResolution(value)
    }
    get resolution(): number {
        return this.resolution_
    }
}

