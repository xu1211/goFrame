package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type Login struct {
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 将request的body中的数据，自动按照json格式解析到结构体
		var loginJson Login
		if err := c.ShouldBindJSON(&loginJson); err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H封装了生成json数据的工具
			return
		}

		// 判断用户名密码是否正确
		if loginJson.User != "root" || loginJson.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8000")
}
