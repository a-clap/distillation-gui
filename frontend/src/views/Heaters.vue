<template>
    <main>
        <h1>{{ $t('heaters.title') }}</h1>
        <div v-for="(heater, index) in heaters" :key="index">
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
import { parameters } from "../../wailsjs/go/models";

const heaters = ref<Heater[]>([]);

onMounted(() => {
    reload()
    HeaterListener.subscribe(updateHeater)
})
onUnmounted(() => {
    HeaterListener.unsubscribe(updateHeater)
})

function reload() {
    HeatersGet().then((got) => {
        let newHeaters: Heater[] = []
        got.forEach((heater: parameters.Heater) => {
            let newHeater = new Heater(heater.ID, heater.enabled)
            newHeaters.push(newHeater)
        })

        heaters.value = newHeaters.sort((a: Heater, b: Heater) => {
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

function updateHeater(h: parameters.Heater) {
    heaters.value.some(function (item: Heater, i: number) {
        if (item.name == h.ID) {
            console.log("got")
            let heater = new Heater(h.ID, h.enabled)
            heaters.value[i] = heater
        }
    });
}



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

