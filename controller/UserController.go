package controller

import (
	"crypto/md5"
	"fmt"
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
	passwd = fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

	model.User{}.Enter(ctx, telephone, passwd)
}
