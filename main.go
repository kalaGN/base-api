package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/index/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":    "pong",
			"status": 2006,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
