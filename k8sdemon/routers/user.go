package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/k8sdemon/dao"
	"go_code/k8sdemon/models"
	"net/http"
)

// UserRouter 用户路由
type UserRouter struct {
}

// RegisterUserRouter 注册用户路由
func (u *UserRouter) RegisterUserRouter(router *gin.Engine) {
	// 路由组
	userRouter := router.Group("/user")
	// 具体请求路径 /user/xxx
	userRouter.GET("/hello", u.Hello)
	userRouter.GET("/info", u.GetAllUserInfo)
	userRouter.POST("/add", u.AddOneUser)
	userRouter.POST("/delete", u.DeleteOneUserById)
	userRouter.POST("/update", u.UpdateOneUserInfo)
}

// Hello 请求方法
func (u *UserRouter) Hello(context *gin.Context) {
	// 返回Json
	user := dao.FindFistUser()
	context.JSON(http.StatusOK, gin.H{"message": user})
}

// GetAllUserInfo 获取所有用户
func (u *UserRouter) GetAllUserInfo(context *gin.Context) {
	// 返回Json
	userSlice, err := dao.GetAllUserInfo()
	if err != nil {
		fmt.Println(err)
	}
	context.JSON(http.StatusOK, userSlice)
}

// AddOneUser 添加一个用户
func (u *UserRouter) AddOneUser(context *gin.Context) {
	// 获取请求参数,前端发送的是Jons类型
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		fmt.Println("参数解析失败：", err)
		context.JSON(http.StatusInternalServerError, 500)
		return
	}
	fmt.Println("接收到的参数：", user.UserId, user.UserName, user.Password)
	// 判空
	if user.UserName == "" || user.Password == "" {
		context.JSON(http.StatusInternalServerError, 500)
		return
	}
	// 写入数据库
	num, err := dao.AddOneUser(user)
	if err != nil {
		fmt.Println("插入失败！")
		context.JSON(http.StatusInternalServerError, 500)
		return
	}
	// 写入成功返回插入条数
	context.JSON(http.StatusOK, num)
}

// DeleteOneUserById 删除一个用户根据用户id值
func (u *UserRouter) DeleteOneUserById(context *gin.Context) {
	// 获取用户id
	var userid int
	err := context.ShouldBindJSON(&userid)
	if err != nil {
		fmt.Println("参数解析失败：", err)
		context.JSON(http.StatusInternalServerError, 0)
		return
	}
	// 判断
	if userid == 0 {
		fmt.Println("传入删除的 userid 值为空")
		context.JSON(http.StatusInternalServerError, 0)
		return
	}
	// 删除
	row, err := dao.DeleteOneUserByUserID(userid)
	if err != nil {
		fmt.Println("DeleteOneUserById()函数调用失败")
		return
	}
	// 返回影响的条数
	context.JSON(http.StatusOK, row)
}

// UpdateOneUserInfo 更新用户信息根据用户ID
func (u *UserRouter) UpdateOneUserInfo(context *gin.Context) {
	// 获取参数
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		fmt.Println("UpdateOneUserInfo函数参数解析失败：", err)
		return
	}
	// 判断数据
	if user.UserId == 0 {
		fmt.Println("UpdateOneUserInfo函数中获取的用户id值不能为0")
		return
	}
	// 更新数据
	row, err := dao.UpdateOneUserInfoById(user)
	if err != nil {
		fmt.Println("UpdateOneUserInfo 更新用户数据失败", err)
		return
	}
	// 返回结果
	context.JSON(http.StatusOK, row)
}
