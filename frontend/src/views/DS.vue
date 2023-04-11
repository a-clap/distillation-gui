<script setup>
import Keyboard from "../components/Keyboard.vue"
import { ref, onMounted } from "vue";

const visible = ref(false)
const input = ref("test")

function toggle() {
    console.log(visible.value)
    visible.value = !visible.value
}

</script>

<template>
    <main class="ds-page">
        <h1>{{ $t('ds.title') }}</h1>
        <button @click="toggle">add 1 </button>
        <Keyboard layout="normal" v-bind:visible="visible" />
    </main>
</template>