package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if !token.Valid || err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		db := common.GetDB()

		user := &model.User{}

		db.Where("id = ?", claims.UserId).First(user)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})

			ctx.Abort()
			return
		}

		// 将用户的信息写进上下文中
		ctx.Set("user", user)
		ctx.Next()
	}

}
