package gitlab

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
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
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetTodos(client, "some.gitlab.instance", "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal("some.gitlab.instance", req.Host)
			s.Equal("/api/v4/todos", req.URL.Path)
			s.EqualValues([]string{"application/json"}, req.Header["Content-Type"])
			s.EqualValues([]string{"__access_token"}, req.Header["Private-Token"])
			return nil
		},
		[]byte("[]"),
	)
	s.Nil(systemError)
}

func (s AccountTests) Test_GetTodos_withSchemaError() {
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetTodos(client, "some.gitlab.instance", "__access_token")
			return err
		},
		tests.HTTPRequestAsserterNoOp,
		[]byte("{}"),
	)
	s.NotNil(systemError)
	s.Contains(systemError.Error(), "failed to unmarshal response")
}
