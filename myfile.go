package main

import (
	. "control"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", safeHandler(ListHandler))
	http.HandleFunc("/register", safeHandler(RegisterHandler))
	http.HandleFunc("/login", safeHandler(LoginHandler))
	http.HandleFunc("/list", safeHandler(ListHandler))
	http.HandleFunc("/create", safeHandler(CreateHandler))
	http.HandleFunc("/delfolder", safeHandler(DelFolderHandler))
	http.HandleFunc("/upload", safeHandler(UploadHandler))
	http.HandleFunc("/download", safeHandler(DownloadHandler))
	http.HandleFunc("/delfile", safeHandler(DelFileHandler))
	http.ListenAndServe(":8080", nil)
}

func safeHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); err != nil && ok {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		if r.RequestURI == "/login" || r.RequestURI == "/register" {
			function(w, r)
		} else {
			_, err := r.Cookie("AccountId")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
			} else {
				function(w, r)
			}
		}
	}
}
