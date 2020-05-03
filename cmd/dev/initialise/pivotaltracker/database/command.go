package database

import (
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/internal/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.DatabaseCanonicalNoun,
		Aliases: constants.DatabaseAliases,
		Short:   "Initialises tables for Pivotal Tracker stuff in the local SQLite3 database",
		Run: func(command *cobra.Command, args []string) {
			pathToDatabaseFile := constants.DefaultPathToSQLite3DB
			if len(config.Global.Dev.Client.Database.Path) > 0 {
				pathToDatabaseFile = config.Global.Dev.Client.Database.Path
			}
			if strings.Index(pathToDatabaseFile, "~") == 0 {
				currentUser, err := user.Current()
				if err != nil {
					panic(err)
				} else if len(currentUser.HomeDir) > 0 {
					pathToDatabaseFile = strings.Replace(pathToDatabaseFile, "~", currentUser.HomeDir, 1)
				}
			}
			log.Debugf("initialising database at %s...", pathToDatabaseFile)
			if initError := pivotaltracker.InitSQLite3Database(pathToDatabaseFile); initError != nil {
				log.Errorf("migration failed: %s", initError)
				os.Exit(1)
			}
			log.Tracef("initialised database at %s", pathToDatabaseFile)
		},
	}
	return &cmd
}
