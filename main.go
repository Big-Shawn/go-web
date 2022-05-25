package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// 获取post的传递参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		}
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

func RandomString(n int) string {

	var strings = []byte("asdfaqetryuioplamznbvASDFGHQWERTYUIOPLZMNCBV")

	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = strings[rand.Intn(len(strings))]
	}

	return string(result)
}
