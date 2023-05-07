<template>
  <main class="process-page">
    <h1>{{ $t('process.title') }}</h1>
    <el-button text @click="() => {
      err.errorCallback(3);
      process.enable.is_enabled = !process.enable.is_enabled;
      process.moveToNext.is_enabled = !process.moveToNext.is_enabled;
      process.disable.is_enabled = !process.disable.is_enabled;
    }">
      click to open the Dialog
    </el-button>
    <section class="process-box">
      <el-row justify="space-between">
        <el-col :span="6">
          <el-button type="primary" size="large" :disabled="process.enable.is_enabled"
            @click="() => process.enable.enable()">
            {{ $t('process.enable') }}
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button type="warning" size="large" :disabled="process.moveToNext.is_enabled"
            @click="() => process.moveToNext.enable()">
            {{ $t('process.moveToNext') }}
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button type="danger" size="large" :disabled="process.disable.is_enabled"
            @click="() => process.disable.enable()">
            {{ $t('process.disable') }}
          </el-button>
        </el-col>
      </el-row>
      <el-row>
        <el-col v-if="!is_valid">
          <label> {{ $t('process.config_not_valid') }}</label>
        </el-col>
      </el-row>
    </section>
  </main>
</template>
<script setup lang="ts">


import { onMounted } from 'vue';
import { useErrorStore } from '../stores/errors';
import { useProcessStore } from '../stores/process';
import { storeToRefs } from 'pinia';

const err = useErrorStore()
const process = useProcessStore()
const { is_valid } = storeToRefs(process)


onMounted(() => {
  process.reload()
})

</script>
<style lang="scss" scoped>
h1 {
  margin-bottom: 2rem;
}

.el-row {
  margin-bottom: 20px;
}

.el-row:last-child {
  margin-bottom: 0;
}

.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>

