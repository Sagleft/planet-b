package notifications

// GodNotificator - simulation alert gizmo
type GodNotificator struct {
	MessengerService Messenger
}

// Messenger for notifications
type Messenger struct {
	Name     string
	BotToken string
	ChatID   string
}

func newGodNotificator() *GodNotificator {
	noty := GodNotificator{
		MessengerService: Messenger{
			Name:     "Telegram",
			BotToken: "",
			ChatID:   "",
		},
	}
	return &noty
}
