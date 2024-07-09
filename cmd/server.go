package main

import (
	"log"
	"net/http"
	"os"
)

func Server() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/newGitActionUpdate", GitUpdate)
	mux.HandleFunc("/slackMessage", Slack)

	//server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:     ":" + port,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//start web server
	infoLog.Printf("starting server on port %s", port)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
