package open

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/open/repo"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "open",
		Aliases: []string{"op", "o"},
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(repo.GetCommand())
	return &cmd
}
