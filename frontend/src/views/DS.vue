<template>
    <main>
        <h1>{{ $t('ds.title') }}</h1>
        <div v-for="(ds, index) in getDses" :key="index">
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
                        <Keyboard v-bind="ds.correction" :write="(e: number) => ds.correction.write(e)"
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
                        <Keyboard v-bind="ds.samples" :write="(e: number) => ds.samples.write(e)"
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
import { reactive, onMounted, computed } from "vue"
import { DS } from '../types/DS';

const dses: DS[] = reactive([])

onMounted(() => {
    dses.push(new DS("ds_1", 1, 2, 9, 3.0))
    dses.push(new DS("ds_2", 3, 4, 10, 5.0))
    dses.push(new DS("ds_3", 6, 7, 11, 8.0))
})

const getDses = computed(() => {
    return dses;
})
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