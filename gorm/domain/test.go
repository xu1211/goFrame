package domain

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)
//映射表结构的struct
type User struct {
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	gorm.Model
}
// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}