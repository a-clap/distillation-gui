<template>
    <main>
        <h1>{{ $t('pt100.title') }}</h1>
        <div v-for="(pt, index) in ptStore.pt" :key="index">
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
import Keyboard from "../components/Keyboard.vue"
import { usePTStore } from "../stores/pt";

const ptStore = usePTStore()

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

