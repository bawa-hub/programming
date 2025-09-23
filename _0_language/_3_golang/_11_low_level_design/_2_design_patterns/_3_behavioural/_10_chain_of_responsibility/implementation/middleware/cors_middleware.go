package middleware

import "fmt"

type CORSMiddleware struct {
	AbstractMiddleware
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (cm *CORSMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("CORSMiddleware: Adding CORS headers")
	response.Headers["Access-Control-Allow-Origin"] = "*"
	response.Headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE"
	response.Headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
	return cm.AbstractMiddleware.Process(request, response)
}