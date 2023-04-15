<script setup>
import Keyboard from "../components/Keyboard.vue"
import { reactive } from "vue";
import { PT100 } from "../types/PT100";

const param = reactive(new PT100("name", 130.0, 10, 13.0))

</script>
<template>
    <main class="ds-page">
        <input v-model="param.correction.value" @click="param.correction.showKeyboard">
        <Keyboard v-bind="param.correction" :write="(e) => param.correction.write(e)"
            :cancel="() => param.correction.cancel()" />
    </main>
</template>
<style lang="scss" scoped></style>