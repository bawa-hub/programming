package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// BASIC PROXY PATTERN
// =============================================================================

// Subject interface
type Subject interface {
	Request() string
}

// Real Subject - the actual object
type RealSubject struct{}

func (rs *RealSubject) Request() string {
	return "RealSubject: Handling request"
}

// Proxy - controls access to the real subject
type Proxy struct {
	realSubject *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		fmt.Println("Proxy: Creating real subject")
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy: Forwarding request to real subject")
	return p.realSubject.Request()
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. VIRTUAL PROXY (Lazy Loading)
type Image interface {
	Display()
	GetWidth() int
	GetHeight() int
}

type RealImage struct {
	filename string
	width    int
	height   int
}

func NewRealImage(filename string) *RealImage {
	fmt.Printf("Loading image from disk: %s\n", filename)
	// Simulate expensive loading operation
	time.Sleep(100 * time.Millisecond)
	
	// Simulate different image sizes
	var width, height int
	switch filename {
	case "large.jpg":
		width, height = 1920, 1080
	case "medium.jpg":
		width, height = 800, 600
	case "small.jpg":
		width, height = 400, 300
	default:
		width, height = 200, 150
	}
	
	return &RealImage{
		filename: filename,
		width:    width,
		height:   height,
	}
}

func (ri *RealImage) Display() {
	fmt.Printf("Displaying image: %s (%dx%d)\n", ri.filename, ri.width, ri.height)
}

func (ri *RealImage) GetWidth() int {
	return ri.width
}

func (ri *RealImage) GetHeight() int {
	return ri.height
}

type ImageProxy struct {
	filename   string
	realImage  *RealImage
	mu         sync.RWMutex
}

func NewImageProxy(filename string) *ImageProxy {
	return &ImageProxy{filename: filename}
}

func (ip *ImageProxy) Display() {
	ip.mu.RLock()
	if ip.realImage == nil {
		ip.mu.RUnlock()
		ip.mu.Lock()
		if ip.realImage == nil {
			ip.realImage = NewRealImage(ip.filename)
		}
		ip.mu.Unlock()
		ip.mu.RLock()
	}
	ip.realImage.Display()
	ip.mu.RUnlock()
}

func (ip *ImageProxy) GetWidth() int {
	ip.mu.RLock()
	defer ip.mu.RUnlock()
	if ip.realImage == nil {
		return 0 // Not loaded yet
	}
	return ip.realImage.GetWidth()
}

func (ip *ImageProxy) GetHeight() int {
	ip.mu.RLock()
	defer ip.mu.RUnlock()
	if ip.realImage == nil {
		return 0 // Not loaded yet
	}
	return ip.realImage.GetHeight()
}

// 2. PROTECTION PROXY (Access Control)
type Database interface {
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
	GetConnectionInfo() string
}

type RealDatabase struct {
	host     string
	port     int
	username string
	password string
}

func NewRealDatabase(host string, port int, username, password string) *RealDatabase {
	return &RealDatabase{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (rd *RealDatabase) Query(sql string) ([]map[string]interface{}, error) {
	fmt.Printf("Executing query: %s\n", sql)
	return []map[string]interface{}{{"id": 1, "name": "John"}}, nil
}

func (rd *RealDatabase) Execute(sql string) error {
	fmt.Printf("Executing command: %s\n", sql)
	return nil
}

func (rd *RealDatabase) GetConnectionInfo() string {
	return fmt.Sprintf("%s:%d@%s:%d", rd.username, rd.port, rd.host, rd.port)
}

type DatabaseProxy struct {
	realDatabase *RealDatabase
	userRole     string
	allowedRoles map[string]bool
}

func NewDatabaseProxy(host string, port int, username, password, userRole string) *DatabaseProxy {
	allowedRoles := map[string]bool{
		"admin":    true,
		"user":     true,
		"readonly": true,
	}
	
	return &DatabaseProxy{
		realDatabase: NewRealDatabase(host, port, username, password),
		userRole:     userRole,
		allowedRoles: allowedRoles,
	}
}

func (dp *DatabaseProxy) Query(sql string) ([]map[string]interface{}, error) {
	if !dp.allowedRoles[dp.userRole] {
		return nil, fmt.Errorf("access denied: insufficient permissions")
	}
	fmt.Printf("User with role '%s' executing query\n", dp.userRole)
	return dp.realDatabase.Query(sql)
}

func (dp *DatabaseProxy) Execute(sql string) error {
	if dp.userRole == "readonly" {
		return fmt.Errorf("access denied: readonly users cannot execute commands")
	}
	if !dp.allowedRoles[dp.userRole] {
		return fmt.Errorf("access denied: insufficient permissions")
	}
	fmt.Printf("User with role '%s' executing command\n", dp.userRole)
	return dp.realDatabase.Execute(sql)
}

func (dp *DatabaseProxy) GetConnectionInfo() string {
	return dp.realDatabase.GetConnectionInfo()
}

// 3. CACHING PROXY
type ExpensiveService interface {
	Compute(input string) string
	GetComputationTime() time.Duration
}

type RealExpensiveService struct {
	computationTime time.Duration
}

func NewRealExpensiveService() *RealExpensiveService {
	return &RealExpensiveService{
		computationTime: 0,
	}
}

func (res *RealExpensiveService) Compute(input string) string {
	fmt.Printf("Computing expensive operation for input: %s\n", input)
	// Simulate expensive computation
	time.Sleep(200 * time.Millisecond)
	res.computationTime = 200 * time.Millisecond
	
	result := fmt.Sprintf("Result for %s: %d", input, len(input)*42)
	return result
}

func (res *RealExpensiveService) GetComputationTime() time.Duration {
	return res.computationTime
}

type CachingProxy struct {
	realService *RealExpensiveService
	cache       map[string]string
	mu          sync.RWMutex
}

func NewCachingProxy() *CachingProxy {
	return &CachingProxy{
		realService: NewRealExpensiveService(),
		cache:       make(map[string]string),
	}
}

func (cp *CachingProxy) Compute(input string) string {
	cp.mu.RLock()
	if result, exists := cp.cache[input]; exists {
		cp.mu.RUnlock()
		fmt.Printf("Cache hit for input: %s\n", input)
		return result
	}
	cp.mu.RUnlock()
	
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	// Double-check locking
	if result, exists := cp.cache[input]; exists {
		fmt.Printf("Cache hit for input: %s\n", input)
		return result
	}
	
	// Compute and cache result
	result := cp.realService.Compute(input)
	cp.cache[input] = result
	fmt.Printf("Cached result for input: %s\n", input)
	return result
}

func (cp *CachingProxy) GetComputationTime() time.Duration {
	return cp.realService.GetComputationTime()
}

func (cp *CachingProxy) GetCacheSize() int {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return len(cp.cache)
}

func (cp *CachingProxy) ClearCache() {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.cache = make(map[string]string)
	fmt.Println("Cache cleared")
}

// 4. LOGGING PROXY
type APIService interface {
	GetData(endpoint string) (map[string]interface{}, error)
	PostData(endpoint string, data map[string]interface{}) error
}

type RealAPIService struct {
	baseURL string
	apiKey  string
}

func NewRealAPIService(baseURL, apiKey string) *RealAPIService {
	return &RealAPIService{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (ras *RealAPIService) GetData(endpoint string) (map[string]interface{}, error) {
	fmt.Printf("Making GET request to %s%s\n", ras.baseURL, endpoint)
	// Simulate API call
	time.Sleep(50 * time.Millisecond)
	return map[string]interface{}{
		"data":    "response data",
		"status":  "success",
		"endpoint": endpoint,
	}, nil
}

func (ras *RealAPIService) PostData(endpoint string, data map[string]interface{}) error {
	fmt.Printf("Making POST request to %s%s with data: %v\n", ras.baseURL, endpoint, data)
	// Simulate API call
	time.Sleep(50 * time.Millisecond)
	return nil
}

type LoggingProxy struct {
	realService *RealAPIService
	logs        []string
	mu          sync.RWMutex
}

func NewLoggingProxy(baseURL, apiKey string) *LoggingProxy {
	return &LoggingProxy{
		realService: NewRealAPIService(baseURL, apiKey),
		logs:        make([]string, 0),
	}
}

func (lp *LoggingProxy) GetData(endpoint string) (map[string]interface{}, error) {
	startTime := time.Now()
	lp.log("GET", endpoint, nil, startTime)
	
	result, err := lp.realService.GetData(endpoint)
	
	duration := time.Since(startTime)
	lp.log("GET", endpoint, result, startTime)
	lp.logDuration("GET", endpoint, duration)
	
	return result, err
}

func (lp *LoggingProxy) PostData(endpoint string, data map[string]interface{}) error {
	startTime := time.Now()
	lp.log("POST", endpoint, data, startTime)
	
	err := lp.realService.PostData(endpoint, data)
	
	duration := time.Since(startTime)
	lp.logDuration("POST", endpoint, duration)
	
	return err
}

func (lp *LoggingProxy) log(method, endpoint string, data interface{}, timestamp time.Time) {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	
	logEntry := fmt.Sprintf("[%s] %s %s - Data: %v", 
		timestamp.Format("2006-01-02 15:04:05"), method, endpoint, data)
	lp.logs = append(lp.logs, logEntry)
}

func (lp *LoggingProxy) logDuration(method, endpoint string, duration time.Duration) {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	
	logEntry := fmt.Sprintf("[%s] %s %s - Duration: %v", 
		time.Now().Format("2006-01-02 15:04:05"), method, endpoint, duration)
	lp.logs = append(lp.logs, logEntry)
}

func (lp *LoggingProxy) GetLogs() []string {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return append([]string{}, lp.logs...)
}

func (lp *LoggingProxy) ClearLogs() {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	lp.logs = make([]string, 0)
}

// 5. REMOTE PROXY (Network Communication)
type RemoteService interface {
	ProcessRequest(request string) (string, error)
	GetServiceInfo() string
}

type RealRemoteService struct {
	serviceURL string
	timeout    time.Duration
}

func NewRealRemoteService(serviceURL string, timeout time.Duration) *RealRemoteService {
	return &RealRemoteService{
		serviceURL: serviceURL,
		timeout:    timeout,
	}
}

func (rrs *RealRemoteService) ProcessRequest(request string) (string, error) {
	fmt.Printf("Sending request to remote service: %s\n", rrs.serviceURL)
	fmt.Printf("Request: %s\n", request)
	
	// Simulate network delay
	time.Sleep(rrs.timeout)
	
	response := fmt.Sprintf("Response from %s: %s processed", rrs.serviceURL, request)
	return response, nil
}

func (rrs *RealRemoteService) GetServiceInfo() string {
	return fmt.Sprintf("Remote service at %s with timeout %v", rrs.serviceURL, rrs.timeout)
}

type RemoteProxy struct {
	realService *RealRemoteService
	cache       map[string]string
	mu          sync.RWMutex
}

func NewRemoteProxy(serviceURL string, timeout time.Duration) *RemoteProxy {
	return &RemoteProxy{
		realService: NewRealRemoteService(serviceURL, timeout),
		cache:       make(map[string]string),
	}
}

func (rp *RemoteProxy) ProcessRequest(request string) (string, error) {
	// Check cache first
	rp.mu.RLock()
	if response, exists := rp.cache[request]; exists {
		rp.mu.RUnlock()
		fmt.Printf("Cache hit for request: %s\n", request)
		return response, nil
	}
	rp.mu.RUnlock()
	
	// Process request
	response, err := rp.realService.ProcessRequest(request)
	if err != nil {
		return "", err
	}
	
	// Cache response
	rp.mu.Lock()
	rp.cache[request] = response
	rp.mu.Unlock()
	
	return response, nil
}

func (rp *RemoteProxy) GetServiceInfo() string {
	return rp.realService.GetServiceInfo()
}

func (rp *RemoteProxy) GetCacheSize() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()
	return len(rp.cache)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== PROXY PATTERN DEMONSTRATION ===\n")

	// 1. BASIC PROXY
	fmt.Println("1. BASIC PROXY:")
	proxy := NewProxy()
	fmt.Println(proxy.Request())
	fmt.Println(proxy.Request()) // Should reuse real subject
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Virtual Proxy (Lazy Loading)
	fmt.Println("Virtual Proxy (Lazy Loading):")
	images := []Image{
		NewImageProxy("large.jpg"),
		NewImageProxy("medium.jpg"),
		NewImageProxy("small.jpg"),
	}
	
	for _, image := range images {
		fmt.Printf("Image dimensions: %dx%d\n", image.GetWidth(), image.GetHeight())
		image.Display()
	}
	fmt.Println()

	// Protection Proxy (Access Control)
	fmt.Println("Protection Proxy (Access Control):")
	adminDB := NewDatabaseProxy("localhost", 5432, "admin", "password", "admin")
	userDB := NewDatabaseProxy("localhost", 5432, "user", "password", "user")
	readonlyDB := NewDatabaseProxy("localhost", 5432, "readonly", "password", "readonly")
	
	// Test different access levels
	databases := []Database{adminDB, userDB, readonlyDB}
	roles := []string{"admin", "user", "readonly"}
	
	for i, db := range databases {
		fmt.Printf("Testing %s role:\n", roles[i])
		if result, err := db.Query("SELECT * FROM users"); err != nil {
			fmt.Printf("  Query failed: %v\n", err)
		} else {
			fmt.Printf("  Query result: %v\n", result)
		}
		
		if err := db.Execute("INSERT INTO users (name) VALUES ('John')"); err != nil {
			fmt.Printf("  Execute failed: %v\n", err)
		} else {
			fmt.Printf("  Execute successful\n")
		}
	}
	fmt.Println()

	// Caching Proxy
	fmt.Println("Caching Proxy:")
	cachingProxy := NewCachingProxy()
	
	inputs := []string{"input1", "input2", "input1", "input3", "input2"}
	for _, input := range inputs {
		start := time.Now()
		result := cachingProxy.Compute(input)
		duration := time.Since(start)
		fmt.Printf("Input: %s, Result: %s, Time: %v\n", input, result, duration)
	}
	
	fmt.Printf("Cache size: %d\n", cachingProxy.GetCacheSize())
	cachingProxy.ClearCache()
	fmt.Printf("Cache size after clear: %d\n", cachingProxy.GetCacheSize())
	fmt.Println()

	// Logging Proxy
	fmt.Println("Logging Proxy:")
	loggingProxy := NewLoggingProxy("https://api.example.com", "api_key_123")
	
	// Make some API calls
	loggingProxy.GetData("/users")
	loggingProxy.PostData("/users", map[string]interface{}{"name": "John", "email": "john@example.com"})
	loggingProxy.GetData("/products")
	
	// Display logs
	fmt.Println("API Logs:")
	for _, log := range loggingProxy.GetLogs() {
		fmt.Printf("  %s\n", log)
	}
	fmt.Println()

	// Remote Proxy
	fmt.Println("Remote Proxy:")
	remoteProxy := NewRemoteProxy("https://remote-service.com/api", 100*time.Millisecond)
	
	requests := []string{"request1", "request2", "request1", "request3"}
	for _, request := range requests {
		start := time.Now()
		response, err := remoteProxy.ProcessRequest(request)
		duration := time.Since(start)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
		} else {
			fmt.Printf("Request: %s, Response: %s, Time: %v\n", request, response, duration)
		}
	}
	
	fmt.Printf("Remote service info: %s\n", remoteProxy.GetServiceInfo())
	fmt.Printf("Cache size: %d\n", remoteProxy.GetCacheSize())
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
