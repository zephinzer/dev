package check

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/check/network"
	"github.com/zephinzer/dev/cmd/dev/check/repositories"
	"github.com/zephinzer/dev/cmd/dev/check/software"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.CheckCanonicalVerb,
		Aliases: constants.CheckAliases,
		Short:   "perform system checks using provided configuration",
		Run: func(command *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Print("running network checks...\n")
				network.GetCommand().Run(command, args)
				log.Print("running software checks...\n")
				software.GetCommand().Run(command, args)
				log.Print("running repository checks...\n")
				repositories.GetCommand().Run(command, args)
				return
			}
			command.Help()
		},
	}
	cmd.AddCommand(repositories.GetCommand())
	cmd.AddCommand(software.GetCommand())
	cmd.AddCommand(network.GetCommand())
	return &cmd
}
