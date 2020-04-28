package pivotaltracker

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.PivotalTrackerCanonicalNoun,
		Aliases: constants.PivotalTrackerAliases,
		Short:   "Retrieves notifications from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			notifs, err := pivotaltracker.GetNotifs(config.Global.Platforms.PivotalTracker.AccessToken)
			if err != nil {
				panic(err)
			}
			log.Printf("notifications from pivotal tracker\n%s", notifs.String())
		},
	}
	return &cmd
}
