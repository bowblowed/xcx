<script setup lang='ts'>
import type { baseModel } from '@/type';
import { ref } from 'vue';
import { post } from '@/api';
import { ElMessageBox } from 'element-plus'; // 导入 ElMessageBox

type User = baseModel & {
    Type: string
    Name: string
    WxNumber: string
    WxOpenID: string
    PhoneNumber: string
    Address: string
}
const userData = ref<User[]>([])

post('/user/list').then(res=>{
    userData.value = res as any
}).catch(err=>{
    ElMessageBox.alert(err)
})

</script>
<template>
    <el-table
    :data="userData" style="width: 100%">
        <el-table-column prop="ID" label='ID' />
        <el-table-column prop="Name" label='姓名' />
        <el-table-column prop="WxNumber" label='微信号' />
        <el-table-column prop="WxOpenId" label='微信openid' />
        <el-table-column prop="PhoneNumber" label='电话号' />
        <el-table-column prop="Address" label='地址' />
    </el-table>
</template>
<style lang='scss' scoped></style>