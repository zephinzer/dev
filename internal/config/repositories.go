package config

import (
	"sort"
	"strings"

	"github.com/zephinzer/dev/pkg/repository"
)

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

// Len implements `sort.Interface`
func (r Repositories) Len() int {
	return len(r)
}

// Less implements `sort.Interface`
func (r Repositories) Less(i, j int) bool {
	if len(r[i].Name) > 0 && len(r[j].Name) > 0 {
		return strings.Compare(r[i].Name, r[j].Name) <= 0
	}
	if len(r[i].URL) > 0 && len(r[j].URL) > 0 {
		iPath, getPathError := r[i].GetPath()
		if getPathError != nil {
			return false
		}
		jPath, getPathError := r[j].GetPath()
		if getPathError != nil {
			return false
		}
		return strings.Compare(iPath, jPath) <= 0
	}
	return false
}

// Swap implements `sort.Interface`
func (r Repositories) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Sort sorts the repositories in alphabetical order
func (r *Repositories) Sort() {
	sort.Sort(r)
}
