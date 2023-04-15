import { Parameter } from "./Parameter";

export class PT100 {
    constructor(name, correction, samples, temperature) {
        this.name = name
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, true, this.writeSamples)
        this.temperature = temperature;
        this.enable = false
    }

    writeCorrection(value) {
        console.log("write correction " + value)
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
}

