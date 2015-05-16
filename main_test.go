package main

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/cybertk/worktile-events-to-slack/worktile"
)

func TestFormat(t *testing.T) {

	Convey("Given a CreateTaskEvent", t, func() {

		event := worktile.CreateTaskEvent{
			Name:      "Create Task",
			EntryName: "Create Entry",
			CreateBy:  worktile.User{Name: "Create User"},
			Project:   worktile.Project{Name: "Create Project", Id: "110"},
		}

		Convey("When format()", func() {

			attachment := format(&event)

			Convey("Should contains corresponding struct info", func() {
				matched, _ := regexp.MatchString("#[0-9A-Fa-f]6", attachment.Color)
				So(matched, ShouldBeTrue)
				So(attachment.Title, ShouldEqual, "Create Project")
				So(attachment.TitleLink, ShouldEqual, "https://worktile.com/project/110")
				So(attachment.Text, ShouldContainSubstring, "Create Task")
				So(attachment.Text, ShouldContainSubstring, "Create Entry")
				So(attachment.Text, ShouldContainSubstring, "Create User")
			})
		})
	})
	Convey("Given a CompleteTaskEvent", t, func() {

		event := worktile.CompleteTaskEvent{
			Name:      "Complete Task",
			EntryName: "Complete Entry",
			CreateBy:  worktile.User{Name: "Complete User"},
			Project:   worktile.Project{Name: "Complete Project", Id: "1"},
		}

		Convey("When format()", func() {

			attachment := format(&event)

			Convey("Should contains corresponding struct info", func() {
				matched, _ := regexp.MatchString("#[0-9A-Fa-f]6", attachment.Color)
				So(matched, ShouldBeTrue)
				So(attachment.Title, ShouldEqual, "Complete Project")
				So(attachment.TitleLink, ShouldEqual, "https://worktile.com/project/1")
				So(attachment.Text, ShouldContainSubstring, "Complete Task")
				So(attachment.Text, ShouldContainSubstring, "Complete Entry")
				So(attachment.Text, ShouldContainSubstring, "Complete User")
			})
		})
	})
	Convey("Given a ExpireTaskEvent", t, func() {

		event := worktile.ExpireTaskEvent{
			Name:       "Expire Task",
			EntryName:  "Expire Entry",
			ExpireDate: 1432396799999,
			CreateBy:   worktile.User{Name: "Expire User"},
			Project:    worktile.Project{Name: "Expire Project", Id: "id"},
		}

		Convey("When format()", func() {

			attachment := format(&event)

			Convey("Should contains corresponding struct info", func() {
				matched, _ := regexp.MatchString("#[0-9A-Fa-f]6", attachment.Color)
				So(matched, ShouldBeTrue)
				So(attachment.Title, ShouldEqual, "Expire Project")
				So(attachment.TitleLink, ShouldEqual, "https://worktile.com/project/id")
				So(attachment.Text, ShouldContainSubstring, "Expire Task")
				So(attachment.Text, ShouldContainSubstring, "Expire Entry")
				So(attachment.Text, ShouldContainSubstring, "May 23")
				So(attachment.Text, ShouldContainSubstring, "Expire User")
			})
		})
	})
	Convey("Given a AssignTaskEvent", t, func() {

		event := worktile.AssignTaskEvent{
			Name:       "Assign Task",
			EntryName:  "Assign Entry",
			CreateBy:   worktile.User{Name: "Assign User"},
			Project:    worktile.Project{Name: "Assign Project", Id: "00000000"},
			AssignUser: worktile.User{Name: "Assign To User"},
		}

		Convey("When format()", func() {

			attachment := format(&event)

			Convey("Should contains corresponding struct info", func() {
				matched, _ := regexp.MatchString("#[0-9A-Fa-f]6", attachment.Color)
				So(matched, ShouldBeTrue)
				So(attachment.Title, ShouldEqual, "Assign Project")
				So(attachment.TitleLink, ShouldEqual, "https://worktile.com/project/00000000")
				So(attachment.Text, ShouldContainSubstring, "Assign Task")
				So(attachment.Text, ShouldContainSubstring, "Assign Entry")
				So(attachment.Text, ShouldContainSubstring, "Assign User")
				So(attachment.Text, ShouldContainSubstring, "Assign To User")
			})
		})
	})
	Convey("Given a CommentTaskEvent", t, func() {

		event := worktile.CommentTaskEvent{
			Name:      "Comment Task",
			EntryName: "Comment Entry",
			Comment: worktile.Comment{
				Message: "Comment Message", CreateBy: worktile.User{Name: "Comment User"}},
			Project: worktile.Project{Name: "Comment Project", Id: "1x"},
		}

		Convey("When format()", func() {

			attachment := format(&event)

			Convey("Should contains corresponding struct info", func() {
				matched, _ := regexp.MatchString("#[0-9A-Fa-f]6", attachment.Color)
				So(matched, ShouldBeTrue)
				So(attachment.Title, ShouldEqual, "Comment Project")
				So(attachment.TitleLink, ShouldEqual, "https://worktile.com/project/1x")
				So(attachment.Text, ShouldContainSubstring, "Comment Task")
				So(attachment.Text, ShouldContainSubstring, "Comment Entry")
				So(attachment.Text, ShouldContainSubstring, "Comment Message")
				So(attachment.Text, ShouldContainSubstring, "Comment User")
			})
		})
	})
}
