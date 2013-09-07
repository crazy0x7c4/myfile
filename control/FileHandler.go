package control

import (
	"io"
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "my files...")
}
