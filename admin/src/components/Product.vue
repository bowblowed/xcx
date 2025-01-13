<template>
    <el-table :data="productData" style="width: 100%">
        <el-table-column prop="ID" label="ID" />
        <el-table-column prop="Name" label="名字" />
        <el-table-column prop="CategoryName" label="种类" />
        <el-table-column prop="DefaultPrice" label="默认价格(分)" />
        <el-table-column prop="Description" label="描述" />
        <el-table-column prop="Count" label="库存" />
        <el-table-column prop="Description" label="描述" />
        <el-table-column prop="PicUrl" label="图片" />
        <el-table-column prop="DescriptionPicUrls" label="细节图片" />
        <el-table-column prop="Props" label="属性" />
        <el-table-column label="操作" width="120">
            <template #default="scope">
                <el-button type="primary" size="small" @click="handleUpdate(scope.row)">
                    更新
                </el-button>
            </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
            <template #default="scope">
                <el-button type="danger" size="small" @click="handleDelete(scope.row)">
                    删除
                </el-button>
            </template>
        </el-table-column>
    </el-table>
    <el-button @click="createProduct"> 添加商品 </el-button>
    <el-dialog v-model="isCreateProduct" title="添加商品">
        <el-form :model="productCreate">
            <el-form-item label="名字">
                <el-input v-model="productCreate.Name" />
            </el-form-item>
            <el-form-item label="种类">
                <el-select v-model="productCreate.CategoryId" placeholder="种类">
                    <el-option v-for="category in categories" :key="category.ID" :value="category.ID"
                        :label="category.Name" />
                </el-select>
            </el-form-item>
            <el-form-item label="默认价格(分)">
                <el-input v-model="productCreate.DefaultPrice" />
            </el-form-item>
            <el-form-item label="库存">
                <el-input v-model="productCreate.Count" />
            </el-form-item>
            <el-form-item label="描述">
                <el-input v-model="productCreate.Description" />
            </el-form-item>
            <el-form-item label="图片">
                <el-upload :on-change="createProductUploadChange" action="#" :show-file-list="false"
                    :auto-upload="false">
                    <img v-if="createProductPicUrl" style="width: 200px; height: 200px;" :src="createProductPicUrl" />
                    <el-icon v-else>
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item label="细节图片">
                <el-upload v-model:file-list="productCreate.DescriptionFiles" list-type="picture-card" action="#"
                    :auto-upload="false">
                    <el-icon>
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item label="属性">
                <div>
                    <el-tree 
                    :data="createProductProps?.children" 
                    show-checkbox node-key="id" 
                    default-expand-all
                    :expand-on-click-node="false">
                        <template #default="{ node, data }">
                            <el-input v-model="data.Name">  </el-input>
                            <span>
                                <a @click="treeInsert(data)"> 添加子属性 </a>
                                <a style="margin-left: 8px" @click="treeRemove(node, data)"> 删除 </a>
                            </span>
                        </template>
                    </el-tree>
                    <el-button @click="treeInsert(createProductProps!)"> 添加属性 </el-button>
                </div>
            </el-form-item>
            <el-button type="primary" @click="submitCreateProduct">
                提交
            </el-button>
        </el-form>

    </el-dialog>
    <el-dialog v-model="isUpdateProduct" title="更新商品">
        <el-form :model="productUpdate">
            <el-form-item label="名字">
                <el-input v-model="productUpdate.Name" />
            </el-form-item>
            <el-form-item label="种类">
                <el-select v-model="productUpdate.CategoryId">
                    <el-option v-for="category in categories" :key="category.ID" :value="category.ID"
                        :label="category.Name" />
                </el-select>
            </el-form-item>
            <el-form-item label="默认价格(分)">
                <el-input v-model="productUpdate.DefaultPrice" />
            </el-form-item>
            <el-form-item label="库存">
                <el-input v-model="productUpdate.Count" />
            </el-form-item>
            <el-form-item label="描述">
                <el-input v-model="productUpdate.Description" />
            </el-form-item>
            <el-form-item label="图片">
                <el-upload :on-change="updateProductUploadChange" action="#" :show-file-list="false"
                    :auto-upload="false">
                    <img v-if="updateProductPicUrl" style="width: 200px; height: 200px;" :src="updateProductPicUrl" />
                    <el-icon v-else>
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item label="细节图片">
                <el-upload v-model:file-list="productUpdate.DescriptionFiles" list-type="picture-card" action="#"
                    :auto-upload="false">
                    <el-icon>
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item label="属性">
                <div>
                    <el-tree :data="updateProductProps?.children" show-checkbox node-key="id" default-expand-all
                        :expand-on-click-node="false">
                        <template #default="{ node, data }">
                            <el-input v-model="data.Name">  </el-input>
                            <span>
                                <a @click="treeInsert(data)"> 添加子属性 </a>
                                <a style="margin-left: 8px" @click="treeRemove(node, data)"> 删除 </a>
                            </span>
                        </template>
                    </el-tree>
                    <el-button @click="treeInsert(updateProductProps!)"> 添加属性 </el-button>
                </div>
            </el-form-item>
            <el-button type="primary" @click="submitUpdateProduct">
                提交
            </el-button>
        </el-form>
    </el-dialog>
