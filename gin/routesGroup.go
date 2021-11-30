package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 创建路由组/v1 ，处理GET请求
	v1 := r.Group("/v1")
	// v1路由组添加路由。 {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}

	// 创建路由组/v2 ，处理POST请求
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
}

// 处理类
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
