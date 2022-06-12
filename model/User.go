package model

import (
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/util"
	"gorm.io/gorm"
	"net/http"
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

	db.AutoMigrate(user)
	db.Create(user)

	c.JSON(200, gin.H{
		"msg":  "注册成功",
		"user": user,
	})

	return
}
