package link

type Link struct {
	Label      string   `json:"label" yaml:"label,omitempty"`
	Categories []string `json:"categories" yaml:"categories,omitempty,flow"`
	URL        string   `json:"url" yaml:"url,omitempty"`
}
