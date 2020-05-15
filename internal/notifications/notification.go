package notifications

func New(title, message string) Notification {
	return Notification{title, message}
}

type Notification struct {
	title   string
	message string
}

func (n Notification) GetTitle() string   { return n.title }
func (n Notification) GetMessage() string { return n.message }
