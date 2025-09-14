package main

import (
	"errors"
	"fmt"
	"time"
)

// 🏗️ CLEAN ARCHITECTURE MASTERY
// Understanding domain-driven design and clean architecture principles

func main() {
	fmt.Println("🏗️ CLEAN ARCHITECTURE MASTERY")
	fmt.Println("==============================")

	// 1. Domain-Driven Design
	fmt.Println("\n1. Domain-Driven Design:")
	domainDrivenDesign()

	// 2. Hexagonal Architecture
	fmt.Println("\n2. Hexagonal Architecture:")
	hexagonalArchitecture()

	// 3. Dependency Injection
	fmt.Println("\n3. Dependency Injection:")
	dependencyInjection()

	// 4. Clean Code Principles
	fmt.Println("\n4. Clean Code Principles:")
	cleanCodePrinciples()

	// 5. Testable Architecture
	fmt.Println("\n5. Testable Architecture:")
	testableArchitecture()

	// 6. Real-world Example
	fmt.Println("\n6. Real-world Example:")
	realWorldExample()
}

// DOMAIN-DRIVEN DESIGN: Understanding DDD principles
func domainDrivenDesign() {
	fmt.Println("Understanding domain-driven design...")
	
	// Create domain entities
	user := NewUser("john_doe", "john@example.com")
	order := NewOrder(user.ID, 100.00)
	
	// Add order items
	order.AddItem("laptop", 1, 800.00)
	order.AddItem("mouse", 2, 25.00)
	
	fmt.Printf("  📊 User: %s (%s)\n", user.Username, user.Email)
	fmt.Printf("  📊 Order: %s, Total: $%.2f\n", order.ID, order.CalculateTotal())
	
	// Domain services
	orderService := NewOrderService()
	
	// Apply business rules
	if err := orderService.ValidateOrder(order); err != nil {
		fmt.Printf("  ❌ Order validation failed: %v\n", err)
	} else {
		fmt.Printf("  ✅ Order is valid\n")
	}
	
	// Process order
	if err := orderService.ProcessOrder(order); err != nil {
		fmt.Printf("  ❌ Order processing failed: %v\n", err)
	} else {
		fmt.Printf("  ✅ Order processed successfully\n")
	}
}

// HEXAGONAL ARCHITECTURE: Understanding ports and adapters
func hexagonalArchitecture() {
	fmt.Println("Understanding hexagonal architecture...")
	
	// Create ports (interfaces)
	var userRepo UserRepository = NewInMemoryUserRepository()
	var emailService EmailService = &SMTPEmailService{}
	var paymentService PaymentService = &StripePaymentService{}
	
	// Create application service (hexagon)
	userService := NewUserService(userRepo, emailService, paymentService)
	
	// Use case: Register new user
	user, err := userService.RegisterUser("jane_doe", "jane@example.com")
	if err != nil {
		fmt.Printf("  ❌ Registration failed: %v\n", err)
	} else {
		fmt.Printf("  ✅ User registered: %s\n", user.Username)
	}
	
	// Use case: Process payment
	payment := Payment{
		UserID:  user.ID,
		Amount:  50.00,
		Currency: "USD",
	}
	
	if err := userService.ProcessPayment(payment); err != nil {
		fmt.Printf("  ❌ Payment failed: %v\n", err)
	} else {
		fmt.Printf("  ✅ Payment processed: $%.2f\n", payment.Amount)
	}
}

// DEPENDENCY INJECTION: Understanding DI patterns
func dependencyInjection() {
	fmt.Println("Understanding dependency injection...")
	
	// Create dependencies
	logger := &ConsoleLogger{}
	metrics := &PrometheusMetrics{}
	cache := &RedisCache{}
	
	// Inject dependencies
	service := NewProductService(logger, metrics, cache)
	
	// Use service
	product := Product{
		ID:    "prod-123",
		Name:  "Laptop",
		Price: 999.99,
	}
	
	service.CreateProduct(product)
	fmt.Printf("  📊 Product created: %s\n", product.Name)
	
	// Test with mock dependencies
	mockLogger := &MockLogger{}
	mockMetrics := &MockMetrics{}
	mockCache := &MockCache{}
	
	testService := NewProductService(mockLogger, mockMetrics, mockCache)
	testService.CreateProduct(product)
	fmt.Printf("  📊 Test service used mock dependencies\n")
}

