export class Parameter {

    constructor(name, value, isFloat, writeCallback) {
        this.name_ = name
        this.value = value
        this.isFloat = isFloat
        this.show = false
        this.writeCallback_ = writeCallback
    }

    showKeyboard() {
        this.show = true
    }

    cancel() {
        this.show = false
    }

    write(value) {
        this.value = value
        this.writeCallback_(value)
        this.cancel()
    }
}
