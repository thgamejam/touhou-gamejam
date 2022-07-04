import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  // pages
  {
    path: '/',
    name: 'Index',
    component: () => import('../pages/Index.vue')
  },
  {
    path: '/signin',
    name: 'SignIn',
    component: () => import('../pages/LoginRegister.vue')
  },
  {
    path: '/blog/:blogId',
    name: 'Blog',
    component: () => import('../pages/Blog.vue')
  },
  {
    path: '/blog/:blogId/edit',
    name: 'BlogEdit',
    component: () => import('../pages/BlogEdit.vue')
  },
  {
    path: '/blog/add',
    name: 'BlogAdd',
    component: () => import('../pages/BlogEdit.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
