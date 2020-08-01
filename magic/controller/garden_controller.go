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
		return ErrorResponse{-1, err.Error()}
	}
	err = service.InitGarden(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// UpdateGarden update
func UpdateGarden(c *gin.Context) interface{} {
	var u = db.Garden{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.UpdateGarden(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// GetGardenByID  get xxx by id
func GetGardenByID(c *gin.Context) interface{} {
	var u = db.Garden{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.GetGardenByID(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// ----背包-----

// ListGardenKnapsack  分页查询背包
func ListGardenKnapsack(c *gin.Context) interface{} {
	var u = db.GardenFlowerKnapsack{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.ListGardenKnapsack(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

//-----花园帮助----

// GetGardenHelpByID  get 花园帮助 by id
func GetGardenHelpByID(c *gin.Context) interface{} {
	var u = db.GardenHelp{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.GetGardenHelpByID(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// GetGardenHelpTitles  get 花园帮助 标题列表
func GetGardenHelpTitles(c *gin.Context) interface{} {

	data, err := service.GetGardenHelpTitles()
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}
