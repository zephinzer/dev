package prompt

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SelectTests struct {
	suite.Suite
}

func TestSelect(t *testing.T) {
	suite.Run(t, &SelectTests{})
}

func (s *SelectTests) Test_ToSelect() {
	selections := []string{"a", "b", "c"}
	selectIndex := "1"
	selected, err := ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.Nil(err)
	s.Equal(0, selected, "selection should utilise a one-based index")
}

func (s *SelectTests) Test_ToSelect_isSkippableWithZero() {
	selections := []string{"a", "b", "c"}
	selectIndex := "0"
	selected, err := ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.Nil(err)
	s.EqualValues(ErrorSkipped, selected)
}

func (s *SelectTests) Test_ToSelect_isSkippableWithLineFeed() {
	selections := []string{"a", "b", "c"}
	selectIndex := "\n"
	selected, err := ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.Nil(err)
	s.EqualValues(ErrorSkipped, selected)
}

func (s *SelectTests) Test_ToSelect_failsGracefullyWithOutOfRange() {
	selections := []string{"a", "b", "c"}
	selectIndex := "-1"
	selected, err := ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.NotNil(err)
	s.EqualValues(ErrorInput, selected, "should return ErrorInput with value %v when a negative number is specified", ErrorInput)

	selectIndex = strconv.Itoa(len(selections) + 1)
	selected, err = ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.NotNil(err)
	s.EqualValues(ErrorInput, selected, "should return ErrorInput with value %v when one-based index selection exceeds number of options", ErrorInput)
}

func (s *SelectTests) Test_ToSelect_failsGracefullyWithNonNumericInput() {
	selections := []string{"a", "b", "c"}
	selectIndex := "a"
	selected, err := ToSelect(InputOptions{
		Reader:            bytes.NewBuffer([]byte(selectIndex)),
		SerializedOptions: selections,
	})
	s.NotNil(err)
	s.EqualValues(ErrorInput, selected, "should return ErrorInput with value %v on non-numeric input", ErrorInput)
}
