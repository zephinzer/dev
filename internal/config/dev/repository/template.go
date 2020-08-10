package repository

import (
	"fmt"
	"strings"
)

// Template holds configurations related to repository templates
// that a user can use to initialise repositories on their machine
type Template struct {
	Name string `json:"name" yaml:"name,omitempty"`
	URL  string `json:"url" yaml:"url,omitempty"`
	Path string `json:"path" yaml:"path,omitempty"`
}

// GetKey returns a (hopefully) unique identifer to use for de-duplicating
// multiple instances of Templates
func (drt Template) GetKey() string {
	return fmt.Sprintf("%s-%s", drt.URL, drt.Path)
}

func (drt Template) String() string {
	var templateString strings.Builder
	templateString.WriteString(drt.Name)
	templateString.WriteString(fmt.Sprintf(" (from %s", drt.URL))
	if len(drt.Path) > 0 {
		templateString.WriteString(fmt.Sprintf(" at %s", drt.Path))
	}
	templateString.WriteByte(')')
	return templateString.String()
}
