package merge_request

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/git"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils/str"
	"github.com/zephinzer/dev/pkg/utils/system"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.MergeRequestNoun,
		Short:   "opens the browser to a merge/pull request at this repository's url",
		Aliases: constants.MergeRequestAliases,
		RunE: func(command *cobra.Command, args []string) error {
			gitRepoRoot := cmdutils.GetGitRepoRootFromWorkingDirectory()
			branchName, err := git.GetCurrentBranch(gitRepoRoot)
			if err != nil {
				return fmt.Errorf("failed to get repository info: %s", err)
			}
			remote, err := git.GetRemote(gitRepoRoot)
			if err != nil {
				return fmt.Errorf("failed to get repository info: %s", err)
			}
			gitUrl, err := str.ParseGitUrl(remote.URL)
			if err != nil {
				return fmt.Errorf("failed to get repository url: %s", err)
			}
			newMergeRequestUrl := fmt.Sprintf("%s/-/merge_requests/new?merge_request%%5Bsource_branch%%5D=%s", gitUrl.GetBrowserUrl(), url.QueryEscape(branchName))

			log.Infof("opening url '%s' in the default browser application...", newMergeRequestUrl)
			system.OpenURIWithDefaultApplication(newMergeRequestUrl)
			return nil
		},
	}
	return &cmd
}