</template>
<script setup lang='ts'>
import type { baseModel } from '@/type';
import { ref } from 'vue';
import type { Category } from './Category.vue';
import { post } from '@/api';
import { Plus } from '@element-plus/icons-vue'
import type { UploadFile, UploadFiles } from 'element-plus';
type Product = baseModel & {
    CategoryId  : number 
    CategoryName: string
    Name: string
    DefaultPrice: number
    Count: number
    Description: string
    PicUrl: string
    DescriptionPicUrls: string[]
    Props: Map<string, string[]>
}

type ProductCreateFrom = {
    CategoryId: number
    Name: string
    DefaultPrice: number
    Count: number
    Description: string
    PicFile: UploadFile | null
    DescriptionFiles: UploadFiles
    Props:  Map<string, string[]>
}

type ProductUpdateFrom = ProductCreateFrom & {
    ID: number
}

type TreeNode = {
    id: number
    Name: string
    Parent?: TreeNode
    children: TreeNode[]
}
function treeInsert(data : TreeNode) {
    data.children.push({
        Parent: data,
        id: data.children.length,
        Name: '',
        children: []
    })
}
function treeRemove(node: TreeNode, data: TreeNode) {
    console.log(data)
    const p = data.Parent
    var i = 0 
    for( ; i < p!.children.length; i++ ) {
        if (p!.children[i] === data) {
            p!.children.splice(i, 1)
            return
        }
    }
}
function TreeToPropMap(node: TreeNode): Map<string, string[]> {
    const m = new Map()
    for (let i = 0; i < node.children.length; i++) {
        var tmp = []
        for (let j = 0; j < node.children[i].children.length; j++) {
            tmp.push(node.children[i].children[j].Name)
        }
        m.set(node.children[i].Name, tmp)
    }
    return m
}

function PropMapToTree(m: Map<string, string[]>): TreeNode {
    console.log(m)
    var id = 0
    var root: TreeNode = {
        id: 0,
        Name: '',
        children: []
    }
    for (let [key, value] of m) {
        var tmp: TreeNode = {
            Parent: root,
            id: id++,
            Name: key ,
            children: []
        }
        for (let j = 0; j < value.length; j++) {
            tmp.children.push({
                Parent: tmp,
                id: id++,
                Name: value[j],
                children: []
            })
        }
        root.children.push(tmp)
    }
    console.log(root)
    return root
}


const productData = ref<Product[]>([])
const productCreate = ref<ProductCreateFrom>({
    CategoryId: 0,
    Name: '',
    DefaultPrice: 0,
    Count: 0,
    Description: '',
    PicFile: null,
    DescriptionFiles: [],
    Props: new Map()
})
const isCreateProduct = ref(false)
const createProductPicUrl = ref("")
const createProductProps = ref<TreeNode>({
    id: 0,
    Name: '',
    children: []
})

const productUpdate = ref<ProductUpdateFrom>({
    ID: 0,
    CategoryId: 0,
    Name: '',
    DefaultPrice: 0,
    Count: 0,
    Description: '',
    PicFile: null,
    DescriptionFiles: [],
    Props: new Map()
})
const isUpdateProduct = ref(false)
const updateProductPicUrl = ref("")
const updateProductProps = ref<TreeNode>({
    id: 0,
    Name: '',
    children: []
})

const categories = ref<Category[]>([])

function getCategories() {
    post("/category/list").then(res => {
        categories.value = res as any
        categories.value.unshift({
            ID: 0,
            Name: '无',
        } as any)
    }).catch(err => {
        ElMessageBox.alert(err)
    })
}
function getProducts() {
    post("/product/list").then(res => {
        productData.value = res as any
    }).catch(err => {
        ElMessageBox.alert(err)
    })
}

function createProduct() {
    isCreateProduct.value = true
}

