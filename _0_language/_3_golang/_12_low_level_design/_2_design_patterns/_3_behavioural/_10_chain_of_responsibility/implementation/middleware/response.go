package middleware

type HTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}