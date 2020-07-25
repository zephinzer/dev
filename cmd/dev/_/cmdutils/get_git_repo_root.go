package cmdutils

import (
	"fmt"
	"os"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils"
)

// GetGitRepoRootFromWorkingDirectory returns the full path of the child directory
// containing a ./.git directory from the current working directory
func GetGitRepoRootFromWorkingDirectory() string {
	cwd, getCwdErr := os.Getwd()
	if getCwdErr != nil {
		ExitWithError(
			fmt.Sprintf(
				"failed to retrieve current working directory: %s",
				getCwdErr,
			),
			constants.ExitErrorSystem,
		)
	}
	gitRepoRoot, findGitRepoRootError := utils.FindParentContainingChildDirectory(".git", cwd)
	if findGitRepoRootError != nil {
		ExitWithError(
			fmt.Sprintf(
				"failed to detect if current directory resides in a git repository: %s",
				findGitRepoRootError,
			),
			constants.ExitErrorSystem,
		)
	} else if len(gitRepoRoot) == 0 {
		ExitWithError(
			fmt.Sprintf(
				"current directory does not seem to reside in a git repository: %s",
				findGitRepoRootError,
			),
			constants.ExitErrorUser,
		)
	}
	return gitRepoRoot
}
