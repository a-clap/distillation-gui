<template>
    <main>
        <h1>{{ $t('ds.title') }}</h1>
        <div v-for="(ds, index) in dses" :key="index">
            <section class="ds-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-switch v-model="ds.enable" :active-text="ds.name" size="large" />
                    </el-col>
                    <el-col :span="4" :offset="1" v-if="ds.enable">
                        <label>{{ $t('ds.correction') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="ds.enable">
                        <input v-model="ds.correction.value" @click="() => ds.correction.showKeyboard()">
                        <Keyboard v-bind="ds.correction" :write="(e: number) => ds.writeCorrection(e)"
                            :cancel="() => ds.correction.cancel()" />
                    </el-col>
                    <el-col :span="5" :offset="1" v-if="ds.enable">
                        <label>{{ $t('ds.temperature') }}</label>
                    </el-col>
                    <el-col :span="6" v-if="ds.enable">
                        <input v-model="ds.temperature">
                    </el-col>
                </el-row>
                <el-row :gutter="20" align="middle">
                    <el-col :span="4" :offset="4" v-if="ds.enable">
                        <label>{{ $t('ds.samples') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="ds.enable">
                        <input v-model="ds.samples.value" @click="() => ds.samples.showKeyboard()">
                        <Keyboard v-bind="ds.samples" :write="(e: number) => ds.writeSamples(e)"
                            :cancel="() => ds.samples.cancel()" />
                    </el-col>
                    <el-col :span="2" :offset=1 v-if="ds.enable">
                        <label>{{ $t('ds.resolution') }}</label>
                    </el-col>
                    <el-col :span="6" :offset=2 v-if="ds.enable">
                        <el-select v-model="ds.resolution" size="large">
                            <el-option :label="$t('ds.resolution_9')" value="9" />
                            <el-option :label="$t('ds.resolution_10')" value="10" />
                            <el-option :label="$t('ds.resolution_11')" value="11" />
                            <el-option :label="$t('ds.resolution_12')" value="12" />
                        </el-select>
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>
<script setup lang="ts">

import Keyboard from "../components/Keyboard.vue"
import { ref, onMounted, onUnmounted } from "vue"
import { DS } from '../types/DS';
import { DSListener } from "../types/DSListener";
import { DSGet } from "../../wailsjs/go/backend/Backend";
import { parameters } from "../../wailsjs/go/models";

const dses = ref<DS[]>([]);

onMounted(() => {
    reload()
    DSListener.subscribeConfig(updateConfig)
    DSListener.subscribeTemperature(updateTemperature)
})

onUnmounted(() => {
    DSListener.unsubscribeConfig(updateConfig)
    DSListener.unsubscribeTemperature(updateTemperature)
})

function reload() {
    DSGet().then((got) => {
        let newDses: DS[] = []
        got.forEach((d: parameters.DS) => {
            let ds = new DS(d.name, d.id, d.enabled, d.correction, d.samples, d.resolution)
            newDses.push(ds)
        })

        dses.value = newDses.sort((a: DS, b: DS) => {
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

function updateConfig(d: parameters.DS) {
    dses.value.some(function (item: DS, i: number) {
        if (item.id == d.id) {
            let ds = new DS(d.name, d.id, d.enabled, d.correction, d.samples, d.resolution)
            ds.temperature = dses.value[i].temperature
            dses.value[i] = ds
        }
    });
}

function updateTemperature(t: parameters.Temperature) {
    dses.value.some(function (item: DS, i: number) {
        if (item.id == t.ID) {
            dses.value[i].temperature = t.temperature
        }
    });

}



</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.ds-box {
    margin-bottom: 1.5rem;
}

label {
    text-align: right;
    margin-left: 0;
}

.el-row {
    margin-bottom: 0.5rem;
}
</style>