package server

import (
	"encoding/json"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/zephinzer/dev/internal/config"
)

func getSoftwaresHandler(res http.ResponseWriter, req *http.Request) {
	var softwares []byte
	// TODO: write error handling if config.Global.Softwares ain't available
	switch req.Header.Get("Accept") {
	case "application/yaml":
		softwares, _ = yaml.Marshal(config.Global.Softwares)
		res.Header().Add("Content-Type", "application/yaml")
	default:
		softwares, _ = json.Marshal(config.Global.Softwares)
		res.Header().Add("Content-Type", "application/json")
	}
	res.Write(softwares)
}
