package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/zephinzer/dev/pkg/utils/request"
)

func GetBoards(accessKey, accessToken string) (*APIv1BoardsResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: "https://api.trello.com/1/members/me/boards",
		Queries: map[string]string{
			"key":    accessKey,
			"token":  accessToken,
			"lists":  "open",
			"fields": "id,name,lists,shortLink,url,desc,dateLastActivity,dateLastView",
		},
	})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv1BoardsResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

func GetListCards(accessKey, accessToken, listID string) (*APIv1ListCardsResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: fmt.Sprintf("https://api.trello.com/1/lists/%s", listID),
		Queries: map[string]string{
			"key":    accessKey,
			"token":  accessToken,
			"lists":  "open",
			"fields": "id,name,lists,shortLink,url,desc,dateLastActivity,dateLastView",
		},
	})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv1ListCardsResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

type APIv1ListCardsResponse []APIv1ListCard

type APIv1ListCard struct {
	// TODO get json fields from lists-cards.json
	ID         string
	Closed     bool
	ListID     string
	BoardID    string
	Name       string
	ShortLink  string
	MemberIDs  []interface{}
	URL        string
	Subscribed bool
}

type APIv1BoardsResponse []APIv1Board

type APIv1Board struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	ShortLink        string      `json:"shortLink"`
	Desc             string      `json:"desc"`
	DateLastActivity string      `json:"dateLastActivity"`
	DateLastView     string      `json:"dateLastView"`
	URL              string      `json:"url"`
	Lists            []APIv1List `json:"lists"`
}

type APIv1List struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	BoardID string `json:"idBoard"`
}
