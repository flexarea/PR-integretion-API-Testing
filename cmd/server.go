package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/flexarea/PR-integration-API-Testing/configs"
	"github.com/flexarea/PR-integration-API-Testing/pkg/models"
	_ "github.com/lib/pq"
)

func (app *Application) Server() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/dbRecord", app.dbRecord)
	mux.HandleFunc("/slackMessage", app.Slack)
	mux.HandleFunc("/logPr", app.LogPr)
	//server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:     ":" + port,
		ErrorLog: app.errLog,
		Handler:  mux,
	}

	env, err := configs.Load_config()

	if err != nil {
		app.errLog.Fatalf("Failed to Load .env: %v", err)
	}

	//connection string formatting (development)
	//connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require", env.DB_USERNAME, env.DB_PASSWORD, env.DB_HOST, env.DB_DATABASE)
	//connection string formatting (deployment)
	connStr := env.DB_URL
	db, err := OpenDB(connStr)

	if err != nil {
		app.errLog.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	//populate application struct (which populate models.LogModels struct in models package)

	app.logs = &models.LogsModel{DB: db}

	//start web server
	app.infoLog.Printf("starting server on port %s", port)
	err = server.ListenAndServe()
	app.errLog.Fatal(err)

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
