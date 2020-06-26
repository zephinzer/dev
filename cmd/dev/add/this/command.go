package this

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/add/this/repository"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ThisCanonicalNoun,
		Aliases: constants.ThisAliases,
		Short:   "adds the specified resource to dev's configuration",
		Run:     run,
	}
	cmd.AddCommand(repository.GetCommand())
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
