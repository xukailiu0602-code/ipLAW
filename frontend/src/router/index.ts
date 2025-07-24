import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import Upload from '../pages/Upload.vue'
import QA from '../pages/QA.vue'
import Report from '../pages/Report.vue'
import Admin from '../pages/Admin.vue'

const routes = [
  { path: '/login', component: Login },
  { path: '/upload', component: Upload },
  { path: '/qa', component: QA },
  { path: '/report/:id', component: Report },
  { path: '/admin', component: Admin },
  { path: '/', redirect: '/qa' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router