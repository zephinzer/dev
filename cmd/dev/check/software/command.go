package software

import (
	"fmt"
	"os"
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
						log.Printf("(\033[1mskipped\033[0m: this system (%s) is not one of [%s])", currentOS, strings.Join(targetedPlatforms, ", "))
						if len(software.Description) > 0 {
							log.Printf(" \033[2m(%s)\033[0m ", software.Description)
						}
						log.Printf("\n")
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
					if len(software.Description) > 0 {
						log.Printf("\033[2m(%s)\033[0m", software.Description)
					}
				} else {
					log.Printf(constants.CheckFailureFormat, software.Name)
					if len(software.Description) > 0 {
						log.Printf("\033[2m(%s)\033[0m", software.Description)
					}
					log.Printf("\n%s", softwareCheckLog.String())
					failCount++
				}
				log.Print("\n")
				softwareCheckLog.Reset()
			}
			log.Printf("\ntotal/processed/skipped/failed: %v/%v/%v/%v\n", len(config.Global.Softwares), processedCount, skippedCount, failCount)
			os.Exit(-1 * failCount)
		},
	}
	return &cmd
}
