package repository

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

func askWhichConfigurationToAddTo(repoURL string) string {
	log.Printf("\n\033[1mwhich configuration file should we add '%s' to?\033[0m\n", repoURL)
	loadedIndex := 0
	loadedConfigs := []string{}
	for configPath := range config.Loaded {
		log.Printf("%v. %s\n", loadedIndex+1, configPath)
		loadedConfigs = append(loadedConfigs, configPath)
		loadedIndex++
	}
	log.Print("\033[1myour response (use 0 to skip):\033[0m ")
	answer := "0"
	_, scanlnError := fmt.Scanln(&answer)
	if scanlnError != nil && !strings.Contains(scanlnError.Error(), "unexpected newline") {
		log.Errorf("an unexpected error occurred: %s", scanlnError)
		os.Exit(constants.ExitErrorSystem)
	}
	indexToUse, atoiError := strconv.Atoi(answer)
	if atoiError != nil {
		log.Warnf("that wasn't a valid option: %s", atoiError)
		askWhichConfigurationToAddTo(repoURL)
	} else if indexToUse > loadedIndex {
		log.Warnf("that wasn't a valid option: %v is an unknown configuration", indexToUse)
		askWhichConfigurationToAddTo(repoURL)
	} else if indexToUse < 0 {
		log.Warn("that wasn't a valid option: use 0 to skip adding this repository")
		askWhichConfigurationToAddTo(repoURL)
	}
	if indexToUse == 0 {
		return ""
	}
	return loadedConfigs[indexToUse-1]
}
