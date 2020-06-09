package link

type Link struct {
	Label      string   `json:"label" yaml:"label,omitempty"`
	Categories []string `json:"categories" yaml:"categories,omitempty"`
	URL        string   `json:"url" yaml:"url,omitempty"`
}
