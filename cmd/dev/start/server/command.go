package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	cf "github.com/usvc/go-config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
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
		Run:     run,
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}

func run(command *cobra.Command, _ []string) {
	router := getRouter()

	bindAddress := fmt.Sprintf("%s:%v", conf.GetString("addr"), conf.GetUint("port"))
	server := http.Server{
		Addr:              bindAddress,
		Handler:           applyLoggerMiddleware(router),
		MaxHeaderBytes:    1024 * 16,
		IdleTimeout:       time.Second * 3,
		ReadHeaderTimeout: time.Second * 3,
		ReadTimeout:       time.Second * 3,
		WriteTimeout:      time.Second * 3,
	}

	log.Infof("starting server on %s", bindAddress)
	if listenAndServeError := server.ListenAndServe(); listenAndServeError != nil {
		log.Errorf("failed to start server: %s", listenAndServeError)
		os.Exit(1)
	}
}
