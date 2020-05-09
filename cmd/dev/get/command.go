package get

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account"
	"github.com/usvc/dev/cmd/dev/get/configuration"
	"github.com/usvc/dev/cmd/dev/get/notifications"
	"github.com/usvc/dev/cmd/dev/get/sysinfo"
	"github.com/usvc/dev/cmd/dev/get/work"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GetCanonicalVerb,
		Aliases: constants.GetAliases,
		Short:   "ask you and shall retrieve",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(account.GetCommand())
	cmd.AddCommand(configuration.GetCommand())
	cmd.AddCommand(notifications.GetCommand())
	cmd.AddCommand(sysinfo.GetCommand())
	cmd.AddCommand(work.GetCommand())
	return &cmd
}
