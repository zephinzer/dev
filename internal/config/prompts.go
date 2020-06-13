package config

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

// PromptSelectLoadedConfiguration does a cli prompt to ask the user
// to select a loaded configuration for changes to be made
func PromptSelectLoadedConfiguration(promptMessage string) string {
	loadedIndex := 0
	loadedConfigs := []string{}
	for configPath := range Loaded {
		loadedConfigs = append(loadedConfigs, configPath)
		loadedIndex++
	}
	if len(loadedConfigs) == 1 {
		return loadedConfigs[0]
	}
	sort.Strings(loadedConfigs)
	log.Printf("\n\033[1m%s\033[0m\n", promptMessage)
	for index, configPath := range loadedConfigs {
		log.Printf("%v. %s\n", index+1, configPath)
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
		PromptSelectLoadedConfiguration(promptMessage)
	} else if indexToUse > loadedIndex {
		log.Warnf("that wasn't a valid option: %v is an unknown configuration", indexToUse)
		PromptSelectLoadedConfiguration(promptMessage)
	} else if indexToUse < 0 {
		log.Warn("that wasn't a valid option: use 0 to skip adding this repository")
		PromptSelectLoadedConfiguration(promptMessage)
	}
	if indexToUse == 0 {
		return ""
	}
	return loadedConfigs[indexToUse-1]
}
