package config

import "github.com/zephinzer/dev/pkg/repository"

// Repositories represents a list of repositories a user should have
// access to on their machines
type Repositories []repository.Repository

// GetWorkspaces returns a list of strings corresponding to all the workspaces
// listed in this instance of Repositories
func (r Repositories) GetWorkspaces() []string {
	workspacesMap := map[string]bool{}
	for _, repository := range r {
		for _, workspace := range repository.Workspaces {
			workspacesMap[workspace] = true
		}
	}
	workspaces := []string{}
	for workspaceName := range workspacesMap {
		workspaces = append(workspaces, workspaceName)
	}
	return workspaces
}

// MergeWith merges the current Repositories instance with a provided
// Repositories instance. The merge strategy is add-only
func (r *Repositories) MergeWith(o Repositories) {
	seen := map[string]bool{}
	for _, repository := range *r {
		repositoryPath, _ := repository.GetPath()
		seen[repositoryPath] = true
	}
	for _, repository := range o {
		repositoryPath, _ := repository.GetPath()
		if value, ok := seen[repositoryPath]; ok && value {
			continue
		}
		*r = append(*r, repository)
		seen[repositoryPath] = true
	}
}
