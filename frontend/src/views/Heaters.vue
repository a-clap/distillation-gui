<template>
    <main>
        <h1>{{ $t('heaters.title') }}</h1>
        <div v-for="(heater, index) in heatersStore.heaters" :key="index">
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

import { useHeatersStore } from "../stores/heaters";
const heatersStore = useHeatersStore()

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

