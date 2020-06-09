package main

import (
	configuration "github.com/usvc/go-config"
	"github.com/zephinzer/dev/internal/constants"
)

var (
	// conf is the local configuration for the root command
	conf = configuration.Map{
		"db-path": &configuration.String{
			Default: constants.DefaultPathToSQLite3DB,
			Shorthand: "D",
			Usage: "path to the sqlite3 database to use",
		},
		"debug": &configuration.Bool{
			Usage: "display up to debug level logs (verbose logging mode)",
		},
		"trace": &configuration.Bool{
			Usage: "display up to trace level logs (very-verbose logging mode)",
		},
	}
)
