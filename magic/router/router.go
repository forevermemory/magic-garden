package router

import (
	"magic/controller"
	"magic/utils/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

// InitRouterV1 初始化gin路由
func InitRouterV1(r *gin.Engine, prefix string) *gin.Engine {
	if r == nil {
		r = gin.Default()
	}
	// TODO user-agent 判断
	r.Use(middleware.Cors()) // 跨域
	// 静态文件  建议前端build之后 交给ngnix管理
	//r.Static(prefix+"/static", "./static")
	//r.LoadHTMLFiles("./static/index.html")
	// r.StaticFS("/static", http.Dir("./static"))
	statikFS, _ := fs.New()
	r.StaticFS("/static", statikFS)
	// r.GET("", func(c *gin.Context) {
	// 	t := template.Must(template.New("index").Parse(INDEXHTML))
	// 	t.ExecuteTemplate(c.Writer, "index", "")
	// })

	r.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	// 生成验证码图片
	r.GET("/captcha", controller.GenerateCaptcha)

	// 用户注册登陆相关
	user := r.Group(prefix + "/user")
	{
		user.GET("/sendsms", route(controller.RegisterUserSendMsg))
		user.POST("/login", controller.UserLogin)
		user.POST("/registe", route(controller.AddUsers))
		user.POST("/update", route(controller.UpdateUsers))
		user.POST("/reset/password", route(controller.UpdateUsersPassword))
		user.GET("/user/get", route(controller.GetUsersByID))
	}

	userGame := r.Group(prefix + "/usergame")
	{
		userGame.POST("/add", route(controller.UserAddGames))
		userGame.POST("/delete", route(controller.UserDeleteGames))
	}
	garden := r.Group(prefix + "/garden")
	{
		garden.GET("/init", route(controller.InitGarden))

		// TODO
		garden.POST("/signin", route(controller.UserAddGames))
	}

	return r
}
