package main

import (
	"log"
	"net/http"

	"github.com/flexarea/PR-integration-API-Testing/cmd/internal/slack"
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
func Slack(w http.ResponseWriter, r *http.Request) {

	newMessage := r.URL.Query().Get("message")
	channelId := r.URL.Query().Get("channelID")
	env, err := Load_config()

	if err != nil {
		log.Fatal(err)
	}

	slack.SendMessage(env.BOT_TOKEN, env.SLACK_MAIN_END_POINT, channelId, newMessage)
}
