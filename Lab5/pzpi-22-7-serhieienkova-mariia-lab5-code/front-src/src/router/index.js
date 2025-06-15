import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../pages/LoginPage.vue'
import RegisterPage from '../pages/RegisterPage.vue'
import DashboardPage from '../pages/DashboardPage.vue'
import PatientsPage from '../pages/PatientsPage.vue'
import MedicinesPage from '../pages/MedicinesPage.vue'
import DiagnosesPage from "../pages/DiagnosesPage.vue";
import PatientsFullInfoPage from "../pages/PatientsFullInfoPage.vue";
import VisitsPage from "../pages/VisitsPage.vue";
import NotificationsPage from "../pages/NotificationsPage.vue";

const routes = [
    { path: '/login', component: LoginPage },
    { path: '/register', component: RegisterPage },
    { path: '/dashboard', component: DashboardPage },
    { path: '/patients', component: PatientsPage },
    { path: '/medicines', component: MedicinesPage },
    { path: '/diagnoses', component: DiagnosesPage },
    { path: '/', redirect: '/login' },
    { path: '/patients/:id/full-info', component: PatientsFullInfoPage },
    {path: '/visits', component: VisitsPage},
    {path: '/notifications', component: NotificationsPage}
]


const router = createRouter({
    history: createWebHistory(),
    routes,
})

// Захист роутів (Navigation Guard)
router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/register']
    const authRequired = !publicPages.includes(to.path)
    const loggedIn = localStorage.getItem('accessToken')

    if (authRequired && !loggedIn) {
        return next('/login')
    }

    next()
})

export default router
