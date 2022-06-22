package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/model"
)

func Register(ctx *gin.Context) {
	var user model.User

	user = model.User{
		Name:      ctx.PostForm("name"),
		Telephone: ctx.PostForm("telephone"),
		Password:  ctx.PostForm("password"),
	}

	user.Register(ctx)
}

func Enter(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	passwd := ctx.PostForm("password")

	model.User{}.Enter(ctx, telephone, passwd)
}

func Info(ctx *gin.Context) {

	model.User{}.Info(ctx)
}
