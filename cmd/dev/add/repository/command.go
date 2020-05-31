package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/validator"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     fmt.Sprintf("%s [flags] <repository_urls...>", constants.RepositoryCanonicalNoun),
		Aliases: constants.RepositoryAliases,
		Short:   "add a repository to your configuration",
		Run:     run,
	}
	return &cmd
}

func run(command *cobra.Command, args []string) {
	if len(args) == 0 {
		command.Help()
		log.Errorf("no repository url specified, see usage above for how to use this command")
		os.Exit(1)
		return
	}
	repositoryURLs := []string{}
	for _, arg := range args {
		log.Tracef("parsing url '%s'...", arg)
		repositoryURL, parseURLError := validator.ParseURL(arg)
		if parseURLError != nil {
			log.Warnf("failed to parse url '%s': %s", arg, parseURLError)
			continue
		}
		repositoryURLs = append(repositoryURLs, repositoryURL.String())
	}
	var repositoriesToWorkOn strings.Builder
	repositoriesToWorkOn.WriteString("repositories to add:\n")
	for _, repositoryURL := range repositoryURLs {
		repositoriesToWorkOn.WriteString(fmt.Sprintf("- %s\n", repositoryURL))
	}
	log.Info(repositoriesToWorkOn.String())
}
