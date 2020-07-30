package workspace

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	iworkspace "github.com/zephinzer/dev/internal/workspace"
	"github.com/zephinzer/dev/pkg/repository"
)

var conf = config.Map{
	"format": &config.String{
		Default:   "vscode",
		Shorthand: "f",
		Usage:     "defines the output format, one of [vscode]",
	},
	"output-directory": &config.String{
		Shorthand: "o",
		Usage:     "when defined, writes the workspace file to %this%/%workspace_name%.%format_extension%",
	},
	"overwrite": &config.Bool{
		Shorthand: "O",
		Usage:     "when active, overwrites the workspace file if it exists",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.WorkspaceCanonicalNoun + " [flags] <workspace name>",
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
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}

func run(command *cobra.Command, args []string) {
	targetWorkspaceName := strings.Join(args, ".")
	if len(targetWorkspaceName) == 0 {
		command.Help()
		workspaces := c.Global.Repositories.GetWorkspaces()
		log.Errorf("no target workspace was defined, found workspaces [%v]", strings.Join(workspaces, ", "))
		os.Exit(constants.ExitErrorInput)
		return
	}

	targetWorkspace := iworkspace.Workspace{
		Name:         targetWorkspaceName,
		Repositories: []repository.Repository{},
	}

	for _, repository := range c.Global.Repositories {
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

	switch conf.GetString("format") {
	case "vscode":
		fallthrough
	default:
		vscodeWorkspace, getWorkspaceError := targetWorkspace.ToVSCode()
		if getWorkspaceError != nil {
			log.Errorf("failed to get workspace for vscode: %s", getWorkspaceError)
			os.Exit(constants.ExitErrorInput | constants.ExitErrorApplication)
		}

		vscodeWorkspaceData, toJSONError := vscodeWorkspace.ToJSON()
		if toJSONError != nil {
			log.Errorf("failed to convert vscode struct '%v' to JSON: %s", vscodeWorkspace, toJSONError)
			os.Exit(constants.ExitErrorApplication)
		}

		outputDirectory := conf.GetString("output-directory")
		if len(outputDirectory) == 0 {
			log.Printf(vscodeWorkspaceData)
		} else {
			if outputDirectory[0] == '~' {
				homeDir := cmdutils.GetHomeDirectory()
				outputDirectory = strings.Replace(outputDirectory, "~", homeDir, 1)
			}
			outputPath := path.Join(outputDirectory, strings.ToLower(targetWorkspace.Name)+iworkspace.VSCodeFileExtension)
			if writeError := vscodeWorkspace.WriteTo(outputPath, conf.GetBool("overwrite")); writeError != nil {
				log.Errorf("failed to write the vscode workspace to '%s': %s", outputPath, writeError)
				os.Exit(constants.ExitErrorSystem | constants.ExitErrorUser | constants.ExitErrorApplication)
			}
			log.Infof("successfully wrote workspace '%s' to file at '%s'", targetWorkspace.Name, outputPath)
		}
	}
}
