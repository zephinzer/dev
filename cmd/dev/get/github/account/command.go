package account

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/github"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/types"
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
					log.Infof("account information for '%s' github account\n", accountName)
					accountInfo, getAccountError := github.GetAccount(
						&http.Client{Timeout: constants.DefaultAPICallTimeout},
						accountAccessToken,
					)
					if getAccountError != nil {
						log.Warnf("failed to retrieve account information for '%s': %s", accountName, getAccountError)
						continue
					}
					log.Printf(types.PrintAccount(accountInfo))
					totalAccountsCount++
				}
			}
			log.Infof("total listed github accounts     : %v", len(config.Global.Platforms.Github.Accounts))
			log.Infof("total accessible github accounts : %v", totalAccountsCount)
		},
	}
	return &cmd
}
