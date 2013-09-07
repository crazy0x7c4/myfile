package main

import (
	. "control"
	"net/http"
)

func main() {
	http.HandleFunc("/register", safeHandler(RegisterHandler))
	http.HandleFunc("/login", safeHandler(LoginHandler))
	http.HandleFunc("/file", safeHandler(FileHandler))
	http.ListenAndServe(":8080", nil)
}

func safeHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); err != nil && ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		function(w, r)
	}
}
