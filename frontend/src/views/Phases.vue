<template>
    <main class="phases-page">
        <h1>{{ $t('phases.title') }}</h1>
        <el-tabs v-model="activeName" type="card" tabPosition="left" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="User" name="first">User</el-tab-pane>
            <el-tab-pane label="Config" name="second">Config</el-tab-pane>
            <el-tab-pane label="Role" name="third">Role</el-tab-pane>
            <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
        </el-tabs>
    </main>
</template>

<script setup lang="ts">

import { ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'

const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}

</script>

<style lang="scss">

.demo-tabs > .el-tabs__content {
  padding: 32px;
  font-weight: 600;
}

</style>