<!-- App.vue文件 -->
<template>
    <div class="text"><h1>K8s 测试项目</h1></div>
    <div>
        <!-- 展示路由组件 -->
		<RouterView></RouterView>
	</div>
</template>
<!-- setup语法糖 -->
<script setup lang="ts">

</script>
<!-- css样式 -->
<style scoped>
.text {
    margin: auto;
    text-align: center;
    padding-top: 10px;
}
</style>