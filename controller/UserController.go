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
