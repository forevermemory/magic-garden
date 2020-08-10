package main

import (
	"fmt"
	"magic/cron"
	"magic/global"
	"magic/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)            // 不显示日志
	global.InitGlobalConn()                 // 初始化连接
	r := router.InitRouterV1(nil, "api/v1") // 设置路由

	go cron.InitCronTask()

	fmt.Println("Listening and serving HTTP on :8000")
	r.Run(":8000")
}
