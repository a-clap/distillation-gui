import { defineConfig } from 'vite'
import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'url'
import vue from '@vitejs/plugin-vue'
import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";
import ElementPlus from 'unplugin-element-plus/vite';

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        ElementPlus(),
        vue(),
        VueI18nPlugin({
            compositionOnly: false,
            runtimeOnly: false,
            include: resolve(dirname(fileURLToPath(import.meta.url)), './src/locales/**'),
        })
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
