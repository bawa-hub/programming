package main

import (
	"fmt"
	"time"
)

// 🎯 INTERFACE COMPOSITION MASTERY
// Understanding interface composition and embedding in Go

func main() {
	fmt.Println("🎯 INTERFACE COMPOSITION MASTERY")
	fmt.Println("=================================")

	// 1. Basic Interface Composition
	fmt.Println("\n1. Basic Interface Composition:")
	basicInterfaceComposition()

	// 2. Interface Embedding
	fmt.Println("\n2. Interface Embedding:")
	interfaceEmbedding()

	// 3. Interface Segregation
	fmt.Println("\n3. Interface Segregation:")
	interfaceSegregation()

	// 4. Interface Composition Patterns
	fmt.Println("\n4. Interface Composition Patterns:")
	interfaceCompositionPatterns()

	// 5. Real-world Interface Design
	fmt.Println("\n5. Real-world Interface Design:")
	realWorldInterfaceDesign()

	// 6. Interface Best Practices
	fmt.Println("\n6. Interface Best Practices:")
	interfaceBestPractices()
}

// BASIC INTERFACE COMPOSITION: Understanding interface composition
func basicInterfaceComposition() {
	fmt.Println("Understanding basic interface composition...")
	
	// Define basic interfaces
	var reader Reader = &FileReader{name: "test.txt"}
	var writer Writer = &FileWriter{name: "output.txt"}
	
	// Use interfaces
	data, _ := reader.Read()
	fmt.Printf("  📖 Read data: %s\n", data)
	
	writer.Write("Hello, World!")
	fmt.Println("  ✏️  Data written")
	
	// Compose interfaces
	var rw ReadWriter = &FileReadWriter{
		Reader: reader,
		Writer: writer,
	}
	
	rw.Write("Composed data")
	data, _ = rw.Read()
	fmt.Printf("  📖 Composed read: %s\n", data)
}

// INTERFACE EMBEDDING: Understanding interface embedding
func interfaceEmbedding() {
	fmt.Println("Understanding interface embedding...")
	
	// Define embedded interfaces
	var logger Logger = &ConsoleLogger{}
	var metrics Metrics = &BasicMetrics{}
	
	// Use embedded interfaces
	logger.Info("Application started")
	metrics.Increment("requests")
	
	// Create composed interface
	service := &UserService{
		Logger:  logger,
		Metrics: metrics,
	}
	
	service.ProcessUser("john_doe")
	service.Info("User processed")
	service.Increment("users_processed")
}

// INTERFACE SEGREGATION: Understanding interface segregation
func interfaceSegregation() {
	fmt.Println("Understanding interface segregation...")
	
	// Create specific interfaces
	userRepo := &InMemoryUserRepository{}
	emailService := &SMTPEmailService{}
	cacheService := &RedisCacheService{}
	
	// Use only what you need
	user, err := userRepo.FindByID("123")
	if err != nil {
		fmt.Printf("  ❌ User not found: %v\n", err)
	} else {
		fmt.Printf("  👤 Found user: %s\n", user.Name)
	}
	
	emailService.SendEmail("user@example.com", "Welcome!")
	fmt.Println("  📧 Email sent")
	
	cacheService.Set("key", "value", 5*time.Minute)
	value, _ := cacheService.Get("key")
	fmt.Printf("  💾 Cached value: %s\n", value)
}

// INTERFACE COMPOSITION PATTERNS: Advanced composition patterns
func interfaceCompositionPatterns() {
	fmt.Println("Understanding interface composition patterns...")
	
	// Pattern 1: Facade Pattern
	var facade PaymentFacade = &PaymentProcessor{
		validator: &CreditCardValidator{},
		notifier:  &EmailNotifier{},
	}
	
	payment := Payment{
		Amount:   100.00,
		Currency: "USD",
		Card:     "4111-1111-1111-1111",
	}
	
	result := facade.ProcessPayment(payment)
	fmt.Printf("  💳 Payment result: %s\n", result.Status)
	
	// Pattern 2: Strategy Pattern
	var strategies []SortStrategy
	strategies = append(strategies, &QuickSort{})
	strategies = append(strategies, &MergeSort{})
	strategies = append(strategies, &BubbleSort{})
	
	data := []int{64, 34, 25, 12, 22, 11, 90}
	
	for i, strategy := range strategies {
		sorted := make([]int, len(data))
		copy(sorted, data)
		strategy.Sort(sorted)
		fmt.Printf("  🔄 Strategy %d result: %v\n", i+1, sorted)
	}
}

