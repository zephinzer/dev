package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinksTest struct {
	suite.Suite
}

func TestLinks(t *testing.T) {
	suite.Run(t, &LinksTest{})
}

func (s *LinksTest) TestMergeWith() {
	input1 := Links{
		{URL: "https://1a.com"},
		{URL: "https://1b.com"},
	}
	input2 := Links{
		{URL: "https://2a.com"},
		{URL: "https://2b.com"},
	}
	input1.MergeWith(input2)
	s.Len(input1, 4)
	s.Len(input2, 2)
}

func (s *LinksTest) TestMergeWith_noDuplicates() {
	input1 := Links{
		{URL: "https://duplicate1.com"},
		{URL: "https://duplicate2.com"},
	}
	input2 := Links{
		{URL: "https://duplicate1.com"},
		{URL: "https://duplicate2.com"},
	}
	input1.MergeWith(input2)
	s.Len(input1, 2)
}
