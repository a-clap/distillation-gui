<script setup>
import Keyboard from "../components/Keyboard.vue"
import { ref } from "vue";

const normal = ref(false)
const numeric = ref(false)

function toggle() {
    normal.value = !normal.value
}
</script>

<template>
    <main class="ds-page">
        <h1>{{ $t('ds.title') }}</h1>
        <button @click="toggle">normal</button>
        <br>
        <button @click="numeric = !numeric">numeric {{ numeric }}</button>
        <Keyboard layout="numeric" v-bind:visible="numeric" />
        <Keyboard layout="normal" v-bind:visible="normal" />
    </main>
</template>