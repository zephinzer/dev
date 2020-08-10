package client

import "errors"

// Notifications holds configurations related to the notifications
// mechanisms of the CLI tool
type Notifications struct {
	Telegram NotificationsTelegram `json:"telegram" yaml:"telegram,omitempty"`
}

func (dcn *Notifications) MergeWith(o Notifications) []error {
	return dcn.Telegram.MergeWith(o.Telegram)
}

// NotificationsTelegram holds configurations related to the
// telegram integration for sending notifications
type NotificationsTelegram struct {
	Token string `json:"token" yaml:"token,omitempty"`
	ID    string `json:"id" yaml:"id,omitempty"`
}

func (dcntg *NotificationsTelegram) MergeWith(o NotificationsTelegram) []error {
	if len(dcntg.Token) > 0 || len(dcntg.ID) > 0 {
		return []error{errors.New("dev.client.notifications.telegram.token already set")}
	}
	dcntg.Token = o.Token
	dcntg.ID = o.ID
	return nil
}
