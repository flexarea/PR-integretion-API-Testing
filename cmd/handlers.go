package main

import (
	"log"
	"net/http"

	"github.com/flexarea/PR-integration-API-Testing/internal/slack"
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
		return
	}

	if app == nil || app.logs == nil {
		http.Error(w, "Server not initialized properly", http.StatusInternalServerError)
		return
	}

	//testing to send data to database
	title := "testing"
	branch := "development"
	destinationBranch := "production"
	pr_comment := "sending data to neon db"
	slackchannel := "C06KPMXQS4U"

	_, err := app.logs.Insert(title, branch, destinationBranch, pr_comment, slackchannel)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

	w.WriteHeader(http.StatusOK)

}
func Slack(w http.ResponseWriter, r *http.Request) {

	newMessage := r.URL.Query().Get("message")
	channelId := r.URL.Query().Get("channelID")
	env, err := Load_config()

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = slack.SendMessage(env.BOT_TOKEN, env.SLACK_MAIN_END_POINT, channelId, newMessage)

	if err != nil {
		log.Print("Error sending message to slack", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent to slack"))
}
