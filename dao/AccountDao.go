package dao

import (
	. "entity"
	"strconv"
)

func AddAccount(name string, password string) {
	sqlStr := "INSERT account SET name=?, password=?"
	Execute(sqlStr, name, password)
}

func GetAccount(name string, password string) (account Account) {
	sqlStr := "SELECT * FROM account WHERE name=? AND password=?"
	data := QueryRow(3, sqlStr, name, password)
	account = setData(data)
	return
}

func GetAccountByName(name string) (account Account) {
	sqlStr := "SELECT * FROM account WHERE name=?"
	data := QueryRow(3, sqlStr, name)
	account = setData(data)
	return
}

func GetAccountById(id uint) (account Account) {
	sqlStr := "SELECT * FROM account WHERE id=?"
	data := QueryRow(3, sqlStr, id)
	account = setData(data)
	return
}

func setData(data []string) (account Account) {
	uinteger32, _ := strconv.ParseUint(data[0], 0, 32)
	account.Id = uint(uinteger32)
	account.Name = data[1]
	account.Password = data[2]
	return
}
