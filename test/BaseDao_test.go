package test

import (
	"dao"
	"log"
	"testing"
)

func TestExecute(t *testing.T) {
	sqlStr := "INSERT account SET name=?, password=?"
	id := dao.Execute(sqlStr, "user1", "12345")
	sqlStr = "DELETE FROM account WHERE id=?"
	dao.Execute(sqlStr, id)
}

func TestQuery(t *testing.T) {
	data := dao.Query("SELECT * FROM account")
	for _, v := range data {
		log.Println(v)
	}
}

func TestQueryRow(t *testing.T) {
	data := dao.QueryRow(3, "SELECT * FROM account WHERE NAME=? AND PASSWORD=?", "userX", "12345")
	log.Println(data)
}
