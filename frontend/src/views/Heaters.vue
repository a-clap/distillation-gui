<template>
    <main>
        <h1>{{ $t('heaters.title') }}</h1>
        <div v-for="(heater, index) in getHeaters" :key="index">
            <section class="heater-box">
                <el-row :gutter="20" align="middle">
                    <el-col :span="3">
                        <el-switch v-model="heater.enable" :active-text="heater.name" size="large" />
                    </el-col>
                </el-row>
            </section>
        </div>
    </main>
</template>

<script setup lang="ts">

import { reactive, onMounted, computed } from "vue"
import { Heater } from '../types/Heater';
import { GetHeaters } from '../../wailsjs/go/main/App'
import { parameters } from "../../wailsjs/go/models";

const heaters: Heater[] = reactive([])

onMounted(() => {
    GetHeaters().then((value: parameters.Heater[]) => {
        value.forEach((heater: parameters.Heater) => {
            console.log(heater.enabled)
            heaters.push(new Heater(heater.ID, heater.enabled))
        })
    })
})

const getHeaters = computed(() => {
    return heaters.sort((n1, n2) => {
        if (n1.name > n2.name) {
            return 1;
        }

        if (n1.name < n2.name) {
            return -1;
        }

        return 0;
    });
})

</script>

<style lang="scss" scoped>
h1 {
    margin-bottom: 2rem;
}

.heater-box {
    font-size: 1.2rem;
    margin-bottom: 1.5rem;

}
</style>

