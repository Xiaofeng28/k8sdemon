<template>
	<div class="main">
		<el-card class="card" style="max-width: 800px">
			<template #header>
				<div class="card-header">
					<span>用户信息展示</span>
				</div>
			</template>
			<el-table :data="tableData" height="400" style="width: 100%">
				<el-table-column fixed prop="userid" label="用户ID" width="150" />
				<el-table-column prop="username" label="用户名" width="200" />
				<el-table-column prop="password" label="用户密码" width="250" />
				<el-table-column fixed="right" label="操作" min-width="120">
					<template #default="scope">
						<el-button link type="primary" size="small" @click="GetTableInfo(scope.row)">
							编辑
						</el-button>
						<el-button link type="danger" size="small" @click="DeleteOneUser(scope.row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
	</div>
	<!-- 编辑弹框 -->
	<div>
		<el-dialog v-model="dialogVisible" title="更新用户的信息" width="700" :before-close="handleClose">
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
							<el-button type="primary" @click="UpdateOneUserInfo">提交</el-button>
							<el-button @click="resetForm">重置</el-button>
						</el-form-item>
					</el-form>
				</el-card>
			</div>
			<template #footer>
				<div class="dialog-footer">
					<el-button @click="dialogVisible = false">取消</el-button>
				</div>
			</template>
		</el-dialog>
	</div>
</template>

<script lang="ts" setup>
import { ref,reactive } from 'vue'
import { ElMessageBox } from 'element-plus'
import { reqDeleteOneUser, reqUpdateOneUserInfo, reqUserInfo } from '@/api/user';
import type { User, UserArray } from '@/type/user';



// 用户信息列表
let tableData = ref<UserArray>([])

// 获取用户信息
const getUserInfos = async () => {
	let result: UserArray = await reqUserInfo();
	tableData.value = result
}
getUserInfos()

// 编辑用户信息
let dialogVisible = ref(false)
// 表单信息
let form = ref({
	userid: 0,
    username: '',
    password: '',
    confirm: '',
})
// 更新用户信息
const UpdateOneUserInfo = async () => {
	//校验
	if (Test(form.value) == false) {
		return
	}
	// 提交更新的信息
	let result = confirm(`您确认要更改该用户的信息为「用户名:${form.value.username};密码:${form.value.password}」吗？`)
	if (result == false) {
		return
	}
	// 填充提交的信息
	let user:User = {userid:0,username:'',password:''}
	user.userid = form.value.userid
	user.username = form.value.username
	user.password = form.value.password
	// 提交
	let updateresult = await reqUpdateOneUserInfo(user)
	if (updateresult == 0) {
		alert("用户信息更新失败！")	
		return
	}
	// 成功
	alert("更新用户信息成功！")
	getUserInfos()
	dialogVisible.value = false
}
// 获取要更改的用户信息
const GetTableInfo = (row: User) => {
	// 弹框
	dialogVisible.value = true
	form.value.userid = row.userid
	form.value.username = row.username
	form.value.password = row.password
}
// 弹框回调
const handleClose = (done: () => void) => {
	ElMessageBox.confirm('您确定要取消编辑吗？')
		.then(() => {
			done()
		})
		.catch(() => {
			// catch error
		})
}
// 更新表单数据的校验
function Test(form: any)  {
	// 校验
    const hasChinese = /[\u4e00-\u9fa5]/.test(form.password);
    if(hasChinese){
        alert("密码不能有中文！")
        return false
    }
    if(form.password == '' || form.username == ''){
        alert("用户名或密码不能为空！")
        return false
    }
    if(form.password != form.confirm){
        alert("两次输入的密码不相同,请检查！")
        return false
    }
}
// 重置
const resetForm = ()  => {
    form.value.username = ''
    form.value.password = ''
    form.value.confirm = ''
}

// 删除用户信息
const DeleteOneUser = async (row: User) => {
	let relust = confirm(`确定要删除以下用户: 「${row.userid}-${row.username}-${row.password}」 吗？`)
	if (relust == false) {
		return
	}
	// 删除
	let result: number = await reqDeleteOneUser(row.userid)
	if (result == 0) {
		alert("删除用户失败！")
		return
	}
	alert("删除成功！")
	getUserInfos()
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