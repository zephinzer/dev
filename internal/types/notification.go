package types

type Notification interface {
	GetTitle() string
	GetMessage() string
}
