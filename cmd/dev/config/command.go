package config

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/config/view"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "config",
		Short: "Configuration related commands",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(view.GetCommand())
	return &cmd
}
