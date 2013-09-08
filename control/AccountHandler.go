package control

import (
	"net/http"
	"service"
)

var accountService = new(service.AccountService)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["register.html"].Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		err := accountService.Register(name, password)
		if err != nil {
			Templates["register.html"].Execute(w, "注册失败")
		} else {
			Templates["register.html"].Execute(w, "注册成功")
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["login.html"].Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		_, err := accountService.Login(name, password)
		if err != nil {
			Templates["login.html"].Execute(w, err.Error())
		} else {
			http.SetCookie(http.Cookie{})
			http.Redirect(w, r, "/file", http.StatusFound)
		}
	}
}
