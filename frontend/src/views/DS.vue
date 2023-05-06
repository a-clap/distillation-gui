<template>
    <main>
        <h1>{{ $t('ds.title') }}</h1>
        <div v-for="(ds, index) in dsStore.ds" :key="index">
            <section class="ds-box">
                <el-row>
                    <el-col :span="2">
                        <el-switch v-model="ds.enable" :active-text="ds.name" size="large" />
                    </el-col>
                    <el-col :span="4" :offset="3" v-if="ds.enable">
                        <label>{{ $t('ds.correction') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="ds.enable">
                        <input v-model="ds.correction.value" @click="() => ds.correction.showKeyboard()">
                        <Keyboard v-bind="ds.correction" :write="(e: number) => ds.writeCorrection(e)"
                            :cancel="() => ds.correction.cancel()" />
                    </el-col>
                    <el-col :span="5" v-if="ds.enable">
                        <label>{{ $t('ds.temperature') }}</label>
                    </el-col>
                    <el-col :span="6" v-if="ds.enable">
                        <input v-model="ds.temperature">
                    </el-col>
                </el-row>
                <el-row align="middle">
                    <el-col :span="4" :offset=5 v-if="ds.enable">
                        <label>{{ $t('ds.samples') }}</label>
                    </el-col>
                    <el-col :span="4" v-if="ds.enable">
                        <input v-model="ds.samples.value" @click="() => ds.samples.showKeyboard()">
                        <Keyboard v-bind="ds.samples" :write="(e: number) => ds.writeSamples(e)"
                            :cancel="() => ds.samples.cancel()" />
                    </el-col>
                    <el-col :span="2" v-if="ds.enable">
                        <label>{{ $t('ds.resolution') }}</label>
                    </el-col>
                    <el-col :span="6" :offset=2 v-if="ds.enable">
                        <el-select v-model="ds.resolution" size="large" class="m-2">
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
import { useDSStore } from "../stores/ds";

const dsStore = useDSStore()

</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.el-select {
    display: block;
    padding: 0;
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