package main

import (
	"fmt"
	"os"
	"path"

	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/log"
)

var (
	// Commit will be set to the commit hash during build time
	Commit string
	// Version will be set to the semantic version during build time
	Version string
	// Timestamp will be set to the timestamp of the build during build time
	Timestamp string
)

func init() {
	log.Init()
	log.Info("loading configuration...")
	log.Debug("loading configuration from ~/dev.yaml...")
	homeDir, getHomeDirError := os.UserHomeDir()
	if getHomeDirError != nil {
		log.Errorf("unable to retrieve user's home directory: %s", getHomeDirError)
		os.Exit(1)
	}
	globalConfigurationFilePath := path.Join(homeDir, "./dev.yaml")
	globalConfiguration, loadConfigurationError := config.NewFromFile(globalConfigurationFilePath)
	if loadConfigurationError != nil {
		log.Debugf("configuration from %s could not be loaded: %s", globalConfigurationFilePath, loadConfigurationError)
	} else {
		config.Global = globalConfiguration
	}

	log.Debug("loading configuration from ./dev.yaml...")
	localConfiguration, loadConfigurationError := config.NewFromFile("./dev.yaml")
	if loadConfigurationError != nil {
		log.Debugf("configurations from ./dev.yaml could not be loaded: %s", loadConfigurationError)
	} else if config.Global == nil {
		config.Global = localConfiguration
	} else {
		config.Global.MergeWith(localConfiguration)
	}
}

func main() {
	cmd := GetCommand()
	cmd.Version = fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp)
	cmd.Execute()
}
