package domain

import "github.com/jinzhu/gorm"

// 封装test数据库的CRUD方法
type TestCRUD struct {
	DB *gorm.DB
}

func (db *TestCRUD) QueryUserById(id uint) (User, error) {
	var (
		user User
	)
	user.ID = id
	err := db.DB.Find(&user).Error
	return user, err
}
func (db TestCRUD) Creat() {

}
func (db TestCRUD) Update() {

}
func (db TestCRUD) Delete() {

}
