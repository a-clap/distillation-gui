<template>
    <main>
        <h1>{{ $t('pt100.title') }}</h1>
        <div v-for="(pt, index) in getPT100s" :key="index">
            <Keyboard v-bind="pt.correction" :write="(e) => pt.correction.write(e)"
                :cancel="() => pt.correction.cancel()" />
            <Keyboard v-bind="pt.samples" :write="(e) => pt.samples.write(e)" :cancel="() => pt.samples.cancel()" />
            <section class="pt-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-checkbox v-model="pt.enable" :label="pt.name" size="large" border />
                    </el-col>
                    <el-col :span="4" :offset="1" v-if="pt.enable">
                        <label>{{ $t('pt100.correction') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="pt.enable">
                        <input v-model="pt.correction.value" @click="() => pt.correction.showKeyboard()">
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
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup>
import Keyboard from "../components/Keyboard.vue"
import { reactive, onMounted, computed } from "vue"
import { PT100 } from '../types/PT100.js';

const pt100s = reactive([])

onMounted(() => {
    pt100s.push(new PT100("pt100_1", 1, 2, 3.0))
    pt100s.push(new PT100("pt100_2", 3, 4, 5.0))
    pt100s.push(new PT100("pt100_3", 6, 7, 8.0))
})

const getPT100s = computed(() => {
    return pt100s;
})

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

