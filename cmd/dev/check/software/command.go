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
					log.Printf("✅ \033[1m%s\033[0m", software.Name)
				} else {
					log.Printf("❌ \033[1m%s\033[0m: %s", software.Name, softwareCheckLog.String())
				}
				log.Print("\n")
				softwareCheckLog.Reset()
			}
		},
	}
	return &cmd
}