// REAL-WORLD INTERFACE DESIGN: Practical interface design
func realWorldInterfaceDesign() {
	fmt.Println("Understanding real-world interface design...")
	
	// Create a real-world service
	var db Database = &PostgreSQLDatabase{}
	var cache Cache = &RedisCache{}
	var logger Logger = &StructuredLogger{}
	
	service := &UserService{
		Logger:  logger,
		Metrics: &BasicMetrics{},
		Database: db,
		Cache:    cache,
	}
	
	// Use the service
	user := &User{
		ID:    "123",
		Name:  "John Doe",
		Email: "john@example.com",
	}
	
	service.CreateUser(user)
	fmt.Println("  👤 User created successfully")
	
	// Test caching
	cachedUser, _ := service.GetUser("123")
	fmt.Printf("  💾 Cached user: %s\n", cachedUser.Name)
}

// INTERFACE BEST PRACTICES: Following Go interface best practices
func interfaceBestPractices() {
	fmt.Println("Understanding interface best practices...")
	
	// 1. Keep interfaces small
	fmt.Println("  📝 Best Practice 1: Keep interfaces small")
	smallInterfaces()
	
	// 2. Accept interfaces, return concrete types
	fmt.Println("  📝 Best Practice 2: Accept interfaces, return concrete types")
	interfaceAcceptance()
	
	// 3. Use interface composition
	fmt.Println("  📝 Best Practice 3: Use interface composition")
	interfaceComposition()
	
	// 4. Design for testability
	fmt.Println("  📝 Best Practice 4: Design for testability")
	testableDesign()
}

func smallInterfaces() {
	// Good: Small, focused interfaces
	var reader Reader = &FileReader{name: "test.txt"}
	var writer Writer = &FileWriter{name: "output.txt"}
	
	// Use small interfaces
	data, _ := reader.Read()
	writer.Write(data)
	
	fmt.Printf("    Small interfaces work: %s\n", data)
}

func interfaceAcceptance() {
	// Good: Accept interfaces
	processData(&FileReader{name: "input.txt"})
	processData(&NetworkReader{url: "https://api.example.com/data"})
}

func processData(reader Reader) {
	data, _ := reader.Read()
	fmt.Printf("    Processed data: %s\n", data)
}

func interfaceComposition() {
	// Good: Compose interfaces
	service := &UserService{
		Logger:  &ConsoleLogger{},
		Metrics: &BasicMetrics{},
	}
	
	service.ProcessUser("test_user")
	fmt.Println("    Composed service works")
}

func testableDesign() {
	// Good: Easy to test with mocks
	var mockRepo UserRepository = &MockUserRepository{}
	service := &UserService{
		Logger:     &ConsoleLogger{},
		Metrics:    &BasicMetrics{},
		Database:   &PostgreSQLDatabase{},
		Cache:      &RedisCache{},
		Repository: mockRepo,
	}
	
	user, _ := service.GetUser("123")
	fmt.Printf("    Testable design: %s\n", user.Name)
}

// INTERFACE DEFINITIONS

// Basic interfaces
type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(data string) error
}

type ReadWriter interface {
	Reader
	Writer
}

// Embedded interfaces
type Logger interface {
	Info(message string)
	Error(message string)
}

type Metrics interface {
	Increment(name string)
	Timing(name string, duration time.Duration)
}

type Service interface {
	Logger
	Metrics
	ProcessUser(userID string)
}

// Segregated interfaces
type UserRepository interface {
	FindByID(id string) (*User, error)
	Save(user *User) error
	Delete(id string) error
}

type EmailService interface {
	SendEmail(to, subject string) error
}

type CacheService interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
}

// Composition patterns
type PaymentFacade interface {
	ProcessPayment(payment Payment) PaymentResult
}

type SortStrategy interface {
	Sort(data []int)
}

type PaymentValidator interface {
	Validate(payment Payment) error
}

type PaymentNotifier interface {
	Notify(payment Payment, result PaymentResult)
}

// Real-world interfaces
type Database interface {
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
}

type Cache interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
}

// IMPLEMENTATIONS

// Basic implementations
type FileReader struct {
	name string
}

func (f *FileReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", f.name), nil
}

type FileWriter struct {
	name string
}

func (f *FileWriter) Write(data string) error {
	fmt.Printf("Writing to %s: %s\n", f.name, data)
	return nil
}

type FileReadWriter struct {
	Reader
	Writer
}

// Embedded implementations
type ConsoleLogger struct{}

