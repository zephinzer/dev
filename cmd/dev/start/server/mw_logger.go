package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zephinzer/dev/internal/log"
)

func applyLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		requestID := uuid.New().String()
		timestamp := time.Now().UTC()
		res.Header().Add("request-id", requestID)
		next.ServeHTTP(res, req)
		log.Debugf(
			"[%s] %s - %s %s %s %s %vms",
			requestID,
			req.RemoteAddr,
			req.URL.User.Username(),
			timestamp.Format(time.RFC3339),
			req.URL.String(),
			req.UserAgent(),
			float64(time.Since(timestamp).Microseconds())/float64(1000),
		)
	})
}
