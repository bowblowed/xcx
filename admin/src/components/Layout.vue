<script setup lang="ts">
import Asideitem from './Asideitem.vue';
import { routes } from '@/router'
import { RouterView, useRouter, type RouteRecord, type RouteRecordRaw } from 'vue-router'
const router = useRouter()

function routeChange(route: RouteRecordRaw) {
  router.push(route.path)
}

</script>
<template>
    <div class="main-layout">
      <el-container class="layout-container">
        <el-aside width="200px" class="layout-aside">
          <el-scrollbar class="layout-aside-scrollbar">
            <Asideitem v-for="route in routes[0].children" :route="route" @click="routeChange(route)">
            </Asideitem>  
          </el-scrollbar> 
        </el-aside>
        <el-container>
          <el-header>Header</el-header>
          <RouterView />
        </el-container>
      </el-container>
    </div>
</template>
<style lang="css" scoped>
.main-layout {
  width: 100%;;
  height: 100%;
}
.layout-container {
  width: 100%;
  height: 100%;
}
.layout-aside-scrollbar {
  width: 100%;
  height: 100%;
}

.layout-aside {
  background-color: #909399;
  display: flex;
  flex-direction: column;
  align-items: start;
}
</style>