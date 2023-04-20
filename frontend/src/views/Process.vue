<template>
  <main class="process-page">
    <h1>{{ $t('process.title') }}</h1>
    <el-button text @click="open">Click to open Message Box</el-button>
  </main>
</template>

<script setup lang="ts">

import { ElMessage, ElMessageBox } from 'element-plus'

const open = () => {
  ElMessageBox.confirm(
    'proxy will permanently delete the file. Continue?',
    'Warning',
    {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning',
      draggable: true,
    }
  )
    .then(() => {
      ElMessage({
        type: 'success',
        message: 'Delete completed',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Delete canceled',
      })
    })
}





</script>
<style lang="scss" scoped>
</style>

