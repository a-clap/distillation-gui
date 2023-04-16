import Parameter from "./Parameter";

export class PT100 {
    name: string;
    correction: Parameter;
    samples: Parameter;
    temperature: number;
    enable_: boolean;

    constructor(name: string, correction: number, samples: number, temperature: number) {
        this.name = name
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, true, this.writeSamples)
        this.temperature = temperature;
        this.enable_ = false
    }

    writeCorrection(value: number) {
        console.log("write correction " + value)
    }

    writeSamples(value: number) {
        console.log("write samples " + value)
    }

    set enable(value) {
        this.enable_ = value
        console.log("enable " + this.enable_)
    }

    get enable(): boolean {
        return this.enable_
    }
}

