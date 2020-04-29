package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	cf "github.com/usvc/go-config"
)

var conf = cf.Map{
	"addr": &cf.String{
		Default: "0.0.0.0",
	},
	"port": &cf.Uint{
		Default: 33835,
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ServerCanonicalNoun,
		Aliases: constants.ServerAliases,
		Short:   "starts the dev server",
		Run: func(command *cobra.Command, _ []string) {
			router := mux.NewRouter()

			// base path for healthchecks
			router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("hello world"))
			})

			// allow for syncing of required softwares
			router.HandleFunc("/softwares", func(res http.ResponseWriter, req *http.Request) {
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
			})

			// allow for syncing of networks
			router.HandleFunc("/networks", func(res http.ResponseWriter, req *http.Request) {
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
			})

			// allow for syncing of accounts
			router.HandleFunc("/platforms", func(res http.ResponseWriter, req *http.Request) {
				var platforms []byte
				// // TODO: write error handling if config.Global.Networks ain't available
				switch req.Header.Get("Accept") {
				case "application/yaml":
					platforms, _ = yaml.Marshal(config.Global.Platforms.GetSanitized())
					res.Header().Add("Content-Type", "application/yaml")
				default:
					platforms, _ = json.Marshal(config.Global.Platforms.GetSanitized())
					res.Header().Add("Content-Type", "application/json")
				}
				res.Write(platforms)
			})
			http.ListenAndServe(fmt.Sprintf("%s:%v", conf.GetString("addr"), conf.GetUint("port")), router)
		},
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}
