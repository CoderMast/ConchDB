import { createRouter, createWebHistory } from 'vue-router';

import Connection from '../components/Connection.vue';
import Execute from '../components/Execute.vue';

const routes = [
    {
        path: '/',
        name: 'Connection',
        component: Connection
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