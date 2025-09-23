package handler

import (
	"fmt"
	"time"
)

// Concrete Handlers
type AuthenticationHandler struct {
	AbstractHandler
}

func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}

func (ah *AuthenticationHandler) CanHandle(request Request) bool {
	return request.GetType() == "AUTH" || request.GetType() == "API"
}

func (ah *AuthenticationHandler) Handle(request Request) bool {
	if ah.CanHandle(request) {
		fmt.Printf("AuthenticationHandler: Processing %s request\n", request.GetType())
		// Simulate authentication logic
		time.Sleep(100 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return ah.AbstractHandler.Handle(request)
}