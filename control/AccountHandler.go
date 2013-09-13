package control

import (
	//"encoding/base64"
	. "dao"
	"net/http"
	"service"
	"strconv"
)

var accountService = new(service.AccountService)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["register.html"].Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")

		account := new(Account)
		account.Name = name
		account.Password = password
		err := accountService.Register(account)
		if err != nil {
			Templates["register.html"].Execute(w, "注册失败")
		} else {
			Templates["login.html"].Execute(w, "注册成功")
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["login.html"].Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		account, err := accountService.Login(name, password)
		if err != nil {
			Templates["login.html"].Execute(w, "登录失败，该账户不存在或密码错误！")
		} else {
			//encodeStr := base64.StdEncoding.EncodeToString([]byte(name + password))
			//keyCookie := http.Cookie{Name: "key", Value: encodeStr}
			accountId := strconv.Itoa(account.Id)
			accountIdCookie := http.Cookie{Name: "AccountId", Value: accountId}
			//http.SetCookie(w, &keyCookie)
			http.SetCookie(w, &accountIdCookie)
			http.Redirect(w, r, "/list", http.StatusFound)
		}
	}
}
