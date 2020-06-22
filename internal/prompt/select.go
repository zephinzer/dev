package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// PromptSelect prompts the user to select an option from the provided
// `promptOptions` argument. If the user skips/declines to respond, a
// -1 is returned. If an error happened, a -2 is returned
func ToSelect(promptMessage string, promptOptions []string) (int, error) {
	fmt.Printf("%s\n", promptMessage)
	for index, option := range promptOptions {
		fmt.Printf("%v. %s\n", index+1, option)
	}
	fmt.Printf("your selection (enter 0 to skip): ")
	var response string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return -2, fmt.Errorf("an unexpected error occurred: %s", scanError)
	}
	selectedOption, atoiError := strconv.Atoi(response)
	if atoiError != nil {
		return -2, fmt.Errorf("an invalid response was provided: %s", atoiError)
	}
	if selectedOption < 0 || selectedOption > len(promptOptions) {
		return -2, fmt.Errorf("an out-of-range response '%v' was provided", selectedOption)
	}
	return selectedOption - 1, nil
}
