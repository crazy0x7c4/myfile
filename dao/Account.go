package dao

import (
	"strconv"
)

type Account struct {
	Id       int
	Name     string
	Password string
}

func (this *Account) Add(account *Account) (int, error) {
	id, err := Add("account", 0, account.Name, account.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *Account) GetOne(condition map[string]interface{}) (*Account, error) {
	data, err := GetOne("account", 3, condition)
	if err != nil {
		return nil, err
	}
	obj := this.parseData(data)
	return obj, nil
}

func (this *Account) parseData(data []string) *Account {
	p := new(Account)
	id, _ := strconv.Atoi(data[0])
	p.Id = id
	p.Name = data[1]
	p.Password = data[2]
	return p
}
