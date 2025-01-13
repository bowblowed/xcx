<template>
    <div>
      <el-table :data="priceTagData" style="width: 100%">
        <el-table-column prop="ProductID" label="产品ID" />
        <el-table-column prop="UserID" label="用户ID" />
        <el-table-column label="属性对" />
        <el-table-column prop="Price" label="价格" />
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="danger" size="small" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-button @click="createPriceTag">添加价格标签</el-button>
      <el-dialog v-model="isCreatePriceTag" title="添加价格标签">
        <el-form :model="priceTagCreate">
          <el-form-item label="产品ID">
            <el-input type="number" v-model.number="priceTagCreate.ProductID" />
          </el-form-item>
          <el-form-item label="用户ID">
            <el-input  type="number" v-model.number="priceTagCreate.UserID" />
          </el-form-item>
          
          <el-form-item label="键值对">
            <el-input v-model="newPropKey" placeholder="键" />
            <el-input v-model="newPropValue" placeholder="值" />
            <el-button @click="addKeyValuePair">添加键值对</el-button>
            <br></br>
            <div v-for="(value, key) in priceTagCreate.PropPair">
              {{ key }}: {{ value }}
              <el-button @click="removeKeyValuePair(value[0])">删除</el-button>
            </div>
          </el-form-item>
          <el-form-item  label="价格">
            <el-input type="number" v-model.number="priceTagCreate.Price" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitHandlePriceTag()">确定</el-button>
            <el-button @click="isCreatePriceTag = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { post } from '@/api';
import { ref } from 'vue';
import User from './User.vue';

const priceTagData = ref([]);
const isCreatePriceTag = ref(false);
const priceTagCreate = ref({
    ProductID: 0,
    UserID: 0,
    PropPair: new Map<string, string>(),
    Price: 0
});

const newPropKey = ref('');
const newPropValue = ref('');
function initpriceTagData() {
    post('/product/tag/list').then(res => {
        priceTagData.value = res as any
    }).catch(err => {
        alert(err)
    })
}

function addKeyValuePair() {
    if (newPropKey.value && newPropValue.value) {
        priceTagCreate.value.PropPair.set(newPropKey.value, newPropValue.value);
        newPropKey.value = '';
        newPropValue.value = '';
    }
}
function removeKeyValuePair(key : string) {
    console.log(key);
    priceTagCreate.value.PropPair.delete(key);
}
function handleUpdate( row ) {
 // 更新逻辑
 }
  
function handleDelete( row ) {
    post('/product/tag/delete', {
        data : {
            ID : row.ID
        }
    }).then(res => {
        initpriceTagData()
    }).catch(err => {
        alert(err)
    })
}
  
function createPriceTag() {
    isCreatePriceTag.value = true;
    priceTagCreate.value = {
        ProductID: 0,
        UserID: 0,
        PropPair: new Map<string, string>(),
        Price: 0 ,
    }
}

async function submitHandlePriceTag() {
    if( !priceTagCreate.value.ProductID ) {
        alert('产品ID不能为空')
        return
    }
    console.log(priceTagCreate.value)
    await post('/product/tag/create', {
        data : {
            ProductId : priceTagCreate.value.ProductID ,
            UserId : priceTagCreate.value.UserID ,
            PropPairs : priceTagCreate.value.PropPair ,
	        Price : priceTagCreate.value.Price , 
        }
    }).then( res =>
        initpriceTagData()
    ).catch(err => {
        alert(err)
    })
    isCreatePriceTag.value = false
}

initpriceTagData()
</script>
  