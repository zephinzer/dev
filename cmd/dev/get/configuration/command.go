package configuration

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
)

var conf = config.Map{
	"links": &config.Bool{
		Usage: "retrieves the links configuration",
	},
	"networks": &config.Bool{
		Usage: "retrieves the networks configuration",
	},
	"platforms": &config.Bool{
		Usage: "retrieves the platforms configuration",
	},
	"repositories": &config.Bool{
		Usage: "retrieves the repositories configuration",
	},
	"software": &config.Bool{
		Usage: "retrieves the software configuration",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ConfigurationCanonicalNoun,
		Aliases: constants.ConfigurationAliases,
		Short:   "retrieves the consumed configuration",
		Run: func(command *cobra.Command, args []string) {
			switch true {
			case conf.GetBool("links"):
				litter.Dump(c.Global.Links)
			case conf.GetBool("network"):
				litter.Dump(c.Global.Networks)
			case conf.GetBool("platforms"):
				litter.Dump(c.Global.Platforms)
			case conf.GetBool("repositories"):
				litter.Dump(c.Global.Repositories)
			case conf.GetBool("software"):
				litter.Dump(c.Global.Softwares)
			default:
				litter.Dump(c.Global)
			}
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
