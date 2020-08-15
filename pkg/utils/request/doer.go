package request

import "net/http"

// Doer is an interface that defines an object like http.Client
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
