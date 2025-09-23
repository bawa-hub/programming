package middleware

import "fmt"

type LoggingMiddleware struct {
	AbstractMiddleware
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (lm *LoggingMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Printf("LoggingMiddleware: %s %s - User: %s\n", 
		request.Method, request.URL, request.User)
	return lm.AbstractMiddleware.Process(request, response)
}