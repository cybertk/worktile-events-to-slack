package worktile

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommentTaskEventUnmarshal(t *testing.T) {

	Convey("Given the `data` field of create_task notification", t, func() {

		data := []byte(`
{
    "tid": "9b4c0bdb901a47f9810158c9b5886996",
    "name": "Comment Task",
    "entry_id": "31d8cfe8bd6d4f11ba381be66d5fb643",
    "entry_name": "TODO",
    "create_date": 1431679138537,
    "comment": {
		"cid": "ab2340d36ec94af79d7c97c12435d609",
		"message": "comment",
		"create_date": 1431709625120,
		"create_by": {
			"uid": "c61889a8603c4d26aae65aaab747d9dd",
			"name": "quanlong",
			"display_name": "Quanlong"
		}
    },
    "project": {
	    "pid": "f5fb3690f0e3403abf3f590a08b3df95",
	    "name": "Foo Project"
	}
}
		`)

		Convey("When decode as struct", func() {

			var event CommentTaskEvent
			err := json.Unmarshal(data, &event)

			Convey("Should parsed succeed", func() {
				So(err, ShouldBeNil)
				So(event.Name, ShouldEqual, "Comment Task")
				So(event.EntryName, ShouldEqual, "TODO")
				So(event.Comment.Message, ShouldEqual, "comment")
				So(event.Comment.CreateBy.Name, ShouldEqual, "quanlong")
				So(event.Project.Name, ShouldEqual, "Foo Project")
			})
		})
	})
}

func TestCommentTaskEventFormat(t *testing.T) {

	Convey("Given a CommentTaskEvent", t, func() {

		event := CommentTaskEvent{
			Name:      "Task Name",
			EntryName: "Entry Name",
			Comment:   Comment{Message: "Comment Message", CreateBy: User{Name: "User Name"}},
			Project:   Project{Name: "Project Name"},
		}

		Convey("When Format()", func() {

			str := event.Format()

			Convey("Should contains corresponding struct info", func() {
				So(str, ShouldContainSubstring, "Task Name")
				So(str, ShouldContainSubstring, "Entry Name")
				So(str, ShouldContainSubstring, "User Name")
				So(str, ShouldContainSubstring, "Comment Message")
				So(str, ShouldContainSubstring, "Project Name")
			})
		})
	})
}
