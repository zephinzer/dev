package account

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/trello"
	"github.com/zephinzer/dev/internal/types"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information from Trello",
		Run: func(command *cobra.Command, args []string) {
			totalAccountsCount := 0
			accountsEncountered := map[string]interface{}{}
			defaultAccessKey := config.Global.Platforms.Trello.AccessKey
			defaultAccessToken := config.Global.Platforms.Trello.AccessToken
			if len(config.Global.Platforms.Trello.Boards) == 0 && (len(defaultAccessToken) > 0 && len(defaultAccessKey) > 0) {
				accountsEncountered[defaultAccessKey+defaultAccessToken] = true
				log.Info("account information for root trello account")
				account, getAccountError := trello.GetAccount(defaultAccessKey, defaultAccessToken)
				if getAccountError != nil {
					log.Warnf("failed to retrieve account information for trello using the default access credentials")
				} else {
					log.Printf(types.PrintAccount(account))
					totalAccountsCount++
				}
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
					log.Infof("account information for trello board '%s' (id: %s)", board.Name, board.ID)
					account, getAccountError := trello.GetAccount(boardAccessKey, boardAccessToken)
					if getAccountError != nil {
						log.Warnf("failed to retrieve account information for associated trello board '%s'", board.Name)
					} else {
						log.Printf(types.PrintAccount(account))
						totalAccountsCount++
					}
				}
			}
			log.Infof("total listed trello boards : %v\n", len(config.Global.Platforms.Trello.Boards))
			log.Infof("total trello accounts      : %v\n", totalAccountsCount)
		},
	}
	return &cmd
}
