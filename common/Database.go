package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init Database collection
func init() {
	// 必须初始化对全局变量赋值？ 但是不是已经在函数里面赋值过了吗?
	DB = InitDB()

}

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "go-web"
	username := "root"
	password := "886600"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	//db, err := gorm.Open(driverName, args)
	DB, err := gorm.Open(mysql.Open(args), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})

	if err != nil {
		panic(gin.H{
			"msg": "failed Connect to database",
			"err": err,
		})
	}

	return DB
}

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}
	return InitDB()
}

// 在项目启动时迁移项目中存在的表
func AutoMigrateTableWhenBoot(model interface{}) {
	if !DB.Migrator().HasTable(model) {
		if err := DB.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}
