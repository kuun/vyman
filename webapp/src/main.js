import './assets/main.css'

import "v3-easyui/dist/themes/default/easyui.css";
import "v3-easyui/dist/themes/icon.css";
import "v3-easyui/dist/themes/vue.css";
import EasyUI from "v3-easyui";

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(EasyUI)
app.use(createPinia())
app.use(router)

app.mount('#app')
