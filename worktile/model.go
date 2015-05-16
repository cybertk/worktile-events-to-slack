package worktile

type Event interface {
}

type User struct {
	Name string
}

type Project struct {
	Name string
	Id   string `json:"pid"`
}
