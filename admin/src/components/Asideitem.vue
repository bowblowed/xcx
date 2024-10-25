<script setup lang='ts'>
import type { RouteRecordRaw } from 'vue-router';

const props = defineProps<{
    route : RouteRecordRaw
}>();
const emits = defineEmits<{
    (e: 'click'): void
}>();
const handleClick = () => {
    emits('click');
}

</script>
<template>
    <el-button type="info" class="item-button" @click="handleClick">
        {{  props.route.name }}
    </el-button>    
</template>
<style lang='css' scoped>
.item-button {
    border-radius: 0px;
    border: none;
    padding: 0px;
    width: 100%;
    height: 80px;
    margin-left: 0px;
}

</style>