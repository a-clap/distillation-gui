<template>
    <main>
        <h1>{{ $t('pt100.title') }}</h1>
        <div v-for="(pt, index) in getPT100s" :key="index">
            <section class="pt-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-checkbox v-model="pt.enable" :label="pt.name" size="large" border />
                    </el-col>
                    <el-col :span="3">
                        <input />
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

input {
    display: block;
    width: 100px;
    height: 34px;
    padding: 6px 12px;
    line-height: 1rem;
    text-align: center;
    color: #555;
    cursor: default;
    caret-color: transparent;
    background-color: #fff;
    background-image: none;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
}
</style>

