package notifications

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Retrieves notifications from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			totalTodoCount := 0
			for _, account := range config.Global.Platforms.Gitlab.Accounts {
				hostname := "gitlab.com"
				if len(account.Hostname) > 0 {
					hostname = account.Hostname
				}
				if len(account.AccessToken) == 0 {
					log.Warnf("no access token found for %s\n", account.Name)
					break
				}
				todos, err := gitlab.GetTodos(hostname, account.AccessToken)
				if err != nil {
					log.Warnf("an error occurred while retrieving notifications from %s\n", hostname)
					continue
				}
				log.Infof("Notifications from gitlab '%s'@'%s' (total: %v)\n\n", account.Name, account.Hostname, len(todos))
				for index, todo := range todos {
					log.Printf("%v. %s\n%s\n\n- - -\n\n", totalTodoCount+index+1, todo.GetTitle(), todo.GetMessage())
				}
				totalTodoCount += len(todos)
			}
			log.Infof("You have a total of %v unread notifications from your linked gitlab accounts\n\n", totalTodoCount)
		},
	}
	return &cmd
}
