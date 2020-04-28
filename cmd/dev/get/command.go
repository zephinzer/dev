package get

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account"
	"github.com/usvc/dev/cmd/dev/get/configuration"
	"github.com/usvc/dev/cmd/dev/get/notifs"
	"github.com/usvc/dev/cmd/dev/get/work"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GetCanonicalVerb,
		Aliases: constants.GetAliases,
		Short:   "Retrieves objects",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(notifs.GetCommand())
	cmd.AddCommand(account.GetCommand())
	cmd.AddCommand(work.GetCommand())
	cmd.AddCommand(configuration.GetCommand())
	return &cmd
}
