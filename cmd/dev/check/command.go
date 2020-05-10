package check

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/check/network"
	"github.com/usvc/dev/cmd/dev/check/repositories"
	"github.com/usvc/dev/cmd/dev/check/software"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.CheckCanonicalVerb,
		Aliases: constants.CheckAliases,
		Short:   "perform system checks using provided configuration",
		Run: func(command *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Print("running network checks...")
				network.GetCommand().Run(command, args)
				log.Print("running software checks...")
				software.GetCommand().Run(command, args)
				log.Print("running repository checks...")
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
