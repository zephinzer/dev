package check

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/check/network"
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
			log.Info("running software checks...")
			software.GetCommand().Run(command, args)
			log.Info("running network checks...")
			network.GetCommand().Run(command, args)
		},
	}
	cmd.AddCommand(software.GetCommand())
	cmd.AddCommand(network.GetCommand())
	return &cmd
}
