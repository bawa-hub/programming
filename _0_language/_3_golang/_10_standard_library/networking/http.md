# http Package - HTTP Client and Server 🌐

The `http` package provides HTTP client and server implementations. It's essential for building web applications, APIs, and HTTP-based services.

## 🎯 Key Concepts

### 1. **HTTP Client**
- `Client` - HTTP client with configurable settings
- `Get()` - GET request
- `Post()` - POST request
- `PostForm()` - POST form data
- `Do()` - Custom HTTP request
- `Head()` - HEAD request
- `NewRequest()` - Create HTTP request

### 2. **HTTP Server**
- `Server` - HTTP server configuration
- `ListenAndServe()` - Start HTTP server
- `ListenAndServeTLS()` - Start HTTPS server
- `Handler` - HTTP handler interface
- `HandlerFunc` - Function-based handler
- `ServeMux` - HTTP request multiplexer

### 3. **HTTP Request/Response**
- `Request` - HTTP request structure
- `Response` - HTTP response structure
- `Header` - HTTP headers
- `Cookie` - HTTP cookies
- `Body` - Request/response body
- `Method` - HTTP method
- `URL` - Request URL

### 4. **Middleware and Handlers**
- `Middleware` - Request processing pipeline
- `HandlerFunc` - Function-based handlers
- `ServeMux` - URL routing
- `StripPrefix` - URL prefix handling
- `TimeoutHandler` - Request timeout
- `FileServer` - Static file serving

### 5. **HTTP Utilities**
- `ParseHTTPVersion()` - Parse HTTP version
- `ParseTime()` - Parse HTTP time
- `CanonicalHeaderKey()` - Canonicalize header keys
- `DetectContentType()` - Detect content type
- `MaxBytesReader()` - Limit request size

### 6. **HTTP/2 and HTTPS**
- `Transport` - HTTP transport configuration
- `TLS` - TLS configuration
- `HTTP/2` - HTTP/2 support
- `Server Push` - HTTP/2 server push
- `WebSockets` - WebSocket support

## 🚀 Common Patterns

### Basic HTTP Client
```go
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
```

### Basic HTTP Server
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

### Custom HTTP Client
```go
client := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns: 100,
        IdleConnTimeout: 90 * time.Second,
    },
}

resp, err := client.Get("https://api.example.com/data")
```

### POST Request with JSON
```go
data := map[string]string{"key": "value"}
jsonData, _ := json.Marshal(data)

resp, err := http.Post("https://api.example.com/data", 
    "application/json", 
    bytes.NewBuffer(jsonData))
```

## ⚠️ Common Pitfalls

1. **Not closing response bodies** - Always close response bodies
2. **Not handling errors** - Check all HTTP errors
3. **Not setting timeouts** - Set appropriate timeouts
4. **Not validating input** - Validate all HTTP input
5. **Not handling context** - Use context for cancellation

## 🎯 Best Practices

1. **Close response bodies** - Use defer to close bodies
2. **Set timeouts** - Configure appropriate timeouts
3. **Handle errors** - Check and handle all errors
4. **Use context** - Use context for cancellation
5. **Validate input** - Validate all HTTP input
6. **Use HTTPS** - Use HTTPS for production
7. **Set headers** - Set appropriate headers

## 🔍 Advanced Features

### Custom Transport
```go
transport := &http.Transport{
    MaxIdleConns: 100,
    IdleConnTimeout: 90 * time.Second,
    TLSHandshakeTimeout: 10 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,
}

client := &http.Client{Transport: transport}
```

### Middleware Pattern
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

### File Server
```go
fs := http.FileServer(http.Dir("./static/"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## 📚 Real-world Applications

1. **Web APIs** - REST API implementation
2. **Microservices** - Service-to-service communication
3. **Web Scraping** - Data extraction from websites
4. **Load Testing** - HTTP performance testing
5. **Proxies** - HTTP proxy servers

## 🧠 Memory Tips

- **http** = **H**yper**T**ext **T**ransfer **P**rotocol
- **Get** = **G**ET request
- **Post** = **P**OST request
- **Do** = **D**o custom request
- **Serve** = **S**erve HTTP
- **Listen** = **L**isten for requests
- **Handle** = **H**andle requests
- **Response** = **R**esponse object

Remember: The http package is your gateway to web development in Go! 🎯
