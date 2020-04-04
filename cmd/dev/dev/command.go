package dev

import (
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use: "dev",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	return &cmd
}
