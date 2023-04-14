<script setup>
import Keyboard from "../components/Keyboard.vue"
import { ref, reactive } from "vue";

const delay = reactive({
    show: false,
    value: 0.1,
    isFloat: true,

    enter(newValue) {
        delay.value = newValue
        delay.cancel()
    },

    cancel() {
        delay.show = false
    },

    toggle() {
        delay.show = !delay.show
    }
})

const delayF = reactive({
    show: false,
    value: 100.1,
    isFloat: true,

    enter(newValue) {
        delayF.value = newValue
        delayF.cancel()
    },

    cancel() {
        delayF.show = false
    },

    toggle() {
        delayF.show = !delayF.show
    }
})
const delayStr = reactive({
    show: false,
    value: "0",

    enter(newValue) {
        delayStr.value = newValue
        delayStr.cancel()
    },

    cancel() {
        delayStr.show = false
    },

    toggle() {
        delayStr.show = !delayStr.show
    }
})

</script>

<template>
    <main class="ds-page">
        <h1>{{ $t('ds.title') }}</h1>

        <input v-model="delayF.value" @click="delayF.toggle">
        <Keyboard v-bind="delayF" @enter="delayF.enter" @cancel="delayF.cancel" />
    </main>
</template>
<style lang="scss">
input {
    display: block;
    width: 100px;
    height: 34px;
    padding: 6px 12px;
    line-height: 1rem;
    text-align: center;
    color: #555;
    cursor: default;
    caret-color: transparent;
    background-color: #fff;
    background-image: none;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
}
</style>