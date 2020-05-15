package database

import (
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/db"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.DatabaseCanonicalNoun,
		Aliases: constants.DatabaseAliases,
		Short:   "Initialises a persistent on-disk sqlite3 database",
		Run:     run,
	}
	return &cmd
}

func run(command *cobra.Command, args []string) {
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
	log.Printf("initialising database at path '%s'...\n", pathToDatabaseFile)
	dbInitError := db.Init(pathToDatabaseFile)
	if dbInitError != nil {
		log.Errorf("failed to initialise database at '%s': %s", pathToDatabaseFile, dbInitError)
		os.Exit(1)
	}
	log.Printf("initialised database at '%s'\n", pathToDatabaseFile)
}
