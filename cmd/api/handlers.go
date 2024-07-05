package api

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//check path

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Use appropriate route for github action integration"))
}
func GitUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
