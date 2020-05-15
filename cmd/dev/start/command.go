package start

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/start/client"
	"github.com/zephinzer/dev/cmd/dev/start/server"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.StartCanonicalVerb,
		Aliases: constants.StartAliases,
		Short:   "starts dev as a long-running background process",
		Run: func(command *cobra.Command, args []string) {
			client.GetCommand().Run(command, args)
		},
	}
	cmd.AddCommand(server.GetCommand())
	cmd.AddCommand(client.GetCommand())
	return &cmd
}
