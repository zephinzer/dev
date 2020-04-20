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
		Use:     "gitlab",
		Aliases: []string{"gl"},
		Short:   "Retrieves notifications from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile(constants.DefaultPathToConfiguration)
			if err != nil {
				panic(err)
			}
			totalTodoCount := 0
			for _, account := range c.Platforms.Gitlab.Accounts {
				hostname := "gitlab.com"
				if len(account.Hostname) > 0 {
					hostname = account.Hostname
				}
				if len(account.AccessToken) == 0 {
					log.Printf("no access token found for %s", account.Name)
					break
				}
				todos, err := gitlab.GetTodos(hostname, account.AccessToken)
				if err != nil {
					panic(err)
				}
				totalTodoCount += len(*todos)
				log.Printf("todos from %s\n%s", account.Name, todos.String())
			}
			log.Printf("total todos from gitlab: %v", totalTodoCount)
		},
	}
	return &cmd
}
