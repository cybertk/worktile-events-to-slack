package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/cybertk/worktile-events-to-slack/worktile"
	"github.com/stretchr/testify/mock"
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

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) sendToSlack(url string, event worktile.Event) error {
	args := m.Called(url, event)
	return args.Error(0)
}

func TestHandler(t *testing.T) {

	Convey("Given a POST request with valid slack_url param and valid notification data body", t, func() {

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

		expectedURL := "https://example.com/AAA/BB/C"

		// TODO: Crash when pass nil as body
		req, _ := http.NewRequest("POST", "http://example.com?slack_url="+expectedURL, bytes.NewBuffer(data))

		var notification worktile.Notification
		json.NewDecoder(bytes.NewBuffer(data)).Decode(&notification)

		Convey("And sendToSlack() always returns nil", func() {

			// Stubs
			w := httptest.NewRecorder()
			m := new(MyMockedObject)

			m.On("sendToSlack", expectedURL, notification.Event()).Return(nil).Once()

			Convey("When invoke hanlder() with them", func() {

				handler(w, req, m.sendToSlack)

				Convey("Then sendToSlack() should be called with correct parameters", func() {

					m.AssertExpectations(t)

					Convey("And send 200 to http.ResponseWriter", func() {
						So(w.Code, ShouldEqual, 200)
					})

				})
			})
		})

		Convey("And sendToSlack() returns error", func() {

			// Stubs
			w := httptest.NewRecorder()
			m := new(MyMockedObject)

			m.On("sendToSlack", expectedURL, notification.Event()).Return(errors.New("Error A")).Once()

			Convey("When invoke hanlder() with them", func() {

				handler(w, req, m.sendToSlack)

				Convey("Then sendToSlack() should be called with correct parameters", func() {

					m.AssertExpectations(t)

					Convey("And send 500 to http.ResponseWriter", func() {
						So(w.Code, ShouldEqual, 500)
					})

				})
			})
		})

	})

	Convey("Given a POST request with valid slack_url param and empty body", t, func() {

		expectedURL := "https://example.com/AAA/BB/C"

		req, _ := http.NewRequest("POST", "http://example.com?slack_url="+expectedURL, nil)

		// Stubs
		m := new(MyMockedObject)
		w := httptest.NewRecorder()

		Convey("When invoke hanlder() with them", func() {

			handler(w, req, m.sendToSlack)

			Convey("Then sendToSlack() should never be called", func() {

				m.AssertNotCalled(t, "sendToSlack")

				Convey("And send 400 to http.ResponseWriter", func() {
					So(w.Code, ShouldEqual, 400)
				})

			})
		})
	})

	Convey("Given a POST request without slack_url param", t, func() {

		req, _ := http.NewRequest("POST", "http://example.com", nil)

		// Stubs
		m := new(MyMockedObject)
		w := httptest.NewRecorder()

		Convey("When invoke hanlder() with them", func() {

			handler(w, req, m.sendToSlack)

			Convey("Then sendToSlack() should never be called", func() {

				m.AssertNotCalled(t, "sendToSlack")

				Convey("And send 400 to http.ResponseWriter", func() {
					So(w.Code, ShouldEqual, 400)

				})
			})
		})
	})
}
