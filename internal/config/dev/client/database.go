package client

import "errors"

// Database holds configurations related to the data persistence
// mechanism of the CLI tool
type Database struct {
	Path string `json:"path" yaml:"path,omitempty"`
}

func (dcdb *Database) MergeWith(o Database) []error {
	if len(dcdb.Path) > 0 {
		return []error{errors.New("dev.client.database.path already set")}
	}
	dcdb.Path = o.Path
	return nil
}
