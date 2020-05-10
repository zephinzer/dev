package configuration

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ConfigurationCanonicalNoun,
		Aliases: constants.ConfigurationAliases,
		Short:   "dumps the current configuration as-is",
		Run: func(command *cobra.Command, args []string) {
			fileConfig, readConfigError := config.NewFromFile(constants.DefaultPathToConfiguration)
			if readConfigError != nil {
				panic(readConfigError)
			}
			if len(fileConfig.Platforms.PivotalTracker.AccessToken) > 0 {
				fileConfig.Platforms.PivotalTracker.AccessToken = "[REDACTED]"
			}
			for i := 0; i < len(fileConfig.Platforms.PivotalTracker.Projects); i++ {
				if len(fileConfig.Platforms.PivotalTracker.Projects[i].AccessToken) > 0 {
					fileConfig.Platforms.PivotalTracker.Projects[i].AccessToken = "[REDACTED]"
				}
			}
			for i := 0; i < len(fileConfig.Platforms.Github.Accounts); i++ {
				if len(fileConfig.Platforms.Github.Accounts[i].AccessToken) > 0 {
					fileConfig.Platforms.Github.Accounts[i].AccessToken = "[REDACTED]"
				}
			}
			for i := 0; i < len(fileConfig.Platforms.Gitlab.Accounts); i++ {
				if len(fileConfig.Platforms.Gitlab.Accounts[i].AccessToken) > 0 {
					fileConfig.Platforms.Gitlab.Accounts[i].AccessToken = "[REDACTED]"
				}
			}
			litter.Dump(fileConfig)
		},
	}
	return &cmd
}
