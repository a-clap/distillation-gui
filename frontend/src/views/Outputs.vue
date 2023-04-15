<template>
    <main>
        <h1 style="margin-bottom: 1rem;">{{ $t('outputs.title') }}</h1>
        <div v-for="(gpio, index) in getGpios" :key="index">
            <section class="gpio-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <div> {{ gpio.name }} </div>
                    </el-col>
                    <el-col :span="4">
                        <span>{{ $t('outputs.active_level') }}</span>
                    </el-col>
                    <el-col :span="9">
                        <el-radio-group v-model="gpio.activeLow" size="large">
                            <el-radio-button label=false> {{ $t('outputs.active_level_low') }}</el-radio-button>
                            <el-radio-button label=true> {{ $t('outputs.active_level_high') }}</el-radio-button> >
                        </el-radio-group>
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup>

import { reactive, onMounted, computed } from "vue"
import { GPIO } from '../types/GPIO.js';

const gpios = reactive([])

onMounted(() => {
    gpios.push(new GPIO("blah", false))
    gpios.push(new GPIO("blah2", true))
})

const getGpios = computed(() => {
    console.log(gpios)
    return gpios;
})

</script>
<style lang="scss" scoped>
.outputs-page {
    margin-bottom: 2rem;
}

.gpio-box {
    font-size: 1.2rem;
    margin-bottom: 2rem;
}
</style>

