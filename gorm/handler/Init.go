package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/xu1211/goFrame/gorm/domain"
)

func Init(db *gorm.DB) {
	// 获取test数据库操作对象
	testCRUD := &domain.TestCRUD{
		DB: db,
	}

	//初始化TestHandler实例
	testHandler := NewTestHandler(testCRUD)

	// 执行服务
	testHandler.QueryUser(1)
	testHandler.handler1()
	testHandler.handler2()
	testHandler.handler3()
}
