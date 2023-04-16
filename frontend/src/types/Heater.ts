export class Heater {
    name_: string;
    enabled_: boolean;
    state_: false;
    constructor(name: string, enabled: boolean) {
        this.name_ = name;
        this.enabled_ = enabled;
        this.state_ = false
    }

    get name() {
        return this.name_
    }

    set enable(value: boolean) {
        this.enabled_ = value
        console.log("enable heater " + this.enabled_)
    }

    get enable(): boolean {
        return this.enabled_
    }

}