package go_to

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
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
	log.Debug("no arguments received, opening links...")
	startFuzzySearchInterface()
	if selectionIndex >= 0 {
		log.Infof("opening %s", searchResults[selectionIndex].Str)
	}
	command.Help()
}