// CLEAN CODE PRINCIPLES: Understanding clean code
func cleanCodePrinciples() {
	fmt.Println("Understanding clean code principles...")
	
	// 1. Single Responsibility Principle
	fmt.Println("  📝 Single Responsibility Principle:")
	singleResponsibilityPrinciple()
	
	// 2. Open/Closed Principle
	fmt.Println("  📝 Open/Closed Principle:")
	openClosedPrinciple()
	
	// 3. Liskov Substitution Principle
	fmt.Println("  📝 Liskov Substitution Principle:")
	liskovSubstitutionPrinciple()
	
	// 4. Interface Segregation Principle
	fmt.Println("  📝 Interface Segregation Principle:")
	interfaceSegregationPrinciple()
	
	// 5. Dependency Inversion Principle
	fmt.Println("  📝 Dependency Inversion Principle:")
	dependencyInversionPrinciple()
}

func singleResponsibilityPrinciple() {
	// Good: Each class has one reason to change
	userRepo := NewInMemoryUserRepository()
	emailService := &SMTPEmailService{}
	
	// User service only handles user operations
	userService := NewUserService(userRepo, emailService, nil)
	
	// Email service only handles email operations
	user, _ := userService.RegisterUser("alice", "alice@example.com")
	emailService.SendWelcomeEmail(user.Email)
	
	fmt.Printf("    ✅ Each service has single responsibility\n")
}

func openClosedPrinciple() {
	// Good: Open for extension, closed for modification
	var processors []PaymentProcessor
	
	// Add different payment processors
	processors = append(processors, &CreditCardProcessor{})
	processors = append(processors, &PayPalProcessor{})
	processors = append(processors, &BankTransferProcessor{})
	
	// Process payments without modifying existing code
	payment := Payment{Amount: 100.00, Currency: "USD"}
	
	for i, processor := range processors {
		result := processor.Process(payment)
		fmt.Printf("    Processor %d: %s\n", i+1, result.Status)
	}
}

func liskovSubstitutionPrinciple() {
	// Good: Subtypes must be substitutable for their base types
	var shapes []Shape
	
	shapes = append(shapes, &Circle{Radius: 5.0})
	shapes = append(shapes, &Rectangle{Width: 10, Height: 8})
	shapes = append(shapes, &Triangle{Base: 6, Height: 4})
	
	// All shapes can be used interchangeably
	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()
		fmt.Printf("    Shape %d: Area=%.2f, Perimeter=%.2f\n", 
			i+1, area, perimeter)
	}
}

func interfaceSegregationPrinciple() {
	// Good: Clients should not depend on interfaces they don't use
	var readers []Reader
	var writers []Writer
	
	// Readers only need Read method
	readers = append(readers, &FileReader{name: "input.txt"})
	readers = append(readers, &NetworkReader{url: "https://api.example.com"})
	
	// Writers only need Write method
	writers = append(writers, &FileWriter{name: "output.txt"})
	writers = append(writers, &ConsoleWriter{})
	
	// Use only what you need
	for _, reader := range readers {
		data, _ := reader.Read()
		fmt.Printf("    Read: %s\n", data)
	}
	
	for _, writer := range writers {
		writer.Write("Hello, World!")
	}
}

func dependencyInversionPrinciple() {
	// Good: Depend on abstractions, not concretions
	var logger Logger = &ConsoleLogger{}
	var metrics Metrics = &PrometheusMetrics{}
	
	// Service depends on abstractions
	service := NewProductService(logger, metrics, nil)
	
	// Can easily swap implementations
	service = NewProductService(&FileLogger{}, &InMemoryMetrics{}, nil)
	_ = service
	
	fmt.Printf("    ✅ Service depends on abstractions\n")
}

// TESTABLE ARCHITECTURE: Understanding testable design
func testableArchitecture() {
	fmt.Println("Understanding testable architecture...")
	
	// Create testable service with mock dependencies
	mockRepo := NewMockUserRepository()
	mockEmail := &MockEmailService{}
	mockPayment := &MockPaymentService{}
	
	service := NewUserService(mockRepo, mockEmail, mockPayment)
	
	// Test user registration
	user, err := service.RegisterUser("test_user", "test@example.com")
	if err != nil {
		fmt.Printf("  ❌ Test failed: %v\n", err)
	} else {
		fmt.Printf("  ✅ Test passed: %s\n", user.Username)
	}
	
	// Verify mock interactions
	if mockRepo.SaveCalled {
		fmt.Printf("  ✅ Repository.Save was called\n")
	}
	if mockEmail.SendCalled {
		fmt.Printf("  ✅ EmailService.Send was called\n")
	}
}

