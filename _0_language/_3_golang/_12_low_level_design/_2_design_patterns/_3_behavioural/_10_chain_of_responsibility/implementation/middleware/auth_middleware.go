package middleware

import "fmt"

type AuthMiddleware struct {
	AbstractMiddleware
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("AuthMiddleware: Checking authentication")
	if request.Headers["Authorization"] == "" {
		response.StatusCode = 401
		response.Body = "Unauthorized"
		return false
	}
	return am.AbstractMiddleware.Process(request, response)
}