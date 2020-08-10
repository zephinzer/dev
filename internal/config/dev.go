package config

import (
	"github.com/zephinzer/dev/internal/config/dev"
)

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Repository dev.Repository `json:"repository" yaml:"repository,omitempty"`
	Client     dev.Client     `json:"client" yaml:"client,omitempty"`
}

func (d *Dev) MergeWith(o Dev) []error {
	warnings := []error{}
	warnings = append(warnings, d.Client.MergeWith(o.Client)...)
	warnings = append(warnings, d.Repository.MergeWith(o.Repository)...)
	return warnings
}
