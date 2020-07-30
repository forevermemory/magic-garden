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
	err = service.UserAddGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
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
