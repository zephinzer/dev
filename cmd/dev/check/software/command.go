package software

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.SoftwareCanonicalNoun,
		Aliases: constants.SoftwareAliases,
		Short:   "verifies that required software specified in the configuration is installed",
		Run: func(command *cobra.Command, args []string) {
			var softwareCheckLog strings.Builder
			currentOS := runtime.GOOS
			processedCount := 0
			skippedCount := 0
			failCount := 0
			for _, software := range config.Global.Softwares {
				if len(software.Platforms) > 0 {
					targetedPlatforms := []string{}
					isTargetedPlatform := false
					for _, p := range software.Platforms {
						if p.String() == currentOS {
							isTargetedPlatform = true
						}
						targetedPlatforms = append(targetedPlatforms, p.String())
					}
					if !isTargetedPlatform {
						log.Printf(constants.CheckSkippedFormat, software.Name)
						log.Printf("(skipping: %s is not one of [%s])\n", currentOS, strings.Join(targetedPlatforms, ", "))
						skippedCount++
						continue
					}
				}
				processedCount++
				runError := software.Check.Run()
				if runError != nil {
					softwareCheckLog.WriteString(fmt.Sprintf("%s", runError))
					if len(software.Install.Link) > 0 {
						softwareCheckLog.WriteString(fmt.Sprintf("\n> see installation instructions at %s", software.Install.Link))
					}
				} else if verifyError := software.Check.Verify(); verifyError != nil {
					softwareCheckLog.WriteString(fmt.Sprintf("%s", verifyError))
					if len(software.Install.Link) > 0 {
						softwareCheckLog.WriteString(fmt.Sprintf("\n> see installation instructions at %s", software.Install.Link))
					}
				}
				if softwareCheckLog.Len() == 0 {
					log.Printf(constants.CheckSuccessFormat, software.Name)
				} else {
					log.Printf(constants.CheckFailureFormat, software.Name)
					log.Printf(softwareCheckLog.String())
					failCount++
				}
				log.Print("\n")
				softwareCheckLog.Reset()
			}
			log.Printf("\ntotal/processed/skipped/failed: %v/%v/%v/%v\n", len(config.Global.Softwares), processedCount, skippedCount, failCount)
		},
	}
	return &cmd
}
