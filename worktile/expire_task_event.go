package worktile

type ExpireTaskEvent struct {
	Name       string
	EntryName  string `json:"entry_name"`
	ExpireDate int64  `json:"expire_date"`
	CreateBy   User   `json:"create_by"`
	Project    Project
}
