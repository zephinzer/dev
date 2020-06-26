package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapsTests struct {
	suite.Suite
}

func TestMaps(t *testing.T) {
	suite.Run(t, &MapsTests{})
}

func (s *MapsTests) TestGetNKeyValuePairsStringMap() {
	s.Equal(1, GetNKeyValuePairsStringMap(map[string]string{
		"a": "A",
	}))
	s.Equal(2, GetNKeyValuePairsStringMap(map[string]string{
		"a": "A",
		"b": "B",
	}))
	s.Equal(3, GetNKeyValuePairsStringMap(map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}))
}
