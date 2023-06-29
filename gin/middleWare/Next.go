package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/* 执行顺序:
中间件1 Next前
中间件2 Next前
中间件3 Next前
调用路由处理函数
中间件3 Next后
中间件2 Next后
中间件1 Next后
*/
func nextMiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件1 Next前")
		// 调用 Abort，终止执行后续的中间件
		c.Next()
		// 当前的函数不会停止
		fmt.Println("中间件1 Next后")
	}
}

func nextMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件2 Next前")
		c.Next()
		fmt.Println("中间件2 Next后")
	}
}

func nextMiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件3 Next前")
		c.Next()
		fmt.Println("中间件3 Next后")
	}
}

func main() {
	r := gin.New()

	r.Use(nextMiddleWare1(), nextMiddleWare2(), nextMiddleWare3())

	r.GET("/test", func(c *gin.Context) {
		fmt.Println("调用路由处理函数")
		// 页面接收
		c.JSON(200, gin.H{"request": "gin框架"})
	})

	r.Run(":8000")
}
