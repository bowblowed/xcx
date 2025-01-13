<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { routes } from '../router/index'
const router = useRouter()
function ItemClick(item: any) {
  router.push(item.path)
}

</script>

<template>
  <div class="layout">
  <el-row class="layout-row">
    <el-col :span="3">
      <el-menu class="mune">
        <el-menu-item v-for="(item, index) in routes[0].children" :key="index" @click="ItemClick(item)">
          {{ item.name }}
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="21">
      <router-view class="main-view" />
    </el-col>
  </el-row>
  </div >
</template>


<style>
.layout{
  height: 100%;
  width: 100%;
}

.layout-row{
  height: 100%;
}
.mune{
  height: 100%;
}
.main-view {
  height: 100%;
}
</style>