package pivotaltracker

type APILabel struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// APIAccount defines the structure for an Account object in responses to API queries
type APIAccount struct {
	Kind   string `json:"kind"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Plan   string `json:"plan"`
}

// APIProject defines the structure for a project object in a response to API queries
type APIProject struct {
	Kind         string `json:"kind"`
	ID           int    `json:"id"`
	ProjectID    int    `json:"project_id"`
	ProjectName  string `json:"project_name"`
	ProjectColor string `json:"project_color"`
	Favorite     bool   `json:"favorite"`
	Role         string `json:"role"`
	LastViewedAt string `json:"last_viewed_at"`
}

type APITimezone struct {
	Kind      string `json:"kind"`
	OlsonName string `json:"olson_name"`
	Offset    string `json:"offset"`
}
