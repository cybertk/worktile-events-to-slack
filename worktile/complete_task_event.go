package worktile

import (
	"fmt"
)

type CompleteTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	CreateBy  User   `json:"create_by"`
	Project   Project
}

// Interface Event
func (t *CompleteTaskEvent) Format() string {

	// Slack Markdown, see https://api.slack.com/docs/formatting
	return fmt.Sprintf("%s完成了_%s_, *%s*中的任务: %s",
		t.CreateBy.Name, t.Project.Name, t.EntryName, t.Name)
}
