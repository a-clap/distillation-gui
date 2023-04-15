<template>
    <main class="pt100-page">
        <h1>{{ $t('pt100.title') }}</h1>

        <el-row class="mb-4">
            <el-button>Default</el-button>
            <el-button type="primary">Primary</el-button>
            <el-button type="success">Success</el-button>
            <el-button type="info">Info</el-button>
            <el-button type="warning">Warning</el-button>
            <el-button type="danger">Danger</el-button>
        </el-row>
    </main>
</template>

<script setup>

import { PT100 } from '../types/PT100.js';



</script>

