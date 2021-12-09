package main

import (
	"github.com/xu1211/goFrame/gorm/handler"
	"github.com/xu1211/goFrame/gorm/repository"
)

func main() {
	// 初始化 数据库连接
	repository.Init()
	// 获取数据库连接
	db := repository.GetDB()

	// 初始化 handler服务
	handler.Init(db)

	// 释放 数据库连接
	defer func() {
		repository.Release()
	}()

}
