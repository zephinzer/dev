package configuration

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "configuration",
		Aliases: []string{"config", "conf", "cf", "c"},
		Short:   "Displays the current configuration",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile(constants.DefaultPathToConfiguration)
			if err != nil {
				panic(err)
			}
			litter.Dump(c)
		},
	}
	return &cmd
}
