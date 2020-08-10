package controller

import (
	"magic/db"
	"magic/global"
	"magic/service"
	"strconv"

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
	id, err := strconv.Atoi(_id)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GetGardenByID(id)
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
