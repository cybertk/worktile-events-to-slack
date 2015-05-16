package worktile

type CompleteTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	CreateBy  User   `json:"create_by"`
	Project   Project
}
