package pivotaltracker

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/pkg/config"
	"github.com/usvc/dev/pkg/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "pivotaltracker",
		Aliases: []string{"pivotal", "pt"},
		Short:   "Retrieves account information from Pivotal Tracker",
		Run: func(command *cobra.Command, args []string) {
			c, err := config.NewFromFile("./dev.yaml")
			if err != nil {
				panic(err)
			}
			if len(c.Platforms.PivotalTracker.AccessToken) == 0 {
				panic("nuuuuuuuuuu")
			}
			user, err := pivotaltracker.GetAccount(c.Platforms.PivotalTracker.AccessToken)
			if err != nil {
				panic(err)
			}
			litter.Dump(user)
		},
	}
	return &cmd
}
