package workspace

// Workspace represents a project-scoped set of code repos
type Workspace struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}
