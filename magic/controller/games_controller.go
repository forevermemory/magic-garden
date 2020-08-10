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
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.AddGames(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// UpdateGames update
func UpdateGames(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.UpdateGames(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// GetGamesByID  get xxx by id
func GetGamesByID(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GetGamesByID(u.ID)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return OKResponse{0, data}
}

// ListGames // list by page condition
func ListGames(c *gin.Context) interface{} {
	var u = db.Games{PageSize: 10, PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListGames(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return OKResponse{0, data}
}

// DeleteGames Delete
func DeleteGames(c *gin.Context) interface{} {
	var u = db.Games{}
	err := c.ShouldBind(&u)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	err = service.DeleteGames(u.ID)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok"}
}
