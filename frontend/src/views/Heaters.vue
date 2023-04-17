<template>
    <main>
        <h1>{{ $t('heaters.title') }}</h1>
        <div v-for="(heater, index) in getHeaters" :key="index">
            <section class="heater-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-switch v-model="heater.enable" :active-text="heater.name" size="large" />
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup lang="ts">

import { ref, computed, onMounted, onUnmounted, } from "vue"
import { Heater } from '../types/Heater';
import { HeaterListener } from '../types/HeaterListener';

onMounted(() => {
    heaterCallback()
    HeaterListener.subscribe(heaterCallback)
})
onUnmounted(() => {
    HeaterListener.unsubscribe(heaterCallback)
})

function heaterCallback() {
    let newHeaters: Heater[] = []
    newHeaters.push(new Heater("heater_1", true))
    newHeaters.push(new Heater("heater_2", false))
    newHeaters.push(new Heater("heater_3", true))
    heaters.value = newHeaters
}

const heaters = ref<Heater[]>([]);
const getHeaters = computed(() => { return heaters.value })

</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.heater-box {
    font-size: 1.2rem;
    margin-bottom: 1.5rem;
}
</style>

