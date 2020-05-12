package github

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/github"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GithubCanonicalNoun,
		Aliases: constants.GithubAliases,
		Short:   "Retrieves account information from Github",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			for _, account := range config.Global.Platforms.Github.Accounts {
				accountAccessToken := account.AccessToken
				if len(accountAccessToken) == 0 {
					log.Infof("skipping '%s': access token was not specified", account.Name)
					continue
				}
				if accountsEncountered[accountAccessToken] == nil {
					accountsEncountered[accountAccessToken] = true
					log.Infof("account information for '%s'\n", account.Name)
					accountInfo, err := github.GetAccount(accountAccessToken)
					if err != nil {
						log.Warnf("failed to retrieve account information for '%s'", account.Name)
						continue
					}
					log.Info(accountInfo.String())
					totalAccountsCount++
				}
			}
			log.Infof("total listed accounts    : %v", len(config.Global.Platforms.Github.Accounts))
			log.Infof("total accessible accounts: %v", totalAccountsCount)
		},
	}
	return &cmd
}
