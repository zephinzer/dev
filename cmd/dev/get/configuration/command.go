package configuration

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ConfigurationCanonicalNoun,
		Aliases: constants.ConfigurationAliases,
		Short:   "dumps the current configuration as-is",
		Run: func(command *cobra.Command, args []string) {
			litter.Dump(config.Global)
		},
	}
	return &cmd
}
