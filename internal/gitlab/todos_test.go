package gitlab

import (
	"testing"

	"github.com/stretchr/testify/suite"
	pkggitlab "github.com/zephinzer/dev/pkg/gitlab"
)

type TodosTests struct {
	suite.Suite
}

func TestTodos(t *testing.T) {
	suite.Run(t, &TodosTests{})
}

func (s TodosTests) Test_TodoSerializer_GetMessage_assigned() {
	ts := TodoSerializer{
		ActionName: "assigned",
		TargetURL:  "__url",
	}
	s.NotPanics(func() {
		message := ts.GetMessage()
		s.NotZero(message)
		s.Contains(message, "__url", "message should contain the url for user to click")
	})
}

func (s TodosTests) Test_TodoSerializer_GetMessage_mentioned() {
	ts := TodoSerializer{
		ActionName: "mentioned",
		Body:       "__body",
		Target: pkggitlab.APIv4Target{
			Title: "__title",
		},
		TargetURL: "__url",
	}
	s.NotPanics(func() {
		message := ts.GetMessage()
		s.NotZero(message)
		s.Contains(message, "__body")
		s.Contains(message, "__title")
		s.Contains(message, ": \"__body\"")
		s.Contains(message, "__url", "should contain the url for user to click")
	})

	ts.Target.Title = "__body"
	s.NotPanics(func() {
		message := ts.GetMessage()
		s.NotZero(message)
		s.Contains(message, "__body")
		s.Contains(message, "__body")
		s.NotContains(message, ": \"__body\"")
		s.Contains(message, "__url", "should contain the url for user to click")
	})
}
