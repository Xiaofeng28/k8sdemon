package main

import (
	"github.com/gin-gonic/gin"
	"go_code/k8sdemon/config"
	"go_code/k8sdemon/routers"
	"strconv"
)

func main() {
	// 1.配置初始化
	myConfig, err := config.GetConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	// 连接数据库
	err = config.InitMysqlConnect(*myConfig)
	if err != nil {
		panic(err)
	}

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 注册路由
	routers.RegisterRouter(router)

	// 启动服务
	router.Run(myConfig.System.Host + ":" + strconv.Itoa(myConfig.System.Port))
}
