package logging

import "fmt"

type ConsoleLogHandler struct {
	AbstractLogHandler
}

func NewConsoleLogHandler() *ConsoleLogHandler {
	return &ConsoleLogHandler{}
}

func (clh *ConsoleLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= INFO
}

func (clh *ConsoleLogHandler) Handle(entry LogEntry) bool {
	if clh.CanHandle(entry) {
		fmt.Printf("[%s] %s: %s\n", 
			entry.Timestamp.Format("15:04:05"), 
			clh.getLevelString(entry.Level), 
			entry.Message)
	}
	return clh.AbstractLogHandler.Handle(entry)
}

func (clh *ConsoleLogHandler) getLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
