package main

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine {

	r.POST("/sign", controller.Register)
	r.GET("/", controller.Index)

	return r
}
