package prompt

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InputHelperTests struct {
	suite.Suite
}

func TestInputHelper(t *testing.T) {
	suite.Run(t, &InputHelperTests{})
}

func (s *InputHelperTests) TestPrintBeforeMessage() {

}
