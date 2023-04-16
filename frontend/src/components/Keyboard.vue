<template>
    <main :class="getWindowClass" v-if="props.show">
        <input class="text" v-model="keyboardValue.value">
        <div v-for="(keys, index) in keySet" :key="index">
            <section class="line">
                <div v-for="(key, index) in keys" :key="index" :class="getClassesOfKey(key)" @click="e => clickKey(e, key)">
                    {{ getCaptionOfKey(key) }}
                </div>
            </section>
        </div>
    </main>
</template>

<script setup lang="ts">

import { computed, watch, ref, reactive } from "vue";
import Layouts from "./KeyboardLayouts";
import isObject from "lodash/isObject";

const props = defineProps({
    value: [String, Number],
    isFloat: Boolean,
    show: Boolean,
    write: Function,
    cancel: Function,
    options: {
        type: Object,
        default() {
            return {};
        }
    }
})

interface valueType {
    value: any;
    layout: string;
    add: (ch: string) => void;
    clr: () => void;
    get: () => any;
}

const intValue = reactive({
    value: 0,
    layout: "numeric",
    justCleared: true,

    update(val: number) {
        intValue.value = val
    },
    add(ch: string) {
        let val = intValue.value.toString()

        if (!intValue.justCleared && val !== "0") {
            switch (ch) {
                case '.':
                    return
                case '-':
                    if (val.startsWith('-')) {
                        val = val.slice(-(val.length - 1))
                    } else {
                        val = "-" + val
                    }
                    intValue.update(Number(val))
                    return
            }
            val += ch
        } else {
            intValue.justCleared = false
            val = ch
        }
        intValue.update(Number(val))
    },
    clr() {
        intValue.justCleared = true
        intValue.value = 0
    },
    get() {
        return Number(intValue.value)
    }
})

const stringValue = reactive({
    value: "",
    layout: "normal",
    justCleared: true,
    add(ch: string) {
        if (!stringValue.justCleared) {
            stringValue.value += ch
        } else {
            stringValue.justCleared = false
            stringValue.value = ch
        }
    },
    clr() {
        stringValue.justCleared = true
        stringValue.value = ""
    },
    get() {
        return stringValue.value
    }
})

const floatValue = reactive({
    value: 0.1,
    layout: "numeric",
    justCleared: true,

    update(val: number) {
        floatValue.value = val
    },
    add(ch: string) {
        let val = floatValue.value.toString()
        if (!floatValue.justCleared) {
            switch (ch) {
                case '.':
                    if (val.includes('.')) {
                        return
                    }
                    break
                case '-':
                    if (val.startsWith('-')) {
                        val = val.slice(-(val.length - 1))
                    } else {
                        val = "-" + val
                    }
                    floatValue.update(Number(val))
                    return
            }
            val += ch
        } else {
            floatValue.justCleared = false
            val = ch
        }
        floatValue.update(Number(val))
    },

    clr() {
        floatValue.justCleared = true
        floatValue.value = 0
    },
    get() {
        return Number(floatValue.value)
    }
})

const keyboardValue: valueType = ref(stringValue)
const isShifted = ref(false)

watch(() => props.show, (trigger) => {
    // Will get called on each show change
    if (trigger) {
        // Check type of input
        if (typeof props.value === 'number') {
            if (props.isFloat) {
                keyboardValue.value = floatValue
            } else {
                keyboardValue.value = intValue
            }
        } else {
            // So string
            keyboardValue.value = stringValue
        }
        keyboardValue.value.value = props.value
    }
});

const getWindowClass = computed(() => {
    return keyboardValue.value.layout + " keyboard-window"
})

