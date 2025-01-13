<script setup lang='ts'>
import type { baseModel } from '@/type';
import { ref } from 'vue';
import { serverUrl, post } from '@/api'
import { ElMessageBox } from 'element-plus'
import type { UploadFile } from 'element-plus';

type Banner = baseModel & {
    baseModel: baseModel;
    PicUrl: string
}
type BannerInput = {
    file : UploadFile | null
}
const bannerData = ref<Banner[]>([])
const isAddBanner = ref(false)
const isUpdateBanner = ref(false)
const bannerAddFormUrl = ref("")
const bannerAddForm = ref<BannerInput>( {
        file :  null
    }
)
function getBannerList() {
    post('/banner/list').then(res => {
        bannerData.value = res as any
        console.log(res)
    }).catch(err => {
        ElMessageBox.alert(err)
    })
}
function submitAddBanner() {
    const formData = new FormData()
    formData.append('file', bannerAddForm.value.file!.raw!)
    post('/banner/create' , {
        headers : {
            'Content-Type': 'multipart/form-data'
        } , 
        data : formData
    }).then(res=>{
        getBannerList()
        isAddBanner.value = false
        bannerAddForm.value.file = null
        bannerAddFormUrl.value = ""
    }).catch(err=>{
    })
}

function deleteBanner(ID : number) {
    post('/banner/delete',{
        data : {
            ID : ID
        }
    }).then(res=>{
        getBannerList()
    }).catch(err=>{
        ElMessageBox.alert(err)
    })
}
function AddBanner() {
    isAddBanner.value = true
}

function handleFileUpload(file : UploadFile) {
    console.log(file)
    bannerAddForm.value.file = file
    bannerAddFormUrl.value = URL.createObjectURL(file!.raw!)
    console.log(bannerAddFormUrl)
}

getBannerList()

</script>
<template>
    <el-table :data="bannerData" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="200px" />
        <el-table-column prop="PicUrl" label="图片" width="500px">
            <template #default="scope">
                <img style="width: 200px;height: 200px;" :src="scope.row.PicUrl" />
            </template>
        </el-table-column>
        <el-table-column>
            <template #default="scope">
                <el-button type="danger" @click="deleteBanner(scope.row.ID)">
                    删除
                </el-button>
            </template>
        </el-table-column>
    </el-table>
    <el-button @click="AddBanner"> 
        添加行 
    </el-button>
    <el-dialog v-model="isAddBanner" title="添加轮播图">
        <el-form  :model="bannerAddForm">
            <el-form-item label="图片">
                <el-upload
                    action="#"
                    :on-change="handleFileUpload"
                    :show-file-list="false"
                    :auto-upload="false"
                >
                <img v-if="bannerAddForm.file" :src="bannerAddFormUrl" style="width: 200px; height: 200px">
                <el-button v-else type="primary">点击上传</el-button>    
                </el-upload>
            </el-form-item>
            <el-button type="primary" @click="submitAddBanner">确定</el-button>
        </el-form> 
    </el-dialog>
</template>
<style lang='scss' scoped></style>