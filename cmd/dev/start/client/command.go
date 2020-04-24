package client

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
)

var conf = config.Map{
	"daemon": &config.Bool{
		Shorthand: "d",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "client",
		Aliases: []string{"c"},
		Short:   "starts the dev client as a background process to provide notifications",
		Run: func(command *cobra.Command, _ []string) {
			log.Println("starting dev client...")
			if conf.GetBool("daemon") {
				var waiter sync.WaitGroup
				waiter.Add(1)
				go func(tick <-chan time.Time) {
					log.Printf("starting long-running placeholder")
					for {
						<-tick
						fmt.Print(".")
					}
				}(time.Tick(3 * time.Second))
				waiter.Wait()
				command.Help()
			} else {
				invocation := os.Args[0]
				log.Printf("converting to background mode with <%s &>", invocation)
				background := exec.Command(invocation, "--daemon")
				background.Start()
			}
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
