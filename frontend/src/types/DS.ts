import Parameter from './Parameter'
import { DSEnable, DSGet, DSSetCorrection, DSSetResolution, DSSetSamples, } from '../../wailsjs/go/backend/Backend'

export class DS {

    name: string;
    id: string;
    correction: Parameter;
    samples: Parameter;
    resolution: number;
    temperature: number;
    enabled: boolean;

    constructor(name: string, id: string, enabled: boolean, correction: number, samples: number, resolution: number, temperature: number) {
        this.name = name
        this.id = id
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, false, this.writeSamples)
        this.resolution = resolution
        this.temperature = temperature;
        this.enabled = enabled
    }

    writeCorrection(value: number) {
        this.correction.value = value
        DSSetCorrection(this.id, value)
    }

    writeResolution(value: number) {
        this.resolution = value
        DSSetResolution(this.id, value)
    }

    writeSamples(value: number) {
        this.samples.value = value
        DSSetSamples(this.id, value)
    }

    set enable(value: boolean) {
        this.enabled = value
        DSEnable(this.id, value)
    }

    get enable(): boolean {
        return this.enabled
    }
}

