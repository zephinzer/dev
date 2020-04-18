package get

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account"
	"github.com/usvc/dev/cmd/dev/get/notifs"
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
	return &cmd
}
