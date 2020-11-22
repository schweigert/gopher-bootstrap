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
    path: '/fibo',
    component: () => import(/* webpackChunkName: "fibo" */ '../views/Fibo.vue')
  },
  {
    path: '/anagram',
    component: () => import(/* webpackChunkName: "anagram" */ '../views/Anagram.vue')
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
