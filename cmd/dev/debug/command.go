package debug

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/debug/notifications"
	systemtray "github.com/usvc/dev/cmd/dev/debug/system-tray"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.DebugCanonicalVerb,
		Aliases: constants.DebugAliases,
		Short:   "Debugs stuff",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(notifications.GetCommand())
	cmd.AddCommand(systemtray.GetCommand())
	return &cmd
}
