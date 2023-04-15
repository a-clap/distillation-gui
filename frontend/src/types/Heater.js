export class Heater {
    constructor(name, enabled) {
        this.name_ = name;
        this.enabled_ = enabled;
        this.state_ = false
    }

    get name() {
        return this.name_
    }

    set enable(value) {
        this.enable_ = value
        console.log("enable heater " + this.enable_)
    }
    get enable() {
        return this.enable_
    }

}