package notifications

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/github"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Retrieves notifications from Github",
		Run: func(command *cobra.Command, args []string) {
			totalNotificationsCount := 0
			log.Infof("retrieving notifications from %v linked github account(s)...", len(config.Global.Platforms.Github.Accounts))
			for _, account := range config.Global.Platforms.Github.Accounts {
				if len(account.AccessToken) == 0 {
					log.Warnf("no access token found for %s\n", account.Name)
					break
				}
				notifs, getNotificationsError := github.GetNotifications(
					&http.Client{Timeout: constants.DefaultAPICallTimeout},
					account.AccessToken,
				)
				if getNotificationsError != nil {
					log.Warnf("an error occurred while retrieving notifications from '%s': %s\n", account.Name, getNotificationsError)
					continue
				}
				log.Infof("notifications from github account '%s' (total: %v)\n\n", account.Name, len(notifs))
				for index, notif := range notifs {
					log.Printf("%v. %s\n%s\n\n- - -\n\n", totalNotificationsCount+index+1, notif.GetTitle(), notif.GetMessage())
				}
				totalNotificationsCount += len(notifs)
			}
			log.Infof("you have a total of %v unread notifications from your linked github accounts\n\n", totalNotificationsCount)
		},
	}
	return &cmd
}
