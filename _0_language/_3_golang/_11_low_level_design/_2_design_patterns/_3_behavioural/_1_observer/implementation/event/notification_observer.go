package event

import "fmt"

type NotificationObserver struct {
	id string
}

func NewNotificationObserver(id string) *NotificationObserver {
	return &NotificationObserver{id: id}
}

func (no *NotificationObserver) HandleEvent(event *Event) {
	fmt.Printf("Notification Observer %s: Sending notification for - %s\n", no.id, event)
}

func (no *NotificationObserver) GetID() string {
	return no.id
}

func (no *NotificationObserver) GetEventTypes() []string {
	return []string{"user_login", "error", "warning", "system_alert"}
}