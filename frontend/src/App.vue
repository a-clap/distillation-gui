<script setup>
import Sidebar from "./components/Sidebar.vue";
</script>

<template>
  <html class="dark">
  <div class="app">
    <Sidebar />
    <router-view />
  </div>

  </html>
</template>

<style lang="scss">
:root {
  --sidebar-dark: #1e293b;
  --sidebar-dark-alt: #334155;
  --sidebar-width: 200px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Fira sans', sans-serif;
}


.app {
  display: flex;
  background-color: rgb(51, 65, 85);
  ;

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