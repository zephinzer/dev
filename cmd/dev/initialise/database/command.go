package database

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/db"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/pivotaltracker"
	"github.com/zephinzer/dev/pkg/utils"
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
	var databasePath string
	databasePath, _ = command.Flags().GetString("db-path")
	if len(config.Global.Dev.Client.Database.Path) > 0 {
		databasePath = config.Global.Dev.Client.Database.Path
	}
	absolutePath, resolvePathError := utils.ResolvePath(databasePath)
	if resolvePathError != nil {
		log.Errorf("failed to resolve database path '%s': %s", databasePath, resolvePathError)
		os.Exit(1)
		return
	}

	log.Debugf("checking database at path '%s'...\n", absolutePath)
	dbCheckError := db.Check(absolutePath)
	if dbCheckError != nil {
		log.Warnf("failed to check the database at path '%s': %s", absolutePath, dbCheckError)
		dbInitError := db.Init(absolutePath)
		if dbInitError != nil {
			log.Errorf("failed to initialise database at '%s': %s", absolutePath, dbInitError)
			os.Exit(1)
		}
		log.Debugf("a new database was initialised at path '%s'", absolutePath)
	} else {
		log.Debugf("database checks for path '%s' succeeded", absolutePath)
	}
	log.Infof("database is initialised at '%s'\n", absolutePath)

	log.Debug("initialising tables for gitlab data storage...")
	if gitlabInitError := gitlab.InitSQLite3Database(absolutePath); gitlabInitError != nil {
		log.Errorf("migration failed: %s", gitlabInitError)
		os.Exit(1)
	}
	log.Info("tables for gitlab data storage have been initialised")

	log.Debug("initialising tables for pivotal tracker data storage...")
	if pivotalInitError := pivotaltracker.InitSQLite3Database(absolutePath); pivotalInitError != nil {
		log.Errorf("migration failed: %s", pivotalInitError)
		os.Exit(1)
	}
	log.Info("tables for pivotal tracker data storage have been initialised")
}
