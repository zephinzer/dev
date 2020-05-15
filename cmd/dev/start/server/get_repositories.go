package server

import (
	"encoding/json"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/zephinzer/dev/internal/config"
)

func getRepositoriesHandler(res http.ResponseWriter, req *http.Request) {
	var repositories []byte
	// TODO: write error handling if config.Global.Repositories ain't available
	switch req.Header.Get("Accept") {
	case "application/yaml":
		repositories, _ = yaml.Marshal(config.Global.Repositories)
		res.Header().Add("Content-Type", "application/yaml")
	default:
		repositories, _ = json.Marshal(config.Global.Repositories)
		res.Header().Add("Content-Type", "application/json")
	}
	res.Write(repositories)
}
