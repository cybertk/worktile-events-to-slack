package worktile

type CreateTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	CreateBy  User   `json:"create_by"`
	Project   Project
}
