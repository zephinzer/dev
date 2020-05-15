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
		Short:   "Retrieves notifications from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			log.Printf("# notifications from gitlab\n\n")
			totalTodoCount := 0
			for _, account := range config.Global.Platforms.Gitlab.Accounts {
				hostname := "gitlab.com"
				if len(account.Hostname) > 0 {
					hostname = account.Hostname
				}
				if len(account.AccessToken) == 0 {
					log.Printf("no access token found for %s\n", account.Name)
					break
				}
				log.Printf("## notifications from %s\n\n", account.Name)
				todos, err := gitlab.GetTodos(hostname, account.AccessToken)
				if err != nil {
					log.Errorf("an error occurred while retrieving notifications from %s\n", hostname)
				} else {
					log.Printf("%s", todos.String())
				}
				todoCount := 0
				if todos != nil {
					todoCount = len(*todos)
				}
				totalTodoCount += todoCount
				log.Printf("> you have a total of %v unread notifications from %s (%s)\n\n", todoCount, account.Name, hostname)
			}
			log.Printf("> you have a total of %v unread notifications from your gitlab accounts\n\n", totalTodoCount)
		},
	}
	return &cmd
}
