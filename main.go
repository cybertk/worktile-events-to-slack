package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/cybertk/worktile-events-to-slack/worktile"
)

// Slack message attachments, see https://api.slack.com/docs/attachments
type SlackMessage struct {
	Attachments [1]SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color     string `json:"color"`
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
}

// export DEBUG=1
func isDebug() bool {
	return len(os.Getenv("DEBUG")) != 0
}

func format(event worktile.Event) SlackAttachment {

	switch e := event.(type) {
	case *worktile.CreateTaskEvent:
		return SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      e.Format(),
		}
	case *worktile.CompleteTaskEvent:
		return SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      e.Format(),
		}
	case *worktile.ExpireTaskEvent:
		return SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      e.Format(),
		}
	case *worktile.AssignTaskEvent:
		return SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      e.Format(),
		}
	case *worktile.CommentTaskEvent:
		return SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      e.Format(),
		}
	default:
		return SlackAttachment{}
	}
}

func sendToSlack(webhookUrl string, event worktile.Event) (*http.Response, error) {
	slackMessage := SlackMessage{
		Attachments: [...]SlackAttachment{format(event)},
	}

	payload, err := json.Marshal(slackMessage)
	if err != nil {
		return nil, err
	}
	payloadStr := string(payload)

	if isDebug() {
		fmt.Println(payloadStr)
	}

	// Slack incoming webhooks API, see https://api.slack.com/incoming-webhooks
	return http.PostForm(webhookUrl, url.Values{"payload": {payloadStr}})
}

func handler(w http.ResponseWriter, r *http.Request, slackUrl string) {
	var notification worktile.Notification

	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		fmt.Println("Decode error")
		w.WriteHeader(500)
		return
	}

	if isDebug() {
		fmt.Println(string(notification.Data))
	}

	if _, err := sendToSlack(slackUrl, notification.Event()); err != nil {
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
