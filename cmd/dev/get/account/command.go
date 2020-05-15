package account

import (
	"github.com/spf13/cobra"
	github "github.com/zephinzer/dev/cmd/dev/get/github/account"
	gitlab "github.com/zephinzer/dev/cmd/dev/get/gitlab/account"
	pivotaltracker "github.com/zephinzer/dev/cmd/dev/get/pivotaltracker/account"
	trello "github.com/zephinzer/dev/cmd/dev/get/trello/account"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information",
		Run: func(command *cobra.Command, args []string) {
			if len(args) == 0 {
				github.GetCommand().Run(command, args)
				gitlab.GetCommand().Run(command, args)
				pivotaltracker.GetCommand().Run(command, args)
				trello.GetCommand().Run(command, args)
				return
			}
			command.Help()
		},
	}

	gitlabAccount := gitlab.GetCommand()
	gitlabAccount.Use = constants.GitlabCanonicalNoun
	gitlabAccount.Aliases = constants.GitlabAliases
	cmd.AddCommand(gitlabAccount)

	githubAccount := github.GetCommand()
	githubAccount.Use = constants.GithubCanonicalNoun
	githubAccount.Aliases = constants.GithubAliases
	cmd.AddCommand(githubAccount)

	pivotaltrackerAccount := pivotaltracker.GetCommand()
	pivotaltrackerAccount.Use = constants.PivotalTrackerCanonicalNoun
	pivotaltrackerAccount.Aliases = constants.PivotalTrackerAliases
	cmd.AddCommand(pivotaltrackerAccount)

	trelloAccount := trello.GetCommand()
	trelloAccount.Use = constants.TrelloCanonicalNoun
	trelloAccount.Aliases = constants.TrelloAliases
	cmd.AddCommand(trelloAccount)

	return &cmd
}
