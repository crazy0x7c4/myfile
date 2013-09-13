package service

import (
	. "dao"
	"errors"
	"log"
)

var accountDao = new(Account)

type AccountService struct {
}

func (this *AccountService) Register(account *Account) error {
	condition := make(map[string]interface{})
	condition["a_name"] = account.Name
	oldAccount, err := accountDao.GetOne(condition)

	if oldAccount == nil && err != nil {
		_, err := accountDao.Add(account)
		if err == nil {
			return nil
		}
		return err
	} else {
		return errors.New("该用户名已经存在")
	}
}

func (this *AccountService) Login(name, password string) (*Account, error) {
	condition := make(map[string]interface{})
	condition["a_name"] = name
	condition["a_password"] = password
	account, err := accountDao.GetOne(condition)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return account, nil
}