const keySet = computed(() => {
    let layout = getLayout();
    if (!layout) {
        return;
    }

    let keys = layout[isShifted.value ? "shifted" : "default"];
    if (!keys) {
        return;
    }

    let res: string[] = [];
    let meta = Layouts["_meta"] || {};
    keys.forEach((line) => {
        let row = [];
        line.split(" ").forEach((item) => {
            if (isObject(item)) {
                row.push(item);
                return
            }
            if (isSpecial(item)) {
                row.push(meta[item]);
            } else {
                row.push({
                    key: item,
                    text: item
                });
            }
        });
        res.push(row);
    });
    return res;
})

function isSpecial(name: string) {
    return Layouts["_meta"][name]
}


function getLayout(): string[] {
    return Layouts[keyboardValue.value.layout]
}

function getCaptionOfKey(key: any) {
    return key.text || key.key || "";
}

function getClassesOfKey(key) {
    let classes = "key " + (key.func || "") + " " + (key.classes || "");
    if (key.size) {
        classes += " size-" + key.size.toString() + " "
    }
    return classes;
}

function clickKey(_: any, key: any) {
    if (key.func) {
        switch (key.func) {
            case "enter":
                enter()
                break
            case "shift":
                shift()
                break
            case "esc":
                esc()
                break
            case "backspace":
                backspace()
                break
            case "clr":
                keyboardValue.value.clr()
                break

        }
        return
    }
    let ch = key.key
    keyboardValue.value.add(ch)
}

function shift() {
    isShifted.value = !isShifted.value
}

function backspace() {
    let val = keyboardValue.value.get().toString()
    if (val.length > 1) {
        val = val.slice(0, val.length - 1)
        keyboardValue.value.value = val
    } else {
        keyboardValue.value.clr()
    }
}

function esc() {
    props.cancel()
}

function enter() {
    let value = keyboardValue.value.get()
    props.write(value)
}


</script>
<style lang="scss">
$height: 2.2rem;
$margin: 0.5rem;
$radius: 0.35rem;

.numeric {
    width: 500px;
}

.normal {
    width: 900px;
}

.keyboard-window {
    z-index: 1000;
    position: fixed;
    margin: auto auto auto auto;
    top: 25%;
    left: 25%;

    padding: 1rem;
    background-color: #EEE;
    box-shadow: 0px 0px 20px rgba(black, 0.3);

    border-radius: 10px;

    .text {
        text-align: center;
        font-size: 1.5rem;
    }

    input {
        width: 100%;
        margin: 0.5rem auto 0.5rem;
    }

    $sizes: 10;

    @for $i from 1 through $sizes {
        .size-#{$i} {
            flex: $i;
        }
    }

    .keyboard {
        width: 100%;
        margin: 0;
    }

    .line {
        display: flex;
        justify-content: space-around;
        margin: 0 auto 7px;
        line-height: $height;

        &:not(:last-child) {
            margin-bottom: $margin;
        }
    }

    .key {
        &:not(:last-child) {
            margin-right: $margin;
        }

        flex: 1;
        overflow: hidden;

        vertical-align: middle;
        border: 1px solid #ccc;
        color: #333;
        background-color: #fff;
        box-shadow: 0px 2px 2px rgba(0, 0, 0, .6);
        border-radius: $radius;
        font-size: 1.25em;
        text-align: center;
        white-space: nowrap;
        user-select: none;
        cursor: pointer;


        &.icons {
            font-family: "Material Icons";
            position: relative;
            vertical-align: middle;
            font-size: 2rem;
        }

        &.backspace:before {
            content: "\e14a";
        }

        &.esc:before {
            content: "\e879";
        }

        &.enter:before {
            content: "\e5ca";
        }


        &.control {
            color: #fff;
            background-color: #7d7d7d;
            border-color: #656565;
        }

        &.featured {
            color: #fff;
            background-color: #337ab7;
            border-color: #2e6da4;
        }

        &:hover {
            color: #333;
            background-color: #d6d6d6;
            border-color: #adadad;
        }

        &:active {
            transform: scale(.90); // translateY(1px);
            color: #333;
            background-color: #d4d4d4;
            border-color: #8c8c8c;
        }
    }
}
</style>