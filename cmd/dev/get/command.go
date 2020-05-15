package get

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/account"
	"github.com/zephinzer/dev/cmd/dev/get/configuration"
	"github.com/zephinzer/dev/cmd/dev/get/notifications"
	"github.com/zephinzer/dev/cmd/dev/get/sysinfo"
	"github.com/zephinzer/dev/cmd/dev/get/work"
	"github.com/zephinzer/dev/cmd/dev/get/workspace"
	"github.com/zephinzer/dev/internal/constants"
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
	cmd.AddCommand(workspace.GetCommand())
	return &cmd
}
