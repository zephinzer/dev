package main

import (
	"time"

	"github.com/zephinzer/dev/internal/constants"
)

var (
	// Commit will be set to the commit hash during build time
	Commit = "<commit-hash>"
	// Version will be set to the semantic version during build time
	Version = "<semver-version>"
	// Timestamp will be set to the timestamp of the build during build time
	Timestamp = time.Now().Format(constants.DevTimeFormat)
)
