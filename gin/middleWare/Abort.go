package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func abortMiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("调用中间件1")
		// 调用 Abort，终止执行后续的中间件
		c.Abort()
		// 当前的函数不会停止
		fmt.Println("Abort中间件1")
	}
}
func abortMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("调用中间件2")
		c.Abort()
		fmt.Println("Abort中间件2")
	}
}
func main() {
	r := gin.New()

	r.Use(abortMiddleWare1())
	r.Use(abortMiddleWare2())

	r.GET("/test", func(c *gin.Context) {
		fmt.Println("调用路由处理函数")
		// 页面接收
		c.JSON(200, gin.H{"request": "gin框架"})
	})

	r.Run(":8000")
}
