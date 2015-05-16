package worktile

type AssignTaskEvent struct {
	Name       string
	EntryName  string `json:"entry_name"`
	CreateBy   User   `json:"create_by"`
	Project    Project
	AssignUser User `json:"assign_user"`
}
