import { createRouter, createWebHistory } from 'vue-router';

import ConnectionVue from '../components/ConnectionVue.vue';
import Execute from '../components/Execute.vue';

const routes = [
    {
        path: '/',
        name: 'Connection',
        component: ConnectionVue
    },
    {
        path: '/execute',
        name: 'Execute',
        component: Execute
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;