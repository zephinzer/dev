//go:generate go run ./generators/htmlimport
package server

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zephinzer/dev/cmd/dev/start/server/assets"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/github"
	"github.com/zephinzer/dev/pkg/oauth"
)

func respondWithError(requestID, platform, err string, res http.ResponseWriter) {
	errorText := fmt.Sprintf("[requestID:%s] - %s", requestID, err)
	log.Error(errorText)
	page := strings.Replace(assets.OAuthTokenErrorPage, "${OAUTH_PLATFORM}", platform, -1)
	page = strings.Replace(page, "${ERROR_TEXT}", errorText, -1)
	res.Header().Add("X-Request-ID", requestID)
	res.Header().Add("Content-Type", "text/html")
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(page))
}

func respondWithToken(requestID, platform, token string, res http.ResponseWriter) {
	log.Infof("[requestID:%s] succeeded", requestID)
	page := strings.Replace(assets.OAuthTokenPage, "${OAUTH_PLATFORM}", platform, -1)
	page = strings.Replace(page, "${OAUTH_ACCESS_TOKEN}", token, -1)
	res.Header().Add("X-Request-ID", requestID)
	res.Header().Add("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(page))
}

func getOAuthCallbackHandler(res http.ResponseWriter, req *http.Request) {
	requestID := uuid.New().String()

	params := mux.Vars(req)

	// platform
	platform := strings.SplitN(params["state"], "-", 2)[0]
	var platformLabel string
	switch platform {
	case constants.GithubCanonicalNoun:
		platformLabel = "Github"
	default:
		respondWithError(requestID, "Unknown", "invalid oauth platform provider specified", res)
		return
	}

	// client id
	var clientID string
	switch platform {
	case constants.GithubCanonicalNoun:
		clientID = conf.GetString("github-client-id")
		if len(clientID) == 0 {
			clientID = c.Global.Dev.Client.Platforms.Github.ClientID
		}
	}
	if len(clientID) == 0 {
		respondWithError(requestID, platformLabel, "client id is required but was not specified in the configuration", res)
		return
	}

	// client secret
	var clientSecret string
	switch platform {
	case constants.GithubCanonicalNoun:
		clientSecret = conf.GetString("github-client-secret")
		if len(clientSecret) == 0 {
			clientSecret = c.Global.Dev.Client.Platforms.Github.ClientSecret
		}
	}
	if len(clientSecret) == 0 {
		respondWithError(requestID, platformLabel, "client secret is required but was not specified in the configuration", res)
		return
	}

	// redirect uri
	var redirectURIFromConfiguration string
	var redirectURI *url.URL
	var parseError error
	switch platform {
	case constants.GithubCanonicalNoun:
		redirectURIFromConfiguration = conf.GetString("github-redirect-uri")
		redirectURI, parseError = url.Parse(redirectURIFromConfiguration)
	}
	if parseError != nil {
		respondWithError(requestID, platformLabel, fmt.Sprintf("failed to parse redirect uri '%s' - %s", redirectURIFromConfiguration, parseError), res)
		return
	}
	redirectURI.RawPath = path.Join(redirectURI.Path, params["state"])

	// get query params
	authorizationResponse := oauth.AuthorizationResponse{}
	if loadFromQueryError := authorizationResponse.LoadFromQuery(req.URL.Query()); loadFromQueryError != nil {
		respondWithError(requestID, platformLabel, fmt.Sprintf("failed to get authorization code (error code: '%s'): %s", authorizationResponse.Error, authorizationResponse.ErrorDescription), res)
		return
	}
	if params["state"] != authorizationResponse.State {
		respondWithError(requestID, platformLabel, fmt.Sprintf("oauth state mismatch: expected '%s' but got '%s'", params["state"], authorizationResponse.State), res)
		return
	}

	// get access token and respond
	switch platform {
	case constants.GithubCanonicalNoun:
		grantRequest := oauth.GrantRequest{
			BaseURL:      github.DefaultOAuthGrantURL,
			ClientID:     c.Global.Dev.Client.Platforms.Github.ClientID,
			ClientSecret: clientSecret,
			Code:         authorizationResponse.Code,
			RedirectURI:  redirectURI.String(),
			State:        params["state"],
		}
		grantResponse, doError := grantRequest.Do()
		if doError != nil {
			respondWithError(requestID, platformLabel, fmt.Sprintf("failed to request for access token: %s", doError), res)
			return
		}
		if grantResponse.Error != nil {
			respondWithError(requestID, platformLabel, fmt.Sprintf("failed to get access token: %s", grantResponse.Error), res)
			return
		}
		respondWithToken(requestID, platformLabel, grantResponse.AccessToken, res)
	}
}
