<template>
    <main>
        <h1>{{ $t('pt100.title') }}</h1>
        <div v-for="(pt, index) in pt100s" :key="index">
            <section class="pt-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-switch v-model="pt.enable" :active-text="pt.name" size="large" />
                    </el-col>
                    <el-col :span="4" :offset="1" v-if="pt.enable">
                        <label>{{ $t('pt100.correction') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="pt.enable">
                        <input v-model="pt.correction.value" @click="() => pt.correction.showKeyboard()">
                        <Keyboard v-bind="pt.correction" :write="(e: number) => pt.writeCorrection(e)"
                            :cancel="() => pt.correction.cancel()" />
                    </el-col>
                    <el-col :span="5" :offset="1" v-if="pt.enable">
                        <label>{{ $t('pt100.temperature') }}</label>
                    </el-col>
                    <el-col :span="6" v-if="pt.enable">
                        <input v-model="pt.temperature">
                    </el-col>
                </el-row>
                <el-row :gutter="20" align="middle">
                    <el-col :span="4" :offset="4" v-if="pt.enable">
                        <label>{{ $t('pt100.samples') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="pt.enable">
                        <input v-model="pt.samples.value" @click="() => pt.samples.showKeyboard()">
                        <Keyboard v-bind="pt.samples" :write="(e: number) => pt.writeSamples(e)"
                            :cancel="() => pt.samples.cancel()" />
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ElRow, ElCol, ElSwitch } from "element-plus";
import Keyboard from "../components/Keyboard.vue"
import { ref, onMounted, onUnmounted } from "vue"
import { PT100 } from '../types/PT100';
import { PTListener } from "../types/PTListener";
import { PTGet } from "../../wailsjs/go/backend/Backend";
import { parameters } from "../../wailsjs/go/models";

const pt100s = ref<PT100[]>([]);

onMounted(() => {
    reload()
    PTListener.subscribeConfig(updateConfig)
    PTListener.subscribeTemperature(updateTemperature)
})

onUnmounted(() => {
    PTListener.unsubscribeConfig(updateConfig)
    PTListener.unsubscribeTemperature(updateTemperature)
})

function reload() {
    PTGet().then((got: parameters.PT[]) => {
        let newPT: PT100[] = []
        got.forEach((p: parameters.PT) => {
            let ds = new PT100(p.name, p.id, p.enabled, p.correction, p.samples)
            newPT.push(ds)
        })

        pt100s.value = newPT.sort((a: PT100, b: PT100) => {
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

function updateConfig(p: parameters.PT) {
    pt100s.value.some(function (item: PT100, i: number) {
        if (item.id == p.id) {
            let pt = new PT100(p.name, p.id, p.enabled, p.correction, p.samples)
            pt.temperature = pt100s.value[i].temperature
            pt100s.value[i] = pt
        }
    });
}

function updateTemperature(t: parameters.Temperature) {
    pt100s.value.some(function (item: PT100, i: number) {
        if (item.id == t.ID) {
            pt100s.value[i].temperature = t.temperature.toFixed(2)
        }
    });

}


</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.pt-box {
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

