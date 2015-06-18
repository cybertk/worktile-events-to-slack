package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cybertk/worktile-events-to-slack/worktile"
)

// Slack message attachments, see https://api.slack.com/docs/attachments
type SlackMessage struct {
	Attachments [1]SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color      string   `json:"color"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	MarkdownIn []string `json:"mrkdwn_in"`
}

// export DEBUG=1
func isDebug() bool {
	return len(os.Getenv("DEBUG")) != 0
}

func format(event worktile.Event) SlackAttachment {

	var attachment SlackAttachment
	switch e := event.(type) {
	case *worktile.CreateTaskEvent:
		attachment = SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      fmt.Sprintf("%s created task *%s* in _%s_", e.CreateBy.Name, e.Name, e.EntryName),
		}
	case *worktile.CompleteTaskEvent:
		attachment = SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      fmt.Sprintf("%s finished task *%s* in _%s_", e.CreateBy.Name, e.Name, e.EntryName),
		}
	case *worktile.ExpireTaskEvent:
		dueDate := time.Unix(e.ExpireDate/1000, 0).Format("Jan _2")
		attachment = SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      fmt.Sprintf("%s set task *%s* in _%s_ due time to %s", e.CreateBy.Name, e.Name, e.EntryName, dueDate),
		}
	case *worktile.AssignTaskEvent:
		attachment = SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      fmt.Sprintf("%s assigned task *%s* in _%s_ to %s", e.CreateBy.Name, e.Name, e.EntryName, e.AssignUser.Name),
		}
	case *worktile.CommentTaskEvent:
		attachment = SlackAttachment{
			Color:     "#36a64f",
			Title:     e.Project.Name,
			TitleLink: "https://worktile.com/project/" + e.Project.Id,
			Text:      fmt.Sprintf("%s add comments to task *%s* in _%s_\n%s", e.Comment.CreateBy.Name, e.Name, e.EntryName, e.Comment.Message),
		}
	default:
		attachment = SlackAttachment{}
	}

	// Enable markdown in text field, see https://api.slack.com/docs/formatting
	attachment.MarkdownIn = []string{"text"}
	return attachment
}

func sendToSlack(webhookUrl string, event worktile.Event) error {
	slackMessage := SlackMessage{
		Attachments: [...]SlackAttachment{format(event)},
	}

	payload, err := json.Marshal(slackMessage)
	if err != nil {
		return err
	}
	payloadStr := string(payload)

	if isDebug() {
		fmt.Println(payloadStr)
	}

	// Slack incoming webhooks API, see https://api.slack.com/incoming-webhooks
	_, err = http.PostForm(webhookUrl, url.Values{"payload": {payloadStr}})
	return err
}

func handler(w http.ResponseWriter, r *http.Request, send func(string, worktile.Event) error) {
	var notification worktile.Notification

	var slackUrl = r.FormValue("slack_url")

	if _, err := url.Parse(slackUrl); err != nil {
		fmt.Println("No body received")
		w.WriteHeader(400)
		return
	}

	if r.Body == nil {
		fmt.Println("No body received")
		w.WriteHeader(400)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		fmt.Println("Decode error")
		w.WriteHeader(500)
		return
	}

	if isDebug() {
		fmt.Println(string(notification.Data))
	}

	if err := send(slackUrl, notification.Event()); err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		// Fallback to default port 3000
		port = "3000"
	}

	fmt.Println("Port: " + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, sendToSlack)
	})
	http.ListenAndServe(":"+port, nil)
}
