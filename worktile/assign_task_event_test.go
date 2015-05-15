package worktile

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAssignTaskEventUnmarshal(t *testing.T) {

	Convey("Given the `data` field of create_task notification", t, func() {

		data := []byte(`
{
	"tid": "9b4c0bdb901a47f9810158c9b5886996",
	"name": "Assign Task",
	"labels": [],
	"assign": [
		{
		"uid": "c61889a8603c4d26aae65aaab747d9dd",
		"name": "quanlong",
		"display_name": "Quanlong",
		"avatar": "default_avatar.png",
		"desc": "",
		"status": 1,
		"online": 0
		}
	],
	"entry_id": "31d8cfe8bd6d4f11ba381be66d5fb643",
	"entry_name": "TODO",
	"expire_date": 0,
	"create_by": {
		"uid": "c61889a8603c4d26aae65aaab747d9dd",
		"name": "quanlong",
		"display_name": "Quanlong",
		"email": "kyan.ql.he@gmail.com"
	},
	"create_date": 1431679138537,
	"project": {
		"pid": "f5fb3690f0e3403abf3f590a08b3df95",
		"name": "Foo Project"
	},
	"assign_user": {
		"uid": "c61889a8603c4d26aae65aaab747d9dd",
		"name": "quanlong",
		"display_name": "Quanlong",
		"email": "kyan.ql.he@gmail.com"
	}
}
		`)

		Convey("When decode as struct", func() {

			var event AssignTaskEvent
			err := json.Unmarshal(data, &event)

			Convey("Should parsed succeed", func() {
				So(err, ShouldBeNil)
				So(event.Name, ShouldEqual, "Assign Task")
				So(event.EntryName, ShouldEqual, "TODO")
				So(event.CreateBy.Name, ShouldEqual, "quanlong")
				So(event.Project.Name, ShouldEqual, "Foo Project")
				So(event.AssignUser.Name, ShouldEqual, "quanlong")
			})
		})
	})
}

func TestAssignTaskEventFormat(t *testing.T) {

	Convey("Given a AssignTaskEvent", t, func() {

		event := AssignTaskEvent{
			Name:       "Task Name",
			EntryName:  "Entry Name",
			CreateBy:   User{Name: "User Name"},
			Project:    Project{Name: "Project Name"},
			AssignUser: User{Name: "Assign Name"},
		}

		Convey("When Format()", func() {

			str := event.Format()

			Convey("Should contains corresponding struct info", func() {
				So(str, ShouldContainSubstring, "Task Name")
				So(str, ShouldContainSubstring, "Entry Name")
				So(str, ShouldContainSubstring, "User Name")
				So(str, ShouldContainSubstring, "Project Name")
				So(str, ShouldContainSubstring, "Assign Name")
			})
		})
	})
}
