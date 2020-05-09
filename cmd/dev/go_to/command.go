package go_to

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/validator"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GotoCanonicalVerb,
		Aliases: constants.GotoAliases,
		Short:   "go somewhere that's been bookmarked",
		Run:     Run,
	}
	return &cmd
}

func Run(command *cobra.Command, args []string) {
	if len(args) == 0 {
		// links reference
		log.Debug("no arguments received, opening links...")
		startFuzzySearchInterface()
		if selectionIndex >= 0 {
			log.Infof("opening %s", searchResults[selectionIndex].Str)
		}
		return
	}
	argument := strings.Join(args, " ")
	log.Infof("received argument: %s", argument)
	switch true {
	case validator.IsGitHTTPUrl(argument):
		log.Debug("this should be a git http url")
	case validator.IsGitSSHUrl(argument):
		log.Debug("this should be a git ssh url")
	}
	command.Help()
}
