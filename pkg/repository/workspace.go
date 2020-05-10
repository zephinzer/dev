package repository

// Repository represents a code repo
type Repository struct {
	Name        string   `json:"name" yaml:"name"`
	Workspaces  []string `json:"workspaces" yaml:"workspaces"`
	Description string   `json:"description" yaml:"description"`
	CloneURL    string   `json:"cloneURL" yaml:"cloneURL"`
	Path        string   `json:"path" yaml:"path"`
}
