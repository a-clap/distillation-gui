import { createApp } from 'vue'
import { i18n } from "./i18n.js";

import ElementPlus from 'element-plus'
// import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/index.css'

import App from './App.vue'
import router from "./router/index.js"



const app = createApp(App)

app.use(router)
app.use(i18n)
app.use(ElementPlus)
app.mount('#app')
