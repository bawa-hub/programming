package event


type EventObserver interface {
	HandleEvent(event *Event)
	GetID() string
	GetEventTypes() []string
}