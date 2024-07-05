package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Server Integration Home"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	//start web server
	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
