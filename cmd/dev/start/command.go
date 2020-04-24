package start

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/start/client"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Short:   "starts dev as a long-running background process",
		Run: func(command *cobra.Command, _ []string) {
			command.Help()
		},
	}
	cmd.AddCommand(client.GetCommand())
	return &cmd
}
