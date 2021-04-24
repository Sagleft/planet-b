package notifications

import (
	"encoding/json"
	"io/ioutil"
)

const (
	notificationsConfigPath = "config/notifications.json"
)

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

func newGodNotificator() (*GodNotificator, error) {
	jsonBytes, err := ioutil.ReadFile(notificationsConfigPath)
	if err != nil {
		return nil, err
	}
	notyMessenger := Messenger{}
	parseErr := json.Unmarshal(jsonBytes, &notyMessenger)
	if parseErr != nil {
		return nil, parseErr
	}

	noty := GodNotificator{
		MessengerService: notyMessenger,
	}
	return &noty, nil
}
