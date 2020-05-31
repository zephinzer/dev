package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	configuration "github.com/usvc/go-config"
	"github.com/zephinzer/dev/cmd/dev/add"
	"github.com/zephinzer/dev/cmd/dev/check"
	"github.com/zephinzer/dev/cmd/dev/debug"
	"github.com/zephinzer/dev/cmd/dev/get"
	"github.com/zephinzer/dev/cmd/dev/go_to"
	"github.com/zephinzer/dev/cmd/dev/initialise"
	"github.com/zephinzer/dev/cmd/dev/install"
	"github.com/zephinzer/dev/cmd/dev/open"
	"github.com/zephinzer/dev/cmd/dev/start"
	"github.com/zephinzer/dev/internal/constants"
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
			loadConfiguration()
		},
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(add.GetCommand())
	cmd.AddCommand(check.GetCommand())
	cmd.AddCommand(debug.GetCommand())
	cmd.AddCommand(get.GetCommand())
	cmd.AddCommand(go_to.GetCommand())
	cmd.AddCommand(initialise.GetCommand())
	cmd.AddCommand(install.GetCommand())
	cmd.AddCommand(open.GetCommand())
	cmd.AddCommand(start.GetCommand())
	conf.ApplyToFlagSet(cmd.PersistentFlags())
	return &cmd
}
