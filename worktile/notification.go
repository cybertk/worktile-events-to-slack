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

	switch notification.Action {
	case "create_task":
		var data CreateTaskEvent
		if err := json.Unmarshal(notification.Data, &data); err != nil {
			fmt.Println("ck-error", err, data.Format())
			return nil
		}
		fmt.Println("ck-error", data.Format())
		return &data
	default:
		fmt.Println(notification.Action, "is of a type I don't know how to handle")
		return nil
	}
}
