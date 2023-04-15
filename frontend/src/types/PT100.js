import { Parameter } from "./Parameter";

export class PT100 {
    constructor(name, correction, samples, temperature) {
        this.name = name
        this.correction_ = new Parameter("correction", correction, true, this.writeCorrection)
        this.samples_ = samples;
        this.temperature_ = temperature;
        this.enable_ = false
    }

    writeCorrection(value) {
        console.log("write correction " + value)
    }


    set enable(value) {
        this.enable_ = value
        console.log("enable " + this.enable_)
    }

    get enable() {
        return this.enable_
    }

    get correction() {
        return this.correction_
    }
}

