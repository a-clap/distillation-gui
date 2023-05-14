export interface KeyboardValue {
    value: string;
    layout: string

    add(ch: string): void;
    clr(): void;
    get(): any;

}

export class FloatValue implements KeyboardValue {
    layout: string;
    private firstKey: boolean;
    private value_: string;
    constructor(v: number) {
        this.value_ = v.toFixed(2)
        if (!this.value.includes('.')) {
            this.value_ += ".0"
        }
        this.firstKey = true
        this.layout = "numeric"
    }

    get value(): string {
        return this.value_
    }
    set value(v: string) {
        this.value_ = v
    }
    
    add(ch: string): void {
        let v = this.value
        if(this.firstKey) {
            this.firstKey = false;
            v = ""
        } else if (ch === '.') {
            if (v.includes('.')) {
                return
            }
        } else if (ch === '.') {
            if (v.startsWith('-')) {
                v = v.slice(-(v.length - 1))
            } else {
                v = "-" + v
            }
            this.value = v
            return
        }
        v += ch
        this.value = v
    }

    clr(): void {
        this.firstKey = true;
        this.value = "0.00"
    }

    get(): number {
        return Number(this.value)
    }
}

export class IntValue implements KeyboardValue {
    layout: string;
    private firstKey: boolean;
    private value_: string;
    constructor(v: number) {
        this.value_ = v.toString()
        this.firstKey = true
        this.layout = "numeric"
    }

    get value(): string {
        return this.value_
    }
    set value(v: string) {
        this.value_ = v
    }
    
    add(ch: string): void {
        if (ch === '.') {
            return
        }

        let v = this.value
        if(this.firstKey) {
            this.firstKey = false;
            v = ""
        } else if (ch === '-'){
            if (v.startsWith('-')) {
                v = v.slice(-(v.length - 1))
            } else {
                v = "-" + v
            }
            this.value = v
            return
        }

        v += ch
        this.value = v
    }

    clr(): void {
        this.firstKey = true;
        this.value = "0"
    }

    get(): number {
        return Number(this.value)
    }
}

export class StringValue implements KeyboardValue {
    layout: string;
    private firstKey: boolean;
    private value_: string;
    constructor(v: string) {
        this.value_ = v.toString()
        this.firstKey = true
        this.layout = "normal"
    }

    get value(): string {
        return this.value_
    }
    set value(v: string) {
        this.value_ = v
    }
    
    add(ch: string): void {
        if(this.firstKey) {
            this.firstKey = false
            this.value = ""
        }
        this.value += ch
    }

    clr(): void {
        this.firstKey = true;
        this.value = "0"
    }

    get(): string {
        return this.value
    }
}