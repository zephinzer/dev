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
			c, err := config.NewFromFile(constants.DefaultPathToConfiguration)
			if err != nil {
				panic(err)
			}
			if len(c.Platforms.PivotalTracker.AccessToken) == 0 {
				panic("nuuuuuuuuuu")
			}
			notifs, err := pivotaltracker.GetNotifs(c.Platforms.PivotalTracker.AccessToken)
			if err != nil {
				panic(err)
			}
			log.Printf("notifications from pivotal tracker\n%s", notifs.String())
		},
	}
	return &cmd
}
