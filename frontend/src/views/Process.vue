<template>
  <main class="process-page">
    <h1>{{ $t('process.title') }}</h1>
    <el-button text @click="() => { err.errorCallback(3) }">
      click to open the Dialog
    </el-button>

  </main>
</template>

<script setup lang="ts">


import { useErrorStore } from '../stores/errors';
const err = useErrorStore()

</script>
<style lang="scss" scoped></style>

