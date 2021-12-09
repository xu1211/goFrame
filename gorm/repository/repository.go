package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	mysqlDsn = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"
	Host     = "127.0.0.1:3306"
	UserName = "root"
	Password = "molotest123"
	Database = "test"
)

var (
	db *gorm.DB
	err error
)

// 获取db
func GetDB() *gorm.DB {
	return db
}

// 初始化 数据库连接db
func Init() {
	dsn := fmt.Sprintf(mysqlDsn, UserName, Password, Host, Database)
	// 连接数据库
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)       // 日志打印模式
	db.SingularTable(true) // 表名映射: gorm会在创建表的时候去掉”s“的后缀
	//连接池
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
}

// 释放 数据库连接db
func Release() {
	err := db.Close()
	if err != nil {
		fmt.Errorf("repository close err: %s", err.Error())
	}
}
