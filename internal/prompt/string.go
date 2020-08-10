package prompt

import (
	"fmt"

	"github.com/zephinzer/dev/pkg/utils/str"
)

// ForString represents a standardised way of requesting for a string input
func ForString(options InputOptions, defaultTo ...string) (string, error) {
	helper := InputHelper(options)
	fmt.Printf("\n")
	helper.PrintBeforeMessage()
	if readInputErr := helper.ReadInput(); readInputErr != nil {
		return "", readInputErr
	}
	response := helper.GetData()
	if str.IsEmpty(response) {
		if len(defaultTo) > 0 {
			response = defaultTo[0]
		}
	}
	return response, nil
}
