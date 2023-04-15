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

<script setup>

import { reactive, onMounted, computed } from "vue"
import { Heater } from '../types/Heater.js';

const heaters = reactive([])

onMounted(() => {
    heaters.push(new Heater("heater_1", false))
    heaters.push(new Heater("heater_2", true))
})

const getHeaters = computed(() => {
    return heaters;
})

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

