package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zephinzer/dev/internal/log"
)

func getRouter() http.Handler {
	router := mux.NewRouter()

	// base path
	log.Debug("registering GET / handler...")
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello world"))
	}).Methods("GET")

	// for liveness checks
	log.Debug("registering GET /healthz handler...")
	router.HandleFunc("/healthz", getLivenessProbeHandler).Methods("GET")

	// for readiness checks
	log.Debug("registering GET /readyz handler...")
	router.HandleFunc("/readyz", getReadinessProbeHandler).Methods("GET")

	// for metrics
	log.Debug("registering GET /metrics handler...")
	router.Handle("/metrics", getMetricsHandler).Methods("GET")

	// allow for syncing of networks
	log.Debug("registering GET /networks handler...")
	router.HandleFunc("/networks", getNetworksHandler).Methods("GET")

	// allow for syncing of accounts
	log.Debug("registering GET /accounts handler...")
	router.HandleFunc("/accounts", getAccountsHandler).Methods("GET")

	// allow for syncing of required softwares
	log.Debug("registering GET /softwares handler...")
	router.HandleFunc("/softwares", getSoftwaresHandler).Methods("GET")

	// allow for syncing of required repositories
	log.Debug("registering GET /repositories handler...")
	router.HandleFunc("/repositories", getRepositoriesHandler).Methods("GET")

	// for platform integration callbacks
	log.Debug("registering GET /oauth/callback handler...")
	router.HandleFunc("/oauth/callback/{state}", getOAuthCallbackHandler).Methods("GET")

	return router
}
