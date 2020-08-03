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

func (s *UtilsTests) Test_FilterConfigurations_rejectedFileNames() {
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

func (s *UtilsTests) Test_FilterConfigurations_acceptedFileNames() {
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

func (s *UtilsTests) Test_NewFromFile() {
	c, newFromFileError := NewFromFile("../../tests/config/dev.yaml")
	s.Nil(newFromFileError)
	if newFromFileError != nil {
		return
	}
	s.Equal("__expected_database_path", c.Dev.Client.Database.Path)
	s.Equal("__expected_telegram_token", c.Dev.Client.Notifications.Telegram.Token)
	s.Equal("__expected_telegram_chat_id", c.Dev.Client.Notifications.Telegram.ID)
	s.Equal("__expected_platform_github_client_id", c.Dev.Client.Platforms.Github.ClientID)
	s.Equal("__expected_platform_github_client_secret", c.Dev.Client.Platforms.Github.ClientSecret)
	s.Equal("__expected_platform_github_redirect_uri", c.Dev.Client.Platforms.Github.RedirectURI)
}
