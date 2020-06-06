package config

import (
	"github.com/zephinzer/dev/pkg/software"
)

// Softwares represents a list of software the user should have on
// their machine
type Softwares []software.Software

// MergeWith merges the current Softwares instance with a provided
// Softwares instance. The merge strategy is add-only
func (s *Softwares) MergeWith(o Softwares) {
	seen := map[string]bool{}
	for _, sw := range *s {
		seen[sw.Check.Command[0]] = true
	}
	for _, sw := range o {
		if value, ok := seen[sw.Check.Command[0]]; ok && value {
			continue
		}
		*s = append(*s, sw)
		seen[sw.Check.Command[0]] = true
	}
}
