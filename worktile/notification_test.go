package worktile

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotificationUnmarshal(t *testing.T) {

	Convey("Given a Worktile Webhook notification", t, func() {

		data := []byte(`
{
	"action": "create_task",
	"data": {
		"tid": "12caaedfe54c471abec4dc54081a135e",
		"name": "Create Task",
		"labels": [],
		"assign": [],
		"entry_id": "31d8cfe8bd6d4f11ba381be66d5fb643",
		"entry_name": "TODO",
		"expire_date": 0,
		"create_by": {
			"uid": "c61889a8603c4d26aae65aaab747d9dd",
			"name": "quanlong",
			"display_name": "Quanlong",
			"email": "kyan.ql.he@gmail.com"
		},
		"create_date": 1431594004446,
		"project": {
			"pid": "f5fb3690f0e3403abf3f590a08b3df95",
			"name": "Foo Project"
		}
	}
}
		`)

		Convey("When decode as struct", func() {

			var notification Notification
			err := json.Unmarshal(data, &notification)

			Convey("Should contains corresponding struct info", func() {
				So(err, ShouldBeNil)
				So(notification.Action, ShouldEqual, "create_task")
				So(notification.Data, ShouldHaveSameTypeAs, json.RawMessage{})
				So(string(notification.Data), ShouldStartWith, "{")
				So(string(notification.Data), ShouldEndWith, "}")
			})
		})

	})
}

func TestNotificationEvent(t *testing.T) {

	Convey("Given a simple notification", t, func() {

		notification := Notification{
			Action: "create_task",
			Data: []byte(`
{
	"tid": "12caaedfe54c471abec4dc54081a135e",
	"name": "Create Task",
	"labels": [],
	"assign": [],
	"entry_id": "31d8cfe8bd6d4f11ba381be66d5fb643",
	"entry_name": "TODO",
	"expire_date": 0,
	"create_by": {
		"uid": "c61889a8603c4d26aae65aaab747d9dd",
		"name": "quanlong",
		"display_name": "Quanlong",
		"email": "kyan.ql.he@gmail.com"
	},
	"create_date": 1431594004446,
	"project": {
		"pid": "f5fb3690f0e3403abf3f590a08b3df95",
		"name": "Foo Project"
	}
}
		`),
		}

		Convey("When parse event", func() {

			event := notification.Event()

			Convey("Should parsed succeed", func() {
				So(event, ShouldHaveSameTypeAs, new(CreateTaskEvent))
			})
		})

	})

	Convey("Given a complete_task notification", t, func() {

		notification := Notification{Action: "complete_task", Data: []byte(`{}`)}

		Convey("When parse event", func() {

			event := notification.Event()

			Convey("Should parsed succeed", func() {
				So(event, ShouldHaveSameTypeAs, new(CompleteTaskEvent))
			})
		})

	})

	Convey("Given a expire_task notification", t, func() {

		notification := Notification{Action: "expire_task", Data: []byte(`{}`)}

		Convey("When parse event", func() {

			event := notification.Event()

			Convey("Should parsed succeed", func() {
				So(event, ShouldHaveSameTypeAs, new(ExpireTaskEvent))
			})
		})

	})

	Convey("Given a assign_task notification", t, func() {

		notification := Notification{Action: "assign_task", Data: []byte(`{}`)}

		Convey("When parse event", func() {

			event := notification.Event()

			Convey("Should parsed succeed", func() {
				So(event, ShouldHaveSameTypeAs, new(AssignTaskEvent))
			})
		})

	})

	Convey("Given a comment_task notification", t, func() {

		notification := Notification{Action: "comment_task", Data: []byte(`{}`)}

		Convey("When parse event", func() {

			event := notification.Event()

			Convey("Should parsed succeed", func() {
				So(event, ShouldHaveSameTypeAs, new(CommentTaskEvent))
			})
		})

	})
}
