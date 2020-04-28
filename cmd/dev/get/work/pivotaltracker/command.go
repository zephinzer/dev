package pivotaltracker

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.PivotalTrackerCanonicalNoun,
		Aliases: constants.PivotalTrackerAliases,
		Short:   "Retrieves your work from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			totalWorkCount := 0
			for _, project := range config.Global.Platforms.PivotalTracker.Projects {
				stories, err := pivotaltracker.GetStories(
					config.Global.Platforms.PivotalTracker.AccessToken,
					project.ProjectID,
				)
				if err != nil {
					panic(err)
				}
				log.Printf("stories from pivotal tracker project [%s] (count: %v)\n%s", project.Name, len(*stories), stories.String())
				totalWorkCount += len(*stories)
			}
			log.Printf("you have a total of %v active item(s)", totalWorkCount)
		},
	}
	return &cmd
}
