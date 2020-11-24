import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/why',
    component: () => import(/* webpackChunkName: "why" */ '../views/Why.vue')
  },
  {
    path: '/install',
    component: () => import(/* webpackChunkName: "install" */ '../views/Install.vue')
  },
  {
    path: '/compile',
    component: () => import(/* webpackChunkName: "compile" */ '../views/Compile.vue')
  },
  {
    path: '/hello',
    component: () => import(/* webpackChunkName: "hello" */ '../views/Hello.vue')
  },
  {
    path: '/funcs',
    component: () => import(/* webpackChunkName: "funcs" */ '../views/Funcs.vue')
  },
  {
    path: '/objects',
    component: () => import(/* webpackChunkName: "objects" */ '../views/Objects.vue')
  },
  {
    path: '/packages',
    component: () => import(/* webpackChunkName: "packages" */ '../views/Packages.vue')
  },
  {
    path: '/tests',
    component: () => import(/* webpackChunkName: "tests" */ '../views/Tests.vue')
  },
  {
    path: '/fibo',
    component: () => import(/* webpackChunkName: "fibo" */ '../views/Fibo.vue')
  },
  {
    path: '/anagram',
    component: () => import(/* webpackChunkName: "anagram" */ '../views/Anagram.vue')
  },
  {
    path: '/gin',
    component: () => import(/* webpackChunkName: "gin" */ '../views/Gin.vue')
  },
  {
    path: '/*',
    component: () => import(/* webpackChunkName: "404" */ '../views/404.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
