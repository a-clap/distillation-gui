<script setup>
import Keyboard from "../components/Keyboard.vue"
import { GetValues } from "../../wailsjs/go/main/App"
import { reactive } from "vue";

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
        GetValues().then((result) => {
            result.forEach(element => {
                console.log(element.name)
                console.log(element.value)
                // console.log(element.Name)
            });
        })
        delayF.show = !delayF.show
    }
})

</script>
<template>
    <main class="ds-page">
        <vm-page-header>{{ $t('ds.title') }}</vm-page-header>

        <input v-model="delayF.value" @click="delayF.toggle">
        <Keyboard v-bind="delayF" @enter="delayF.enter" @cancel="delayF.cancel" />
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