package handler

import (
	"fmt"
	"time"
)



type ValidationHandler struct {
	AbstractHandler
}

func NewValidationHandler() *ValidationHandler {
	return &ValidationHandler{}
}

func (vh *ValidationHandler) CanHandle(request Request) bool {
	return request.GetType() == "API" || request.GetType() == "DATA"
}

func (vh *ValidationHandler) Handle(request Request) bool {
	if vh.CanHandle(request) {
		fmt.Printf("ValidationHandler: Processing %s request\n", request.GetType())
		// Simulate validation logic
		time.Sleep(75 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return vh.AbstractHandler.Handle(request)
}