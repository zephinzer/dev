package config

import "github.com/zephinzer/dev/pkg/network"

// Networks represents a list of networks that the user's machine should
// be able to connect to
type Networks []network.Network

// MergeWith merges the current Networks instance with a provided
// Networks instance. The merge strategy is add-only
func (n *Networks) MergeWith(o Networks) {
	seen := map[string]bool{}
	for _, nw := range *n {
		seen[nw.Check.URL] = true
	}
	for _, nw := range o {
		if seen[nw.Check.URL] == true {
			continue
		}
		*n = append(*n, nw)
		seen[nw.Check.URL] = true
	}
}
