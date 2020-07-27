package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}
