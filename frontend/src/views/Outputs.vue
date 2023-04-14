<template>
    <main class="outputs-page">
        <h1>{{ $t('outputs.title') }}</h1>
    </main>
</template>

<script setup>

import { reactive, onMounted } from "vue"
import { Rectangle } from '../types/GPIO.js';

// const gp = reactive({
// value: GPIO,
// })

onMounted(() => {
    let r = new Rectangle(1, 2)
    r.callGo()
})

</script>

