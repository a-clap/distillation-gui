<script setup>
import Keyboard from "../components/Keyboard.vue"
import { reactive } from "vue";
import { PT100 } from "../types/PT100";

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
        // GetValues().then((result) => {
        //     result.forEach(element => {
        //         console.log(element.name)
        //         console.log(element.value)
        //         // console.log(element.Name)
        //     });
        // })
        delayF.show = !delayF.show
    }
})

const param = reactive(new PT100("name", 130.0, 10, 13.0))

function toggle() {
    console.log("toggle")
}

</script>
<template>
    <main class="ds-page">
        <input v-model="param.correction.value" @click="param.correction.showKeyboard">
        <Keyboard v-bind="param.correction" :write="(e) => param.correction.write(e)"
            :cancel="() => param.correction.cancel()" />
        <br>
        <!-- < input v - model=" param.value" @click="delayF.toggle"> -->
        <!-- <Keyboard v-bind="delayF" @enter="delayF.enter" @cancel="delayF.cancel" /> -->
    </main>
</template>
<style lang="scss" scoped>
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