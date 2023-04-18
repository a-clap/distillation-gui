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
import { HeatersGet } from "../../wailsjs/go/backend/Backend";

onMounted(() => {
    heaterCallback()
    HeaterListener.subscribe(heaterCallback)
})
onUnmounted(() => {
    HeaterListener.unsubscribe(heaterCallback)
})

function heaterCallback() {
    HeatersGet().then((got) => {
        let newHeaters: Heater[] = []
        got.forEach((heater) => {
            let newHeater = new Heater(heater.ID, heater.enabled)
            newHeaters.push(newHeater)
        })

        heaters.value = newHeaters.sort((a: Heater, b:Heater) => {
        console.log(a.name + " " + b.name)
        if (a.name > b.name) {
            return 1
        }
        if (a.name < b.name) {
            return -1
        }
        return 0
    })
    })
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

