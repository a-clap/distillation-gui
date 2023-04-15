export class Parameter {

    constructor(value, isFloat, writeCallback) {
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
