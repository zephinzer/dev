package github

import (
	"net/url"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	cf "github.com/usvc/go-config"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/github"
	"github.com/zephinzer/dev/pkg/oauth"
)

var conf = cf.Map{
	"client-id": &cf.String{
		Shorthand: "i",
		Usage:     "define a custom github app client ID to use",
	},
	"redirect-uri": &cf.String{
		Default:   "http://localhost:33835/oauth/callback",
		Shorthand: "r",
		Usage:     "defines the redirect uri to use",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GithubCanonicalNoun,
		Aliases: constants.GithubAliases,
		Run:     run,
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	// client id
	clientID := conf.GetString("client-id")
	if len(clientID) == 0 {
		clientID = c.Global.Dev.Client.Platforms.Github.ClientID
	}
	if len(clientID) == 0 {
		clientID = constants.GithubClientID
	}

	// state
	stateString := "github-" + uuid.New().String()

	// redirect uri
	githubRedirectURI := conf.GetString("redirect-uri")
	redirectURI, parseError := url.Parse(githubRedirectURI)
	if parseError != nil {
		log.Errorf("failed to parse redirect uri '%s': %s", githubRedirectURI, parseError)
		os.Exit(constants.ExitErrorInput | constants.ExitErrorConfiguration)
	}
	redirectURI.Path = path.Join(redirectURI.Path, stateString)

	authReq := oauth.AuthorizationRequest{
		BaseURL:     github.DefaultOAuthAuthorizationURL,
		ClientID:    c.Global.Dev.Client.Platforms.Github.ClientID,
		RedirectURI: redirectURI.String(),
		Scopes:      []string{"user", "notifications"},
		State:       stateString,
	}
	authReqURL, getURLError := authReq.GetURL()
	if getURLError != nil {
		log.Errorf("failed to get the authorization url: %s", getURLError)
		os.Exit(constants.ExitErrorApplication)
	}
	log.Infof("opening '%s' in the default browser", authReqURL)
	authReq.OpenInBrowser()
}
