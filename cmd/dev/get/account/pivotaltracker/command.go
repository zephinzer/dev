package pivotaltracker

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/pivotaltracker"
	cf "github.com/usvc/go-config"
)

var conf = cf.Map{
	"format": &cf.String{
		Shorthand: "f",
	},
}

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
				printAccountInfo(defaultAccessToken, conf.GetString("format"))
				totalAccountsCount++
			}
			for _, project := range config.Global.Platforms.PivotalTracker.Projects {
				projectAccessToken := project.AccessToken
				if len(projectAccessToken) == 0 && len(defaultAccessToken) > 0 {
					projectAccessToken = defaultAccessToken
				}
				if accountsEncountered[projectAccessToken] == nil {
					accountsEncountered[projectAccessToken] = true
					log.Printf("\n# account information for project '%s' (id: %s)\n\n", project.Name, project.ProjectID)
					printAccountInfo(projectAccessToken, conf.GetString("format"))
					totalAccountsCount++
				}
			}
			log.Printf("> total listed projects: %v\n", len(config.Global.Platforms.PivotalTracker.Projects))
			log.Printf("> total accounts: %v\n", totalAccountsCount)
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}

func printAccountInfo(accessToken string, format ...string) {
	accountInfo, err := pivotaltracker.GetAccount(accessToken)
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", accountInfo.String(format...))
}
