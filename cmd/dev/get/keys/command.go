package keys

import (
	"path"

	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils/keys"
)

var conf = config.Map{
	"path": &config.String{
		Shorthand: "p",
	},
}

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Retrieves user keys available on this machine",
		Run:   run,
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	homeDir := cmdutils.GetHomeDirectory()
	sshKeysDirectory := path.Join(homeDir, "/.ssh")

	keys, err := keys.GetSSH(sshKeysDirectory)
	if err != nil {
		cmd.Help()
	}

	for _, key := range keys {
		log.Infof(key.String())
	}
}
