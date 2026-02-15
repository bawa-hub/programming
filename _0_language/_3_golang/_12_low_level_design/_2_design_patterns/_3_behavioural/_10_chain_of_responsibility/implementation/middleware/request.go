package middleware

type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
	User    string
	Role    string
}