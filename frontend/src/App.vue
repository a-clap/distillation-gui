<script setup>
import Sidebar from "./components/Sidebar.vue";
</script>

<template>
  <div class="app">
    <Sidebar />
    <router-view />
  </div>
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

  main {
    flex: 1 1 0;
    padding-left: 0.5rem;
    padding-top: 0.5rem;
  }
}
</style>