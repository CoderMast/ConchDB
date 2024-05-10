import {createApp} from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style.css';
import router from './router'

var app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.mount('#app')
