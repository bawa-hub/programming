package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

// Custom types for demonstration
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Server struct {
	Port   int
	Routes map[string]http.HandlerFunc
}

func NewServer(port int) *Server {
	return &Server{
		Port:   port,
		Routes: make(map[string]http.HandlerFunc),
	}
}

func (s *Server) HandleFunc(pattern string, handler http.HandlerFunc) {
	s.Routes[pattern] = handler
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	
	for pattern, handler := range s.Routes {
		mux.HandleFunc(pattern, handler)
	}
	
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	return server.ListenAndServe()
}

func main() {
	fmt.Println("🚀 Go http Package Mastery Examples")
	fmt.Println("====================================")

	// 1. Basic HTTP Client Operations
	fmt.Println("\n1. Basic HTTP Client Operations:")
	
	// Simple GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Printf("Error making GET request: %v", err)
	} else {
		defer resp.Body.Close()
		
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Content Length: %d\n", resp.ContentLength)
		
		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
		} else {
			fmt.Printf("Response Body (first 200 chars): %s\n", string(body[:min(200, len(body))]))
		}
	}

	// 2. HTTP Client with Custom Configuration
	fmt.Println("\n2. HTTP Client with Custom Configuration:")
	
	// Create custom HTTP client
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
	
	// Make request with custom client
	req, err := http.NewRequest("GET", "https://httpbin.org/user-agent", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	} else {
		req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
		req.Header.Set("Accept", "application/json")
		
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v", err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("Custom client request successful: %s\n", resp.Status)
		}
	}

	// 3. POST Request with JSON Data
	fmt.Println("\n3. POST Request with JSON Data:")
	
	// Create JSON data
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	} else {
		// Make POST request
		resp, err := http.Post("https://httpbin.org/post", 
			"application/json", 
			bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error making POST request: %v", err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("POST request successful: %s\n", resp.Status)
			
			// Read response
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				fmt.Printf("Response (first 300 chars): %s\n", string(body[:min(300, len(body))]))
			}
		}
	}

	// 4. Form Data POST Request
	fmt.Println("\n4. Form Data POST Request:")
	
	// Create form data
	formData := map[string]string{
		"username": "johndoe",
		"password": "secret123",
		"email":    "john@example.com",
	}
	
	// Convert to form values
	formValues := make([]string, 0, len(formData))
	for key, value := range formData {
		formValues = append(formValues, fmt.Sprintf("%s=%s", key, value))
	}
	formString := strings.Join(formValues, "&")
	
	// Make POST request with form data
	resp, err = http.Post("https://httpbin.org/post", 
		"application/x-www-form-urlencoded", 
		strings.NewReader(formString))
	if err != nil {
		log.Printf("Error making form POST request: %v", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Form POST request successful: %s\n", resp.Status)
	}

	// 5. HTTP Server Implementation
	fmt.Println("\n5. HTTP Server Implementation:")
	
	// Start HTTP server in goroutine
	go func() {
		server := NewServer(8080)
		
		// Add routes
		server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World! Path: %s", r.URL.Path)
		})
		
		server.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			users := []User{
				{ID: 1, Name: "Alice", Email: "alice@example.com"},
				{ID: 2, Name: "Bob", Email: "bob@example.com"},
			}
			
			json.NewEncoder(w).Encode(APIResponse{
				Success: true,
				Data:    users,
				Message: "Users retrieved successfully",
			})
		})
		
		server.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(APIResponse{
				Success: true,
				Data:    map[string]string{"status": "healthy"},
				Message: "Server is running",
			})
		})
		
		fmt.Println("HTTP server starting on :8080")
		if err := server.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(1 * time.Second)
	
	// Test the server
	testServer := func() {
		// Test root endpoint
		resp, err := http.Get("http://localhost:8080/")
		if err != nil {
			log.Printf("Error testing root endpoint: %v", err)
		} else {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Root endpoint response: %s\n", string(body))
		}
		
		// Test API endpoint
		resp, err = http.Get("http://localhost:8080/api/users")
		if err != nil {
			log.Printf("Error testing API endpoint: %v", err)
		} else {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("API endpoint response: %s\n", string(body))
		}
		
		// Test health endpoint
		resp, err = http.Get("http://localhost:8080/api/health")
		if err != nil {
			log.Printf("Error testing health endpoint: %v", err)
		} else {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Health endpoint response: %s\n", string(body))
		}
	}
	
	testServer()

	// 6. HTTP Middleware
	fmt.Println("\n6. HTTP Middleware:")
	
	// Logging middleware
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			fmt.Printf("Request: %s %s - %v\n", r.Method, r.URL.Path, duration)
		})
	}
	
	// CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
	
	// Create handler with middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Middleware test - Method: %s, Path: %s", r.Method, r.URL.Path)
	})
	
	wrappedHandler := loggingMiddleware(corsMiddleware(handler))
	
	// Test middleware
	req, err = http.NewRequest("GET", "http://localhost:8080/test", nil)
	if err == nil {
		rr := httptest.NewRecorder()
		wrappedHandler.ServeHTTP(rr, req)
		fmt.Printf("Middleware test response: %s\n", rr.Body.String())
	}

	// 7. HTTP Headers and Cookies
	fmt.Println("\n7. HTTP Headers and Cookies:")
	
	// Test headers
	req, err = http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err == nil {
		req.Header.Set("X-Custom-Header", "MyValue")
		req.Header.Set("Authorization", "Bearer token123")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Headers response: %s\n", string(body[:min(300, len(body))]))
		}
	}
	
	// Test cookies
	req, err = http.NewRequest("GET", "https://httpbin.org/cookies", nil)
	if err == nil {
		req.AddCookie(&http.Cookie{Name: "session", Value: "abc123"})
		req.AddCookie(&http.Cookie{Name: "theme", Value: "dark"})
		
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Cookies response: %s\n", string(body[:min(300, len(body))]))
		}
	}

	// 8. HTTP File Server
	fmt.Println("\n8. HTTP File Server:")
	
	// Create a simple file server
	fileServer := http.FileServer(http.Dir("."))
	
	// Test file server (this would normally serve files)
	req, err = http.NewRequest("GET", "/", nil)
	if err == nil {
		rr := httptest.NewRecorder()
		fileServer.ServeHTTP(rr, req)
		fmt.Printf("File server response status: %d\n", rr.Code)
	}

	// 9. HTTP Context and Timeouts
	fmt.Println("\n9. HTTP Context and Timeouts:")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Make request with context
	req, err = http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	if err == nil {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Request with context error: %v\n", err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("Request with context successful: %s\n", resp.Status)
		}
	}
	
	// Test timeout
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	req, err = http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
	if err == nil {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Timeout request error (expected): %v\n", err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("Timeout request successful: %s\n", resp.Status)
		}
	}

	// 10. HTTP Redirect Handling
	fmt.Println("\n10. HTTP Redirect Handling:")
	
	// Create client that follows redirects
	redirectClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("Redirecting to: %s\n", req.URL.String())
			return nil
		},
	}
	
	// Test redirect
	resp, err = redirectClient.Get("https://httpbin.org/redirect/2")
	if err != nil {
		fmt.Printf("Redirect test error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Redirect test successful: %s\n", resp.Status)
	}

	// 11. HTTP Basic Authentication
	fmt.Println("\n11. HTTP Basic Authentication:")
	
	// Test basic auth
	req, err = http.NewRequest("GET", "https://httpbin.org/basic-auth/user/pass", nil)
	if err == nil {
		req.SetBasicAuth("user", "pass")
		
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Basic auth response: %s\n", string(body[:min(200, len(body))]))
		}
	}

	// 12. HTTP Error Handling
	fmt.Println("\n12. HTTP Error Handling:")
	
	// Test various HTTP status codes
	statusCodes := []int{200, 201, 400, 401, 403, 404, 500, 502, 503}
	
	for _, code := range statusCodes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", code)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Status %d: Error - %v\n", code, err)
		} else {
			resp.Body.Close()
			fmt.Printf("Status %d: %s\n", code, resp.Status)
		}
	}

	// 13. HTTP Performance Testing
	fmt.Println("\n13. HTTP Performance Testing:")
	
	// Simple performance test
	start := time.Now()
	successCount := 0
	totalRequests := 10
	
	for i := 0; i < totalRequests; i++ {
		resp, err := http.Get("https://httpbin.org/get")
		if err == nil {
			resp.Body.Close()
			successCount++
		}
	}
	
	duration := time.Since(start)
	fmt.Printf("Performance test: %d/%d requests successful in %v\n", 
		successCount, totalRequests, duration)
	fmt.Printf("Average response time: %v\n", duration/time.Duration(totalRequests))

	// 14. HTTP Custom Transport
	fmt.Println("\n14. HTTP Custom Transport:")
	
	// Create custom transport
	customTransport := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		DisableKeepAlives:   false,
		DisableCompression:  false,
		MaxIdleConnsPerHost: 2,
	}
	
	customClient := &http.Client{
		Transport: customTransport,
		Timeout:   10 * time.Second,
	}
	
	// Test custom transport
	resp, err = customClient.Get("https://httpbin.org/get")
	if err == nil {
		defer resp.Body.Close()
		fmt.Printf("Custom transport test successful: %s\n", resp.Status)
	}

	// 15. HTTP WebSocket-like Communication
	fmt.Println("\n15. HTTP WebSocket-like Communication:")
	
	// Simulate long polling
	longPollingHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		
		// Simulate waiting for data
		time.Sleep(2 * time.Second)
		
		response := APIResponse{
			Success: true,
			Data:    map[string]string{"message": "Long polling response"},
			Message: "Data received",
		}
		
		json.NewEncoder(w).Encode(response)
	}
	
	// Test long polling with a test server
	longPollServer := httptest.NewServer(http.HandlerFunc(longPollingHandler))
	defer longPollServer.Close()
	
	req, err = http.NewRequest("GET", longPollServer.URL+"/longpoll", nil)
	if err == nil {
		start := time.Now()
		resp, err := client.Do(req)
		duration := time.Since(start)
		
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Long polling response (%v): %s\n", duration, string(body))
		}
	}

	fmt.Println("\n🎉 http Package Mastery Complete!")
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
