<template>
    <main>
        <h1>{{ $t('outputs.title') }}</h1>
        <div v-for="(gpio, index) in gpios" :key="index">
            <section class="gpio-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <div> {{ gpio.name }} </div>
                    </el-col>
                    <el-col :span="4">
                        <span>{{ $t('outputs.active_level') }}</span>
                    </el-col>
                    <el-col :span="7">
                        <el-radio-group v-model="gpio.activeLevel" size="large">
                            <el-radio-button label=false> {{ $t('outputs.active_level_low') }}</el-radio-button>
                            <el-radio-button label=true> {{ $t('outputs.active_level_high') }}</el-radio-button>
                        </el-radio-group>
                    </el-col>
                    <el-col :span="6">
                        <span>{{ $t('outputs.manual_control') }}</span>
                    </el-col>
                    <el-col :span="4">
                        <el-checkbox v-model="gpio.state" :label="$t('outputs.manual_control_force')" size="large" border />
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup lang="ts">

import { ref, onMounted, onUnmounted, computed } from "vue"
import { GPIO } from '../types/GPIO';
import { GPIOListener } from "../types/GPIOListener";
import { parameters } from "../../wailsjs/go/models";
import { GPIOGet } from "../../wailsjs/go/backend/Backend";

const gpios = ref<GPIO[]>([]);

onMounted(() => {
    reload()
    GPIOListener.subscribe(update)
})

onUnmounted(() => {
    GPIOListener.unsubscribe(update)
})

function reload() {
    console.log("reload")
    GPIOGet().then((got : parameters.GPIO[]) => {
        let newGPIO: GPIO[] = []
        got.forEach((p: parameters.GPIO) => {
            let gp = new GPIO(p.id, p.active_level, p.value)
            newGPIO.push(gp)
        })

        gpios.value = newGPIO.sort((a: GPIO, b: GPIO) => {
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

function update(p: parameters.GPIO) {
    gpios.value.some(function (item: GPIO, i: number) {
        if (item.name == p.id) {
            let gp = new GPIO(p.id, p.active_level, p.value)
            gpios.value[i] = gp
        }
    });
}

</script>
<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.gpio-box {
    font-size: 1.2rem;
    margin-bottom: 2rem;
}

.el-checkbox {
    --el-color-primary: var(--el-color-danger);
    --el-checkbox-text-color: var(--el-color-danger);
}
</style>