func (c *ConsoleLogger) Info(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func (c *ConsoleLogger) Error(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

type BasicMetrics struct{}

func (b *BasicMetrics) Increment(name string) {
	fmt.Printf("Metric incremented: %s\n", name)
}

func (b *BasicMetrics) Timing(name string, duration time.Duration) {
	fmt.Printf("Timing: %s = %v\n", name, duration)
}

type UserService struct {
	Logger
	Metrics
	Database
	Cache
	Repository UserRepository
}

func (u *UserService) ProcessUser(userID string) {
	u.Info(fmt.Sprintf("Processing user: %s", userID))
	u.Increment("users_processed")
}

func (u *UserService) LogInfo(message string) {
	u.Info(message)
}

func (u *UserService) IncrementMetric(name string) {
	u.Increment(name)
}

func (u *UserService) CreateUser(user *User) error {
	u.Info(fmt.Sprintf("Creating user: %s", user.Name))
	return nil
}

func (u *UserService) GetUser(id string) (*User, error) {
	u.Info(fmt.Sprintf("Getting user: %s", id))
	return &User{ID: id, Name: "John Doe"}, nil
}

// Segregated implementations
type InMemoryUserRepository struct {
	users map[string]*User
}

func (r *InMemoryUserRepository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(user *User) error {
	if r.users == nil {
		r.users = make(map[string]*User)
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	delete(r.users, id)
	return nil
}

type SMTPEmailService struct{}

func (s *SMTPEmailService) SendEmail(to, subject string) error {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}

type RedisCacheService struct{}

func (r *RedisCacheService) Set(key, value string, ttl time.Duration) error {
	fmt.Printf("Caching %s = %s (TTL: %v)\n", key, value, ttl)
	return nil
}

func (r *RedisCacheService) Get(key string) (string, error) {
	return fmt.Sprintf("cached_value_for_%s", key), nil
}

// Composition pattern implementations
type PaymentProcessor struct {
	validator PaymentValidator
	notifier  PaymentNotifier
}

func (p *PaymentProcessor) ProcessPayment(payment Payment) PaymentResult {
	// Validate payment
	if err := p.validator.Validate(payment); err != nil {
		return PaymentResult{Status: "failed", Error: err.Error()}
	}
	
	// Process payment
	result := PaymentResult{Status: "success", TransactionID: "txn_123"}
	
	// Notify user
	p.notifier.Notify(payment, result)
	
	return result
}

type CreditCardValidator struct{}

func (c *CreditCardValidator) Validate(payment Payment) error {
	if payment.Card == "" {
		return fmt.Errorf("card number required")
	}
	return nil
}

type StripeProcessor struct{}

func (s *StripeProcessor) Process(payment Payment) PaymentResult {
	return PaymentResult{Status: "success", TransactionID: "txn_123"}
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(payment Payment, result PaymentResult) {
	fmt.Printf("Payment notification: %s\n", result.Status)
}

// Strategy pattern implementations
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) {
	// Simplified quicksort
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

type MergeSort struct{}

func (m *MergeSort) Sort(data []int) {
	// Simplified merge sort
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) {
	// Bubble sort
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// Real-world implementations
type PostgreSQLDatabase struct{}

func (p *PostgreSQLDatabase) Query(sql string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{{"id": "123", "name": "John"}}, nil
}

func (p *PostgreSQLDatabase) Execute(sql string) error {
	fmt.Printf("Executing SQL: %s\n", sql)
	return nil
}

type RedisCache struct{}

func (r *RedisCache) Set(key, value string, ttl time.Duration) error {
	fmt.Printf("Redis SET %s = %s\n", key, value)
	return nil
}

func (r *RedisCache) Get(key string) (string, error) {
	return fmt.Sprintf("redis_value_for_%s", key), nil
}

type StructuredLogger struct{}

func (s *StructuredLogger) Info(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func (s *StructuredLogger) Error(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

// Mock implementations
type MockUserRepository struct{}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
	return &User{ID: id, Name: "Mock User"}, nil
}

func (m *MockUserRepository) Save(user *User) error {
	return nil
}

func (m *MockUserRepository) Delete(id string) error {
	return nil
}

type NetworkReader struct {
	url string
}

func (n *NetworkReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", n.url), nil
}

// Data structures
type User struct {
	ID    string
	Name  string
	Email string
}

type Payment struct {
	Amount   float64
	Currency string
	Card     string
}

type PaymentResult struct {
	Status        string
	TransactionID string
	Error         string
}

