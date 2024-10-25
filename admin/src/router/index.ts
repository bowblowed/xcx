import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import type { App } from 'vue'
import Banner from '@/components/Banner.vue'
const routes = [
  { path : '/' , component : import('@/components/Layout.vue') , 
    redirect : '/dashbord' , 
    children : [ 
    { path : '/dashbord' , component : HelloWorld , name : "home"} ,
    { path : '/banner' , component : Banner , name : "banner"} ,
  ]}
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes 
})

export  {routes , router} ;
