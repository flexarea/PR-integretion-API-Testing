# **Integrated Automation System for GitHub PR Actions, Trello Cards, and Slack**
-------------------------

```
API-Server-Integration/
├── cmd/
│   ├── api/
│   │   └── server.go          # HTTP server setup and route handling
│   │   └── handlers.go        # Entry point of the application
│   │ 
│   ├── main.go            # Entry point of the application
│   └── config.go          # configuration and environment variable handling 
├── internal/
│   ├── trello/
│   │   ├── trello.go          # Functions for interacting with Trello API
│   │   ├── trello_test.go      
│   ├── slack/
│   │   └── slack.go           # Functions for interacting with Slack API
│   │   ├── slack_test.go 
├── pkg/
│   ├── models/
│   │   └── models.go          # Data models and types used in the application
├── .gitignore                 # Git ignore file (ignore .env)
├── .env               # Environment variables file ( API keys, main endpoints, Bot and api token)
├── go.mod                     # Go module file
├── go.sum                     # Go module dependencies file
└── README.md                  # Project documentation
```
### To run current state


```bash
go run cmd/*.go
```

Or can run it directly in cmd but will have to load environment with relative path

```go
//config.go
func Load_config() (Configs, error) {
	err := godotenv.Load("../.env")
    /* rest of the code below */
```

```bash
go run *.go
```
