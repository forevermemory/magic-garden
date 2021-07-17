package controller

/*
date:2021-07-17 14:20:14
*/

import (
	"magic/db"
	"magic/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddUserLevel add
func AddUserLevel(c *gin.Context) interface{} {
	var req = db.UserLevel{}
	err := c.ShouldBind(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.AddUserLevel(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "添加成功", Data: data}
}

// UpdateUserLevel update
func UpdateUserLevel(c *gin.Context) interface{} {
	var req = db.UserLevel{}
	err := c.ShouldBind(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.UpdateUserLevel(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "更新成功", Data: data}
}

// GetUserLevelByID  get xxx by id
func GetUserLevelByID(c *gin.Context) interface{} {
	_id := c.Param("oid")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.GetUserLevelByID(id)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// ListUserLevel // list by page condition
func ListUserLevel(c *gin.Context) interface{} {
	var req = db.UserLevel{}
	err := c.ShouldBind(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	data, err := service.ListUserLevel(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "ok", Data: data}
}

// DeleteUserLevel Delete
func DeleteUserLevel(c *gin.Context) interface{} {
	var req = db.UserLevel{}
	err := c.ShouldBind(&req)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	err = service.DeleteUserLevel(req.ID)
	if err != nil {
		return Response{Code: -1, Msg: err.Error()}
	}
	return Response{Code: 0, Msg: "删除成功"}
}
