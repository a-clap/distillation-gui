<template>
    <main :class="getBaseWindow" v-if="props.visible">
        <input class="text">
        <div v-for="(keys, index) in keySet" :key="index">
            <section class="line">
                <div v-for="(key, index) in keys" :key="index" :class="getClassesOfKey(key)"> {{ getCaptionOfKey(key) }}
                </div>
            </section>
        </div>
    </main>
</template>

<script setup>
import { computed, watch } from "vue";
import { Layouts } from "./KeyboardLayouts.js";
import isString from "lodash/isString";
import isObject from "lodash/isObject";

const props = defineProps({
    visible: Boolean,
    input: [HTMLInputElement, HTMLTextAreaElement],
    layout: [String, Object],
    currentKeySet: {
        type: String,
        default: "default",
    },
    accept: Function,
    cancel: Function,
    change: Function,
    next: Function,
    options: {
        type: Object,
        default() {
            return {};
        }
    }
})


watch(() => props.visible, (first) => {
    console.log(
        "Watch props.selected function called with args:",
        first
    );
});

const getBaseWindow = computed(() => {
    let base = props.layout + " keyboard-window"
    console.log(base)
    return base
})


const keySet = computed(() => {
    let layout = getLayout();
    if (!layout) {
        return;
    }

    let keys = layout[props.currentKeySet];
    if (!keys) {
        return;
    }

    let res = [];
    let meta = layout["_meta"] || {};
    keys.forEach((line) => {
        let row = [];
        line.split(" ").forEach((item) => {
            if (isObject(item)) {
                row.push(item);
                return
            }

            if (!isString(item)) {
                return
            }

            if (item.length > 2 && item[0] == "{" && item[item.length - 1] == "}") {
                let name = item.substring(1, item.length - 1);
                if (meta[name])
                    row.push(meta[name]);
                else
                    console.error("Missing named key from meta: " + name);
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


function getLayout() {
    if (isString(props.layout)) {
        return Layouts[props.layout];
    }
    return props.layout;
}

function changeKeySet(name) {
    let layout = getLayout();
    if (layout[name] != null)
        props.currentKeySet = name;
}

function toggleKeySet(name) {
    props.currentKeySet = props.currentKeySet == name ? "default" : name;
}

function getCaptionOfKey(key) {
    return key.text || key.key || "";
}

function getClassesOfKey(key) {
    let classes = "key " + (key.func || "") + " " + (key.classes || "");
    if (key.size) {
        classes += " size-" + key.size.toString() + " "
    }
    if (key.keySet && props.currentKeySet == key.keySet) {
        classes += " activated";
    }
    return classes;

}

function supportsSelection() {
    return (/text|password|search|tel|url/).test(props.input.type);
}

function getCaret() {
    if (supportsSelection()) {
        let pos = {
            start: props.input.selectionStart || 0,
            end: props.input.selectionEnd || 0
        };
        if (pos.end < pos.start)
            pos.end = pos.start;
        return pos;
    } else {
        let val = props.input.value;
        return {
            start: val.length,
            end: val.length
        };
    }
}

function backspace(caret, text) {
    if (caret.start < caret.end) {
        text = text.substring(0, caret.start) + text.substring(caret.end);
    } else {
        text = text.substring(0, caret.start - 1) + text.substring(caret.start);
        caret.start -= 1;
    }
    caret.end = caret.start;
    return text;
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

function clickKey(e, key) {
    if (!props.input) return;
    if (props.options.preventClickEvent) e.preventDefault();
    let caret = getCaret();
    let text = props.input.value;

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
            if (props.options.useKbEvents) {
                let e = document.createEvent("Event");
                e.initEvent("keydown", true, true);
                e.which = e.keyCode = addChar.charCodeAt();
                if (props.input.dispatchEvent(e)) {
                    text = insertChar(caret, text, addChar);
                }
            } else {
                text = insertChar(caret, text, addChar);
            }
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

function setFocusToInput(caret) {
    props.input.focus();
    if (caret && supportsSelection()) {
        props.input.selectionStart = caret.start;
        props.input.selectionEnd = caret.end;
    }
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
    width: 1400px;
}

.keyboard-window {
    position: fixed;
    top: 25%;
    left: 25%;

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

        &.backspace {
            position: relative;
        }

        &.backspace:before {
            font-family: "Material Icons";
            content: "\e14a";
            position: absolute;
            top: 5%;
            left: 38%;
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
            transform: scale(.98); // translateY(1px);
            color: #333;
            background-color: #d4d4d4;
            border-color: #8c8c8c;
        }

        &.activated {
            color: #fff;
            background-color: #5bc0de;
            border-color: #46b8da;
        }
    }
}
</style>