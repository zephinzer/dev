package main

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/check"
	"github.com/usvc/dev/cmd/dev/debug"
	"github.com/usvc/dev/cmd/dev/get"
	"github.com/usvc/dev/cmd/dev/go_to"
	"github.com/usvc/dev/cmd/dev/initialise"
	"github.com/usvc/dev/cmd/dev/open"
	"github.com/usvc/dev/cmd/dev/start"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "dev",
		Short: "The ultimate developer experience tool",
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
	return &cmd
}
