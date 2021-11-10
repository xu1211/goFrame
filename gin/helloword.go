package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
1. 创建路由
2. 路由.GET/POST/PUT/DELETE(http路径, 处理函数(*gin.Context) )
	gin.Context 封装了request和response
3. 启动路由，监听端口
*/
func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	r.GET("/get", func(c *gin.Context) {
		// *gin.Context.String() 输出response
		c.String(http.StatusOK, "hello World!")
	})
	r.POST("/post")
	r.PUT("/put")

	// 3.监听端口，Run("里面不指定端口号默认为8080")
	fmt.Println("Please Visit -  http://localhost:8000")
	r.Run(":8000")

}
