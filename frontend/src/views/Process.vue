<template>
  <main class="process-page">
    <h1>{{ $t('process.title') }}</h1>
    <section class="process-box">
      <el-row justify="space-between">
        <el-col :span="6">
          <el-button type="success" size="large" :disabled="!process.enable.is_enabled"
            @click="() => process.enable.enable()">
            {{ $t('process.enable') }}
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button type="warning" size="large" :disabled="!process.moveToNext.is_enabled"
            @click="() => process.moveToNext.enable()">
            {{ $t('process.moveToNext') }}
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button type="danger" size="large" :disabled="!process.disable.is_enabled"
            @click="() => process.disable.enable()">
            {{ $t('process.disable') }}
          </el-button>
        </el-col>
      </el-row>
      <el-row justify="center">
        <el-col :span="25" class="bold-text">
          <label v-if="!is_valid"> {{ $t('process.config_not_valid') }}</label>
          <label v-if="process.show_status && process.running"> {{ $t('process.running') }}</label>
          <label v-if="process.show_status && !process.running"> {{ $t('process.done') }}</label>
        </el-col>
      </el-row>
      <template v-if="process.show_status">
        <el-row>
          <el-col :span="5">
            {{ $t('process.start_time') }}
          </el-col>
          <el-col :span="4">
            {{ process.start_time }}
          </el-col>
          <el-col :span="5" :offset="6">
            {{ $t('process.end_time') }}
          </el-col>
          <el-col :span="4">
            {{ process.end_time }}
          </el-col>
        </el-row>
        <el-row justify="center">
          <el-col :span="5">
            {{ $t('process.current_phase') }}
          </el-col>
          <el-col :span="1" :offset="1">
            <input class="small-input" v-model="process.current_phase" />
          </el-col>
        </el-row>
        <el-row justify="center" align="middle">
          <el-col :span="4">
            {{ $t('process.move_to_next_type') }}
          </el-col>
          <el-col :span="3" class="bold-text">
            <label v-if="process.current_type_time">{{ $t('process.move_to_next_type_time') }}</label>
            <label v-else>{{ $t('process.move_to_next_type_temperature') }}</label>
          </el-col>
          <el-col :span="5" :offset="3">
            {{ $t('process.timeleft') }}
          </el-col>
          <el-col :span="3">
            <input class="small-input" v-model="process.phase_timeleft" />
          </el-col>
        </el-row>
        <el-row v-if="!process.current_type_time" justify="center">
          <el-col :span="2">
            <label>{{ $t('process.sensor') }}</label>
          </el-col>
          <el-col :span="3" :offset="1">
            <label> {{ process.phase_sensor }}</label>
          </el-col>
          <el-col :span="8">
            <label>{{ $t('process.sensor_threshold') }}</label>
          </el-col>
          <el-col :span="3" :offset="1">
            <input class="small-input" v-model="process.phase_sensor_threshold" />
          </el-col>
        </el-row>
        <el-row justify="center" align="middle">
          <el-col class="bold-text" span="25">
            {{ $t('process.heaters') }}
          </el-col>
        </el-row>
        <el-row align="middle">
          <template v-for="heater in process.heaters">
            <el-col :span="3" :offset=1>
              {{ heater.id }}
            </el-col>
            <el-col :span="3">
              <input class="small-input" v-model="heater.pwr" />
              {{ $t('process.power_sign') }}
            </el-col>
          </template>
        </el-row>
        <el-row justify="center" align="middle">
          <el-col class="bold-text" span="25">
            {{ $t('process.sensors') }}
          </el-col>
        </el-row>
        <el-row align="middle">
          <template v-for="sensor in process.sensors">
            <el-col :span="3" :offset=1>
              {{ sensor.id }}
            </el-col>
            <el-col :span="3">
              <input class="small-input" v-model="sensor.temperature" />
              {{ $t('process.temperature_sign') }}
            </el-col>
          </template>
        </el-row>
        <el-row justify="center" align="middle">
          <el-col class="bold-text" span="25">
            {{ $t('process.outputs') }}
          </el-col>
        </el-row>
        <el-row align="middle">
          <template v-for="output in process.outputs">
            <el-col :span="3" :offset=1>
              {{ output.id }}
            </el-col>
            <el-col :span="3">
              <el-checkbox v-model="output.state" :label="$t('outputs.manual_control_force')" size="large" border />
            </el-col>
          </template>
        </el-row>
      </template>
    </section>
  </main>
</template>
<script setup lang="ts">


import { onMounted } from 'vue';
import { useProcessStore } from '../stores/process';
import { storeToRefs } from 'pinia';

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

.bold-text {
  font-weight: 800;
}

.small-input {
  width: 75px
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

.el-checkbox {
  --el-color-primary: var(--el-color-danger);
  --el-checkbox-text-color: var(--el-color-danger);
}
</style>

