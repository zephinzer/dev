package pivotaltracker

// Projects defines a structure to store a list of Pivotal Tracker projects
type Projects []Project

// GetSanitized returns a new instance of Projects with the secrets redacted
func (p Projects) GetSanitized() Projects {
	projects := []Project{}
	for _, project := range p {
		if project.Public {
			projects = append(projects, project.GetSanitized())
		}
	}
	return projects
}

// Project defines a Pivotal Tracker project
type Project struct {
	// Name is the custom label given to this project as the user would like to define it
	Name string `json:"name" yaml:"name"`
	// Description is a user-defined description of what this project is for
	Description string `json:"description" yaml:"description"`
	// ProjectID of the project in PivotalTracker that's obtainable by going to a board and getting it from the URL
	ProjectID string `json:"projectID" yaml:"projectID"`
	// AccessToken for this project, if not specified, defaults to use parent access token.
	// This value can be retrieved from the section named API Token at the URL:L
	//   https://www.pivotaltracker.com/profile
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	// Public indicates whether this account should be public, if so, the /platforms
	// endpoint on the dev server will expose this account; this is done to accomodate
	// using both personal and work accounts, note that even if this is set to true,
	// the AccessToken should not be exposed
	Public bool `json:"public" yaml:"public"`
}

// GetSanitized returns a new instance of this Project with the secrets redacted
func (p Project) GetSanitized() Project {
	return Project{
		Name:        p.Name,
		Description: p.Description,
		AccessToken: "[REDACTED]",
		ProjectID:   p.ProjectID,
		Public:      p.Public,
	}
}
