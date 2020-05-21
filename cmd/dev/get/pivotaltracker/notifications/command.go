package notifications

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Retrieves notifications from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			accessTokensMap := map[string]string{
				config.Global.Platforms.PivotalTracker.AccessToken: "root",
			}
			for _, project := range config.Global.Platforms.PivotalTracker.Projects {
				if len(project.AccessToken) > 0 {
					if _, ok := accessTokensMap[project.AccessToken]; !ok {
						accessTokensMap[project.AccessToken] = project.Name
					}
				}
			}
			accounts := map[string]string{}
			for accessToken, projectName := range accessTokensMap {
				accounts[projectName] = accessToken
			}
			errorCount := 0
			notificationCount := 0
			for name, accessToken := range accounts {
				notifications, getNotificationsError := pivotaltracker.GetNotifications(accessToken, time.Now().Add(-1*time.Hour*24*31))
				if getNotificationsError != nil {
					log.Warnf("failed to get notifications from pivotal tracker project/account '%s': %s", name, getNotificationsError)
					errorCount++
					continue
				}
				log.Infof("Notifications from account/project '%s' (total: %v)\n\n", name, len(notifications))
				for index, notification := range notifications {
					log.Printf("%v. %s\n%s\n\n- - -\n\n", notificationCount+index+1, notification.GetTitle(), notification.GetMessage())
				}
				notificationCount += len(notifications)
			}
			log.Infof("You have a total of %v unread notifications on your linked pivotal accounts", notificationCount)
			os.Exit(errorCount)
		},
	}
	return &cmd
}
