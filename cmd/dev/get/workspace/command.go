package workspace

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	iworkspace "github.com/zephinzer/dev/internal/workspace"
	"github.com/zephinzer/dev/pkg/repository"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.WorkspaceCanonicalNoun,
		Aliases: constants.WorkspaceAliases,
		Short:   "retrieve code that defines a workspace",
		Example: strings.TrimRight(`
  retrieve repositories in the 'dev' workspace

    dev get workspace dev > dev.code-workspace

  Retrieve repositories in the 'ops' workspace

    dev get workspace ops > ops.code-workspace
`, "\n"),
		Run: run,
	}
	return &cmd
}

func run(command *cobra.Command, args []string) {
	targetWorkspaceName := strings.Join(args, ".")
	if len(targetWorkspaceName) == 0 {
		command.Help()
		workspaces := map[string]bool{}
		for _, repository := range config.Global.Repositories {
			for _, workspaceName := range repository.Workspaces {
				workspaces[workspaceName] = true
			}
		}
		validWorkspaces := []string{}
		for workspaceName := range workspaces {
			validWorkspaces = append(validWorkspaces, workspaceName)
		}
		log.Errorf("no target workspace was defined, found workspaces [%v]", strings.Join(validWorkspaces, ", "))
		os.Exit(1)
		return
	}

	targetWorkspace := iworkspace.Workspace{
		Name:         targetWorkspaceName,
		Repositories: []repository.Repository{},
	}

	for _, repository := range config.Global.Repositories {
		log.Tracef("processing repository '%s'...", repository.Name)
		isInWorkspace := false
		for _, workspace := range repository.Workspaces {
			if workspace == targetWorkspaceName {
				isInWorkspace = true
				break
			}
		}
		if !isInWorkspace {
			log.Debugf("skipped repository '%s'", repository.Name)
			continue
		}
		targetWorkspace.Repositories = append(targetWorkspace.Repositories, repository)
	}
	vscodeWorkspace, getWorkspaceError := targetWorkspace.GetVSCode()
	if getWorkspaceError != nil {
		log.Errorf("failed to get code for vscode: %s", getWorkspaceError)
		os.Exit(1)
	}
	log.Print(vscodeWorkspace)
	log.Infof("if you haven't, you can use `%s %s > \"~/%s.code-workspace\"` to place this in your root directory",
		command.CommandPath(),
		targetWorkspaceName,
		targetWorkspaceName,
	)
}
