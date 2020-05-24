package account

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/types"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}

			for _, gitlabAccount := range config.Global.Platforms.Gitlab.Accounts {
				name := gitlabAccount.Name
				if len(name) == 0 {
					name = "unnamed"
				}
				hostname := gitlabAccount.Hostname
				if len(gitlabAccount.Hostname) == 0 {
					hostname = constants.DefaultGitlabHostname
				}
				accessToken := gitlabAccount.AccessToken
				if len(accessToken) == 0 {
					log.Infof("skipping '%s'@'%s': access token was not specified", name, hostname)
					continue
				}

				if accountsEncountered[accessToken] == nil {
					accountsEncountered[accessToken] = true
					log.Infof("account information for '%s'@'%s' gitlab account", name, hostname)
					accountInfo, err := gitlab.GetAccount(hostname, accessToken)
					if err != nil {
						log.Warnf("failed to retrieve account information for '%s'@'%s'", name, hostname)
						continue
					}
					log.Printf(types.PrintAccount(accountInfo))
					totalAccountsCount++
				}
			}
			log.Infof("total listed gitlab accounts     : %v", len(config.Global.Platforms.Gitlab.Accounts))
			log.Infof("total accessible gitlab accounts : %v", totalAccountsCount)
		},
	}
	return &cmd
}
