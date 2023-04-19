<template>
    <main class="process-page">
        <h1>{{ $t('process.title') }}</h1>
    </main>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { Listener } from '../types/Listener';
import { NotifyHeaters } from '../../wailsjs/go/backend/Events';


let v = new Listener(NotifyHeaters)

onMounted(() => {
    v.subscribe(testing)
})

onUnmounted(() => {
    v.unsubscribe(testing)
})


function testing(...args: any) {
    console.log("from func type: " + typeof args)
    console.log("from func: " + args)
}


</script>
<style lang="scss" scoped></style>

