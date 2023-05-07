import { defineConfig } from 'vite'
import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'url'
import vue from '@vitejs/plugin-vue'
import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        VueI18nPlugin({
            compositionOnly: false,
            runtimeOnly: false,
            include: resolve(dirname(fileURLToPath(import.meta.url)), './src/locales/**'),
        }),
        {
            name: "singleHMR",
            handleHotUpdate({ modules }) {
                modules.map((m) => {
                    m.importedModules = new Set();
                    m.importers = new Set();
                });

                return modules;
            },
        },
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
