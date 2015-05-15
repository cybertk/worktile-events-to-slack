package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/cybertk/worktile-events-to-slack/worktile"
)

// Slack incoming webhooks API, see https://api.slack.com/incoming-webhooks
type SlackMessage struct {
	Text       string `json:"text"`
	IsMarkdown bool   `json:"mrkdwn"`
}

func sendToSlack(webhookUrl string, message string) (*http.Response, error) {
	slackMessage := SlackMessage{Text: message, IsMarkdown: true}

	payload, err := json.Marshal(slackMessage)
	if err != nil {
		return nil, err
	}
	payloadStr := string(payload)
	fmt.Println(payloadStr)
	return http.PostForm(webhookUrl, url.Values{"payload": {payloadStr}})
}

func handler(w http.ResponseWriter, r *http.Request, slackUrl string) {
	var message string
	var notification worktile.Notification

	err := json.NewDecoder(r.Body).Decode(&notification)

	if debug := os.Getenv("DEBUG"); len(debug) != 0 {
		fmt.Println(string(notification.Data))
	}

	if err != nil {
		message = fmt.Sprintln("Cannot decode notification: ", err)
	} else {

		if event := notification.Event(); event == nil {
			message = "Cannot decode event"
		} else {
			message = event.Format()
		}
	}

	if _, err = sendToSlack(slackUrl, message); err != nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func main() {

	incomingWebhookUrl := os.Getenv("SLACK_URL")
	port := os.Getenv("PORT")

	if len(port) == 0 {
		// Fallback to default port 3000
		port = "3000"
	}
	if len(incomingWebhookUrl) == 0 {
		fmt.Println("environment variables SLACK_URL is not set correctly")
		return
	}

	fmt.Println("Slack Incoming Webhook URL: " + incomingWebhookUrl)
	fmt.Println("Port: " + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, incomingWebhookUrl)
	})
	http.ListenAndServe(":"+port, nil)
}
