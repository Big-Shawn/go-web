package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// User 这里创建了一个结构体，好像也创建了一个数据表
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	r := gin.Default()
	r.POST("/sign", func(c *gin.Context) {
		// 获取post的传递参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不得少于六位"})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)
		}
		// 这里直接是显示到服务器端上了
		log.Print(name, telephone, password)

		c.JSON(200, gin.H{"msg": "注册成功"})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

// RandomString  生成一个随机的字符串
func RandomString(n int) string {

	var strings = []byte("asdfaqetryuioplamznbvASDFGHQWERTYUIOPLZMNCBV")

	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = strings[rand.Intn(len(strings))]
	}

	return string(result)
}

func isExists(db *gorm.DB, field, value string) bool {
	var res User
	db.Where(field+" = ? ", value).First(&res)

	if res.ID != 0 {
		return true
	}
	return false

}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := 3306
	database := "go-web"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed Connect to database")
	}

	return db

}
