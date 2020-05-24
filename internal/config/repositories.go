package config

import "github.com/zephinzer/dev/pkg/repository"

// Repositories represents a list of repositories a user should have
// access to on their machines
type Repositories []repository.Repository

// MergeWith merges the current Repositories instance with a provided
// Repositories instance. The merge strategy is add-only
func (r *Repositories) MergeWith(o Repositories) {
	seen := map[string]bool{}
	for _, rp := range *r {
		repoPath, _ := rp.GetPath()
		seen[repoPath] = true
	}
	for _, rp := range o {
		repoPath, _ := rp.GetPath()
		if value, ok := seen[repoPath]; ok && value {
			continue
		}
		*r = append(*r, rp)
		seen[repoPath] = true
	}
}
