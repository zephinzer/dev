package types

type Notifications []Notification

type Notification interface {
	GetTitle() string
	GetMessage() string
}
