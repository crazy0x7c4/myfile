package test

import (
	"log"
	"service"
	"testing"
)

func TestLogin(t *testing.T) {
	accountService := new(service.AccountService)
	account, err := accountService.Login("crazy", "11111")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(account)
}
