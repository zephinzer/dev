package notifications

import (
	"github.com/spf13/cobra"
	github "github.com/zephinzer/dev/cmd/dev/get/github/notifications"
	gitlab "github.com/zephinzer/dev/cmd/dev/get/gitlab/notifications"
	pivotaltracker "github.com/zephinzer/dev/cmd/dev/get/pivotaltracker/notifications"
	trello "github.com/zephinzer/dev/cmd/dev/get/trello/notifications"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Retrieves notifications",
		Run: func(command *cobra.Command, args []string) {
			if len(args) == 0 {
				github.GetCommand().Run(command, args)
				gitlab.GetCommand().Run(command, args)
				pivotaltracker.GetCommand().Run(command, args)
				trello.GetCommand().Run(command, args)
			}
			command.Help()
		},
	}

	gitlabCommand := gitlab.GetCommand()
	gitlabCommand.Use = constants.GitlabCanonicalNoun
	gitlabCommand.Aliases = constants.GitlabAliases
	cmd.AddCommand(gitlabCommand)

	githubCommand := github.GetCommand()
	githubCommand.Use = constants.GithubCanonicalNoun
	githubCommand.Aliases = constants.GithubAliases
	cmd.AddCommand(githubCommand)

	pivotaltrackerCommand := pivotaltracker.GetCommand()
	pivotaltrackerCommand.Use = constants.PivotalTrackerCanonicalNoun
	pivotaltrackerCommand.Aliases = constants.PivotalTrackerAliases
	cmd.AddCommand(pivotaltrackerCommand)

	trelloCommand := trello.GetCommand()
	trelloCommand.Use = constants.TrelloCanonicalNoun
	trelloCommand.Aliases = constants.TrelloAliases
	cmd.AddCommand(trelloCommand)

	return &cmd
}
