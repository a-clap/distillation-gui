<template>
  <main class="wifi-page">
    <h1>{{ $t('wifi.title') }}</h1>
    <el-progress :style="{'opacity': wifi.busy ? 1 : 0}"
                 :percentage="100"
                 :format="() => ''" :stroke-width=4
                 :indeterminate="true"/>

    <section class="header">
      <el-switch v-model="wifi.enabled" :active-text="$t('wifi.enable')" size="large" @change="wifi.enable()"/>
      <el-button v-if="wifi.enabled" @click="() => { wifi.getAP() }" type="primary" round>{{
          $t('wifi.reload_ap')
        }}
      </el-button>
    </section>

    <section class="status">
      <section class="wifi-status">
        <div> {{ $t('wifi.status') }}:</div>
        <el-button v-if="connected" size="small" type="success" :icon="Check" circle/>
        <el-button v-else size="small" type="danger" :icon="Close" circle/>
      </section>
      <section class="ap-status">
        <div v-if="connected"> {{ $t('wifi.ap_status') }}:</div>
        <div v-if="connected"> {{ ap }}</div>
      </section>
    </section>

    <section class="status" v-if="connected">

    </section>
    <section class="aplist" v-if="wifi.enabled">
      <el-scrollbar>
        <el-table :data="wifi.apList"
                  max-height=400
                  highlight-current-row style="width: 100%"
                  @current-change="onChange">
          <el-table-column type="index" label="" width="50"/>
          <el-table-column property="ssid" label="SSID" width="500"/>
        </el-table>
      </el-scrollbar>
    </section>
    <Keyboard v-bind="wifi.password" :write="(e: string) => wifi.password.write(e)"
              :cancel="() => wifi.password.cancel()"/>
  </main>
</template>

<script setup lang="ts">

import Keyboard from '../components/Keyboard.vue';
import {AP, useWIFIStore} from '../stores/wifi';
import {onMounted, ref} from "vue";
import {WifiIsConnected} from "../../wailsjs/go/backend/Backend";
import {backend} from "../../wailsjs/go/models";
import {Check, Close} from "@element-plus/icons-vue";

const wifi = useWIFIStore()

const connected = ref(false)
const ap = ref('')

onMounted(() => {
  WifiIsConnected().then((conn: backend.WifiConnected) => {
    connected.value = conn.connected
    ap.value = conn.AP
  })
})

const onChange = (row: AP) => {
  wifi.password.showKeyboard()
  wifi.ssid = row.ssid
}

</script>

<style lang="scss" scoped>
h1 {
  margin-bottom: 0.5rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  * {
    margin-left: 2rem;
    margin-right: 2rem;
    margin-bottom: 1rem;
  }
}

.aplist {
  display: flex;
  justify-content: center;
}

.el-scrollbar {
  display: flex;
}

.status {
  display: flex;
  align-items: center;
  justify-content: center;

  .wifi-status {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .ap-status {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  * {
    margin-left: 2rem;
    margin-bottom: 0.5rem;
  }

}


.el-header {
  display: flex;
  justify-content: start;

  .el-switch {
    flex: 7;
  }

  .el-button {
    flex: 1;
    margin-right: 5rem;
  }
}

.el-progress {
  margin-left: 2rem;
  margin-bottom: 0.5rem;
}
</style>