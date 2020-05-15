package main

import (
	"os"
	"path"

	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
)

func initialiseConfiguration() {
	log.Debug("loading configuration...")
	log.Trace("loading configuration from ~/dev.yaml...")
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

	log.Trace("loading configuration from ./dev.yaml...")
	localConfiguration, loadConfigurationError := config.NewFromFile("./dev.yaml")
	if loadConfigurationError != nil {
		log.Debugf("configurations from ./dev.yaml could not be loaded: %s", loadConfigurationError)
	} else if config.Global == nil {
		config.Global = localConfiguration
	} else {
		config.Global.MergeWith(localConfiguration)
	}
}

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
