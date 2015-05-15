package worktile

type Event interface {
	Format() string
}

type User struct {
	Name string
}

type Project struct {
	Name string
}
