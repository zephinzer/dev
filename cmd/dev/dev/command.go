package dev

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get"
	"github.com/usvc/dev/cmd/dev/initialise"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "dev",
		Short: "The ultimate developer experience tool",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(get.GetCommand())
	cmd.AddCommand(initialise.GetCommand())
	return &cmd
}
