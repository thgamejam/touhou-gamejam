import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  // blog-venja-cc
  {
    path: '/',
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