function createProductUploadChange(file: UploadFile) {
    productCreate.value.PicFile = file
    createProductPicUrl.value = URL.createObjectURL(file.raw!)
}

function submitCreateProduct() {
    const formData = new FormData()
    if(productCreate.value.Name == "") {
        return ElMessageBox.alert("名字不能为空")
    }
    formData.append("Name" ,productCreate.value.Name)
    formData.append("CategoryId", productCreate.value.CategoryId.toString())
    formData.append("DefaultPrice", productCreate.value.DefaultPrice.toString())
    formData.append("Count", productCreate.value.Count.toString())
    formData.append("Description", productCreate.value.Description)
    formData.append("PicFile", productCreate.value.PicFile ? productCreate.value.PicFile.raw! : "")
    console.log(productCreate.value.DescriptionFiles)
    if (productCreate.value.DescriptionFiles.length > 0) {
        formData.append("DescriptionFilesLen", productCreate.value.DescriptionFiles.length.toString())
        for(let i = 0; i < productCreate.value.DescriptionFiles.length; i++) {
            formData.append("DescriptionFiles_" + i.toString(), productCreate.value.DescriptionFiles[i].raw!)
        }
    }
    const PropMap = TreeToPropMap(createProductProps.value!)
    let obj : any = {};
    PropMap.forEach((value, key) => {
       obj[key] = value;
     });
    formData.append("Props", JSON.stringify(obj))
    post("/product/create" , {
        headers : {
            'Content-Type': 'multipart/form-data'
        } , 
        data : formData
    }).then(res=>{
        getProducts()
    }).catch(err=>{
        ElMessageBox.alert(err)
    }).finally(()=> {
        isCreateProduct.value = false
    })

}

function submitUpdateProduct() {
    const formData = new FormData()
    formData.append("ID", productUpdate.value.ID.toString())
    formData.append("Name", productUpdate.value.Name)
    formData.append("CategoryId", productUpdate.value.CategoryId.toString())
    formData.append("DefaultPrice", productUpdate.value.DefaultPrice.toString())
    formData.append("Count", productUpdate.value.Count.toString())
    formData.append("Description", productUpdate.value.Description)
    formData.append("PicFile", productUpdate.value.PicFile ? productUpdate.value.PicFile.raw! : "")
    if (productUpdate.value.DescriptionFiles.length > 0) {
        formData.append("DescriptionFilesLen", productUpdate.value.DescriptionFiles.length.toString())
        for(let i = 0; i < productUpdate.value.DescriptionFiles.length; i++) {
            formData.append("DescriptionFiles_" + i.toString(), productUpdate.value.DescriptionFiles[i].raw!)
        }
    }
    const PropMap = TreeToPropMap(updateProductProps.value!)
    let obj : any = {};
    PropMap.forEach((value, key) => {
       obj[key] = value;
     });
    formData.append("Props", JSON.stringify(obj))
    post("/product/update", {
        headers : {
            'Content-Type': 'multipart/form-data'
        } , 
        data : formData
    }).then(res=>{
        getProducts()
    }).catch(err=>{
        ElMessageBox.alert(err)
    }).finally(()=> {
        isUpdateProduct.value = false
    })
}   

function handleUpdate(row: Product) {
    isUpdateProduct.value = true
    var descriptionFiles = []
    for (let i = 0; i < row.DescriptionPicUrls.length; i++) {
        descriptionFiles.push({
            url : row.DescriptionPicUrls[i]
        } as UploadFile)
    }
    productUpdate.value = {
        ID : row.ID,
        CategoryId: row.CategoryId,
        Name: row.Name,
        DefaultPrice: row.DefaultPrice,
        Count: row.Count,
        Description: row.Description,
        PicFile: null,
        DescriptionFiles: descriptionFiles ,
        Props: new Map(Object.entries(row.Props))
    }
    updateProductProps.value = PropMapToTree(productUpdate.value.Props)
    updateProductPicUrl.value = row.PicUrl
    console.log(productUpdate.value)
}

function handleDelete(row: Product) {
    ElMessageBox.confirm("确定删除该商品？").then(()=>{
        post("/product/delete", {
            data: {
                ID: row.ID
            }
        }).then(res=>{
            getProducts()
        }).catch(err=>{
            ElMessageBox.alert(err)
        })
    })
}

function updateProductUploadChange(file: UploadFile) {
    productUpdate.value.PicFile = file
    updateProductPicUrl.value = URL.createObjectURL(file.raw!)
}

getProducts()
getCategories()
</script>
<style lang='scss' scoped></style>