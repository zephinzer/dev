package trello

import (
	"log"

	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/pkg/trello"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "trello",
		Aliases: []string{"tr"},
		Short:   "Retrieves account information from Trello",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			defaultAccessKey := config.Global.Platforms.Trello.AccessKey
			defaultAccessToken := config.Global.Platforms.Trello.AccessToken
			if len(config.Global.Platforms.Trello.Boards) == 0 && (len(defaultAccessToken) > 0 && len(defaultAccessKey) > 0) {
				accountsEncountered[defaultAccessKey+defaultAccessToken] = true
				printAccountInfo(defaultAccessKey, defaultAccessToken)
				totalAccountsCount++
			}
			for _, board := range config.Global.Platforms.Trello.Boards {
				boardAccessKey := board.AccessKey
				boardAccessToken := board.AccessToken
				if len(boardAccessKey) == 0 && len(defaultAccessKey) > 0 {
					boardAccessKey = defaultAccessKey
				}
				if len(boardAccessToken) == 0 && len(defaultAccessToken) > 0 {
					boardAccessToken = defaultAccessToken
				}
				if accountsEncountered[boardAccessKey+boardAccessToken] == nil {
					accountsEncountered[boardAccessKey+boardAccessToken] = true
					log.Printf("\n# account information for board '%s' (id: %s)\n\n", board.Name, board.ID)
					printAccountInfo(boardAccessKey, boardAccessToken)
					totalAccountsCount++
				}
			}
			log.Printf("> total listed boards: %v\n", len(config.Global.Platforms.Trello.Boards))
			log.Printf("> total accounts: %v\n", totalAccountsCount)
		},
	}
	return &cmd
}

func printAccountInfo(accessKey, accessToken string) {
	accountInfo, err := trello.GetAccount(accessKey, accessToken)
	if err != nil {
		panic(err)
	}
	litter.Dump(accountInfo)
	// log.Println(accountInfo.String())
}
