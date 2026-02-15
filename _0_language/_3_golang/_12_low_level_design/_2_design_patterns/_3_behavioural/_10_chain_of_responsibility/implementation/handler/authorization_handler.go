package handler

import (
	"fmt"
	"time"
)



type AuthorizationHandler struct {
	AbstractHandler
}

func NewAuthorizationHandler() *AuthorizationHandler {
	return &AuthorizationHandler{}
}

func (azh *AuthorizationHandler) CanHandle(request Request) bool {
	return request.GetType() == "AUTH" || request.GetType() == "API" || request.GetType() == "ADMIN"
}

func (azh *AuthorizationHandler) Handle(request Request) bool {
	if azh.CanHandle(request) {
		fmt.Printf("AuthorizationHandler: Processing %s request\n", request.GetType())
		// Simulate authorization logic
		time.Sleep(50 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return azh.AbstractHandler.Handle(request)
}
