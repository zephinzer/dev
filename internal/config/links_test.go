package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinksTest struct {
	suite.Suite
}

func TestLinks(t *testing.T) {
	suite.Run(t, &LinksTest{})
}

func (s *LinksTest) TestMarshal() {
	c, newFromFileError := NewFromFile("../../tests/config/links.yaml")
	s.Nil(newFromFileError)
	if newFromFileError != nil {
		return
	}
	s.Len(c.Links, 2)
	for index, link := range c.Links {
		s.Equal(fmt.Sprintf("__expected_label_%v", index+1), link.Label)
		s.Equal(fmt.Sprintf("https://expected%v.url.com", index+1), link.URL)
		s.Equal(fmt.Sprintf("__expected_category_%v", index+1), link.Categories[0])
	}
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
