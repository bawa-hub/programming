package event

type LoggingObserver struct {
	id string
}

func NewLoggingObserver(id string) *LoggingObserver {
	return &LoggingObserver{id: id}
}

func (lo *LoggingObserver) HandleEvent(event *Event) {
	fmt.Printf("Logging Observer %s: Logging event - %s\n", lo.id, event)
}

func (lo *LoggingObserver) GetID() string {
	return lo.id
}

func (lo *LoggingObserver) GetEventTypes() []string {
	return []string{"user_login", "user_logout", "error", "warning"}
}