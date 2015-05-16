package worktile

import (
	"encoding/json"
	"fmt"
)

// See https://worktile.com/club/thread/563c7312756645aaa549f7ede92bc087
type Notification struct {
	Action string
	Data   json.RawMessage
}

func (notification *Notification) Event() Event {

	var event Event

	switch notification.Action {
	case "create_task":
		event = new(CreateTaskEvent)
	case "complete_task":
		event = new(CompleteTaskEvent)
	case "expire_task":
		event = new(ExpireTaskEvent)
	case "assign_task":
		event = new(AssignTaskEvent)
	case "comment_task":
		event = new(CommentTaskEvent)
	default:
		event = nil
	}

	if event == nil {
		fmt.Println(notification.Action, " is of a type I don't know how to handle")
		return nil
	}

	if err := json.Unmarshal(notification.Data, event); err != nil {
		return nil
	} else {
		return event
	}
}
