package cmdutils

import (
	"fmt"

	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils/str"
	"github.com/zephinzer/dev/pkg/validator"
)

func AddRepositoryToConfig(repoURL string) error {
	parsedURL, err := validator.ParseURL(repoURL)
	if err != nil {
		return fmt.Errorf("failed to get ssh git clone url for repository with url '%s': %s", repoURL, err)
	}

	configurationPath, err := config.PromptSelectLoadedConfiguration(
		fmt.Sprintf("which configuration file should we add '%s' to?", repoURL),
	)
	if err != nil {
		return fmt.Errorf("failed to get a valid configuration file: %s", err)
	}
	if str.IsEmpty(configurationPath) {
		return fmt.Errorf("configuration was skipped for repo '%s'", repoURL)
	}
	log.Infof("adding repo '%s' to configuration at '%s'...", repoURL, configurationPath)

	// TODO: complete this using code from ~/cmd/dev/add/repository/command.go
	log.Infof(parsedURL.String())
	log.Infof(configurationPath)
	return nil
}
