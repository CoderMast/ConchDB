// 导入 Vue 组件
import { createApp } from 'vue'
import App from './App.vue'

// 导入根样式表
import './style.css';

// 注册路由
import router from './router'

// 注册 NaiveUI 组件库
import NaiveUI from 'naive-ui'


let app = createApp(App)
app.use(NaiveUI)
app.use(router)
app.mount('#app')
