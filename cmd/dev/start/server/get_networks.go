package server

import (
	"encoding/json"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/usvc/dev/internal/config"
)

func getNetworksHandler(res http.ResponseWriter, req *http.Request) {
	var networks []byte
	// TODO: write error handling if config.Global.Networks ain't available
	switch req.Header.Get("Accept") {
	case "application/yaml":
		networks, _ = yaml.Marshal(config.Global.Networks)
		res.Header().Add("Content-Type", "application/yaml")
	default:
		networks, _ = json.Marshal(config.Global.Networks)
		res.Header().Add("Content-Type", "application/json")
	}
	res.Write(networks)
}
