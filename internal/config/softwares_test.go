package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/pkg/software"
)

type SoftwaresTest struct {
	suite.Suite
}

func TestSoftwares(t *testing.T) {
	suite.Run(t, &SoftwaresTest{})
}

func (s *SoftwaresTest) TestMergeWith() {
	input1 := Softwares{
		{Check: software.Check{Command: []string{"command_a"}}},
		{Check: software.Check{Command: []string{"command_b"}}},
	}
	input2 := Softwares{
		{Check: software.Check{Command: []string{"command_c"}}},
		{Check: software.Check{Command: []string{"command_d"}}},
	}
	input1.MergeWith(input2)
	s.Len(input1, 4)
	s.Len(input2, 2)
}

func (s *SoftwaresTest) TestMergeWith_noDuplicates() {
	input1 := Softwares{
		{Check: software.Check{Command: []string{"command_1"}}},
		{Check: software.Check{Command: []string{"command_2"}}},
	}
	input2 := Softwares{
		{Check: software.Check{Command: []string{"command_1"}}},
		{Check: software.Check{Command: []string{"command_2"}}},
	}
	input1.MergeWith(input2)
	s.Len(input1, 2)
}
