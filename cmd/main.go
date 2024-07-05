package main

import (
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Server Integration Home"))
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	//server configuration
	server := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//start web server
	infoLog.Println("starting server on :4000")
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
