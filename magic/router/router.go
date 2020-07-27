package router

import (
	"html/template"
	"magic/utils/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

// InitRouter 初始化gin路由
func InitRouter(r *gin.Engine, prefix string) *gin.Engine {
	if r == nil {
		r = gin.Default()
	}
	r.Use(middleware.Cors()) // 跨域
	// 静态文件  建议前端build之后 交给ngnix管理
	//r.Static(prefix+"/static", "./static")
	//r.LoadHTMLFiles("./static/index.html")
	// r.StaticFS("/static", http.Dir("./static"))
	statikFS, _ := fs.New()
	r.StaticFS("/static", statikFS)
	r.GET("", func(c *gin.Context) {
		t := template.Must(template.New("index").Parse(INDEXHTML))
		t.ExecuteTemplate(c.Writer, "index", "")
	})

	r.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// aaa := r.Group(prefix + "/aaa")
	// {
	// 	// aaa.GET("/delete", route(controller.DeleteApplicationUpgradeHistory))
	// }

	return r
}
