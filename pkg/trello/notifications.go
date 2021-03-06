package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils/request"
)

func GetNotifications(accessKey, accessToken string) (*APIv1MemberNotificationResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: "https://api.trello.com/1/members/me/notifications",
		Queries: map[string]string{
			"key":         accessKey,
			"token":       accessToken,
			"read_filter": "unread",
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
	var response APIv1MemberNotificationResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

type APIv1MemberNotificationResponse []APIv1MemberNotification

// String converts the notifications object into a CLI-friendly block of text
func (n APIv1MemberNotificationResponse) String() string {
	var output strings.Builder
	for i := 0; i < len(n); i++ {
		output.WriteString(n[i].String())
		output.WriteString("\n\n")
	}
	return output.String()
}

type APIv1MemberNotificationData struct {
	Board      *APIv1MemberNotificationDataBoard `json:"board"`
	Card       *APIv1MemberNotificationDataCard  `json:"card"`
	ListBefore *APIv1MemberNotificationDataList  `json:"listBefore"`
	ListAfter  *APIv1MemberNotificationDataList  `json:"listAfter"`
	Text       string                            `json:"text"`
}

type APIv1MemberNotificationDataBoard struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortLink string `json:"shortLink"`
}

type APIv1MemberNotificationDataCard struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IDShort   int    `json:"idShort"`
	ShortLink string `json:"shortLink"`
}

type APIv1MemberNotificationDataList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type APIv1MemberNotification struct {
	ID            string                      `json:"id"`
	Type          string                      `json:"type"`
	Date          string                      `json:"date"`
	Data          APIv1MemberNotificationData `json:"data"`
	Unread        bool                        `json:"unread"`
	MemberCreator *APIv1MemberResponse        `json:"memberCreator"`
}

func (n APIv1MemberNotification) String() string {
	datetime := n.Date
	timestamp, err := time.Parse(constants.PivotalTrackerAPITimeFormat, n.Date)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}
	switch n.Type {
	case "commentCard":
		return fmt.Sprintf("[%s/%s]\n  - comment was made by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
	case "removedFromCard":
		return fmt.Sprintf("[%s/%s]\n  - you were removed by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
	case "addedToCard":
		return fmt.Sprintf("[%s/%s]\n  - you were added to this by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
	case "mentionedOnCard":
		return fmt.Sprintf("[%s/%s]\n  - you were mentioned by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
	case "changeCard":
		if n.Data.ListAfter == nil {
			return fmt.Sprintf("[%s/%s]\n  - card was changed by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
		}
		return fmt.Sprintf("[%s/%s]\n  - card was shifted to %s by %s (@%s) about %s\n  - link: https://trello.com/c/%s/", n.Data.Board.Name, n.Data.Card.Name, n.Data.ListAfter.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Card.ShortLink)
	case "removedFromBoard":
		return fmt.Sprintf("[%s/**]\n  - you were removed by %s (@%s) about %s\n  - link: https://trello.com/b/%s/", n.Data.Board.Name, n.MemberCreator.FullName, n.MemberCreator.Username, datetime, n.Data.Board.ShortLink)
	default:
		var unknownNotification strings.Builder
		if n.Data.Board != nil {
			if n.Data.Card != nil {
				unknownNotification.WriteString(fmt.Sprintf("[%s/%s]\n", n.Data.Board.Name, n.Data.Card.Name))
			} else {
				unknownNotification.WriteString(fmt.Sprintf("[%s]\n", n.Data.Board.Name))
			}
		}
		unknownNotification.WriteString(fmt.Sprintf("  - notification type: %s\n", n.Type))
		unknownNotification.WriteString(fmt.Sprintf("  - happened on: %s\n", datetime))
		if n.MemberCreator != nil {
			unknownNotification.WriteString(fmt.Sprintf("\n  - triggered by: %s (@%s)", n.MemberCreator.FullName, n.MemberCreator.Username))
		}
		return unknownNotification.String()
	}
}
