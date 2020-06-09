package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zephinzer/dev/cmd/dev/add"
	"github.com/zephinzer/dev/cmd/dev/check"
	"github.com/zephinzer/dev/cmd/dev/debug"
	"github.com/zephinzer/dev/cmd/dev/get"
	"github.com/zephinzer/dev/cmd/dev/go_to"
	"github.com/zephinzer/dev/cmd/dev/initialise"
	"github.com/zephinzer/dev/cmd/dev/install"
	"github.com/zephinzer/dev/cmd/dev/open"
	"github.com/zephinzer/dev/cmd/dev/start"
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
