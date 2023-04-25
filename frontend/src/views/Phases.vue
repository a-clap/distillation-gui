<template>
  <main class="phases-page">
    <h1>{{ $t('phases.title') }}</h1>
    <el-tabs v-model="activated" type="card" tabPosition="left" class="demo-tabs">
      <el-tab-pane :label="$t('phases.main')" name="main" class="main-tab">
        <label>{{ $t('phases.count') }}</label>
        <input v-model="phases.phaseCount.value" @click="() => phases.phaseCount.showKeyboard()">
        <Keyboard v-bind="phases.phaseCount" :write="(e: number) => phases.phaseCount.write(e)"
          :cancel="() => phases.phaseCount.cancel()" />
      </el-tab-pane>
      <el-tab-pane v-for="(phase, index) in  phases.phases " :key="index" :label="$t('phases.phase') + ' ' + (index + 1)">
        <el-row :gutter="20" align="middle">
          <el-col :span="5">
            {{ $t('phases.next_type') }}
          </el-col>
          <el-col :span="7">
            <el-switch v-model="phase.next_type" :inactive-text="$t('phases.next_type_time')"
              :active-text="$t('phases.next_type_temp')" size="large" />
          </el-col>
          <el-col :span="7" :offset="1">
            <label v-if="!phase.next_type">{{ $t('phases.timeleft') }}</label>
            <label v-else>{{ $t('phases.timeleft_temp') }}</label>
          </el-col>
          <el-col :span="3">
            <input v-model="phase.next_timeleft.value" @click="() => phase.next_timeleft.showKeyboard()">
            <Keyboard v-bind="phase.next_timeleft" :write="(e: number) => phase.next_timeleft.write(e)"
              :cancel="() => phase.next_timeleft.cancel()" />
          </el-col>
        </el-row>
        <el-row :gutter="20" align="middle" v-if="phase.next_type">
          <el-col :span="4">
            <label>{{ $t('phases.next_sensor') }}</label>
          </el-col>
          <el-col :span="5">
            <el-select v-model="phase.next_sensor" size="large" class="m-2">
              <el-option v-for="sensor in phase.next_avail_sensors" :label="sensor" :value="sensor" />
            </el-select>
          </el-col>
          <el-col :span="6" :offset="4">
            <label v-if="phase.next_type">{{ $t('phases.sensor_threshold') }}</label>
          </el-col>
          <el-col :span="3" :offset="1">
            <input v-model="phase.next_sensor_threshold.value" @click="() => phase.next_sensor_threshold.showKeyboard()">
            <Keyboard v-bind="phase.next_sensor_threshold" :write="(e: number) => phase.next_sensor_threshold.write(e)"
              :cancel="() => phase.next_sensor_threshold.cancel()" />
          </el-col>
        </el-row>
        <el-row :gutter="20" align="middle">
          <el-col :span="15" :offset="8" :style="'font-weight: 800'">
            <label v-if="phase.heaters.length == 0">{{ $t('phases.no_heaters') }}</label>
            <label v-else>{{ $t('phases.heaters') }}</label>
          </el-col>
        </el-row>
        <el-row :gutter="20" v-for="(heater) in  phase.heaters ">
          <el-col :span="5" :style="'font-weight: 800'">
            <label>{{ heater.id }}</label>
          </el-col>
          <el-col :span="4">
            <input v-model="heater.power.value" @click="() => heater.power.showKeyboard()">
            <Keyboard v-bind="heater.power" :write="(e: number) => heater.power.write(e)"
              :cancel="() => heater.power.cancel()" />
          </el-col>
          <el-col :span="5">
            <label>{{ $t('phases.heater_power') }}</label>
          </el-col>
        </el-row>
        <el-row :gutter="20" align="middle">
          <el-col :span="15" :offset="8" :style="'font-weight: 800'">
            <label v-if="phase.gpios.length == 0">{{ $t('phases.no_gpio') }}</label>
            <label v-else>{{ $t('phases.gpio') }}</label>
          </el-col>
        </el-row>
        <template v-for="( gpio ) in  phase.gpios ">
          <el-row :gutter="20">
            <el-col :span="5">
              <label :style="'font-weight: 800'">{{ gpio.id }}</label>
            </el-col>
            <el-col :span="4">
              <el-checkbox v-model="gpio.inverted" :label="$t('phases.gpio_inverted')" size="large" border />
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="4">
              <label>{{ $t('phases.next_sensor') }}</label>
            </el-col>
            <el-col :span="6">
              <el-select v-model="gpio.sensor_id" size="large" class="m-2">
                <el-option v-for="sensor in phase.next_avail_sensors" :label="sensor" :value="sensor" />
              </el-select>
            </el-col>
            <el-col :span="5" :offset="2">
              <label>{{ $t('phases.gpio_hysteresis') }}</label>
            </el-col>
            <el-col :span="3" :offset="1">
              <input v-model="gpio.hysteresis.value" @click="() => gpio.hysteresis.showKeyboard()">
              <Keyboard v-bind="gpio.hysteresis" :write="(e: number) => gpio.hysteresis.write(e)"
                :cancel="() => gpio.hysteresis.cancel()" />
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="7">
              <label>{{ $t('phases.gpio_temp_min') }}</label>
            </el-col>
            <el-col :span="3">
              <input v-model="gpio.t_low.value" @click="() => gpio.t_low.showKeyboard()">
              <Keyboard v-bind="gpio.t_low" :write="(e: number) => gpio.t_low.write(e)"
                :cancel="() => gpio.t_low.cancel()" />
            </el-col>
            <el-col :span="7" :offset=1>
              <label>{{ $t('phases.gpio_temp_max') }}</label>
            </el-col>
            <el-col :span="3">
              <input v-model="gpio.t_high.value" @click="() => gpio.t_high.showKeyboard()">
              <Keyboard v-bind="gpio.t_high" :write="(e: number) => gpio.t_high.write(e)"
                :cancel="() => gpio.t_high.cancel()" />
            </el-col>
          </el-row>
        </template>
      </el-tab-pane>
    </el-tabs>
  </main>
