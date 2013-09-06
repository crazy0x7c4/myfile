package service

import (
	"dao"
	. "entity"
	"errors"
)

func Register(name string, password string) error {
	account := dao.GetAccountByName(name)
	if account.Id == 0 && account.Name == "" {
		dao.AddAccount(name, password)
		return nil
	}
	return errors.New("用户名已被注册！")
}

func Login(name string, password string) (account Account) {
	account = dao.GetAccount(name, password)
	return
}
