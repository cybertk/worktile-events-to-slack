package worktile

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCompleteTaskEventUnmarshal(t *testing.T) {

	Convey("Given the `data` field of create_task notification", t, func() {

		data := []byte(`
{
    "tid": "0ca720f640f84386bc9e8b6f0b3c5abb",
    "name": "Complete Task",
    "labels": [],
    "assign": [],
    "entry_id": "31d8cfe8bd6d4f11ba381be66d5fb643",
    "entry_name": "TODO",
    "expire_date": 1432396799999,
    "create_by": {
      "uid": "c61889a8603c4d26aae65aaab747d9dd",
      "name": "quanlong",
      "display_name": "Quanlong",
      "email": "kyan.ql.he@gmail.com"
    },
    "create_date": 1431679622563,
    "project": {
      "pid": "f5fb3690f0e3403abf3f590a08b3df95",
      "name": "Foo Project"
    }
}
		`)

		Convey("When decode as struct", func() {

			var event CompleteTaskEvent
			err := json.Unmarshal(data, &event)

			Convey("Should parsed succeed", func() {
				So(err, ShouldBeNil)
				So(event.Name, ShouldEqual, "Complete Task")
				So(event.EntryName, ShouldEqual, "TODO")
				So(event.CreateBy.Name, ShouldEqual, "quanlong")
				So(event.Project.Name, ShouldEqual, "Foo Project")
			})
		})
	})
}

func TestCompleteTaskEventFormat(t *testing.T) {

	Convey("Given a CompleteTaskEvent", t, func() {

		event := CompleteTaskEvent{
			Name:      "Task Name",
			EntryName: "Entry Name",
			CreateBy: User{
				Name: "User Name",
			},
			Project: Project{
				Name: "Project Name",
			},
		}

		Convey("When Format()", func() {

			str := event.Format()

			Convey("Should contains corresponding struct info", func() {
				So(str, ShouldContainSubstring, "Task Name")
				So(str, ShouldContainSubstring, "Entry Name")
				So(str, ShouldContainSubstring, "User Name")
				So(str, ShouldContainSubstring, "Project Name")
			})
		})
	})
}
