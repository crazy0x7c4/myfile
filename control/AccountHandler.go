package control

import (
	"io"
	"net/http"
	"service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["register.html"].Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		err := service.Register(name, password)
		if err != nil {
			Templates["register.html"].Execute(w, err.Error())
		} else {
			Templates["register.html"].Execute(w, "注册成功")
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "login...")
}
