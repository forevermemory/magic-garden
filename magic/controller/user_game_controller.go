package controller

import (
	"magic/global"
	"magic/service"

	"github.com/gin-gonic/gin"
)

// UserAddGames 用户添加游戏
func UserAddGames(c *gin.Context) interface{} {
	var u = global.UserAddGamesParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.UserAddGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// UserDeleteGames 用户删除游戏
func UserDeleteGames(c *gin.Context) interface{} {
	var u = global.UserAddGamesParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.UserDeleteGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// UserGamesList 用户添加游戏列表
func UserGamesList(c *gin.Context) interface{} {
	var u = global.UserAddGamesParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.UserGamesList(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// UserOrderGames 用户添加游戏列表
func UserOrderGames(c *gin.Context) interface{} {
	var u = global.UserAddGamesParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.UserOrderGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}
