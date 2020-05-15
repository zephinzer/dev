package database

import (
	"log"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/db"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.DatabaseCanonicalNoun,
		Aliases: constants.DatabaseAliases,
		Short:   "Initialises a persistent on-disk sqlite3 database",
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
			log.Printf("initialising database at %s...", pathToDatabaseFile)
			dbInitError := db.Init(pathToDatabaseFile)
			if dbInitError != nil {
				panic(dbInitError)
			}
			log.Printf("initialised database at %s", pathToDatabaseFile)
		},
	}
	return &cmd
}
