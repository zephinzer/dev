package network

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/network"
)

// GetCommand retrieves the command interface for the parent
// command to consume
func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NetworkCanonicalNoun,
		Aliases: constants.NetworkAliases,
		Short:   "verifies that host machine has access to networks specified in the configuration",
		Run:     Run,
	}
	return &cmd
}

// Run executes this command
func Run(command *cobra.Command, args []string) {
	var waiter sync.WaitGroup

	for _, nw := range config.Global.Networks {
		waiter.Add(1)

		// using goroutines so they can happen asynchronously
		go func(network network.Network) {
			var networkCheckLog strings.Builder
			runError := network.Check.Run()
			if runError != nil {
				networkCheckLog.WriteString(fmt.Sprintf("%s", runError))
			} else if verifyError := network.Check.Verify(); verifyError != nil {
				networkCheckLog.WriteString(fmt.Sprintf("%s", verifyError))
			}

			if networkCheckLog.Len() == 0 {
				log.Printf(constants.CheckSuccessFormat, network.Name)
				log.Printf("connected to %s", network.Check.URL)
			} else {
				log.Printf(constants.CheckFailureFormat, network.Name)
				log.Printf("failed to connect to %s: %s", network.Check.URL, networkCheckLog.String())
			}
			log.Print("\n")
			networkCheckLog.Reset()
			waiter.Done()
		}(nw)
	}

	waiter.Wait()
}
