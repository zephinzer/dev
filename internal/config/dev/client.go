package dev

import (
	"github.com/zephinzer/dev/internal/config/dev/client"
)

// Client holds configurations related to the CLI tool
type Client struct {
	Database      client.Database      `json:"database" yaml:"database,omitempty"`
	Notifications client.Notifications `json:"notifications" yaml:"notifications,omitempty"`
	Platforms     client.Platforms     `json:"platforms" yaml:"platforms,omitempty"`
}

func (dc *Client) MergeWith(o Client) []error {
	warnings := []error{}
	warnings = append(warnings, dc.Database.MergeWith(o.Database)...)
	warnings = append(warnings, dc.Notifications.MergeWith(o.Notifications)...)
	warnings = append(warnings, dc.Platforms.MergeWith(o.Platforms)...)
	return warnings
}
