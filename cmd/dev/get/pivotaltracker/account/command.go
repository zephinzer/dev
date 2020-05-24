package account

import (
	"github.com/spf13/cobra"
	cf "github.com/usvc/go-config"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/pivotaltracker"
	"github.com/zephinzer/dev/internal/types"
)

var conf = cf.Map{
	"format": &cf.String{
		Shorthand: "f",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			defaultAccessToken := config.Global.Platforms.PivotalTracker.AccessToken
			if len(config.Global.Platforms.PivotalTracker.Projects) == 0 && len(defaultAccessToken) > 0 {
				accountsEncountered[defaultAccessToken] = true
				log.Info("account information for root credentials")
				account, err := pivotaltracker.GetAccount(defaultAccessToken)
				if err != nil {
					log.Warnf("failed to retrieve account information for pivotal tracker using the default access credentials")
				} else {
					log.Printf(types.PrintAccount(account))
					totalAccountsCount++
				}
			}
			for _, project := range config.Global.Platforms.PivotalTracker.Projects {
				projectAccessToken := project.AccessToken
				if len(projectAccessToken) == 0 && len(defaultAccessToken) > 0 {
					projectAccessToken = defaultAccessToken
				}
				if accountsEncountered[projectAccessToken] == nil {
					accountsEncountered[projectAccessToken] = true
					log.Infof("account information for project '%s' (id: %s)", project.Name, project.ProjectID)
					account, err := pivotaltracker.GetAccount(projectAccessToken)
					if err != nil {
						log.Warnf("failed to retrieve account information for associated pivotal tracker project '%s'", project.Name)
					} else {
						log.Printf(types.PrintAccount(account))
						totalAccountsCount++
					}
				}
			}
			log.Infof("total listed projects: %v\n", len(config.Global.Platforms.PivotalTracker.Projects))
			log.Infof("total accounts: %v\n", totalAccountsCount)
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
