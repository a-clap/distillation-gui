<template>
  <main class="phases-page">
    <h1>{{ $t('phases.title') }}</h1>
    <el-tabs v-model="activated" type="card" tabPosition="left" class="demo-tabs">
      <el-tab-pane label="Main" name="main">
        <input v-model="phaseConfig.phaseCount.value" @click="() => phaseConfig.phaseCount.showKeyboard()">
        <Keyboard v-bind="phaseConfig.phaseCount" :write="(e: number) => phaseConfig.phaseCount.write(e)"
          :cancel="() => phaseConfig.phaseCount.cancel()" />
      </el-tab-pane>
      <el-tab-pane label="Config" name="second">Config</el-tab-pane>
      <el-tab-pane label="Role" name="third">Role</el-tab-pane>
      <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
    </el-tabs>
  </main>
</template>

<script setup lang="ts">

import Keyboard from "../components/Keyboard.vue"
import { Process } from '../types/Process';
import { onMounted, onUnmounted, ref } from 'vue'
import { ProcessListener } from "../types/ProcessListener";
import { distillation } from "../../wailsjs/go/models";
const activated = ref('main')
const phaseConfig = ref<Process>(new Process());

onMounted(() => {
  ProcessListener.subscribePhaseCount(phaseCountUpdate)
})

onUnmounted(() => {
  ProcessListener.unsubscribePhaseCount(phaseCountUpdate)
})

function phaseCountUpdate(v: distillation.ProcessPhaseCount) {
  console.log(v)
}

</script>

<style lang="scss">
.demo-tabs>.el-tabs__content {
  padding: 32px;
  font-weight: 600;
}
</style>