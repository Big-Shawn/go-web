package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/fs"
	"path/filepath"
	"strings"
)

var DB *gorm.DB

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

func AutoMigrateTableWhenBoot(path string) {
	//fmt.Println(path)
	if err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			modelFile := info.Name()
			modelName := strings.SplitN(".", modelFile, 1)
			// 对文件内容进行正则匹配查看是否存在模型定义
			// 读取文件内容

			fmt.Println(modelName)
		}
		return err
	}); err != nil {
		panic(err)
	}
}
