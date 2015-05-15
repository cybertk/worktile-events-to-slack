package worktile

import (
	"fmt"
)

type AssignTaskEvent struct {
	Name       string
	EntryName  string `json:"entry_name"`
	CreateBy   User   `json:"create_by"`
	Project    Project
	AssignUser User `json:"assign_user"`
}

// Interface Event
func (t *AssignTaskEvent) Format() string {
	// Slack Markdown, see https://api.slack.com/docs/formatting
	return fmt.Sprintf("%s在_%s_, *%s*中将任务: %s分配给了%s",
		t.CreateBy.Name, t.Project.Name, t.EntryName, t.Name, t.AssignUser.Name)
}
