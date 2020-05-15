package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zephinzer/dev/internal/log"
)

func getRouter() http.Handler {
	router := mux.NewRouter()

	// base path for healthchecks
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello world"))
	}).Methods("GET")

	// allow for syncing of networks
	log.Debug("registering GET /networks handler")
	router.HandleFunc("/networks", getNetworksHandler).Methods("GET")

	// allow for syncing of accounts
	log.Debug("registering GET /accounts handler")
	router.HandleFunc("/accounts", getAccountsHandler).Methods("GET")

	// allow for syncing of required softwares
	log.Debug("registering GET /softwares handler")
	router.HandleFunc("/softwares", getSoftwaresHandler).Methods("GET")

	// allow for syncing of required repositories
	log.Debug("registering GET /repositories handler")
	router.HandleFunc("/repositories", getRepositoriesHandler).Methods("GET")

	return router
}
