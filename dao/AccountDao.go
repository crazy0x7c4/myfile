package dao

import (
	. "entity"
	"strconv"
)

type AccountDao struct {
}

func (accountDao *AccountDao) AddAccount(name string, password string) (int, error) {
	sqlStr := "INSERT account SET name=?, password=?"
	id, err := Execute(sqlStr, name, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (accountDao *AccountDao) GetAccount(name string, password string) (*Account, error) {
	sqlStr := "SELECT * FROM account WHERE name=? AND password=?"
	data, err := QueryRow(3, sqlStr, name, password)
	if err != nil {
		return nil, err
	}
	account := accountDao.setData(data)
	return account, nil
}

func (accountDao *AccountDao) GetAccountByName(name string) (*Account, error) {
	sqlStr := "SELECT * FROM account WHERE name=?"
	data, err := QueryRow(3, sqlStr, name)
	if err != nil {
		return nil, err
	}
	account := accountDao.setData(data)
	return account, nil
}

func (accountDao *AccountDao) GetAccountById(id uint) (*Account, error) {
	sqlStr := "SELECT * FROM account WHERE id=?"
	data, err := QueryRow(3, sqlStr, id)
	if err != nil {
		return nil, err
	}
	account := accountDao.setData(data)
	return account, nil
}

func (accountDao *AccountDao) setData(data []string) *Account {
	account := new(Account)
	integer64, _ := strconv.ParseInt(data[0], 0, 32)
	account.Id = int(integer64)
	account.Name = data[1]
	account.Password = data[2]
	return account
}
