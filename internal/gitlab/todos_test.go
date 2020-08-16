package gitlab

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	pkggitlab "github.com/zephinzer/dev/pkg/gitlab"
	"github.com/zephinzer/dev/pkg/utils/request"
	"github.com/zephinzer/dev/tests"
)

type TodosTests struct {
	suite.Suite
}

func TestTodos(t *testing.T) {
	suite.Run(t, &TodosTests{})
}

func (s TodosTests) Test_GetTodos() {
	s.Nil(tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetTodos(client, "__hostname", "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal("__hostname", req.Host)
			s.Equal([]string{"__access_token"}, req.Header["Private-Token"])
			return nil
		},
		[]byte("[]"),
	))
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

func (s TodosTests) Test_TodoSerializer_GetMessage_shouldBeUnique() {
	actionNames := []string{
		"assigned",
		"mentioned",
		"directly_addressed",
		"build_failed",
		"marked",
		"approval_required",
		"unmergeable",
	}
	todoSerializers := []TodoSerializer{}
	for _, actionName := range actionNames {
		todoSerializers = append(
			todoSerializers,
			TodoSerializer{ActionName: actionName},
		)
	}
	seen := map[string]string{}
	for _, todoSerializer := range todoSerializers {
		title := todoSerializer.GetMessage()
		actionName, ok := seen[title]
		s.Falsef(ok, "there should not be duplicate messages, action names '%s' and '%s' have the same title", actionName, todoSerializer.ActionName)
		seen[title] = todoSerializer.ActionName
	}
}

func (s TodosTests) Test_TodoSerializer_GetTitle_shouldBeUnique() {
	actionNames := []string{
		"assigned",
		"mentioned",
		"directly_addressed",
		"build_failed",
		"marked",
		"approval_required",
		"unmergeable",
	}
	todoSerializers := []TodoSerializer{}
	for _, actionName := range actionNames {
		todoSerializers = append(
			todoSerializers,
			TodoSerializer{ActionName: actionName},
		)
	}
	seen := map[string]string{}
	for _, todoSerializer := range todoSerializers {
		title := todoSerializer.GetTitle()
		actionName, ok := seen[title]
		s.Falsef(ok, "there should not be duplicate titles, action names '%s' and '%s' have the same title", actionName, todoSerializer.ActionName)
		seen[title] = todoSerializer.ActionName
	}
}

func (s TodosTests) Test_TodoSerializer_getTargetType() {
	todoSerializer := TodoSerializer{TargetType: "Issue"}
	s.Equal("an issue", todoSerializer.getTargetType())
	todoSerializer = TodoSerializer{TargetType: "MergeRequest"}
	s.Equal("a merge request", todoSerializer.getTargetType())
	todoSerializer = TodoSerializer{TargetType: "AnythingElse"}
	s.Equal("an item", todoSerializer.getTargetType())
}
