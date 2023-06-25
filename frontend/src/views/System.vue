<template>
    <main>
        <h1>{{ $t('system.title') }}</h1>
        <el-row>
            <el-col class="bold-text" span="25">
                {{ $t('system.ds') }}
            </el-col>
        </el-row>
        <template v-for="(ds, index) in dsStore.ds" :key="index">
            <el-row align="middle">
                <el-col :span="5">
                    {{ ds.id }}
                </el-col>
                <el-col :span="4" :offset="2">
                    <input class="name-input" v-model="ds.name.view" @click="() => ds.name.showKeyboard()">
                    <Keyboard v-bind="ds.name" :write="(e: string) => ds.writeName(e)" :cancel="() => ds.name.cancel()" />
                </el-col>
            </el-row>
        </template>
        <el-row>
            <el-col class="bold-text" span="25">
                {{ $t('system.pt') }}
            </el-col>
        </el-row>
        <template v-for="(pt, index) in ptStore.pt" :key="index">
            <el-row align="middle">
                <el-col :span="5">
                    {{ pt.id }}
                </el-col>
                <el-col :span="4" :offset="2">
                    <input class="name-input" v-model="pt.name.view" @click="() => pt.name.showKeyboard()">
                    <Keyboard v-bind="pt.name" :write="(e: string) => pt.writeName(e)" :cancel="() => pt.name.cancel()" />
                </el-col>
            </el-row>
        </template>
    </main>
</template>

<script setup lang="ts">

import Keyboard from "../components/Keyboard.vue"
import { useDSStore } from "../stores/ds";
import { usePTStore } from "../stores/pt";

const dsStore = useDSStore()
const ptStore = usePTStore()

</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 1rem;
}

.name-input {
    width: 150px
}

.el-row {
    margin-bottom: 0.5rem;
}
</style>