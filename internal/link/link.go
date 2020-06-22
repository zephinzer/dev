package link

// Link represents an entry in the link directory that can be used by a
// user to identify and access a link they're looking for
type Link struct {
	Label      string   `json:"label" yaml:"label,omitempty"`
	Categories []string `json:"categories" yaml:"categories,omitempty,flow"`
	URL        string   `json:"url" yaml:"url,omitempty"`
}

// GetKey returns a (hopefully) unique identifier for this Link instance
// for operations such as de-duplication
func (l Link) GetKey() string {
	return l.URL
}
