package logging

import "time"

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LogEntry struct {
	Level     LogLevel
	Message   string
	Timestamp time.Time
	Source    string
}

type LogHandler interface {
	Handle(entry LogEntry) bool
	SetNext(handler LogHandler)
	CanHandle(entry LogEntry) bool
}

type AbstractLogHandler struct {
	next LogHandler
}

func (alh *AbstractLogHandler) SetNext(handler LogHandler) {
	alh.next = handler
}

func (alh *AbstractLogHandler) Handle(entry LogEntry) bool {
	if alh.next != nil {
		return alh.next.Handle(entry)
	}
	return true
}