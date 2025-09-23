package handler

import "fmt"


type LoggingHandler struct {
	AbstractHandler
}

func NewLoggingHandler() *LoggingHandler {
	return &LoggingHandler{}
}

func (lh *LoggingHandler) CanHandle(request Request) bool {
	return true // Logging handler can handle all requests
}

func (lh *LoggingHandler) Handle(request Request) bool {
	fmt.Printf("LoggingHandler: Logging %s request with priority %d\n", 
		request.GetType(), request.GetPriority())
	// Always pass to next handler
	return lh.AbstractHandler.Handle(request)
}