package logging

import "fmt"

type FileLogHandler struct {
	AbstractLogHandler
	filename string
}

func NewFileLogHandler(filename string) *FileLogHandler {
	return &FileLogHandler{filename: filename}
}

func (flh *FileLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= WARN
}

func (flh *FileLogHandler) Handle(entry LogEntry) bool {
	if flh.CanHandle(entry) {
		fmt.Printf("FileLogHandler: Writing to %s - [%s] %s\n", 
			flh.filename, 
			entry.Timestamp.Format("15:04:05"), 
			entry.Message)
	}
	return flh.AbstractLogHandler.Handle(entry)
}
