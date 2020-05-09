package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/check"
	"github.com/usvc/dev/cmd/dev/debug"
	"github.com/usvc/dev/cmd/dev/get"
	"github.com/usvc/dev/cmd/dev/go_to"
	"github.com/usvc/dev/cmd/dev/initialise"
	"github.com/usvc/dev/cmd/dev/open"
	"github.com/usvc/dev/cmd/dev/start"
	"github.com/usvc/dev/internal/constants"
	configuration "github.com/usvc/go-config"
)

var (
	// Commit will be set to the commit hash during build time
	Commit = "<commit-hash>"
	// Version will be set to the semantic version during build time
	Version = "<semver-version>"
	// Timestamp will be set to the timestamp of the build during build time
	Timestamp = time.Now().Format(constants.DevTimeFormat)
	// conf is the local configuration for the root command
	conf = configuration.Map{
		"debug": &configuration.Bool{
			Usage: "display up to debug level logs (verbose logging mode)",
		},
		"trace": &configuration.Bool{
			Usage: "display up to trace level logs (very-verbose logging mode)",
		},
	}
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "dev",
		Short:   "The ultimate developer experience tool",
		Version: fmt.Sprintf("%s-%s built at %s", Version, Commit, Timestamp),
		PersistentPreRun: func(command *cobra.Command, args []string) {
			initialiseLogger()
			initialiseConfiguration()
		},
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(debug.GetCommand())
	cmd.AddCommand(check.GetCommand())
	cmd.AddCommand(get.GetCommand())
	cmd.AddCommand(initialise.GetCommand())
	cmd.AddCommand(start.GetCommand())
	cmd.AddCommand(open.GetCommand())
	cmd.AddCommand(go_to.GetCommand())
	conf.ApplyToFlagSet(cmd.PersistentFlags())
	return &cmd
}
