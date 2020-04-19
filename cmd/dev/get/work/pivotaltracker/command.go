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
		Use:     "pivotaltracker",
		Aliases: []string{"pivotal", "pt"},
		Short:   "Retrieves your work from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile(constants.DefaultPathToConfiguration)
			if err != nil {
				panic(err)
			}
			if len(c.Platforms.PivotalTracker.AccessToken) == 0 {
				panic("nuuuuuuuuuu")
			}
			totalWorkCount := 0
			for _, project := range c.Platforms.PivotalTracker.Projects {
				stories, err := pivotaltracker.GetStories(
					c.Platforms.PivotalTracker.AccessToken,
					project.ProjectID,
				)
				if err != nil {
					panic(err)
				}
				log.Printf("stories from pivotal tracker project %s\n%s", project.Name, stories.String())
				totalWorkCount += len(*stories)
			}
			log.Printf("you have a total of %v active item(s)", totalWorkCount)
		},
	}
	return &cmd
}
