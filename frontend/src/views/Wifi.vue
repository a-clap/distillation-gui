<template>
    <main class="wifi-page">
        <h1>{{ $t('wifi.title') }}</h1>
        <el-container class="wifi-box" style="height: 80%">
            <el-header>
                <el-switch v-model="wifi.enabled" :active-text="$t('wifi.enable')" size="large" @change="wifi.enable()" />
                <el-button v-if="wifi.enabled" @click="() => { wifi.getAP() }" type="primary" round>{{ $t('wifi.reload_ap')
                }}</el-button>
            </el-header>
            <el-main v-if="wifi.enabled">
                <el-scrollbar>
                    <el-table :data="wifi.apList" highlight-current-row style="width: 100%" @current-change="onChange">
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

import Keyboard from '../components/Keyboard.vue';
import { useWIFIStore, AP } from '../stores/wifi';


const wifi = useWIFIStore()

const onChange = (row: AP) => {
    wifi.password.showKeyboard()
    wifi.ssid = row.ssid
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