// vue3框架提供的方法createApp方法,可以用来创建应用实例方法
import { createApp } from 'vue'
// 引入根组件App
import App from '@/App.vue'
// 引入全局组件--顶部、底部都是全局组件
// import MyHome from '@/pages/MyHome.vue'
// import MyAbout from '@/pages/MyAbout.vue'
// 引入 vue-router 核心插件并安装(全路径：@/router/index.ts，可省略)
import router from '@/router'
// 引入 ElementPlus 组件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 利用createApp方法创建应用实例,且将应用实例挂载到挂载点上
const app = createApp(App)
// 注册全局组件
// app.component('MyHome', MyHome)
// app.component('MyAbout', MyAbout)
// 使用我们配置的 vue-router 路由插件
app.use(router)
app.use(ElementPlus)
// 挂载
app.mount('#app')