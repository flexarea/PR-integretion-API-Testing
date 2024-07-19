package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flexarea/PR-integration-API-Testing/configs"
	"github.com/flexarea/PR-integration-API-Testing/internal/slack"
	"github.com/flexarea/PR-integration-API-Testing/internal/trello"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	//check path

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Use appropriate route for github action integration"))
}
func (app *Application) GitUpdate(w http.ResponseWriter, r *http.Request) {

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
	title := "test 2"
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
func (app *Application) Slack(w http.ResponseWriter, r *http.Request) {

	newMessage := r.URL.Query().Get("message")
	channelId := r.URL.Query().Get("channelID")
	env, err := configs.Load_config()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = slack.SendMessage(env.BOT_TOKEN, env.SLACK_MAIN_END_POINT, channelId, newMessage)

	if err != nil {
		app.errLog.Print("Error sending message to slack", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newMessage))
}

func (app *Application) Trello(w http.ResponseWriter, r *http.Request) {

	//get query parameters

	message := r.URL.Query().Get("message")

	list_id := "6671997dea31db576b213fce"

	//loading  .env
	env, err := configs.Load_config()

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err := trello.GettingCardsInList(env, list_id)

	if err != nil {
		app.errLog.Fatal(err)
	}

	cardID := trello.Retrieve(8, *result)

	targetListID := "666c33ed3ac4db04d453eacf"

	code, err := trello.MoveCardtoList(env, cardID, targetListID)

	if code == http.StatusOK && err == nil {

		http.Redirect(w, r, fmt.Sprintf("/slackMessage?channelID=%s&message=%s", env.CHANNELID, message), http.StatusSeeOther)

	} else {
		log.Print("Error in Accessing Trello API", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
}
