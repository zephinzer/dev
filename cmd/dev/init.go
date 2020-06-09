package main

import (
	"fmt"
	"os"
	"strings"

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
	log.Debug("loading configurations...")
	files, getFilesError := config.GetFiles()
	if getFilesError != nil {
		log.Error("failed to retrieve configuration files: %s", getFilesError)
		os.Exit(1)
	}
	for _, file := range files {
		log.Debugf("processing configuration at %s...", file)
		configuration, newFromFileError := config.NewFromFile(file)
		if newFromFileError != nil {
			errorString := fmt.Sprintf("global configuration from %s could not be loaded: %s", file, newFromFileError)
			if strings.Contains(newFromFileError.Error(), "yaml: line") {
				log.Errorf(errorString)
				os.Exit(1)
				return
			}
			log.Warn(errorString)
			continue
		}
		if mergeWarnings := config.Global.MergeWith(configuration); mergeWarnings != nil && len(mergeWarnings) > 0 {
			for _, warning := range mergeWarnings {
				log.Warnf("encountered warning: %s", warning)
			}
		}
		log.Debugf("done processing configuration at %s", file)
	}
}
