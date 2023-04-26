<template>
    <main class="wifi-page">
        <h1>{{ $t('wifi.title') }}</h1>
        <el-container class="wifi-box" style="height: 80%">
            <el-header>
                <el-switch v-model="wifi.enabled" :active-text="$t('wifi.enable')" size="large" @change="getAP" />
                <el-button v-if="wifi.enabled" @click="getAP(true)" type="primary" round>{{ $t('wifi.reload_ap')
                }}</el-button>
            </el-header>
            <el-main v-if="wifi.enabled">
                <el-scrollbar>
                    <el-table :data="wifi.accessPoints" highlight-current-row style="width: 100%"
                        @current-change="onChange">
                        <el-table-column type="index" width="50" />
                        <el-table-column property="ssid" label="SSID" />
                    </el-table>
                </el-scrollbar>
            </el-main>
            <el-footer>
                <el-progress v-if="wifi.busy" :percentage="100" :format="() => ''" :stroke-width=8 :indeterminate="true" />
            </el-footer>
        </el-container>
        <Keyboard v-bind="wifi.password" :write="(e: string) => wifi.password.write(e)"
            :cancel="() => wifi.password.cancel()" />
    </main>
</template>

<script setup lang="ts">
import { ElContainer, ElHeader, ElTable, ElMain, ElSwitch, ElFooter, ElScrollbar, ElButton, ElTableColumn, ElProgress } from 'element-plus';
import { ref } from 'vue';
import Keyboard from '../components/Keyboard.vue';
import { WifiAPList } from '../../wailsjs/go/backend/Backend';
import { onMounted, onUnmounted } from 'vue';
import Parameter from '../types/Parameter';

interface AP {
    ssid: string
}

type Wifi = {
    enabled: boolean;
    busy: boolean;
    currentSSID: string;
    connected: boolean;
    accessPoints: AP[];
    password: Parameter;
}

const wifi = ref<Wifi>({
    enabled: false,
    busy: false,
    currentSSID: "",
    connected: false,
    accessPoints: [],
    password: new Parameter("", false, connect)
})

onMounted(() => {
    let w = localStorage.getItem('wifi')
    if (w) {
        wifi.value = JSON.parse(w)
        getAP(wifi.value.enabled)
    }

})
onUnmounted(() => {
    localStorage.setItem('wifi', JSON.stringify(wifi.value))
})


const onChange = (row: AP) => {
    wifi.value.password.showKeyboard()
    wifi.value.currentSSID = row.ssid
}
function connect(psk: string) {
    console.log("connecting to " + wifi.value.currentSSID + " with password " + psk)
}

function getAP(enabled: boolean) {
    if (enabled == true) {
        wifi.value.accessPoints = []
        let newAps: AP[] = []
        wifi.value.busy = true
        WifiAPList().then(aps => {
            try {
                aps.forEach(element => {
                    newAps.push({ ssid: element })
                });
                wifi.value.accessPoints = newAps
            } catch (e) {
                console.log("something went wrong " + e)
            } finally {
                wifi.value.busy = false
            }
        })
    }
}

</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 1rem;
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
    align-self: end;
    margin-bottom: 1rem;
}
</style>