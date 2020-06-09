package configuration

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"gopkg.in/yaml.v2"
)

var conf = config.Map{
	"dev": &config.Bool{
		Usage: "retrieves the dev configuration",
	},
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
	"softwares": &config.Bool{
		Usage: "retrieves the software configuration",
	},
	"format": &config.String{
		Default: "yaml",
		Usage:   "specify the format of the output, one of [raw, yaml]",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ConfigurationCanonicalNoun,
		Aliases: constants.ConfigurationAliases,
		Short:   "Retrieves the consumed configuration",
		Run: func(command *cobra.Command, args []string) {
			format := conf.GetString("format")
			var output interface{}
			switch true {
			case conf.GetBool("dev"):
				output = c.Global.Dev
			case conf.GetBool("links"):
				output = c.Global.Links
			case conf.GetBool("networks"):
				output = c.Global.Networks
			case conf.GetBool("platforms"):
				output = c.Global.Platforms
			case conf.GetBool("repositories"):
				output = c.Global.Repositories
			case conf.GetBool("softwares"):
				output = c.Global.Softwares
			default:
				output = c.Global
			}
			switch format {
			case "raw":
				litter.Dump(output)
			case "yaml":
				fallthrough
			default:
				yamlOutput, marshalError := yaml.Marshal(output)
				if marshalError != nil {
					log.Errorf("error marshalling configuration into yaml: %s", marshalError)
				}
				log.Print(string(yamlOutput))
				log.Print("\n")
			}
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
