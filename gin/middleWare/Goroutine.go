package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		log.Println("原始 c *gin.Context :" + c.Request.URL.Path)
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("协程-拷贝 c *gin.Context :" + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("原始 c *gin.Context :" + c.Request.URL.Path)
		go func() {
			time.Sleep(5 * time.Second)
			/* 不建议使用原始对象
			可能会导致并发访问和修改上下文对象，从而引发竞态条件和数据不一致等问题。
			*/
			log.Println("协程-原始 c *gin.Context :" + c.Request.URL.Path)
		}()

	})
	r.Run(":8000")
}
