export interface writeCallbackType { (value: any): void }

export default class Parameter {
    value: number | string;
    isFloat: boolean;
    show: boolean;
    writeCallback: writeCallbackType;


    constructor(value: number | string, isFloat: boolean, writeCallback: writeCallbackType) {
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

    write(value: number | string) {
        this.value = value
        this.writeCallback(value)
        this.cancel()
    }
}
