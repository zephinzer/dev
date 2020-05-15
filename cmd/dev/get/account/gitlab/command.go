package gitlab

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/gitlab"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GitlabCanonicalNoun,
		Aliases: constants.GitlabAliases,
		Short:   "Retrieves account information from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			for _, account := range config.Global.Platforms.Gitlab.Accounts {
				name := account.Name
				if len(name) == 0 {
					name = "unnamed"
				}
				hostname := account.Hostname
				if len(account.Hostname) == 0 {
					hostname = constants.DefaultGitlabHostname
				}
				accessToken := account.AccessToken
				if len(accessToken) == 0 {
					log.Infof("skipping '%s'@'%s': access token was not specified", name, hostname)
					continue
				}

				if accountsEncountered[accessToken] == nil {
					accountsEncountered[accessToken] = true
					log.Infof("account information for '%s'@'%s'\n", name, hostname)
					accountInfo, err := gitlab.GetAccount(hostname, accessToken)
					if err != nil {
						log.Warnf("unable to retrieve account information for '%s'@'%s'", name, hostname)
						continue
					}
					log.Infof(accountInfo.String())
					totalAccountsCount++
				}
			}
			log.Infof("total listed accounts     : %v", len(config.Global.Platforms.Gitlab.Accounts))
			log.Infof("total accessible accounts : %v", totalAccountsCount)
		},
	}
	return &cmd
}

func printAccountInfo(hostname, accessToken string) error {
	accountInfo, err := gitlab.GetAccount(hostname, accessToken)
	if err != nil {
		return err
	}
	log.Infof(accountInfo.String())
	return nil
}
