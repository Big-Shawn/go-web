package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//db := common.GetDB()
	//defer db.Close()
	//

	r := gin.Default()
	r = CollectRouter(r)
	// 项目启动时自动迁移编写的模型
	// 1. 遍历获取目录下面的所有的 model 文件，然后逐个进行注册
	// 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run())
}
