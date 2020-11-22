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
    path: '/fibo',
    name: 'Fibonacci',
    component: () => import(/* webpackChunkName: "fibo" */ '../views/Fibo.vue')
  },
  {
    path: '/anagram',
    name: 'Anagram',
    component: () => import(/* webpackChunkName: "anagram" */ '../views/Anagram.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
