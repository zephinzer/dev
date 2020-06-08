package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DevConfigTests struct {
	suite.Suite
}

func TestDevConfig(t *testing.T) {
	suite.Run(t, &DevConfigTests{})
}

func (s *DevConfigTests) TestMarshal() {
	c, newFromFileError := NewFromFile("../../tests/config/dev.yaml")
	s.Nil(newFromFileError)
	if newFromFileError != nil {
		return
	}
	s.Equal("__expected_database_path", c.Dev.Client.Database.Path)
	s.Equal("__expected_telegram_token", c.Dev.Client.Notifications.Telegram.Token)
	s.Equal("__expected_telegram_chat_id", c.Dev.Client.Notifications.Telegram.ID)
}
