package controller

import (
	"magic/db"
	"magic/global"
	"magic/service"

	"github.com/gin-gonic/gin"
)

/*
date:2020-07-30 20:00:37
*/

// InitGarden 初始化花园
func InitGarden(c *gin.Context) interface{} {
	var u = global.UserAddGamesParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	err = service.InitGarden(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok"}
}

// UpdateGarden update
func UpdateGarden(c *gin.Context) interface{} {
	var u = db.Garden{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.UpdateGarden(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GetGardenByID  get xxx by id
func GetGardenByID(c *gin.Context) interface{} {
	_id := c.Param("oid")
	data, err := service.GetGardenByID(_id)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// ----背包-----

// ListGardenKnapsack  分页查询背包
func ListGardenKnapsack(c *gin.Context) interface{} {

	var u = db.GardenFlowerKnapsack{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListGardenKnapsack(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

//-----花园帮助----

// GetGardenHelpByID  get 花园帮助 by id
func GetGardenHelpByID(c *gin.Context) interface{} {
	var u = db.GardenHelp{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GetGardenHelpByID(u.ID)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GetGardenHelpTitles  get 花园帮助 标题列表
func GetGardenHelpTitles(c *gin.Context) interface{} {
	data, err := service.GetGardenHelpTitles()
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// -------花园签到

// GardenEveryDaySignin  get 花园帮助 标题列表
func GardenEveryDaySignin(c *gin.Context) interface{} {
	var u = global.GardenParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardenEveryDaySignin(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}

	return Response{Code: 0, Msg: "ok", Data: data}
}

// return Response{Code: -1,Msg: "ok",Data: data}
// return Response{Code: -1,Msg: err.Error()}

// ListGardenSigninHistory // list by page condition
func ListGardenSigninHistory(c *gin.Context) interface{} {
	var u = db.GardenSigninHistory{PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListGardenSigninHistory(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// 花盆 -----

// GardeFlowerpotList  查看某个花盆
func GardeFlowerpotList(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotList(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotDetail  查看某个花盆
func GardeFlowerpotDetail(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotDetail(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotSow  播种一个或者多个花盆
func GardeFlowerpotSow(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotSow(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotLookAfter  除草浇水除虫操作
func GardeFlowerpotLookAfter(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotLookAfter(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotRemove  移除花盆中成长的花朵🌹
func GardeFlowerpotRemove(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotRemove(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotDyeing  染色
func GardeFlowerpotDyeing(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotDyeing(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardeFlowerpotFertilizer  施肥
func GardeFlowerpotFertilizer(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardeFlowerpotFertilizer(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// HarvestFlower  收获
func HarvestFlower(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.HarvestFlower(u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "收获成功", Data: data}
}

// --------gb获取历史记录

// ListGardenGbDetail // list by page condition
func ListGardenGbDetail(c *gin.Context) interface{} {
	var u = db.GardenGbDetail{PageSize: 10, PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListGardenGbDetail(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// BuyShopSeed // 购买种子
func BuyShopSeed(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.BuyShopSeed(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// BuyShopProp // 购买道具
func BuyShopProp(c *gin.Context) interface{} {
	var u = global.GardenPotParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.BuyShopProp(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// ListGardenMagician // 魔法屋内部可合成列表
func ListGardenMagician(c *gin.Context) interface{} {
	var u = db.GardenSeeds{PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListGardenMagician(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardenMagicianDetail // 魔法屋查询一个种子合成所需材料
func GardenMagicianDetail(c *gin.Context) interface{} {
	var u = global.MagicianParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardenMagicianDetail(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardenMagicianSynthesis // 魔法屋合成
func GardenMagicianSynthesis(c *gin.Context) interface{} {
	var u = global.MagicianParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardenMagicianSynthesis(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardenHouseList // 花房花朵分页查询
func GardenHouseList(c *gin.Context) interface{} {
	var u = global.MagicianParams{Page: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardenHouseList(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GardenHouseStatistics // 花房花朵分页查询 统计
func GardenHouseStatistics(c *gin.Context) interface{} {
	var u = global.MagicianParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GardenHouseStatistics(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}
