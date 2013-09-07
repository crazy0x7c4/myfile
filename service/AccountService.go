package service

import (
	"dao"
	. "entity"
	"errors"
	"log"
)

var accountDao = new(dao.AccountDao)

type AccountService struct {
}

func (accountService *AccountService) Register(name string, password string) error {
	account, err := accountDao.GetAccountByName(name)
	if account == nil && err != nil {
		_, err := accountDao.AddAccount(name, password)
		if err == nil {
			return nil
		}
		return err
	} else {
		return errors.New("账户已经存在")
	}
}

func (accountService *AccountService) Login(name string, password string) (*Account, error) {
	account, err := accountDao.GetAccount(name, password)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return account, nil
}
