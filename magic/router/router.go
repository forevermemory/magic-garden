package router

import (
	"magic/controller"
	"magic/utils/middleware"

	"github.com/gin-gonic/gin"
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

	// r.GET("", func(c *gin.Context) {
	// 	t := template.Must(template.New("index").Parse(INDEXHTML))
	// 	t.ExecuteTemplate(c.Writer, "index", "")
	// })

	// statikFS, _ := fs.New()
	// r.StaticFS("/static", statikFS)

	r.Static("/static", "./static")

	r.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	// 生成验证码图片
	r.GET("/captcha", controller.GenerateCaptcha)

	r.POST("/login", controller.UserLogin)
	r.POST("/register", route(controller.AddUsers))
	// 用户注册登陆相关
	user := r.Group(prefix + "/user")
	{
		user.GET("/sendsms", route(controller.RegisterUserSendMsg))
		user.GET("/registe/isUsernameExist", route(controller.IsUsernameExist))
		user.POST("/update", route(controller.UpdateUsers))
		user.POST("/reset/password", route(controller.UpdateUsersPassword))
		user.GET("/user/get", route(controller.GetUsersByID))
		// 好友系统 TODO
		user.POST("/friends/add", route(controller.UpdateUsersPassword))
		user.POST("/friends/delete", route(controller.UpdateUsersPassword))
		user.POST("/friends/list", route(controller.UpdateUsersPassword))
		user.POST("/friends/list/:oid", route(controller.UpdateUsersPassword))

	}

	// 等级
	levels := r.Group(prefix + "/levels")
	{
		levels.POST("/add", route(controller.AddUserLevel))
		levels.POST("/update", route(controller.UpdateUserLevel))
		levels.GET("/list/:oid", route(controller.GetUserLevelByID))
		levels.GET("/list", route(controller.ListUserLevel))
		levels.GET("/delete", route(controller.DeleteUserLevel))
	}
	// games
	games := r.Group(prefix + "/games")
	{
		games.POST("/add", route(controller.AddGames))
		games.POST("/update", route(controller.UpdateGames))
		games.GET("/list/:oid", route(controller.GetGamesByID))
		games.GET("/list", route(controller.ListGames))
		games.GET("/delete", route(controller.DeleteGames))
		// games.POST("/userAddGames", route(controller.UserAddGames))
	}

	userGame := r.Group(prefix + "/usergame")
	{
		userGame.GET("/list", route(controller.UserGamesList))
		userGame.POST("/add", route(controller.UserAddGames))
		userGame.POST("/delete", route(controller.UserDeleteGames))
		userGame.POST("/order", route(controller.UserOrderGames))
	}
	garden := r.Group(prefix + "/garden")
	{
		// gb 获得历史 购买种子道具
		garden.GET("/gb/history", route(controller.ListGardenGbDetail))
		garden.POST("/gb/shop/buy/seed", route(controller.BuyShopSeed))
		garden.POST("/gb/shop/buy/prop", route(controller.BuyShopProp))
		// 魔法屋
		garden.GET("/magician/list", route(controller.ListGardenMagician))
		garden.POST("/magician/detail", route(controller.GardenMagicianDetail))
		garden.POST("/magician/synthesis", route(controller.GardenMagicianSynthesis))
		// 花房  花篮和花瓶
		garden.GET("/house/list", route(controller.GardenHouseList))
		garden.GET("/house/statistics", route(controller.GardenHouseStatistics))

		// 送花系统 好友系统

		// 偷花 帮助系统  您成功采摘了好友一朵 xxx,花朵已进入您的花篮
		// 采摘失败 已经摘过了,做人不要太贪心哦

		// 初始化花园
		garden.GET("/init", route(controller.InitGarden))
		// 查询花园详情
		garden.GET("/list/:oid", route(controller.GetGardenByID))
		// 更新花园的公告等信息
		garden.POST("/update/baseinfo", route(controller.UpdateGarden))
		// 查看背包
		garden.GET("/knapsack/list", route(controller.ListGardenKnapsack))
		// 花园帮助
		garden.GET("/help/list", route(controller.GetGardenHelpTitles))
		garden.GET("/help/detail", route(controller.GetGardenHelpByID))
		//  花园签到
		garden.GET("/signin", route(controller.GardenEveryDaySignin))
		garden.GET("/signin/list", route(controller.ListGardenSigninHistory))
		// 花盆
		garden.GET("/flowerpot/list", route(controller.GardeFlowerpotList))
		garden.GET("/flowerpot/detail", route(controller.GardeFlowerpotDetail))
		garden.POST("/flowerpot/sow", route(controller.GardeFlowerpotSow))
		garden.POST("/flowerpot/lookafter", route(controller.GardeFlowerpotLookAfter))
		garden.POST("/flowerpot/remove", route(controller.GardeFlowerpotRemove))
		garden.POST("/flowerpot/dyeing", route(controller.GardeFlowerpotDyeing))
		garden.POST("/flowerpot/fertilizer", route(controller.GardeFlowerpotFertilizer))
		garden.POST("/flowerpot/harvest", route(controller.HarvestFlower))
		// garden.GET("/flowerpot/purchase", route(controller.GardeFlowerpotPurchase))
		// 收支统计 TODO
	}

	return r
}
