package routers

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.Engine) {
	// 注册用户路由
	new(UserRouter).RegisterUserRouter(router)
}
