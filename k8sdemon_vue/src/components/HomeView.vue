<template>
    <!-- 前端 -->
    <div class="text">
        <RouterLink to="/user/info"><el-button type="primary">查看用户信息</el-button></RouterLink>
        <p></p>
        <RouterLink to="/user/add"><el-button type="primary">添加用户</el-button></RouterLink>
    </div>
</template>

<script setup lang="ts">
    //javascript or typescript
</script>

<style>
    /* css 样式 */
.text {
    text-align: center;
}
</style>