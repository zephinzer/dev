package get

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account"
	"github.com/usvc/dev/cmd/dev/get/config"
	"github.com/usvc/dev/cmd/dev/get/notifs"
	"github.com/usvc/dev/cmd/dev/get/work"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Retrieves objects",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(notifs.GetCommand())
	cmd.AddCommand(account.GetCommand())
	cmd.AddCommand(work.GetCommand())
	cmd.AddCommand(config.GetCommand())
	return &cmd
}
