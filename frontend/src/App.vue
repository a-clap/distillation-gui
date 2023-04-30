<template>
  <html class="dark">
  <div class="app">
    <Sidebar />
    <router-view />
  </div>

  </html>
</template>
<script setup lang="ts">

import Sidebar from "./components/Sidebar.vue";
import { markRaw } from "vue";
import { ErrorListener } from "./types/ErrorListener";
import { ElMessageBox } from 'element-plus'
import { CloseBold } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n';
import { useGpioStore } from "./stores/gpios";
import { useDSStore } from "./stores/ds";
import { useHeatersStore } from "./stores/heaters";
import { usePTStore } from "./stores/pt";
import { useWIFIStore } from "./stores/wifi";
import { usePhasesStore } from "./stores/phases";
const { t, te } = useI18n();

ErrorListener.subscribe(errorCallback)

function errorCallback(id: number) {
  let err = `errors.${id}`;
  if (te(err)) {
    err = t(err)
  } else {
    err = t('errors.unknown')
    err += id.toString()
  }

  open(t('errors.title'), err)
}

const open = (title: string, msg: string) => {
  ElMessageBox.alert(
    msg,
    title,
    {
      customClass: "message-box",
      type: "error",
      icon: markRaw(CloseBold),
      autofocus: false,
      roundButton: true,
      center: true,
      showClose: false,
      confirmButtonText: 'OK'
    })
}

let initFuncs: any[] = [
  useGpioStore(),
  useDSStore(),
  useHeatersStore(),
  usePTStore(),
  useWIFIStore(),
  usePhasesStore()
]

initFuncs.forEach((store) => {
  setTimeout(() => { store.init() }, 10);
})


</script>
<style lang="scss">
:root {
  --sidebar-dark: #1e293b;
  --sidebar-dark-alt: #334155;
  --sidebar-width: 200px;
  --window-width: 1024px;
  --window-height: 768px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Fira sans', sans-serif;
  transition: none !important;
  transform: none !important;
}

.app {
  display: flex;
  background-color: rgb(51, 65, 85);

  main {
    flex: 1 1 0;
    padding-left: 0.5rem;
    padding-top: 0.5rem;
  }

  input {
    width: 100px;
    height: 34px;
    padding: 6px 12px;
    line-height: 1rem;
    text-align: center;
    color: #555;
    cursor: default;
    caret-color: transparent;
    background-color: #fff;
    background-image: none;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
  }

  input:focus {
    outline: none !important;
    border: 1px solid var(--el-color-primary);
  }
}
</style>