package keys

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
)

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "keys",
		Run: run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	privateKeys, authorizedKeys, err := utils.GetLocalSSHKeys()
	if err != nil {
		cmd.Help()
	}
	log.Info("private keys follow")
	for _, keyPath := range privateKeys {
		log.Infof("- %s", keyPath)
	}
	log.Info("authorized keys follow")
	for keyPath, keyComment := range authorizedKeys {
		log.Infof("- %s (%s)", keyPath, keyComment)
	}
}
