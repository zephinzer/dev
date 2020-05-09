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

// GetSanitized returns a sanitised copy of the Config instance
func (c Config) GetSanitized() Config {
	return Config{
		AccessKey:   "[REDACTED]",
		AccessToken: "[REDACTED]",
		Boards:      c.Boards.GetSanitized(),
	}
}

// Boards is a slice of Board instances
type Boards []Board

// GetSanitized returns a sanitised copy of the Boards instance
func (b Boards) GetSanitized() Boards {
	boards := []Board{}
	for _, board := range b {
		if board.Public {
			boards = append(boards, board.GetSanitized())
		}
	}
	return boards
}

// Board represents a Trello board
type Board struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	AccessKey   string `json:"accessKey" yaml:"accessKey"`
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	Public      bool   `json:"public" yaml:"public"`
}

// GetSanitized returns a sanitised copy of the Board instance
func (b Board) GetSanitized() Board {
	return Board{
		ID:          b.ID,
		Name:        b.Name,
		AccessKey:   "[REDACTED]",
		AccessToken: "[REDACTED]",
		Public:      b.Public,
	}
}
