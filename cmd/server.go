package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flexarea/PR-integration-API-Testing/pkg/models"
	_ "github.com/lib/pq"
)

var app *Application

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

	env, err := Load_config()

	if err != nil {
		errorLog.Fatalf("Failed to Load .env: %v", err)
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require", env.DB_USERNAME, env.DB_PASSWORD, env.DB_HOST, env.DB_DATABASE)

	db, err := OpenDB(connStr)

	if err != nil {
		errorLog.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	//populate application struct (which populate models.LogModels struct in models package)
	app = &Application{
		logs: &models.LogsModel{DB: db},
	}

	//start web server
	infoLog.Printf("starting server on port %s", port)
	err = server.ListenAndServe()
	errorLog.Fatal(err)

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
