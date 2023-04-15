import { Parameter } from "./Parameter";

export class DS {
    constructor(name, correction, samples, resolution, temperature) {
        this.name = name
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, false, this.writeSamples)
        this.resolution_ = resolution
        this.temperature = temperature;
        this.enable_ = false
    }

    writeCorrection(value) {
        console.log("write correction " + value)
    }

    writeResolution(value) {
        console.log("writeResolution " + value)
    }

    writeSamples(value) {
        console.log("write samples " + value)
    }

    set enable(value) {
        this.enable_ = value
        console.log("enable " + this.enable_)
    }

    get enable() {
        return this.enable_
    }

    set resolution(value) {
        this.resolution_ = value
        this.writeResolution(value)
    }
    get resolution() {
        return this.resolution_
    }
}

