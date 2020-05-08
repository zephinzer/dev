package link

type Link struct {
	Label      string   `json:"label" yaml:"label"`
	Categories []string `json:"categories" yaml:"categories"`
	URL        string   `json:"url" yaml:"url"`
}
