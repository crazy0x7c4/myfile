package test

import (
	"dao"
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	dao.AddAccount("user", "12345")
}

func TestGetAccount(t *testing.T) {
	account := dao.GetAccount("userxxxx", "12345")
	log.Println(account)
}
