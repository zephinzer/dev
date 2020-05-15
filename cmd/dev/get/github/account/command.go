package account

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/github"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information from Github",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			for _, githubAccount := range config.Global.Platforms.Github.Accounts {
				accountName := githubAccount.Name
				accountAccessToken := githubAccount.AccessToken
				if len(accountAccessToken) == 0 {
					log.Infof("skipping '%s': access token was not specified", accountName)
					continue
				}
				if accountsEncountered[accountAccessToken] == nil {
					accountsEncountered[accountAccessToken] = true
					log.Infof("account information for '%s'\n", accountName)
					accountInfo, err := github.GetAccount(accountAccessToken)
					if err != nil {
						log.Warnf("failed to retrieve account information for '%s'", accountName)
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
