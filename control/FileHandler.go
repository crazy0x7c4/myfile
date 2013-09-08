package control

import (
	"io"
	"log"
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	baseAuth := r.Header.Get("Authorization")
	log.Println(baseAuth)
	if baseAuth == "" {
		io.WriteString(w, "NO Auth...")
	} else {
		io.WriteString(w, "my files...")
	}
}
