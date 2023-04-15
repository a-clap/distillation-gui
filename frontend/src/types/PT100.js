class PT100 {
    constructor(name, correction, samples, temperature) {
        this.name = name || "";
        this.correction = correction || 0;
        this.samples = samples || 0;
        this.temperature = temperature || 0;
    }   
}

export {PT100}