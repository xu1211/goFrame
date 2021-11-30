package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 自定义中间件
func myMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")

		// 修改gin.Context
		c.Set("request", "中间件")
		status := c.Writer.Status()

		fmt.Println("中间件执行完毕", status)

		t2 := time.Since(t)
		fmt.Println("中间件执行time:", t2)
	}
}
func main() {
	// 1.创建路由
	r := gin.New()

	// 注册全局中间件
	r.Use(myMiddleWare())

	// 添加路由
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取参
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 返回
			c.JSON(200, gin.H{"request": req})
		})

	}

	// 启动监听
	r.Run(":8000")
}
