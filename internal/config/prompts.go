package config

import (
	"io"
	"os"
	"sort"

	"github.com/zephinzer/dev/internal/prompt"
	"github.com/zephinzer/dev/pkg/utils/defaults"
)

// PromptSelectLoadedConfiguration does a cli prompt to ask the user
// to select a loaded configuration for changes to be made
func PromptSelectLoadedConfiguration(promptMessage string, useOtherReader ...io.Reader) (string, error) {
	loadedIndex := 0
	loadedConfigs := []string{}
	for configPath := range Loaded {
		loadedConfigs = append(loadedConfigs, configPath)
		loadedIndex++
	}
	// TODO: add case where there are no loaded configs
	if len(loadedConfigs) == 1 {
		return loadedConfigs[0], nil
	}
	sort.Strings(loadedConfigs)
	reader := defaults.GetIoReader(os.Stdin, useOtherReader...)
	indexToUse, promptToSelectErr := prompt.ToSelect(prompt.InputOptions{
		BeforeMessage:     promptMessage,
		SerializedOptions: loadedConfigs,
		Reader:            reader,
	})
	if promptToSelectErr != nil {
		return "", promptToSelectErr
	}
	if indexToUse == int(prompt.ErrorSkipped) {
		return "", nil
	}
	return loadedConfigs[indexToUse], nil
}
