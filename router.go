package main

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
	"go-web/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.AuthMiddleware())

	r.POST("/enter", controller.Enter)
	r.POST("/sign", controller.Register)
	r.GET("/", controller.Index)
	r.GET("/info", controller.Info)

	return r
}
