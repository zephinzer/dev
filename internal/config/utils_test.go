package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	mockos "github.com/zephinzer/dev/pkg/mocks/os"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) TestFilterConfigurations_rejectedFileNames() {
	input := []os.FileInfo{
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "dev.yml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "dev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "a.dev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "adev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "0.dev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "0dev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "deva.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", "dev0.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.yamla"}),
	}
	output := FilterConfigurations(input)
	s.Len(output, 0)
}

func (s *UtilsTests) TestFilterConfigurations_acceptedFileNames() {
	input := []os.FileInfo{
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.yml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.1.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.a.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.aa.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.A.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.0.a.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.a.0.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.00.aa.yaml"}),
		(&mockos.FileInfo{}).Returns(mockos.ReturnValue{"Name", ".dev.aa.00.yaml"}),
	}
	output := FilterConfigurations(input)
	s.Len(output, len(input))
}
