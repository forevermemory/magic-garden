package controller

/*
date:2020-07-30 20:20:24
*/

import (
	"magic/db"
	"magic/service"

	"github.com/gin-gonic/gin"
)

// AddGames add
func AddGames(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.AddGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// UpdateGames update
func UpdateGames(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.UpdateGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// GetGamesByID  get xxx by id
func GetGamesByID(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.GetGamesByID(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// ListGames // list by page condition
func ListGames(c *gin.Context) interface{} {
	var u = db.Games{PageSize: 10, PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.ListGames(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// DeleteGames Delete
func DeleteGames(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.DeleteGames(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}
