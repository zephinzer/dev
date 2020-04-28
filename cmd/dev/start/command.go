package start

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/start/client"
	"github.com/usvc/dev/cmd/dev/start/server"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.StartCanonicalVerb,
		Aliases: constants.StartAliases,
		Short:   "starts dev as a long-running background process",
		Run: func(command *cobra.Command, _ []string) {
			command.Help()
		},
	}
	cmd.AddCommand(server.GetCommand())
	cmd.AddCommand(client.GetCommand())
	return &cmd
}
