package debug

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/debug/notifications"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.DebugCanonicalVerb,
		Aliases: constants.DebugAliases,
		Short:   "run debug versions of error-prone functionality to verify behaviour",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(notifications.GetCommand())
	return &cmd
}
