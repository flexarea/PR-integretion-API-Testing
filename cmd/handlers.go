package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
func (app *Application) dbRecord(w http.ResponseWriter, r *http.Request) {
	if app == nil || app.logs == nil {
		http.Error(w, "Server not initialized properly", http.StatusInternalServerError)
		return
	}

	//db model start here
	title := r.URL.Query().Get("title")
	branch := r.URL.Query().Get("branch")
	destinationBranch := r.URL.Query().Get("destinationBranch")
	pr_comment := r.URL.Query().Get("comment")
	channelID := r.URL.Query().Get("channelID")
	//db model ends here

	_, err := app.logs.Insert(title, branch, destinationBranch, pr_comment, channelID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.infoLog.Print("PR successfully Logged")
	w.Write([]byte("PR Successfully Logged"))
}
func (app *Application) Slack(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	branch := r.URL.Query().Get("branch")
	destinationBranch := r.URL.Query().Get("destinationBranch")
	pr_comment := r.URL.Query().Get("comment")
	channelID := r.URL.Query().Get("channelID")

	env, err := configs.Load_config()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = slack.SendMessage(env.BOT_TOKEN, env.SLACK_MAIN_END_POINT, channelID, pr_comment)

	if err != nil {
		app.errLog.Print("Error sending message to slack", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/dbRecord?title=%s&branch=%s&destinationBranch=%s&channelID=%s&comment=%s", title, branch, destinationBranch, channelID, pr_comment), http.StatusSeeOther)
}

func (app *Application) LogPr(w http.ResponseWriter, r *http.Request) {

	/*
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	*/

	//get query parameters
	listID := r.URL.Query().Get("title")
	branch := r.URL.Query().Get("branch")
	destinationBranch := r.URL.Query().Get("destinationBranch")
	pr_comment := r.URL.Query().Get("comment")
	channelID := r.URL.Query().Get("channelID")
	idShort, err := strconv.Atoi(r.URL.Query().Get("idShort"))

	if err != nil {
		app.errLog.Print(err)

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//loading  .env
	env, err := configs.Load_config()

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err := trello.GettingCardsInList(env, listID)

	if err != nil {
		app.errLog.Fatal(err)
	}

	cardID := trello.Retrieve(idShort, *result)

	errr := trello.MoveCardtoList(env, cardID, destinationBranch)

	if errr == nil {

		http.Redirect(w, r, fmt.Sprintf("/slackMessage?title=%s&branch=%s&destinationBranch=%s&channelID=%s&comment=%s", listID, branch, destinationBranch, channelID, pr_comment), http.StatusSeeOther)

	} else {
		app.errLog.Print("Error in Accessing Trello API", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
}
