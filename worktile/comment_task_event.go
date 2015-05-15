package worktile

import (
	"fmt"
)

type Comment struct {
	Message  string
	CreateBy User `json:"create_by"`
}

type CommentTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	CreateBy  User   `json:"create_by"`
	Comment   Comment
	Project   Project
}

// Interface Event
func (t *CommentTaskEvent) Format() string {
	// Slack Markdown, see https://api.slack.com/docs/formatting
	return fmt.Sprintf("%s在_%s_, *%s*中的任务: %s中添加了%s",
		t.Comment.CreateBy.Name, t.Project.Name, t.EntryName, t.Name, t.Comment.Message)
}