// REAL-WORLD EXAMPLE: Complete clean architecture example
func realWorldExample() {
	fmt.Println("Understanding real-world clean architecture...")
	
	// Create a complete e-commerce system
	system := NewECommerceSystem()
	
	// Register user
	user, err := system.RegisterUser("customer", "customer@example.com")
	if err != nil {
		fmt.Printf("  ❌ Registration failed: %v\n", err)
		return
	}
	
	// Create product
	product := Product{
		ID:    "prod-001",
		Name:  "Gaming Laptop",
		Price: 1299.99,
	}
	system.CreateProduct(product)
	
	// Add to cart
	cart := system.GetCart(user.ID)
	cart.AddItem(product, 1)
	
	// Checkout
	order, err := system.Checkout(user.ID)
	if err != nil {
		fmt.Printf("  ❌ Checkout failed: %v\n", err)
		return
	}
	
	fmt.Printf("  ✅ Order created: %s, Total: $%.2f\n", 
		order.ID, order.CalculateTotal())
}

// DOMAIN MODELS

type User struct {
	ID       string
	Username string
	Email    string
	Created  time.Time
}

func NewUser(username, email string) *User {
	return &User{
		ID:       generateID(),
		Username: username,
		Email:    email,
		Created:  time.Now(),
	}
}

type Order struct {
	ID     string
	UserID string
	Items  []OrderItem
	Total  float64
	Status OrderStatus
}

type OrderItem struct {
	ProductName string
	Quantity    int
	Price       float64
}

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderConfirmed
	OrderShipped
	OrderDelivered
	OrderCancelled
)

func NewOrder(userID string, total float64) *Order {
	return &Order{
		ID:     generateID(),
		UserID: userID,
		Items:  make([]OrderItem, 0),
		Total:  total,
		Status: OrderPending,
	}
}

func (o *Order) AddItem(productName string, quantity int, price float64) {
	item := OrderItem{
		ProductName: productName,
		Quantity:    quantity,
		Price:       price,
	}
	o.Items = append(o.Items, item)
}

func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

type Payment struct {
	UserID   string
	Amount   float64
	Currency string
}

// DOMAIN SERVICES

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) ValidateOrder(order *Order) error {
	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}
	
	if order.CalculateTotal() <= 0 {
		return errors.New("order total must be greater than zero")
	}
	
	return nil
}

func (s *OrderService) ProcessOrder(order *Order) error {
	// Business logic for processing order
	order.Status = OrderConfirmed
	return nil
}

// PORTS (INTERFACES)

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}

type EmailService interface {
	SendWelcomeEmail(email string) error
	SendOrderConfirmation(email string, order *Order) error
}

type PaymentService interface {
	ProcessPayment(payment Payment) error
}

type Logger interface {
	Info(message string)
	Error(message string)
}

type Metrics interface {
	IncrementCounter(name string)
	RecordDuration(name string, duration time.Duration)
}

type Cache interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
}

// ADAPTERS (IMPLEMENTATIONS)

type InMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

type SMTPEmailService struct{}

func (s *SMTPEmailService) SendWelcomeEmail(email string) error {
	fmt.Printf("  📧 Welcome email sent to %s\n", email)
	return nil
}

func (s *SMTPEmailService) SendOrderConfirmation(email string, order *Order) error {
	fmt.Printf("  📧 Order confirmation sent to %s for order %s\n", email, order.ID)
	return nil
}

type StripePaymentService struct{}

func (s *StripePaymentService) ProcessPayment(payment Payment) error {
	fmt.Printf("  💳 Payment processed: $%.2f %s\n", payment.Amount, payment.Currency)
	return nil
}

// APPLICATION SERVICES

type UserService struct {
	userRepo    UserRepository
	emailService EmailService
	paymentService PaymentService
}

func NewUserService(userRepo UserRepository, emailService EmailService, paymentService PaymentService) *UserService {
	return &UserService{
		userRepo:      userRepo,
		emailService:  emailService,
		paymentService: paymentService,
	}
}

