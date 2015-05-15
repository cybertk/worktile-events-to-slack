package worktile

import (
	"fmt"
)

type CreateTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	CreateBy  User   `json:"create_by"`
	Project   Project
}

// Interface Event
func (t *CreateTaskEvent) Format() string {
	// Slack Markdown, see https://api.slack.com/docs/formatting
	return fmt.Sprintf("%s在_%s_, *%s*中新建了任务: %s", t.CreateBy.Name, t.Project.Name, t.EntryName, t.Name)
}
