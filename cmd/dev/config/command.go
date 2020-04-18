package config

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/pkg/config"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "config",
		Short: "Displays the dev environment configuration",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile("./dev.yaml")
			if err != nil {
				panic(err)
			}
			litter.Dump(c)
		},
	}
	return &cmd
}
