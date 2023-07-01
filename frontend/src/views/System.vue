<template>
  <main>
    <h1>{{ $t('system.title') }}</h1>
    <el-tabs v-model="activated" type="card" tabPosition="left" class="demo-tabs">
      <el-tab-pane :label="$t('system.names')" name="names">
        <div class="name-container">{{ $t('system.ds') }}</div>
        <template v-for="(ds, index) in dsStore.ds" :key="index">
          <section class="name-container">
            <div class="name-elem">{{ ds.id }}</div>
            <input class="name-input" v-model="ds.name.view" @click="() => ds.name.showKeyboard()">
            <Keyboard v-bind="ds.name" :write="(e: string) => ds.writeName(e)" :cancel="() => ds.name.cancel()"/>
          </section>
        </template>
        <div class="name-container">{{ $t('system.pt') }}</div>
        <template v-for="(pt, index) in ptStore.pt" :key="index">
          <section class="name-container">
            <div class="name-elem">{{ pt.id }}</div>
            <input class="name-input" v-model="pt.name.view" @click="() => pt.name.showKeyboard()">
            <Keyboard v-bind="pt.name" :write="(e: string) => pt.writeName(e)" :cancel="() => pt.name.cancel()"/>
          </section>
        </template>
      </el-tab-pane>
      <el-tab-pane :label="$t('system.time')" name="time">
        <div class="ntp-container">
          <div>{{ $t('system.time_now') }}: {{ timeNow }}</div>
          <el-switch v-model="ntp" :active-text="$t('system.ntp_time')" size="large"/>
        </div>

        <div class="time-picker" v-if=!ntpEnabled>
          <el-time-picker
              v-model="currentTime"
              :placeholder="$t('system.pick_time')"
              :clearable=false
              size="large"
          />
          <el-date-picker
              v-model="currentDate"
              type="date"
              :placeholder="$t('system.pick_date')"
              :clearable=false
              size="large"
          />
          <el-button type="primary" size="large" @click="setTime">
            {{ $t('system.set_time') }}
          </el-button>
        </div>
      </el-tab-pane>
      <el-tab-pane :label="$t('system.net')" name="net">
        <section class="net-interface">
          <h3>{{ $t(('system.net_interface')) }}</h3>
          <h3>{{ $t(('system.net_ip')) }}</h3>
        </section>
        <section v-for="(netInterface, index) in netInterfaces" :key="index">
          <section class="net-interface">
            <div> {{ netInterface.name }}</div>
            <div> {{ netInterface.ip_addr }}</div>
          </section>
        </section>
      </el-tab-pane>
    </el-tabs>
  </main>
</template>

<script setup lang="ts">

import {useDSStore} from "../stores/ds";
import {usePTStore} from "../stores/pt";
import {computed, onMounted, ref} from "vue";
import Keyboard from "../components/Keyboard.vue";
import dayjs from 'dayjs'
import {ListInterfaces, NTPGet, NTPSet, TimeSet} from "../../wailsjs/go/backend/Backend";
import {backend} from "../../wailsjs/go/models";
import {Loader} from "../types/Loader";
import {AppErrorCodes} from "../stores/error_codes";
import {i18n} from "../i18n";
import NetInterface = backend.NetInterface;

const currentDate = ref(new Date())
const currentTime = ref(new Date())
const activated = ref('names')
const dsStore = useDSStore()
const ptStore = usePTStore()

const ntpEnabled = ref(false)
const ntp = computed({
  get: () => ntpEnabled.value,
  set: (v: boolean) => {

    let msg = i18n.global.t('system.ntp_loading')
    Loader.show(AppErrorCodes.NTPFailed, 5000, msg)
    if (!v) {
      currentTime.value = new Date()
    }

    NTPSet(v).then((err: any) => {
      if (!err) {
        ntpEnabled.value = v
      }
      Loader.close()
    })
  }
})

const netInterfaces = ref<NetInterface[]>([])
const timeNow = ref('')


onMounted(() => {

  NTPGet().then((value: boolean) => {
    ntp.value = value
  })

  timeNow.value = dayjs().format('HH:mm:ss DD/MM/YYYY')
  setInterval(() => {
    timeNow.value = dayjs().format('HH:mm:ss DD/MM/YYYY')
  }, 1000)

  ListInterfaces().then((interfaces: NetInterface[]) => {
    netInterfaces.value = interfaces
    console.log(interfaces)
  })

})

function setTime() {
  let fullDate = currentTime.value
  fullDate.setFullYear(currentDate.value.getFullYear())
  fullDate.setMonth(currentDate.value.getMonth())
  fullDate.setDate(currentDate.value.getDate())

  TimeSet(fullDate.getTime()).then(() => {

  })
}

</script>

<style lang="scss" scoped>
h1 {
  margin-bottom: 1rem;
}

.name-input {
  width: 150px
}

.name-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 2rem;
  font-size: 2rem;

  .name-input {
    margin-left: 3rem;
  }

  .keyboard-window {
    font-size: initial;
  }
}

.ntp-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.ntp-container > * {
  margin-bottom: 2rem;
}

.net-interface {
  display: flex;
  justify-content: space-around;
  align-items: center;

  * {
    margin-bottom: 1rem;
  }
}

.time-picker {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;

}


</style>