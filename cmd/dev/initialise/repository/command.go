package repository

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/prompt"
	"github.com/zephinzer/dev/pkg/utils"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     fmt.Sprintf("%s [path]", constants.RepositoryCanonicalNoun),
		Aliases: constants.RepositoryAliases,
		Short:   "Initialises a repository using the template",
		Run:     run,
	}
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	pathToInitialiseAt := path.Join(args...)
	pathToInitialiseAt, resolvePathError := utils.ResolvePath(pathToInitialiseAt)
	if resolvePathError != nil {
		log.Errorf("failed to resolve path '%s': %s", pathToInitialiseAt, resolvePathError)
		os.Exit(constants.ExitErrorInput)
	}
	availableTemplates := []string{}
	for _, template := range c.Global.Dev.Repository.Templates {
		var templateString strings.Builder
		templateString.WriteString(template.Name)
		templateString.WriteString(fmt.Sprintf(" (from %s", template.URL))
		if len(template.Path) > 0 {
			templateString.WriteString(fmt.Sprintf(" at %s", template.Path))
		}
		templateString.WriteByte(')')
		availableTemplates = append(availableTemplates, templateString.String())
	}
	log.Infof("initialising a code repository at '%s'...", pathToInitialiseAt)

	selectedIndex, selectError := prompt.ToSelect("choose a repository template", availableTemplates)
	if selectError != nil {
		os.Exit(constants.ExitErrorInput)
	}
	fmt.Println(selectedIndex)
}
