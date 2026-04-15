// 导入依赖
import { createRouter, createWebHistory } from 'vue-router'

// createRouter方法,用于创建路由器实例，可以管理多个路由
export default createRouter({
    // 路由的模模式设置
    history: createWebHistory(),
    // 管理路由
    routes: [
        {
            // 重定向的配置
            path: '/',
            redirect: '/home'
        },
        {
            // 主页
            path: '/home',
            component: () => import('@/components/HomeView.vue')
        },
        {
            // 用户信息展示
            path: '/user/info',
            component: () => import('@/components/UserInfoView.vue')
        },
        {
            // 添加用户
            path: "/user/add",
            component: () => import('@/components/UserAddView.vue')
        }
    ],
    // 滚动行为：控制滚动条的位置
    scrollBehavior() {
        return {
            left: 0,
            top: 0,
        }
    }
})