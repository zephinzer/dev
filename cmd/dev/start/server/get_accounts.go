package server

import (
	"encoding/json"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/zephinzer/dev/internal/config"
)

func getAccountsHandler(res http.ResponseWriter, req *http.Request) {
	var accounts []byte
	// // TODO: write error handling if config.Global.Networks ain't available
	switch req.Header.Get("Accept") {
	case "application/yaml":
		accounts, _ = yaml.Marshal(config.Global.Platforms.GetSanitized())
		res.Header().Add("Content-Type", "application/yaml")
	default:
		accounts, _ = json.Marshal(config.Global.Platforms.GetSanitized())
		res.Header().Add("Content-Type", "application/json")
	}
	res.Write(accounts)
}
