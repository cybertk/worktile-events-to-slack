package worktile

import (
	"fmt"
	"time"
)

type ExpireTaskEvent struct {
	Name       string
	EntryName  string `json:"entry_name"`
	ExpireDate int64  `json:"expire_date"`
	CreateBy   User   `json:"create_by"`
	Project    Project
}

// Interface Event
func (t *ExpireTaskEvent) Format() string {

	date := time.Unix(t.ExpireDate/1000, 0).Format("Jan _2")
	// Slack Markdown, see https://api.slack.com/docs/formatting
	return fmt.Sprintf("%s在_%s_, *%s*中给任务: %s设定截止日期%s",
		t.CreateBy.Name, t.Project.Name, t.EntryName, t.Name, date)
}
