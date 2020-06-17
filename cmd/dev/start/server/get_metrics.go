package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var getMetricsHandler = promhttp.Handler()
