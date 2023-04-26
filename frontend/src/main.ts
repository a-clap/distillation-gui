import { createApp } from 'vue'
import { i18n } from "./i18n"

import 'element-plus/theme-chalk/dark/css-vars.css'

import App from './App.vue'
import router from "./router/index.js"


const app = createApp(App)

app.use(router)
app.use(i18n)
app.mount('#app')
