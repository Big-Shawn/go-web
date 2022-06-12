package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/common"
	"os"
)

func main() {
	//db := common.GetDB()
	//defer db.Close()
	//
	fmt.Println()

	r := gin.Default()
	r = CollectRouter(r)
	if currPath, err := os.Getwd(); err == nil {
		common.AutoMigrateTableWhenBoot(currPath + "/model")
	}
	// 项目启动时自动迁移编写的模型
	// 1. 遍历获取目录下面的所有的 model 文件，然后逐个进行注册
	// 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run())
}
