import Parameter from './Parameter'
import { DSEnable, DSSetCorrection, DSSetName, DSSetResolution, DSSetSamples, } from '../../wailsjs/go/backend/Backend'

export class DS {

    name: Parameter;
    id: string;
    correction: Parameter;
    samples: Parameter;
    resolution_: number;
    private temperature_: string;
    private enabled: boolean;

    constructor(name: string, id: string, enabled: boolean, correction: number, samples: number, resolution: number, temperature: number = 0) {
        this.id = id
        this.name = new Parameter(name, true, this.writeName)
        this.correction = new Parameter(correction, true, this.writeCorrection)
        this.samples = new Parameter(samples, false, this.writeSamples)
        this.resolution_ = resolution
        this.temperature_ = ""
        this.enabled = enabled

        this.temperature = temperature
    }

    set temperature(t: string | number) {
        if (typeof t === 'number') {
            this.temperature_ = t.toFixed(2)
        } else {
            this.temperature_ = t
        }
    }

    get temperature(): string | number {
        return this.temperature_
    }

    writeCorrection(value: number) {
        this.correction.value = value
        DSSetCorrection(this.id, value)
    }

    writeName(value: string) {
        this.name.value = value
        DSSetName(this.id, value)
    }

    writeSamples(value: number) {
        this.samples.value = value
        DSSetSamples(this.id, value)
    }

    set resolution(value: string) {
        this.resolution_ = Number(value)
        DSSetResolution(this.id, this.resolution_)
    }
    get resolution(): string {
        return this.resolution_.toString()
    }
    set enable(value: boolean) {
        this.enabled = value
        DSEnable(this.id, value)
    }

    get enable(): boolean {
        return this.enabled
    }
}
