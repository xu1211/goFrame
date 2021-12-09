package handler

import (
	"github.com/xu1211/goFrame/gorm/domain"
)

func NewTestHandler(do *domain.TestCRUD) *TestHandler {
	return &TestHandler{
		dbAgent: do,
	}
}

// 封装Handler服务
type TestHandler struct {
	dbAgent *domain.TestCRUD
}

func (h TestHandler) QueryUser(id uint) (domain.User, error) {
	user, err := h.dbAgent.QueryUserById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (h TestHandler) handler1() {
	h.dbAgent.Creat()
}

func (h TestHandler) handler2() {
	h.dbAgent.Update()
}

func (h TestHandler) handler3() {
	h.dbAgent.Delete()
}
