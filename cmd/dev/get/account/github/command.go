package github

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/github"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GithubCanonicalNoun,
		Aliases: constants.GithubAliases,
		Short:   "Retrieves account information from Github",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile("./dev.yaml")
			if err != nil {
				panic(err)
			}
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			for _, account := range c.Platforms.Github.Accounts {
				accountAccessToken := account.AccessToken
				if len(accountAccessToken) == 0 {
					log.Println("skipping '%s': access token was not specified", account.Name)
					continue
				}
				if accountsEncountered[accountAccessToken] == nil {
					accountsEncountered[accountAccessToken] = true
					log.Printf("account information for '%s'\n", account.Name)
					printAccountInfo(accountAccessToken)
					totalAccountsCount++
				}
			}
			log.Printf("total listed projects    : %v", len(c.Platforms.Github.Accounts))
			log.Printf("total accessible accounts: %v", totalAccountsCount)
		},
	}
	return &cmd
}

func printAccountInfo(accessToken string) {
	accountInfo, err := github.GetAccount(accessToken)
	if err != nil {
		panic(err)
	}
	// litter.Dump(accountInfo)
	log.Println(accountInfo.String())
}
