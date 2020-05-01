package trello

// Config defines the structure of the configuration for the Trello integration
type Config struct {
	// Boards is the list of boards to monitor
	Boards Boards `json:"boards" yaml:"boards"`
	// AccessKey to be used for all listed boards if their own access token is not defined.
	// This value can be retrieved from the page at https://trello.com/app-key
	AccessKey string `json:"accessKey" yaml:"accessKey"`
	// AccessToken to be used for all listed boards if their own access token is not defined.
	// This value can be retrieved by clicking on Token on the page at https://trello.com/app-key
	AccessToken string `json:"accessToken" yaml:"accessToken"`
}

type Boards []Board

type Board struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	AccessKey   string `json:"accessKey" yaml:"accessKey"`
	AccessToken string `json:"accessToken" yaml:"accessToken"`
}
