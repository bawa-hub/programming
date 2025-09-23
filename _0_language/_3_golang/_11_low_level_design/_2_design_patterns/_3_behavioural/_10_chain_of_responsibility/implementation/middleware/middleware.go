package middleware



type Middleware interface {
	Process(request *HTTPRequest, response *HTTPResponse) bool
	SetNext(middleware Middleware)
}

type AbstractMiddleware struct {
	next Middleware
}

func (am *AbstractMiddleware) SetNext(middleware Middleware) {
	am.next = middleware
}

func (am *AbstractMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	if am.next != nil {
		return am.next.Process(request, response)
	}
	return true
}