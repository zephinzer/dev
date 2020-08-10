package prompt

import (
	"fmt"
	"strconv"

	"github.com/zephinzer/dev/pkg/utils/str"
)

// ToSelect prompts the user to select an option from the provided
// `options.Options` argument. On success, returns the selected index.
// Returns `ErrorSkipped` if the user indicated to skip, returns
// `ErrorInput` on input validation errors, returns `ErrorSystem` on
// system errors
func ToSelect(options InputOptions) (int, error) {
	helper := InputHelper(options)
	helper.PrintBeforeMessage()
	fmt.Printf("\n")
	helper.PrintOptions()
	fmt.Printf("\n")
	helper.PrintAfterMessage()
	if readInputError := helper.ReadInput(); readInputError != nil {
		return int(ErrorSystem), readInputError
	}
	response := helper.GetData()
	if str.IsEmpty(response) {
		return int(ErrorSkipped), nil
	}
	selectedOption, atoiError := strconv.Atoi(response)
	if atoiError != nil {
		return int(ErrorInput), fmt.Errorf("an invalid response was provided: %s", atoiError)
	} else if selectedOption == 0 {
		return int(ErrorSkipped), nil
	} else if selectedOption < 0 || selectedOption > len(options.SerializedOptions) {
		return int(ErrorInput), fmt.Errorf("an out-of-range response '%v' was provided", selectedOption)
	}
	return selectedOption - 1, nil
}
