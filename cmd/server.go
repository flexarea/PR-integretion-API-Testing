package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	env, err := Load_config()

	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", env.DB_USERNAME, env.DB_PASSWORD, env.DB_HOST, env.DB_DATABASE)

	db, err := OpenDB(connStr)

	defer db.Close()

	var version string

	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}
	fmt.Printf("version=%s\n", version)
}

func OpenDB(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
