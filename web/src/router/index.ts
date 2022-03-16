import { createRouter, createWebHistory } from 'vue-router'
import Main from '../views/Main.vue'

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  },
  {
    path: '/login',
    name: 'login',
    component: ()=>import('../views/LoginPage.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },

  // blog-venja-cc
  {
    path: '/blog',
    name: 'Blog',
    component: () => import('../blog-venja-cc/Index.vue')
  },
  {
    path: '/signin',
    name: 'SignIn',
    component: () => import('../blog-venja-cc/LoginRegister.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
