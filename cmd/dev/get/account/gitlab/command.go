package gitlab

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/gitlab"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GitlabCanonicalNoun,
		Aliases: constants.GitlabAliases,
		Short:   "Retrieves account information from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile("./dev.yaml")
			if err != nil {
				panic(err)
			}
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			for _, account := range c.Platforms.Gitlab.Accounts {
				accountAccessToken := account.AccessToken
				accountHostname := account.Hostname
				if len(account.Hostname) == 0 {
					accountHostname = constants.DefaultGitlabHostname
				}
				if len(accountAccessToken) == 0 {

					log.Println("skipping '%s'@'%s': access token was not specified", account.Name, accountHostname)
					continue
				}
				if accountsEncountered[accountAccessToken] == nil {
					accountsEncountered[accountAccessToken] = true
					log.Printf("account information for '%s'@'%s'\n", account.Name, accountHostname)
					printAccountInfo(accountHostname, accountAccessToken)
					totalAccountsCount++
				}
			}
			log.Printf("total listed projects    : %v", len(c.Platforms.Gitlab.Accounts))
			log.Printf("total accessible accounts: %v", totalAccountsCount)
		},
	}
	return &cmd
}

func printAccountInfo(hostname, accessToken string) {
	accountInfo, err := gitlab.GetAccount(hostname, accessToken)
	if err != nil {
		panic(err)
	}
	// litter.Dump(accountInfo)
	log.Println(accountInfo.String())
}
