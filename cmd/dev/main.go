package main

import (
	"fmt"

	"github.com/usvc/dev/cmd/dev/dev"
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
}

func main() {
	var loadConfigurationError error
	log.Debug("loading configurations from ./dev.yaml")
	config.Global, loadConfigurationError = config.NewFromFile("./dev.yaml")
	if loadConfigurationError != nil {
		panic(loadConfigurationError)
	}
	cmd := dev.GetCommand()
	cmd.Version = fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp)
	cmd.Execute()
}
