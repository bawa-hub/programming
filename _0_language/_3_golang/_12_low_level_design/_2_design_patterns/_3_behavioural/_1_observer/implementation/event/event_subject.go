package event

type EventSubject interface {
	Subscribe(observer EventObserver)
	Unsubscribe(observer EventObserver)
	Publish(event *Event)
}