package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) Test_ContainsInt() {
	testCase := []int{1, 2, 3, 4, 5, 7}
	s.True(ContainsInt(5, testCase))
	s.False(ContainsInt(6, testCase))
	s.True(ContainsInt(7, testCase))
}
