package pivotaltracker

// Config defines the structure of the configuration that can be consumed for tracking
// notifications and events from the Pivotal Tracker platform
type Config struct {
	// Projects is the list of projects to monitor
	Projects Projects `yaml:"projects"`
	// AccessToken to be used for all listed projects if their own access token is not defined.
	// This value can be retrieved from the section named API Token at the URL:
	//   https://www.pivotaltracker.com/profile
	AccessToken string `yaml:"accessToken"`
}

// Projects defines a structure to store a list of Pivotal Tracker projects
type Projects []Project

// Project defines a Pivotal Tracker project
type Project struct {
	// Name is the custom label given to this project as the user would like to define it
	Name string `yaml:"name"`
	// ProjectID of the project in PivotalTracker that's obtainable by going to a board and getting it from the URL
	ProjectID string `yaml:"projectID"`
	// AccessToken for this project, if not specified, defaults to use parent access token.
	// This value can be retrieved from the section named API Token at the URL:L
	//   https://www.pivotaltracker.com/profile
	AccessToken string `yaml:"accessToken"`
}
