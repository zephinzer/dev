package add

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/add/repository"
	"github.com/zephinzer/dev/cmd/dev/add/this"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AddCanonicalVerb,
		Aliases: constants.AddAliases,
		Short:   "add to your configuration",
		Run:     run,
	}
	cmd.AddCommand(repository.GetCommand())
	cmd.AddCommand(this.GetCommand())
	return &cmd
}

func run(command *cobra.Command, args []string) {
	command.Help()
}
