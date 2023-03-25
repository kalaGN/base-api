package main

import (
	btsConig "base-api/config"
	pkconfig "base-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}

func main() {
	pkconfig.InitConfig("env")
	r := gin.Default()
	r.POST("/index/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":    "pong",
			"status": 2006,
		})
	})
	r.Run(":" + pkconfig.Get("app.port")) // listen and serve on 0.0.0.0:8080
}
