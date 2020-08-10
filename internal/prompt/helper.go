package prompt

import (
	"bufio"
	"fmt"
	"os"
)

// InputHelper implements helper functions over the InputOptions
// structure
type InputHelper InputOptions

// PrintBeforeMessage prints the message to be displayed before
// the possible options are printed
func (ih InputHelper) PrintBeforeMessage() {
	if len(ih.BeforeMessage) > 0 {
		fmt.Printf("\n\033[1m%s\033[0m", ih.BeforeMessage)
	}
}

// PrintOptions prints the possible options for the user to
// select
func (ih InputHelper) PrintOptions() {
	if len(ih.BeforeMessage) == 0 {
		fmt.Print("\n")
	}
	for index, option := range ih.SerializedOptions {
		fmt.Printf("%v. %s", index+1, option)
		if index != len(ih.SerializedOptions)-1 {
			fmt.Printf("\n")
		}
	}
}

// PrintAfterMessage prints the message to be displayed after
// the possible options have been printed
func (ih InputHelper) PrintAfterMessage() {
	if len(ih.AfterMessage) == 0 {
		ih.AfterMessage = "your selection (enter 0 to skip)"
	}
	fmt.Printf("\033[1m%s\033[0m: ", ih.AfterMessage)
}

// ReadInput reads in a user's input
func (ih *InputHelper) ReadInput() error {
	if ih.Reader == nil {
		ih.Reader = os.Stdin
	}
	scanner := bufio.NewScanner(ih.Reader)
	if scanner.Scan() {
		ih.data = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return fmt.Errorf("failed to get input from tty: %s", scanError)
	}
	return nil
}

// GetData retrieves the read input
func (ih InputHelper) GetData() string {
	return ih.data
}
