<template>
    <main class="home-page">
        <h1>Home</h1>
        <p>This is the home page</p>
        {{$t('message')}}
        <br>
        <LocaleSelect/>
    </main>
</template>

<script setup>
import LocaleSelect from "@/components/LocaleSelect.vue";
</script>