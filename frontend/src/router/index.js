import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: () => import('../views/Process.vue')
        },
        {
            path: '/phases',
            component: () => import('../views/Phases.vue')
        },
        {
            path: '/wifi',
            component: () => import('../views/Wifi.vue')
        },
        {
            path: '/pt100',
            component: () => import('../views/PT100.vue')
        },
        {
            path: '/ds',
            component: () => import('../views/DS.vue')
        },
        {
            path: '/heaters',
            component: () => import('../views/Heaters.vue')
        },
        {
            path: '/outputs',
            component: () => import('../views/Outputs.vue')
        },
        {
            path: '/names',
            component: () => import('../views/Names.vue')
        },
        {
            path: '/logs',
            component: () => import('../views/Logs.vue')
        },
        {
            path: '/status',
            component: () => import('../views/Status.vue')
        }
    ]
})

export default router