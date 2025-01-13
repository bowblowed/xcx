<script setup lang='ts'>
import { post } from '@/api';
import type { baseModel } from '@/type';
import { ElMessageBox, type UploadFile, type UploadFiles } from 'element-plus';
import { ref } from 'vue';

export type Category = baseModel & {
    Name: string
    PicUrl: string
}
type CategoryCreateForm = {
    Name : string
    file : File | null
}
type CategorUpdateForm = {
    ID : number
    Name : string
    file : File | null
}
const categoriesData = ref<Category[]>([])
const isAddCategory = ref(false)
const categoryAddForm = ref<CategoryCreateForm>({
    Name : "",
    file : null
})
const isUpdateCategory = ref(false)
const categoryUpdateForm = ref<CategorUpdateForm>({
    ID : 0 ,
    Name : "",
    file : null
})
const updateUrl = ref("")
const uploadUrl = ref("")
function getCategories() {
    post("/category/list").then(res => {
        categoriesData.value = res as any
    }).catch(err => {
        ElMessageBox.alert(err)
    })
}
function addCategory() {
    isAddCategory.value = true
}
function updateCategory(item : Category) {
    categoryUpdateForm.value.ID = item.ID
    categoryUpdateForm.value.Name = item.Name
    updateUrl.value = item.PicUrl
    isUpdateCategory.value = true
}
function updateChange(file : UploadFile) {
    categoryUpdateForm.value.file = file.raw!
    updateUrl.value = URL.createObjectURL(categoryUpdateForm.value.file!)
    console.log(updateUrl)
}

function uploadChange(file : UploadFile) {
    categoryAddForm.value.file = file.raw!
    uploadUrl.value = URL.createObjectURL(file.raw!)
}
function submitAddCategory() {
    if( categoryAddForm.value.Name == "" ) {
        return ElMessageBox.alert("名字不能为空")
    }    
    const formData = new FormData()
    formData.append("file", categoryAddForm.value.file!)
    formData.append("Name", categoryAddForm.value.Name)
    post("/category/create", {
        headers : {
            'Content-Type': 'multipart/form-data'
        } , 
        data : formData
    }).then(res=>{
        getCategories()
        isAddCategory.value = false
    }).catch(err=>{
        console.log(err)
        ElMessageBox.alert(err)
        isAddCategory.value = false
    })
}
function deleteCategory(id: number) {
    post("/category/delete", {
        data : {
            ID : id
        }
    }).then(res=>{
        getCategories()
    }).catch(err=>{
        ElMessageBox.alert(err)
    })
}
function submitUpdateCategory() {
    const formData = new FormData()
    formData.append("file", categoryUpdateForm.value.file!)
    formData.append("ID", categoryUpdateForm.value.ID.toString())
    post("/category/uploadPic" , {
        headers : {
            'Content-Type': 'multipart/form-data'
        } , 
        data : formData
    }).then(res=>{
        getCategories()
    }).catch(err=>{
        ElMessageBox.alert(err)
    }).finally(()=> {
        isUpdateCategory.value = false
        categoryUpdateForm.value.file = null
        updateUrl.value = ""
        categoryUpdateForm.value.file = null
    })
}

getCategories()

</script>
<template>
    <el-table :data="categoriesData" style="width: 100%">
        <el-table-column prop="ID" label="ID" />
        <el-table-column prop="Name" label="名字" />
        <el-table-column prop="PicUrl" label="图片">
            <template #default="scope">
                <img :src="scope.row.PicUrl" style="height: 200px; width: 200px;">
            </template>
        </el-table-column>
        <el-table-column>
            <template #default="scope">
                <el-button type="warning" @click="deleteCategory(scope.row.ID)">删除</el-button>
            </template>
        </el-table-column>
        <el-table-column >
            <template #default="scope">
                <el-button type="primary" @click="updateCategory(scope.row)"> 修改</el-button>
            </template>
        </el-table-column>
    </el-table>
    <el-button type="primary" @click="addCategory">
        新增类别
    </el-button>
    <el-dialog v-model="isAddCategory">
        <el-form :model="categoryAddForm">
            <el-form-item label="名字">
                <el-input  v-model="categoryAddForm.Name" />
            </el-form-item>
            <el-form-item label="图片">
                <el-upload action="#" 
                :show-file-list="false" 
                :on-change="uploadChange"
                :auto-upload="false">
                <img v-if="uploadUrl" :src="uploadUrl" style="width: 200px; height: 200px">
                <el-button v-else size="small" type="primary">点击上传</el-button>
                </el-upload>
            </el-form-item>
            <el-button type="primary" @click="submitAddCategory">确定</el-button>
        </el-form>
    </el-dialog>
    <el-dialog v-model="isUpdateCategory">
        <el-form :model="categoryAddForm">
            <el-form-item label="名字">
                {{ categoryUpdateForm.Name }}
            </el-form-item>
            <el-form-item label="图片">
                <el-upload action="#" 
                :show-file-list="false" 
                :on-change="updateChange"
                :auto-upload="false">
                <img v-if="updateUrl" :src="updateUrl" style="width: 200px; height: 200px">
                <el-button v-else size="small" type="primary"> 点击上传 </el-button>
                </el-upload>
            </el-form-item>
            <el-button type="primary" @click="submitUpdateCategory">确定</el-button>
        </el-form>
    </el-dialog>
</template>
<style lang='scss' scoped></style>