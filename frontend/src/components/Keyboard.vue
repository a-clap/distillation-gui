<template>
    <div class="vkeyboard">
        <table class="keyboard" v-for="(cline, index) in keySet" :key="index">
            <tbody>
                <td v-for="(key, index) in cline" :key="index" :class="getClassesOfKey(key)" v-text="getCaptionOfKey(key)"
                    :style="getKeyStyle(key)" @mousedown="mousedown" @click="e => clickKey(e, key)" />
            </tbody>
        </table>
    </div>
</template>

<script setup>
import { computed, ref, watch } from "vue";
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
            }
            else if (isString(item)) {
                if (item.length > 2 && item[0] == "{" && item[item.length - 1] == "}") {
                    let name = item.substring(1, item.length - 1);
                    if (meta[name])
                        row.push(meta[name]);
                    else
                        console.warn("Missing named key from meta: " + name);
                } else {
                    if (item == "") {
                        // Placeholder
                        row.push({
                            placeholder: true
                        });

                    } else {
                        // Normal key
                        row.push({
                            key: item,
                            text: item
                        });
                    }
                }
            }
        });
        res.push(row);
    });
    return res;
})

// TODO: add watch
function layout() {
    props.currentKeySet = "default";
}

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
    if (key.placeholder)
        return "placeholder";
    else {
        let classes = "key " + (key.func || "") + " " + (key.classes || "");
        if (key.keySet && props.currentKeySet == key.keySet)
            classes += " activated";
        return classes;
    }
}

function getKeyStyle(key) {
    if (key.width)
        return {
            flex: key.width
        };
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

function mousedown(e) {
    if (!props.input) return;
    if (props.options.preventClickEvent) e.preventDefault();
    props.inputScrollLeft = props.input.scrollLeft;
}

function clickKey(e, key) {
    if (!props.input) return;
    if (props.options.preventClickEvent) e.preventDefault();
    let caret = getCaret();
    let text = props.input.value;

    let addChar = null;
    if (typeof key == "object") {
        log.console(key)
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
$width: 3rem;
$height: 2.2em;
$margin: 0.5em;
$radius: 0.35em;

.vkeyboard {
    .keyboard {
        width: 100%;
        margin: 0;

        .line {
            display: flex;
            justify-content: space-around;

            &:not(:last-child) {
                margin-bottom: $margin;
            }
        }

        .key {
            &:not(:last-child) {
                margin-right: $margin;
            }

            width: $width;
            height: $height;
            line-height: $height;
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
                width: $width;
                background-image: url("./icons/backspace.svg");
                background-position: center center;
                background-repeat: no-repeat;
                background-size: 20%;
            }

            &.half {
                flex: calc($width / 2);
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

        .placeholder {
            flex: calc($width / 2);
            height: $height;
            line-height: $height;

            &:not(:last-child) {
                margin-right: $margin;
            }
        }


        &:before,
        &:after {
            content: "";
            display: table;
        }

        &:after {
            clear: both;
        }
    }

    // .keyboard
}

// .vue-touch-keyboard
</style>