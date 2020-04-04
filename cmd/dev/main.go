package main

import (
	"fmt"

	"github.com/usvc/dev/cmd/dev/dev"
)

var (
	Commit    string
	Version   string
	Timestamp string
)

func main() {
	cmd := dev.GetCommand()
	cmd.Version = fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp)
	cmd.Execute()
}