func (s *UserService) RegisterUser(username, email string) (*User, error) {
	// Check if user already exists
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return nil, errors.New("user already exists")
	}
	
	// Create user
	user := NewUser(username, email)
	
	// Save user
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}
	
	// Send welcome email
	if err := s.emailService.SendWelcomeEmail(user.Email); err != nil {
		// Log error but don't fail registration
		fmt.Printf("  ⚠️  Failed to send welcome email: %v\n", err)
	}
	
	return user, nil
}

func (s *UserService) ProcessPayment(payment Payment) error {
	return s.paymentService.ProcessPayment(payment)
}

// CLEAN CODE IMPLEMENTATIONS

type ProductService struct {
	logger  Logger
	metrics Metrics
	cache   Cache
}

func NewProductService(logger Logger, metrics Metrics, cache Cache) *ProductService {
	return &ProductService{
		logger:  logger,
		metrics: metrics,
		cache:   cache,
	}
}

func (s *ProductService) CreateProduct(product Product) {
	s.logger.Info(fmt.Sprintf("Creating product: %s", product.Name))
	s.metrics.IncrementCounter("products_created")
	
	// Cache product
	if s.cache != nil {
		s.cache.Set(fmt.Sprintf("product:%s", product.ID), product.Name, time.Hour)
	}
}

// PAYMENT PROCESSORS

type PaymentProcessor interface {
	Process(payment Payment) PaymentResult
}

type PaymentResult struct {
	Status string
	Error  string
}

type CreditCardProcessor struct{}

func (p *CreditCardProcessor) Process(payment Payment) PaymentResult {
	return PaymentResult{Status: "success"}
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) Process(payment Payment) PaymentResult {
	return PaymentResult{Status: "success"}
}

type BankTransferProcessor struct{}

func (p *BankTransferProcessor) Process(payment Payment) PaymentResult {
	return PaymentResult{Status: "pending"}
}

// SHAPES FOR LSP DEMONSTRATION

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t *Triangle) Perimeter() float64 {
	return t.Base + t.Height + t.Base
}

// READER/WRITER INTERFACES

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(data string) error
}

type FileReader struct {
	name string
}

func (f *FileReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", f.name), nil
}

type NetworkReader struct {
	url string
}

func (n *NetworkReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", n.url), nil
}

type FileWriter struct {
	name string
}

func (f *FileWriter) Write(data string) error {
	fmt.Printf("Writing to %s: %s\n", f.name, data)
	return nil
}

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(data string) error {
	fmt.Printf("Console: %s\n", data)
	return nil
}

// LOGGER IMPLEMENTATIONS

type ConsoleLogger struct{}

func (l *ConsoleLogger) Info(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func (l *ConsoleLogger) Error(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

type FileLogger struct{}

func (l *FileLogger) Info(message string) {
	fmt.Printf("FILE INFO: %s\n", message)
}

func (l *FileLogger) Error(message string) {
	fmt.Printf("FILE ERROR: %s\n", message)
}

// METRICS IMPLEMENTATIONS

type PrometheusMetrics struct{}

func (m *PrometheusMetrics) IncrementCounter(name string) {
	fmt.Printf("METRIC: %s incremented\n", name)
}

func (m *PrometheusMetrics) RecordDuration(name string, duration time.Duration) {
	fmt.Printf("METRIC: %s duration %v\n", name, duration)
}

type InMemoryMetrics struct{}

func (m *InMemoryMetrics) IncrementCounter(name string) {
	fmt.Printf("MEMORY METRIC: %s incremented\n", name)
}

func (m *InMemoryMetrics) RecordDuration(name string, duration time.Duration) {
	fmt.Printf("MEMORY METRIC: %s duration %v\n", name, duration)
}

// CACHE IMPLEMENTATIONS

type RedisCache struct{}

func (r *RedisCache) Set(key, value string, ttl time.Duration) error {
	fmt.Printf("REDIS: Set %s = %s (TTL: %v)\n", key, value, ttl)
	return nil
}

func (r *RedisCache) Get(key string) (string, error) {
	return fmt.Sprintf("cached_value_for_%s", key), nil
}

// MOCK IMPLEMENTATIONS

type MockUserRepository struct {
	SaveCalled bool
	users      map[string]*User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*User),
	}
}

