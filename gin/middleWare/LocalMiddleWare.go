package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.New()

	// 给路由组 添加中间件, 可以添加多个中间件
	v1 := r.Group("/v1", Validate1())
	{
		// 给路由 添加中间件, 可以添加多个中间件
		v1.GET("/login", login, Validate3())
		v1.GET("submit", submit)
	}

	// 创建路由组/v2 ，处理POST请求
	v2 := r.Group("/v2", Validate2)
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
}

// 中间件
func Validate1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("v1局部中间件")
	}
}

func Validate2(c *gin.Context) {
	fmt.Println("v2局部中间件")
}

func Validate3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("/v1/login局部中间件")
	}
}

// Handler
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
