package cmdutils

import (
	"fmt"

	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
	"github.com/zephinzer/dev/pkg/validator"
)

func AddRepositoryToConfig(repoURL string) error {
	parsedURL, err := validator.ParseURL(repoURL)
	if err != nil {
		return fmt.Errorf("failed to get ssh git clone url for repository with url '%s': %s", repoURL, err)
	}

	configurationPath := config.PromptSelectLoadedConfiguration(
		fmt.Sprintf("which configuration file should we add '%s' to?", repoURL),
	)
	if utils.IsEmptyString(configurationPath) {
		return fmt.Errorf("configuration was skipped for repo '%s'", repoURL)
	}
	log.Infof("adding repo '%s' to configuration at '%s'...", repoURL, configurationPath)

	log.Infof(parsedURL.String())
	log.Infof(configurationPath)
	return nil
}
