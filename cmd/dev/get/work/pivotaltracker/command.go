package pivotaltracker

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/pivotaltracker"
	cf "github.com/usvc/go-config"
)

var conf = cf.Map{
	"format": &cf.String{
		Shorthand: "f",
	},
}

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
				log.Printf("\n# stories from pivotal tracker project [%s] (count: %v)\n\n", project.Name, len(*stories))
				log.Printf("%s", stories.String(conf.GetString("format")))
				totalWorkCount += len(*stories)
			}
			log.Printf("> you have a total of %v item(s)", totalWorkCount)
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
