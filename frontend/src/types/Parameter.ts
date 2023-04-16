export interface writeCallbackType { (value: number): void }

export default class Parameter {
    value: number;
    isFloat: boolean;
    show: boolean;
    writeCallback: writeCallbackType;


    constructor(value: number, isFloat: boolean, writeCallback: writeCallbackType) {
        this.value = value
        this.isFloat = isFloat
        this.show = false
        this.writeCallback = writeCallback
    }

    showKeyboard() {
        this.show = true
    }

    cancel() {
        this.show = false
    }

    write(value: number) {
        this.value = value
        this.writeCallback(value)
        this.cancel()
    }
}
