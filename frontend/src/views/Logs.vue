<template>
  <main class="logs-page">
    <h1>{{ $t('logs.title') }}</h1>
    <el-container class="log-box" style="height: 80%">
      <el-header>
        <el-button type="warning" @click="errStore.reset_skipped">Warning</el-button>
      </el-header>
      <el-main>
        <el-table-v2 :columns="columns" :data="data" :width="750" :height="500" fixed />
      </el-main>
    </el-container>
  </main>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import { useErrorStore } from '../stores/errors';
import { useLogStore, Column } from '../stores/log';

const errStore = useErrorStore()
const logStore = useLogStore()

let tableData = ref<string[][]>([]);
tableData.value = logStore.logData

let tableColumns = ref<Column[]>([])
tableColumns.value = logStore.columns;


const generateColumns = (length = tableColumns.value.length, prefix = 'column-', props?: any) =>
  Array.from({ length }).map((_, columnIndex) => ({
    ...props,
    key: `${prefix}${columnIndex}`,
    dataKey: `${prefix}${columnIndex}`,
    title: `${tableColumns.value[columnIndex].title}`,
    width: `${tableColumns.value[columnIndex].width}`,
  }))

const generateData = (
  columns: ReturnType<typeof generateColumns>,
  length = tableData.value.length,
  prefix = 'row-'
) =>
  Array.from({ length }).map((_, rowIndex) => {
    return columns.reduce(
      (rowData, column, columnIndex) => {
        rowData[column.dataKey] = tableData.value[rowIndex][columnIndex]
        return rowData
      },
      {
        id: `${prefix}${rowIndex}`,
        parentId: null,
      }
    )
  })

const columns = generateColumns()
const data = generateData(columns)

</script>
<style lang="scss" scoped>
h1 {
  margin-bottom: 2rem;
}

.el-table-v2 {
  margin: auto;
}
</style>
  
  