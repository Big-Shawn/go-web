package model

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-web/common"
	"go-web/util"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func init() {
	// 根据import 的性质，一般会先完成依赖包的init 然后再进行其他包的引入
	common.AutoMigrateTableWhenBoot(&User{})
}

// User 这里创建了一个结构体，需要使用迁移才能自动化生成表
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"size:255;not null"`
}

func (user *User) isExists(db *gorm.DB, field, value string) bool {

	db.Where(field+" = ?", value).First(&user)

	if user.ID != 0 {
		return true
	}
	return false

}

func (user *User) Register(c *gin.Context) {
	// 获取post的传递参数

	if len(user.Telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不得少于六位"})
		return
	} else {
		// 这里不能直接使用string 进行类型转换，而只能使用这种方式
		// 将结果转换为16进制
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	}

	if len(user.Name) == 0 {
		user.Name = util.RandomString(10)
	}

	if user.isExists(common.DB, "Telephone", user.Telephone) {
		c.JSON(200, gin.H{
			"code": 422,
			"msg":  "该手机号已注册",
		})
		return
	}

	db := common.GetDB()

	db.Create(user)

	c.JSON(200, gin.H{
		"msg":  "注册成功",
		"user": user,
	})

	return
}

func (user User) Enter(ctx gin.Context, telephone, password string) {
	db := common.GetDB()

	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 && user.Password == password {
		// 返回用户信息和token值

		ttl := jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1)))
		method := "RS256"
		playload := user
		token, err := util.SetToken(ttl, method, playload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "服务器错误")
			return
		}
		ctx.JSON(200, token)

	}

	ctx.JSON(http.StatusUnprocessableEntity, "用户名或密码错误")
	return

}
