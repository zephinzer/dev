package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/pkg/network"
)

type NetworksTest struct {
	suite.Suite
}

func TestNetworks(t *testing.T) {
	suite.Run(t, &NetworksTest{})
}

func (s *NetworksTest) TestMergeWith() {
	input1 := Networks{
		{Check: network.Check{URL: "https://1a.com"}},
		{Check: network.Check{URL: "https://1b.com"}},
	}
	input2 := Networks{
		{Check: network.Check{URL: "https://2a.com"}},
		{Check: network.Check{URL: "https://2b.com"}},
	}
	input1.MergeWith(input2)
	s.Len(input1, 4)
	s.Len(input2, 2)
}

func (s *NetworksTest) TestMergeWith_noDuplicates() {
	input1 := Networks{
		{Check: network.Check{URL: "https://duplicate1.com"}},
		{Check: network.Check{URL: "https://duplicate2.com"}},
	}
	input2 := Networks{
		{Check: network.Check{URL: "https://duplicate1.com"}},
		{Check: network.Check{URL: "https://duplicate2.com"}},
	}
	input1.MergeWith(input2)
	s.Len(input1, 2)
}
