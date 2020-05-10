package software

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.SoftwareCanonicalNoun,
		Aliases: constants.SoftwareAliases,
		Short:   "verifies that required software specified in the configuration is installed",
		Run: func(command *cobra.Command, args []string) {
			var softwareCheckLog strings.Builder
			for _, software := range config.Global.Softwares {
				runError := software.Check.Run()
				if runError != nil {
					softwareCheckLog.WriteString(fmt.Sprintf("%s", runError))
				} else if verifyError := software.Check.Verify(); verifyError != nil {
					softwareCheckLog.WriteString(fmt.Sprintf("%s", verifyError))
				}
				if softwareCheckLog.Len() == 0 {
					log.Printf(constants.CheckSuccessFormat, software.Name)
				} else {
					log.Printf(constants.CheckFailureFormat, software.Name)
					log.Printf(softwareCheckLog.String())
				}
				log.Print("\n")
				softwareCheckLog.Reset()
			}
		},
	}
	return &cmd
}
