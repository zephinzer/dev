package config

import "github.com/zephinzer/dev/internal/link"

// Links represents a list of links which the uesr should be able to
// access
type Links []link.Link

// MergeWith merges the current Links instance with a provided
// Links instance. The merge strategy is add-only
func (l *Links) MergeWith(o Links) {
	seen := map[string]bool{}
	for _, link := range *l {
		seen[link.GetKey()] = true
	}
	for _, link := range o {
		if val, ok := seen[link.GetKey()]; val && ok {
			continue
		}
		*l = append(*l, link)
		seen[link.GetKey()] = true
	}
}
