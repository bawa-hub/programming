package logging

import "fmt"

type EmailLogHandler struct {
	AbstractLogHandler
	email string
}

func NewEmailLogHandler(email string) *EmailLogHandler {
	return &EmailLogHandler{email: email}
}

func (elh *EmailLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= ERROR
}

func (elh *EmailLogHandler) Handle(entry LogEntry) bool {
	if elh.CanHandle(entry) {
		fmt.Printf("EmailLogHandler: Sending email to %s - [%s] %s\n", 
			elh.email, 
			entry.Timestamp.Format("15:04:05"), 
			entry.Message)
	}
	return elh.AbstractLogHandler.Handle(entry)
}