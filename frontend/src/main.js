import {createApp} from 'vue'
import { i18n } from "./i18n.js";
import App from './App.vue'
import router from "./router/index.js"

const app = createApp(App)

app.use(router)
app.use(i18n)
app.mount('#app')
