import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import type { App } from 'vue'
import Banner from '@/components/Banner.vue'
import Layout from '@/components/Layout.vue'
import User from '@/components/User.vue'
import Category from '@/components/Category.vue'
import Product from '@/components/Product.vue'
import PriceTag from '@/components/PriceTag.vue'
const routes : any = [
  { path : '/' , component : Layout , 
    redirect : '/dashbord' , 
    children : [ 
    { path : 'dashbord' , component : HelloWorld , name : "home"} ,
    { path : 'banner' , component : Banner , name : "轮播图"} ,
    { path : 'user' , component : User , name : "用户" } ,
    { path : 'category' , component : Category , name : "类别" } ,
    { path : 'product' , component : Product , name : "产品" } ,
  { path : 'pricetag' , component : PriceTag , name : "价格标签" } ,
  ]}
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes 
})

export  {routes , router} ;
