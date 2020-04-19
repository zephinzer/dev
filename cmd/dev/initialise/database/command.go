package database

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/db"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "database",
		Aliases: []string{"db"},
		Short:   "Initialises a persistent on-disk sqlite3 database",
		Run: func(command *cobra.Command, args []string) {
			dbInitError := db.Init(constants.DefaultPathToSQLite3DB)
			if dbInitError != nil {
				panic(dbInitError)
			}
		},
	}
	return &cmd
}
