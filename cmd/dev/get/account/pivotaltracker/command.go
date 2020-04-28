package pivotaltracker

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.PivotalTrackerCanonicalNoun,
		Aliases: constants.PivotalTrackerAliases,
		Short:   "Retrieves account information from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			defaultAccessToken := config.Global.Platforms.PivotalTracker.AccessToken
			if len(config.Global.Platforms.PivotalTracker.Projects) == 0 && len(defaultAccessToken) > 0 {
				accountsEncountered[defaultAccessToken] = true
				printAccountInfo(defaultAccessToken)
				totalAccountsCount++
			}
			for _, project := range config.Global.Platforms.PivotalTracker.Projects {
				projectAccessToken := project.AccessToken
				if len(projectAccessToken) == 0 && len(defaultAccessToken) > 0 {
					projectAccessToken = defaultAccessToken
				}
				if accountsEncountered[projectAccessToken] == nil {
					accountsEncountered[projectAccessToken] = true
					log.Printf("account information for project '%s' (id: %s)\n", project.Name, project.ProjectID)
					printAccountInfo(projectAccessToken)
					totalAccountsCount++
				}
			}
			log.Printf("total listed projects: %v", len(config.Global.Platforms.PivotalTracker.Projects))
			log.Printf("total accounts: %v", totalAccountsCount)
		},
	}
	return &cmd
}

func printAccountInfo(accessToken string) {
	accountInfo, err := pivotaltracker.GetAccount(accessToken)
	if err != nil {
		panic(err)
	}
	log.Println(accountInfo.String())
}
