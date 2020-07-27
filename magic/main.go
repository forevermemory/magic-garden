package main

import (
	"magic/router"
	mylog "magic/utils/log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := router.InitRouter(nil, "api/v1")
	mylog.Success("Listening and serving HTTP on :8000")
	r.Run(":8000")
}
