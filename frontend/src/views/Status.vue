<template>
  <main class="status-page">
    <div class="header">
      <h1>{{ $t('status.title') }}</h1>
      <router-link to="/">
        <el-button size="large" type="primary" :icon="ArrowLeftBold"/>
      </router-link>
    </div>
    <section class="status-box">
      <template v-if="process.show_status">
        <div class="container">
          <h2>{{ $t('status.current_phase') }} </h2>
          <input v-model="process.current_phase"/>
        </div>
        <template v-for="sensor in process.sensors">
          <section class="container">
            <h2>{{ sensor.id }}</h2>
            <input v-model="sensor.temperature"/>
            <h2>{{ $t('status.temperature_sign') }} </h2>
          </section>
        </template>
        <template v-for="output in process.outputs">
          <section class="container">
            <h2>{{ output.id }}</h2>
            <el-button v-if="output.state" size="large" type="success" :icon="Check" circle/>
            <el-button v-else size="large" type="danger" :icon="Close" circle/>
          </section>
        </template>
      </template>
    </section>
  </main>
</template>

<script setup lang="ts">

import {ArrowLeftBold, Check, Close} from "@element-plus/icons-vue";
import {useProcessStore} from "../stores/process";
import {onMounted} from "vue";

const process = useProcessStore()

onMounted(() => {
  process.reload()
})

</script>

<style lang="scss" scoped>

$text-font-size: 4rem;
$input-font-size: 4rem;
$bottom-margin: 2rem;

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;

  .el-button {
    width: 100px;
    font-size: 20px;
    margin-right: 10px;
  }

  margin-bottom: 1rem;
}

.status-box {
  font-size: $text-font-size;
}

.container {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin-bottom: $bottom-margin;

  .el-button {
    height: 100px;
    width: 100px;
    font-size: $text-font-size;
    margin-right: 10px;
    margin-left: 10px;

  }
}


input {
  margin-left: 2rem;
  font-size: $input-font-size;
  width: $input-font-size + 10rem;
  height: $input-font-size + 2rem;
}

</style>
  
  