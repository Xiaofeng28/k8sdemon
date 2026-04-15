<template>
    <div class="main">
        <!-- 卡片组件 -->
        <el-card class="card" style="max-width: 800px">
            <template #header>
                <div class="card-header">
                    <span>添加用户信息</span>
                </div>
            </template>
            <!-- 表单组件 -->
            <el-form :model="form" label-width="auto" style="max-width: 600px">
                <el-form-item label="用户名：">
                    <el-input v-model="form.username" />
                </el-form-item>
                <el-form-item label="密码：">
                    <el-input v-model="form.password" />
                </el-form-item>
                <el-form-item label="确认密码：">
                    <el-input v-model="form.confirm" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="AddOneUser">提交</el-button>
                    <el-button @click="resetForm">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
// 引入
import { reqAddOneUser } from '@/api/user';
import { reactive } from 'vue'

// 表单信息
const form = reactive({
    username: '',
    password: '',
    confirm: '',
})
// 提交表单
const AddOneUser = async () => {
    // 校验
    const hasChinese = /[\u4e00-\u9fa5]/.test(form.password);
    if(hasChinese){
        alert("密码不能有中文！")
        return
    }
    if(form.password == '' || form.username == ''){
        alert("用户名或密码不能为空！")
        return
    }
    if(form.password != form.confirm){
        alert("两次输入的密码不相同,请检查！")
        return
    }
    //二次确认
    let rel = confirm(`您确认要添加以下用户：「用户名：${form.username};密码：${form.password}」吗？`)
    if (rel == false) {
        return
    }
    // 校验通过并提交
	let result: number = await reqAddOneUser(form.username,form.password);
	console.log(result)
    if(result == 1){
        alert("添加用户成功！")
    }
    resetForm()
}

// 重置
const resetForm = ()  => {
    form.username = ''
    form.password = ''
    form.confirm = ''
}

</script>

<style scoped>
.main {
    padding-top: 50px;
}

.card {
    margin: auto;
}

.text {
    text-align: center;
}
</style>