</template>

<script setup lang="ts">

import Keyboard from "../components/Keyboard.vue"
import { Phases } from '../types/Phases';
import { onMounted, onUnmounted, ref } from 'vue'
import { ProcessListener } from "../types/ProcessListener";
import { ProcessPhaseConfig } from "../types/Phases";
import { distillation, process } from "../../wailsjs/go/models";
import { PhasesGetPhaseConfigs } from "../../wailsjs/go/backend/Backend"
const activated = ref('main')
const phases = ref<Phases>(new Phases());

onMounted(() => {
  reload()
  ProcessListener.subscribePhaseCount(phaseCountUpdate)
  ProcessListener.subscribePhaseConfig(phaseConifgUpdate)
})

onUnmounted(() => {
  ProcessListener.unsubscribePhaseCount(phaseCountUpdate)
  ProcessListener.unsubscribePhaseConfig(phaseConifgUpdate)
})

function reload() {
  PhasesGetPhaseConfigs().then((value: distillation.ProcessPhaseConfig[]) => {
    let size = value.length
    let configs: ProcessPhaseConfig[] = []

    value.forEach((v: distillation.ProcessPhaseConfig, i: number) => {
      let next = new process.MoveToNextConfig()

      next.seconds_to_move = 1
      next.sensor_id = "sensor_1"
      next.sensor_threshold = 13
      next.temperature_hold_seconds = 10
      next.type = 0

      configs.push(new ProcessPhaseConfig(i, next, v.heaters, v.gpio))
    })

    phases.value = new Phases(configs, size)
  })
}


function phaseCountUpdate(v: distillation.ProcessPhaseCount) {
  reload()
}

function phaseConifgUpdate(n: number, v: distillation.ProcessPhaseConfig) {
  phases.value.phases[n] = new ProcessPhaseConfig(n, v.next, v.heaters, v.gpio)
}


</script>

<style lang="scss">
.demo-tabs>.el-tabs__content {
  padding: 32px;
}

.main-tab {
  display: flex;
  align-items: center;
  justify-content: space-evenly;
}

.el-row {
  margin-bottom: 1rem;
  align-items: center;
}
</style>