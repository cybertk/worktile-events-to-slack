package worktile

type Comment struct {
	Message  string
	CreateBy User `json:"create_by"`
}

type CommentTaskEvent struct {
	Name      string
	EntryName string `json:"entry_name"`
	Comment   Comment
	Project   Project
}
