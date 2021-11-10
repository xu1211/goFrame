package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
*gin.Context.Param()		获取url
*gin.Context.Query()		获取param
*gin.Context.PostForm()		获取表单param
*gin.Context.FormFile()		获取multipart文件
*gin.Context.MultipartForm()	获取多个multipart表单参数


 */
func main() {
	r := gin.Default()

	// 获取http的URL
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		// 截取斜杠/
		action = strings.Trim(action, "/")

		c.String(http.StatusOK, name+" is "+action)
	})

	// 获取http的URL param参数
	r.GET("/user", func(c *gin.Context) {
		//http://localhost:8000/user?name=cosmo
		defName := c.DefaultQuery("name", "默认名") // DefaultQuery 没有name参数时会用默认值
		fmt.Println(defName)
		name := c.Query("name") // Query 没有name参数时会用空字符串
		fmt.Println(name)
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// 获取http（application/x-www-form-urlencoded，multipart/form-data）表单的参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post") // 不存在时使用默认值
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	// 获取http multipart/form-data格式的文件
	r.MaxMultipartMemory = 8 << 20 //限制文件上传最大尺寸
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		c.SaveUploadedFile(file, file.Filename) //保存file
		c.String(http.StatusOK, file.Filename)
	})

	//默认为监听8080端口
	r.Run(":8000")
}
