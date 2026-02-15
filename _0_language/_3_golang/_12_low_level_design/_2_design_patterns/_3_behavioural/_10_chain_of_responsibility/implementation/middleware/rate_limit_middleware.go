package middleware

import "fmt"

type RateLimitMiddleware struct {
	AbstractMiddleware
	requests map[string]int
	limit    int
}

func NewRateLimitMiddleware(limit int) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		requests: make(map[string]int),
		limit:    limit,
	}
}

func (rlm *RateLimitMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("RateLimitMiddleware: Checking rate limit")
	clientIP := request.Headers["X-Forwarded-For"]
	if clientIP == "" {
		clientIP = "unknown"
	}
	
	rlm.requests[clientIP]++
	if rlm.requests[clientIP] > rlm.limit {
		response.StatusCode = 429
		response.Body = "Too Many Requests"
		return false
	}
	return rlm.AbstractMiddleware.Process(request, response)
}