func (m *MockUserRepository) Save(user *User) error {
	m.SaveCalled = true
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

type MockEmailService struct {
	SendCalled bool
}

func (m *MockEmailService) SendWelcomeEmail(email string) error {
	m.SendCalled = true
	fmt.Printf("MOCK: Welcome email sent to %s\n", email)
	return nil
}

func (m *MockEmailService) SendOrderConfirmation(email string, order *Order) error {
	m.SendCalled = true
	fmt.Printf("MOCK: Order confirmation sent to %s\n", email)
	return nil
}

type MockPaymentService struct{}

func (m *MockPaymentService) ProcessPayment(payment Payment) error {
	fmt.Printf("MOCK: Payment processed: $%.2f\n", payment.Amount)
	return nil
}

type MockLogger struct{}

func (m *MockLogger) Info(message string) {
	fmt.Printf("MOCK LOG: %s\n", message)
}

func (m *MockLogger) Error(message string) {
	fmt.Printf("MOCK ERROR: %s\n", message)
}

type MockMetrics struct{}

func (m *MockMetrics) IncrementCounter(name string) {
	fmt.Printf("MOCK METRIC: %s incremented\n", name)
}

func (m *MockMetrics) RecordDuration(name string, duration time.Duration) {
	fmt.Printf("MOCK METRIC: %s duration %v\n", name, duration)
}

type MockCache struct{}

func (m *MockCache) Set(key, value string, ttl time.Duration) error {
	fmt.Printf("MOCK CACHE: Set %s = %s\n", key, value)
	return nil
}

func (m *MockCache) Get(key string) (string, error) {
	return fmt.Sprintf("mock_cached_value_for_%s", key), nil
}

// E-COMMERCE SYSTEM

type ECommerceSystem struct {
	userService    *UserService
	productService *ProductService
	orderService   *OrderService
	users          map[string]*User
	products       map[string]*Product
	carts          map[string]*Cart
	orders         map[string]*Order
}

type Cart struct {
	UserID string
	Items  []CartItem
}

type CartItem struct {
	Product  Product
	Quantity int
}

func NewECommerceSystem() *ECommerceSystem {
	userRepo := NewInMemoryUserRepository()
	emailService := &SMTPEmailService{}
	paymentService := &StripePaymentService{}
	
	return &ECommerceSystem{
		userService:    NewUserService(userRepo, emailService, paymentService),
		productService: NewProductService(&ConsoleLogger{}, &PrometheusMetrics{}, &RedisCache{}),
		orderService:   NewOrderService(),
		users:          make(map[string]*User),
		products:       make(map[string]*Product),
		carts:          make(map[string]*Cart),
		orders:         make(map[string]*Order),
	}
}

func (s *ECommerceSystem) RegisterUser(username, email string) (*User, error) {
	user, err := s.userService.RegisterUser(username, email)
	if err != nil {
		return nil, err
	}
	
	s.users[user.ID] = user
	s.carts[user.ID] = &Cart{UserID: user.ID, Items: make([]CartItem, 0)}
	
	return user, nil
}

func (s *ECommerceSystem) CreateProduct(product Product) {
	s.products[product.ID] = &product
	s.productService.CreateProduct(product)
}

func (s *ECommerceSystem) GetCart(userID string) *Cart {
	return s.carts[userID]
}

func (s *ECommerceSystem) Checkout(userID string) (*Order, error) {
	cart := s.carts[userID]
	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}
	
	// Create order
	order := NewOrder(userID, 0)
	for _, item := range cart.Items {
		order.AddItem(item.Product.Name, item.Quantity, item.Product.Price)
	}
	
	// Validate and process order
	if err := s.orderService.ValidateOrder(order); err != nil {
		return nil, err
	}
	
	if err := s.orderService.ProcessOrder(order); err != nil {
		return nil, err
	}
	
	s.orders[order.ID] = order
	
	// Clear cart
	cart.Items = cart.Items[:0]
	
	return order, nil
}

func (c *Cart) AddItem(product Product, quantity int) {
	// Check if item already exists
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += quantity
			return
		}
	}
	
	// Add new item
	c.Items = append(c.Items, CartItem{
		Product:  product,
		Quantity: quantity,
	})
}

// UTILITY FUNCTIONS

func generateID() string {
	return fmt.Sprintf("id-%d", time.Now().UnixNano())
}
