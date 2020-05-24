package main

import (
	"os"
	"path"

	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
)

func initialiseLogger() {
	switch true {
	case conf.GetBool("debug"):
		log.Init(log.LevelDebug)
	case conf.GetBool("trace"):
		log.Init(log.LevelTrace)
	default:
		log.Init()
	}
}

func loadConfiguration() {
	log.Debug("loading configuration...")
	loadGlobalConfiguration()
	loadLocalConfiguration()

	if config.Global.Includes != nil && len(config.Global.Includes) > 0 {
		for _, include := range config.Global.Includes {
			includedConfig, getConfigError := include.GetConfig()
			if getConfigError != nil {
				log.Warnf("failed to include configuration at '%s': %s", include, getConfigError)
				continue
			}
			config.Global.MergeWith(includedConfig)
		}
	}

	if config.Global == nil {
		log.Warn("no stored configuration was loaded, using defaults for all commands")
	}
}

func loadGlobalConfiguration() {
	log.Trace("loading configuration from ~/dev.yaml...")
	homeDir, getHomeDirError := os.UserHomeDir()
	if getHomeDirError != nil {
		log.Debugf("unable to retrieve user's home directory: %s", getHomeDirError)
		return
	}
	globalConfigurationFilePath := path.Join(homeDir, "./dev.yaml")
	globalConfiguration, loadConfigurationError := config.NewFromFile(globalConfigurationFilePath)
	if loadConfigurationError != nil {
		log.Debugf("global configuration from %s could not be loaded: %s", globalConfigurationFilePath, loadConfigurationError)
		return
	}
	config.Global.MergeWith(globalConfiguration)
	log.Debug("loaded glocal configuration from ~/dev.yaml")
}

func loadLocalConfiguration() {
	log.Trace("loading configuration from ./dev.yaml...")
	localConfiguration, loadConfigurationError := config.NewFromFile("./dev.yaml")
	if loadConfigurationError != nil {
		log.Debugf("local configuration from ./dev.yaml could not be loaded: %s", loadConfigurationError)
		return
	}
	log.Debug("loaded local configuration from ./dev.yaml")
	config.Global.MergeWith(localConfiguration)
}
