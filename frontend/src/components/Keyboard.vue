<template>
    <main :class="getBaseWindow" v-if="props.visible">
        <input class="text" v-model="currentValue">
        <div v-for="(keys, index) in keySet" :key="index">
            <section class="line">
                <div v-for="(key, index) in keys" :key="index" :class="getClassesOfKey(key)" @click="e => clickKey(e, key)">
                    {{ getCaptionOfKey(key) }}
                </div>
            </section>
        </div>
    </main>
</template>

<script setup>

import { computed, watch, ref } from "vue";
import { Layouts } from "./KeyboardLayouts.js";
import isObject from "lodash/isObject";

const currentValue = ref("start_value")
const currentLayout = ref("normal")
const currentKeySet = ref("default")

const emit = defineEmits({
    enter: (s) => {
        if (s && (typeof s === 'string' || typeof s === 'number')) {
            return true
        } else {
            console.warn(`Invalid submit event payload!`)
            return false
        }
    },
    cancel: () => {
        return true
    }
})

const props = defineProps({
    value: [String, Number],
    visible: Boolean,
    options: {
        type: Object,
        default() {
            return {};
        }
    }
})

watch(() => props.visible, (first) => {
    // Will get called on each visible change
    // We should update value in placeholder and change currentLayout depending on type of value
    if (first) {
        // Update placeholder value
        currentValue.value = props.value.toString()
        // Reset currentKeySet to default
        currentKeySet.value = "default"
        // Change layout type
        if (Number.isInteger(props.value)) {
            currentLayout.value = "numeric"
        } else {
            currentLayout.value = "normal"
        }

    }
});

// Check what kind of window we should show 
const getBaseWindow = computed(() => {
    let base = currentLayout.value + " keyboard-window"
    return base
})


const keySet = computed(() => {
    let layout = getLayout();
    if (!layout) {
        return;
    }

    let keys = layout[currentKeySet.value];
    if (!keys) {
        return;
    }

    let res = [];
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

function isSpecial(name) {
    return Layouts["_meta"][name]
}


function getLayout() {
    return Layouts[currentLayout.value]
}

function getCaptionOfKey(key) {
    return key.text || key.key || "";
}

function getClassesOfKey(key) {
    let classes = "key " + (key.func || "") + " " + (key.classes || "");
    if (key.size) {
        classes += " size-" + key.size.toString() + " "
    }
    return classes;

}


function insertChar(caret, text, ch) {
    if (caret.start < caret.end) {
        text = text.substring(0, caret.start) + ch.toString() + text.substring(caret.end);
    } else {
        text = text.substr(0, caret.start) + ch.toString() + text.substr(caret.start);
    }
    caret.start += ch.length;
    caret.end = caret.start;
    return text;
}

function clickKey(_, key) {
    if (key.func) {
        switch (key.func) {
            case "enter":
                enter()
                return
            case "shift":
                shift()
                return
            case "esc":
                esc()
                return
        }
    } else {
    }
    return

    let text = "blah"
    let addChar = null;
    if (typeof key == "object") {
        if (key.keySet) {
            toggleKeySet(key.keySet);
        }
        else if (key.func) {
            switch (key.func) {
                case "backspace": {
                    text = backspace(caret, text);
                    break;
                }
                case "accept": {
                    if (props.accept)
                        props.accept(text);
                    return;
                }
                case "cancel": {
                    if (props.cancel)
                        props.cancel();
                    return;
                }
                case "next": {
                    if (props.next)
                        props.next();
                    return;
                }
                default: {
                    props.$emit(key.func);
                }
            }
        } else {
            addChar = key.key;
        }
    } else {
        addChar = key;
    }

    if (addChar) {
        if (props.input.maxLength <= 0 || text.length < props.input.maxLength) {
            text = insertChar(caret, text, addChar);
        }
        if (props.currentKeySet == "shifted")
            changeKeySet("default");
    }

    props.input.value = text;
    setFocusToInput(caret);
    if (props.change)
        props.change(text, addChar);
    if (props.input.maxLength > 0 && text.length >= props.input.maxLength) {
        // The value reached the maxLength
        if (props.next)
            props.next();
    }
    // trigger 'input' Event
    props.input.dispatchEvent(new Event("input", { bubbles: true }));
}

function shift() {
    if (currentKeySet.value === "default") {
        currentKeySet.value = "shifted"
    } else {
        currentKeySet.value = "default"
    }
}

function enter() {
    emit('enter', currentValue.value)
}

function backspace() {
    props.visible = false
}

function clr() {
    emit('cancel')
}

function esc() {
    emit('cancel')
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
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);

    padding: 1rem;
    background-color: #EEE;
    box-shadow: 0px 0px 20px rgba(black, 0.3);

    border-radius: 10px;

    input {
        display: block;
        width: 100%;
        height: 34px;
        padding: 6px 12px;
        font-size: 14px;
        line-height: 1.42857143;
        color: #555;
        background-color: #fff;
        background-image: none;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
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