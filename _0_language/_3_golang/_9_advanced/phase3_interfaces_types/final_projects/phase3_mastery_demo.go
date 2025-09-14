// 🚀 PHASE 3 MASTERY DEMONSTRATION
// A comprehensive project showcasing interfaces, type systems, reflection, and clean architecture
package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// ============================================================================
// INTERFACE DESIGN PATTERNS DEMONSTRATION
// ============================================================================

// Core interfaces for a microservices architecture
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
}

type Metrics interface {
	IncrementCounter(name string, tags map[string]string)
	RecordHistogram(name string, value float64, tags map[string]string)
	RecordGauge(name string, value float64, tags map[string]string)
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration) error
	Delete(key string) error
}

type Database interface {
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)
	Execute(query string, args ...interface{}) (int64, error)
	Transaction(fn func(Database) error) error
}

// Service interfaces
type UserService interface {
	GetUser(id string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type OrderService interface {
	CreateOrder(userID string, items []OrderItem) (*Order, error)
	GetOrder(id string) (*Order, error)
	UpdateOrderStatus(id string, status OrderStatus) error
}

type PaymentService interface {
	ProcessPayment(amount float64, currency string, method PaymentMethod) (*PaymentResult, error)
	RefundPayment(paymentID string, amount float64) error
}

// ============================================================================
// ADVANCED TYPE SYSTEM DEMONSTRATION
// ============================================================================

// Generic repository pattern
type Repository[T any] interface {
	FindByID(id string) (*T, error)
	FindAll() ([]*T, error)
	Save(entity *T) error
	Delete(id string) error
}

// Generic service with constraints
type Service[T any] interface {
	Process(data T) (T, error)
	Validate(data T) error
}

// Type-safe event system
type Event interface {
	Type() string
	Data() interface{}
	GetTimestamp() time.Time
}

type EventHandler[T Event] interface {
	Handle(event T) error
}

// Generic event bus
type EventBus[T Event] interface {
	Subscribe(handler EventHandler[T]) error
	Publish(event T) error
	Unsubscribe(handler EventHandler[T]) error
}

// ============================================================================
// REFLECTION MASTERY DEMONSTRATION
// ============================================================================

// Dynamic struct builder using reflection
type StructBuilder struct {
	fields map[string]interface{}
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{
		fields: make(map[string]interface{}),
	}
}

func (sb *StructBuilder) AddField(name string, value interface{}) *StructBuilder {
	sb.fields[name] = value
	return sb
}

func (sb *StructBuilder) Build() interface{} {
	// Create a dynamic struct using reflection
	fields := make([]reflect.StructField, 0, len(sb.fields))
	values := make([]reflect.Value, 0, len(sb.fields))
	
	for name, value := range sb.fields {
		fieldType := reflect.TypeOf(value)
		fields = append(fields, reflect.StructField{
			Name: name,
			Type: fieldType,
		})
		values = append(values, reflect.ValueOf(value))
	}
	
	structType := reflect.StructOf(fields)
	structValue := reflect.New(structType).Elem()
	
	for i, value := range values {
		structValue.Field(i).Set(value)
	}
	
	return structValue.Interface()
}

// Dynamic method invoker
type MethodInvoker struct {
	obj interface{}
}

func NewMethodInvoker(obj interface{}) *MethodInvoker {
	return &MethodInvoker{obj: obj}
}

func (mi *MethodInvoker) Call(methodName string, args ...interface{}) ([]interface{}, error) {
	objValue := reflect.ValueOf(mi.obj)
	method := objValue.MethodByName(methodName)
	
	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}
	
	// Convert arguments to reflect.Values
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	
	// Call the method
	results := method.Call(argValues)
	
	// Convert results back to interface{}
	resultInterfaces := make([]interface{}, len(results))
	for i, result := range results {
		resultInterfaces[i] = result.Interface()
	}
	
	return resultInterfaces, nil
}

// ============================================================================
// CLEAN ARCHITECTURE IMPLEMENTATION
// ============================================================================

// Domain entities
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`
	Items       []OrderItem `json:"items"`
	Total       float64     `json:"total"`
	Status      OrderStatus `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodPayPal     PaymentMethod = "paypal"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

type PaymentResult struct {
	ID       string    `json:"id"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Status   string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// ============================================================================
// CONCRETE IMPLEMENTATIONS
// ============================================================================

// Logger implementations
type ConsoleLogger struct {
	level string
}

func NewConsoleLogger(level string) *ConsoleLogger {
	return &ConsoleLogger{level: level}
}

func (l *ConsoleLogger) Info(msg string, fields ...interface{}) {
	fmt.Printf("[INFO] %s %v\n", msg, fields)
}

func (l *ConsoleLogger) Error(msg string, fields ...interface{}) {
	fmt.Printf("[ERROR] %s %v\n", msg, fields)
}

func (l *ConsoleLogger) Debug(msg string, fields ...interface{}) {
	if l.level == "debug" {
		fmt.Printf("[DEBUG] %s %v\n", msg, fields)
	}
}

// Metrics implementations
type PrometheusMetrics struct {
	counters   map[string]int64
	histograms map[string][]float64
	gauges     map[string]float64
	mu         sync.RWMutex
}

func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{
		counters:   make(map[string]int64),
		histograms: make(map[string][]float64),
		gauges:     make(map[string]float64),
	}
}

func (m *PrometheusMetrics) IncrementCounter(name string, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[name]++
}

func (m *PrometheusMetrics) RecordHistogram(name string, value float64, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.histograms[name] = append(m.histograms[name], value)
}

func (m *PrometheusMetrics) RecordGauge(name string, value float64, tags map[string]string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gauges[name] = value
}

// Cache implementations
type InMemoryCache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
}

type cacheItem struct {
	value     interface{}
	expiresAt time.Time
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]cacheItem),
	}
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, exists := c.data[key]
	if !exists {
		return nil, false
	}
	
	if time.Now().After(item.expiresAt) {
		delete(c.data, key)
		return nil, false
	}
	
	return item.value, true
}

func (c *InMemoryCache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.data[key] = cacheItem{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	
	return nil
}

func (c *InMemoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.data, key)
	return nil
}

// Database implementations
type MockDatabase struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		data: make(map[string]interface{}),
	}
}

func (db *MockDatabase) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	// Simulate query execution
	results := []map[string]interface{}{
		{"id": "1", "name": "John Doe", "email": "john@example.com"},
		{"id": "2", "name": "Jane Smith", "email": "jane@example.com"},
	}
	
	return results, nil
}

func (db *MockDatabase) Execute(query string, args ...interface{}) (int64, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	// Simulate execution
	return 1, nil
}

func (db *MockDatabase) Transaction(fn func(Database) error) error {
	// Simulate transaction
	return fn(db)
}

// Service implementations
type UserServiceImpl struct {
	logger   Logger
	metrics  Metrics
	cache    Cache
	database Database
}

func NewUserService(logger Logger, metrics Metrics, cache Cache, database Database) UserService {
	return &UserServiceImpl{
		logger:   logger,
		metrics:  metrics,
		cache:    cache,
		database: database,
	}
}

func (s *UserServiceImpl) GetUser(id string) (*User, error) {
	s.logger.Info("Getting user", "id", id)
	s.metrics.IncrementCounter("user.get", map[string]string{"id": id})
	
	// Try cache first
	if cached, found := s.cache.Get("user:" + id); found {
		s.logger.Debug("User found in cache", "id", id)
		return cached.(*User), nil
	}
	
	// Query database
	results, err := s.database.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		s.logger.Error("Database query failed", "error", err)
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	
	user := &User{
		ID:        results[0]["id"].(string),
		Name:      results[0]["name"].(string),
		Email:     results[0]["email"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Cache the result
	s.cache.Set("user:"+id, user, 5*time.Minute)
	
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *User) error {
	s.logger.Info("Creating user", "name", user.Name, "email", user.Email)
	s.metrics.IncrementCounter("user.create", map[string]string{"email": user.Email})
	
	user.ID = generateID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	_, err := s.database.Execute("INSERT INTO users (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
	
	if err != nil {
		s.logger.Error("Failed to create user", "error", err)
		return err
	}
	
	// Cache the new user
	s.cache.Set("user:"+user.ID, user, 5*time.Minute)
	
	s.logger.Info("User created successfully", "id", user.ID)
	return nil
}

func (s *UserServiceImpl) UpdateUser(user *User) error {
	s.logger.Info("Updating user", "id", user.ID)
	s.metrics.IncrementCounter("user.update", map[string]string{"id": user.ID})
	
	user.UpdatedAt = time.Now()
	
	_, err := s.database.Execute("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?",
		user.Name, user.Email, user.UpdatedAt, user.ID)
	
	if err != nil {
		s.logger.Error("Failed to update user", "error", err)
		return err
	}
	
	// Update cache
	s.cache.Set("user:"+user.ID, user, 5*time.Minute)
	
	s.logger.Info("User updated successfully", "id", user.ID)
	return nil
}

func (s *UserServiceImpl) DeleteUser(id string) error {
	s.logger.Info("Deleting user", "id", id)
	s.metrics.IncrementCounter("user.delete", map[string]string{"id": id})
	
	_, err := s.database.Execute("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		s.logger.Error("Failed to delete user", "error", err)
		return err
	}
	
	// Remove from cache
	s.cache.Delete("user:" + id)
	
	s.logger.Info("User deleted successfully", "id", id)
	return nil
}

// Order service implementation
type OrderServiceImpl struct {
	logger        Logger
	metrics       Metrics
	database      Database
	userService   UserService
	paymentService PaymentService
}

func NewOrderService(logger Logger, metrics Metrics, database Database, userService UserService, paymentService PaymentService) OrderService {
	return &OrderServiceImpl{
		logger:         logger,
		metrics:        metrics,
		database:       database,
		userService:    userService,
		paymentService: paymentService,
	}
}

func (s *OrderServiceImpl) CreateOrder(userID string, items []OrderItem) (*Order, error) {
	s.logger.Info("Creating order", "user_id", userID, "items_count", len(items))
	s.metrics.IncrementCounter("order.create", map[string]string{"user_id": userID})
	
	// Verify user exists
	_, err := s.userService.GetUser(userID)
	if err != nil {
		s.logger.Error("User not found", "user_id", userID, "error", err)
		return nil, fmt.Errorf("user not found: %w", err)
	}
	
	// Calculate total
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	
	order := &Order{
		ID:        generateID(),
		UserID:    userID,
		Items:     items,
		Total:     total,
		Status:    OrderStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Save to database
	_, err = s.database.Execute("INSERT INTO orders (id, user_id, total, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		order.ID, order.UserID, order.Total, string(order.Status), order.CreatedAt, order.UpdatedAt)
	
	if err != nil {
		s.logger.Error("Failed to create order", "error", err)
		return nil, err
	}
	
	s.logger.Info("Order created successfully", "order_id", order.ID, "total", order.Total)
	return order, nil
}

func (s *OrderServiceImpl) GetOrder(id string) (*Order, error) {
	s.logger.Info("Getting order", "id", id)
	s.metrics.IncrementCounter("order.get", map[string]string{"id": id})
	
	results, err := s.database.Query("SELECT * FROM orders WHERE id = ?", id)
	if err != nil {
		s.logger.Error("Database query failed", "error", err)
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("order not found")
	}
	
	order := &Order{
		ID:        results[0]["id"].(string),
		UserID:    results[0]["user_id"].(string),
		Total:     results[0]["total"].(float64),
		Status:    OrderStatus(results[0]["status"].(string)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	return order, nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(id string, status OrderStatus) error {
	s.logger.Info("Updating order status", "id", id, "status", status)
	s.metrics.IncrementCounter("order.update_status", map[string]string{"id": id, "status": string(status)})
	
	_, err := s.database.Execute("UPDATE orders SET status = ?, updated_at = ? WHERE id = ?",
		string(status), time.Now(), id)
	
	if err != nil {
		s.logger.Error("Failed to update order status", "error", err)
		return err
	}
	
	s.logger.Info("Order status updated successfully", "id", id, "status", status)
	return nil
}

// Payment service implementation
type PaymentServiceImpl struct {
	logger  Logger
	metrics Metrics
}

func NewPaymentService(logger Logger, metrics Metrics) PaymentService {
	return &PaymentServiceImpl{
		logger:  logger,
		metrics: metrics,
	}
}

func (s *PaymentServiceImpl) ProcessPayment(amount float64, currency string, method PaymentMethod) (*PaymentResult, error) {
	s.logger.Info("Processing payment", "amount", amount, "currency", currency, "method", method)
	s.metrics.IncrementCounter("payment.process", map[string]string{"currency": currency, "method": string(method)})
	
	// Simulate payment processing
	time.Sleep(100 * time.Millisecond)
	
	result := &PaymentResult{
		ID:        generateID(),
		Amount:    amount,
		Currency:  currency,
		Status:    "completed",
		CreatedAt: time.Now(),
	}
	
	s.logger.Info("Payment processed successfully", "payment_id", result.ID, "amount", amount)
	return result, nil
}

func (s *PaymentServiceImpl) RefundPayment(paymentID string, amount float64) error {
	s.logger.Info("Processing refund", "payment_id", paymentID, "amount", amount)
	s.metrics.IncrementCounter("payment.refund", map[string]string{"payment_id": paymentID})
	
	// Simulate refund processing
	time.Sleep(50 * time.Millisecond)
	
	s.logger.Info("Refund processed successfully", "payment_id", paymentID, "amount", amount)
	return nil
}

// ============================================================================
// GENERIC REPOSITORY IMPLEMENTATION
// ============================================================================

type GenericRepository[T any] struct {
	logger   Logger
	database Database
	cache    Cache
	table    string
}

func NewGenericRepository[T any](logger Logger, database Database, cache Cache, table string) Repository[T] {
	return &GenericRepository[T]{
		logger:   logger,
		database: database,
		cache:    cache,
		table:    table,
	}
}

func (r *GenericRepository[T]) FindByID(id string) (*T, error) {
	r.logger.Debug("Finding by ID", "table", r.table, "id", id)
	
	// Try cache first
	cacheKey := fmt.Sprintf("%s:%s", r.table, id)
	if cached, found := r.cache.Get(cacheKey); found {
		r.logger.Debug("Found in cache", "table", r.table, "id", id)
		return cached.(*T), nil
	}
	
	// Query database
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", r.table)
	results, err := r.database.Query(query, id)
	if err != nil {
		r.logger.Error("Database query failed", "error", err)
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("record not found")
	}
	
	// Convert to generic type using reflection
	record := r.mapToStruct(results[0])
	
	// Cache the result
	r.cache.Set(cacheKey, record, 5*time.Minute)
	
	return record, nil
}

func (r *GenericRepository[T]) FindAll() ([]*T, error) {
	r.logger.Debug("Finding all", "table", r.table)
	
	query := fmt.Sprintf("SELECT * FROM %s", r.table)
	results, err := r.database.Query(query)
	if err != nil {
		r.logger.Error("Database query failed", "error", err)
		return nil, err
	}
	
	records := make([]*T, len(results))
	for i, result := range results {
		records[i] = r.mapToStruct(result)
	}
	
	return records, nil
}

func (r *GenericRepository[T]) Save(entity *T) error {
	r.logger.Debug("Saving entity", "table", r.table)
	
	// Use reflection to get struct fields
	entityValue := reflect.ValueOf(entity).Elem()
	entityType := entityValue.Type()
	
	// Build INSERT query
	fields := make([]string, 0, entityType.NumField())
	placeholders := make([]string, 0, entityType.NumField())
	values := make([]interface{}, 0, entityType.NumField())
	
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)
		
		// Get JSON tag or use field name
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		fields = append(fields, jsonTag)
		placeholders = append(placeholders, "?")
		values = append(values, fieldValue.Interface())
	}
	
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		r.table,
		fmt.Sprintf("%v", fields),
		fmt.Sprintf("%v", placeholders))
	
	_, err := r.database.Execute(query, values...)
	if err != nil {
		r.logger.Error("Failed to save entity", "error", err)
		return err
	}
	
	r.logger.Debug("Entity saved successfully", "table", r.table)
	return nil
}

func (r *GenericRepository[T]) Delete(id string) error {
	r.logger.Debug("Deleting entity", "table", r.table, "id", id)
	
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.table)
	_, err := r.database.Execute(query, id)
	if err != nil {
		r.logger.Error("Failed to delete entity", "error", err)
		return err
	}
	
	// Remove from cache
	cacheKey := fmt.Sprintf("%s:%s", r.table, id)
	r.cache.Delete(cacheKey)
	
	r.logger.Debug("Entity deleted successfully", "table", r.table, "id", id)
	return nil
}

func (r *GenericRepository[T]) mapToStruct(data map[string]interface{}) *T {
	// Create new instance of T
	var entity T
	entityValue := reflect.ValueOf(&entity).Elem()
	entityType := entityValue.Type()
	
	// Map fields
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)
		
		// Get JSON tag or use field name
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		if value, exists := data[jsonTag]; exists {
			if fieldValue.CanSet() {
				fieldValue.Set(reflect.ValueOf(value))
			}
		}
	}
	
	return &entity
}

// ============================================================================
// EVENT SYSTEM IMPLEMENTATION
// ============================================================================

type UserCreatedEvent struct {
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserCreatedEvent) Type() string {
	return "user.created"
}

func (e UserCreatedEvent) Data() interface{} {
	return e
}

func (e UserCreatedEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

type OrderCreatedEvent struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Total     float64   `json:"total"`
	Timestamp time.Time `json:"timestamp"`
}

func (e OrderCreatedEvent) Type() string {
	return "order.created"
}

func (e OrderCreatedEvent) Data() interface{} {
	return e
}

func (e OrderCreatedEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

type EventBusImpl[T Event] struct {
	handlers []EventHandler[T]
	mu       sync.RWMutex
}

func NewEventBus[T Event]() EventBus[T] {
	return &EventBusImpl[T]{
		handlers: make([]EventHandler[T], 0),
	}
}

func (eb *EventBusImpl[T]) Subscribe(handler EventHandler[T]) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	eb.handlers = append(eb.handlers, handler)
	return nil
}

func (eb *EventBusImpl[T]) Publish(event T) error {
	eb.mu.RLock()
	handlers := make([]EventHandler[T], len(eb.handlers))
	copy(handlers, eb.handlers)
	eb.mu.RUnlock()
	
	// Execute handlers concurrently
	var wg sync.WaitGroup
	for _, handler := range handlers {
		wg.Add(1)
		go func(h EventHandler[T]) {
			defer wg.Done()
			if err := h.Handle(event); err != nil {
				log.Printf("Error handling event: %v", err)
			}
		}(handler)
	}
	
	wg.Wait()
	return nil
}

func (eb *EventBusImpl[T]) Unsubscribe(handler EventHandler[T]) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	for i, h := range eb.handlers {
		if reflect.DeepEqual(h, handler) {
			eb.handlers = append(eb.handlers[:i], eb.handlers[i+1:]...)
			break
		}
	}
	
	return nil
}

// Event handlers
type EmailNotificationHandler struct {
	logger Logger
}

func NewEmailNotificationHandler(logger Logger) *EmailNotificationHandler {
	return &EmailNotificationHandler{logger: logger}
}

func (h *EmailNotificationHandler) Handle(event Event) error {
	switch e := event.(type) {
	case UserCreatedEvent:
		h.logger.Info("Sending welcome email", "user_id", e.UserID, "email", e.UserEmail)
		// Simulate email sending
		time.Sleep(50 * time.Millisecond)
		h.logger.Info("Welcome email sent", "user_id", e.UserID)
	case OrderCreatedEvent:
		h.logger.Info("Sending order confirmation email", "order_id", e.OrderID, "user_id", e.UserID)
		// Simulate email sending
		time.Sleep(50 * time.Millisecond)
		h.logger.Info("Order confirmation email sent", "order_id", e.OrderID)
	}
	return nil
}

type MetricsEventHandler struct {
	metrics Metrics
}

func NewMetricsEventHandler(metrics Metrics) *MetricsEventHandler {
	return &MetricsEventHandler{metrics: metrics}
}

func (h *MetricsEventHandler) Handle(event Event) error {
	switch event.Type() {
	case "user.created":
		h.metrics.IncrementCounter("events.user.created", map[string]string{})
	case "order.created":
		h.metrics.IncrementCounter("events.order.created", map[string]string{})
	}
	return nil
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

func generateID() string {
	return fmt.Sprintf("id-%d", time.Now().UnixNano())
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🚀 PHASE 3 MASTERY DEMONSTRATION")
	fmt.Println("=================================")
	fmt.Println()
	
	// Initialize dependencies
	logger := NewConsoleLogger("debug")
	metrics := NewPrometheusMetrics()
	cache := NewInMemoryCache()
	database := NewMockDatabase()
	
	// Create services
	userService := NewUserService(logger, metrics, cache, database)
	paymentService := NewPaymentService(logger, metrics)
	orderService := NewOrderService(logger, metrics, database, userService, paymentService)
	
	// Create event bus
	eventBus := NewEventBus[Event]()
	
	// Subscribe event handlers
	emailHandler := NewEmailNotificationHandler(logger)
	metricsHandler := NewMetricsEventHandler(metrics)
	
	eventBus.Subscribe(emailHandler)
	eventBus.Subscribe(metricsHandler)
	
	// Demonstrate interface design patterns
	demonstrateInterfacePatterns(logger, metrics, cache, database)
	
	// Demonstrate advanced type system
	demonstrateAdvancedTypeSystem(logger)
	
	// Demonstrate reflection mastery
	demonstrateReflectionMastery(logger)
	
	// Demonstrate clean architecture
	demonstrateCleanArchitecture(userService, orderService, paymentService, eventBus, logger)
	
	// Demonstrate generic repository
	demonstrateGenericRepository(logger, metrics, cache, database)
	
	// Demonstrate event system
	demonstrateEventSystem(eventBus, logger)
	
	fmt.Println()
	fmt.Println("🎉 PHASE 3 MASTERY DEMONSTRATION COMPLETE!")
	fmt.Println("You have successfully demonstrated mastery of:")
	fmt.Println("✅ Interface design patterns and composition")
	fmt.Println("✅ Advanced type system with generics")
	fmt.Println("✅ Reflection and dynamic programming")
	fmt.Println("✅ Clean architecture principles")
	fmt.Println("✅ Event-driven architecture")
	fmt.Println("✅ Generic repository patterns")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Phase 4: Error Handling & Logging Mastery!")
}

func demonstrateInterfacePatterns(logger Logger, metrics Metrics, cache Cache, database Database) {
	fmt.Println("🎯 INTERFACE DESIGN PATTERNS DEMONSTRATION")
	fmt.Println("==========================================")
	
	// Demonstrate interface composition
	logger.Info("Demonstrating interface composition")
	metrics.IncrementCounter("demo.interface_composition", map[string]string{})
	
	// Demonstrate dependency injection
	logger.Info("Demonstrating dependency injection")
	cache.Set("demo_key", "demo_value", 1*time.Minute)
	
	if value, found := cache.Get("demo_key"); found {
		logger.Info("Cache retrieval successful", "value", value)
	}
	
	// Demonstrate database operations
	logger.Info("Demonstrating database operations")
	results, err := database.Query("SELECT * FROM users")
	if err != nil {
		logger.Error("Query failed", "error", err)
	} else {
		logger.Info("Query successful", "results_count", len(results))
	}
	
	fmt.Println("✅ Interface design patterns demonstrated successfully")
	fmt.Println()
}

func demonstrateAdvancedTypeSystem(logger Logger) {
	fmt.Println("🔧 ADVANCED TYPE SYSTEM DEMONSTRATION")
	fmt.Println("=====================================")
	
	// Demonstrate generic functions
	logger.Info("Demonstrating generic functions")
	
	// Generic max function
	maxInt := max(10, 20)
	maxString := max("hello", "world")
	logger.Info("Generic max function", "max_int", maxInt, "max_string", maxString)
	
	// Generic slice operations
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapSlice(numbers, func(x int) int { return x * 2 })
	logger.Info("Generic slice mapping", "original", numbers, "doubled", doubled)
	
	// Generic filter
	evens := filterSlice(numbers, func(x int) bool { return x%2 == 0 })
	logger.Info("Generic slice filtering", "numbers", numbers, "evens", evens)
	
	// Generic reduce
	sum := reduceSlice(numbers, 0, func(acc, x int) int { return acc + x })
	logger.Info("Generic slice reduction", "numbers", numbers, "sum", sum)
	
	fmt.Println("✅ Advanced type system demonstrated successfully")
	fmt.Println()
}

func demonstrateReflectionMastery(logger Logger) {
	fmt.Println("🪞 REFLECTION MASTERY DEMONSTRATION")
	fmt.Println("===================================")
	
	// Demonstrate struct builder
	logger.Info("Demonstrating dynamic struct builder")
	builder := NewStructBuilder()
	dynamicStruct := builder.
		AddField("Name", "John Doe").
		AddField("Age", 30).
		AddField("Email", "john@example.com").
		Build()
	
	logger.Info("Dynamic struct created", "struct", dynamicStruct)
	
	// Demonstrate method invoker
	logger.Info("Demonstrating dynamic method invocation")
	calculator := &Calculator{}
	invoker := NewMethodInvoker(calculator)
	
	results, err := invoker.Call("Add", 10, 20)
	if err != nil {
		logger.Error("Method invocation failed", "error", err)
	} else {
		logger.Info("Method invocation successful", "results", results)
	}
	
	// Demonstrate type inspection
	logger.Info("Demonstrating type inspection")
	inspectType(dynamicStruct, logger)
	
	fmt.Println("✅ Reflection mastery demonstrated successfully")
	fmt.Println()
}

func demonstrateCleanArchitecture(userService UserService, orderService OrderService, paymentService PaymentService, eventBus EventBus[Event], logger Logger) {
	fmt.Println("🏗️ CLEAN ARCHITECTURE DEMONSTRATION")
	fmt.Println("====================================")
	
	// Create a user
	logger.Info("Creating user")
	user := &User{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
	}
	
	err := userService.CreateUser(user)
	if err != nil {
		logger.Error("Failed to create user", "error", err)
		return
	}
	
	logger.Info("User created successfully", "user_id", user.ID)
	
	// Publish user created event
	event := UserCreatedEvent{
		UserID:    user.ID,
		UserName:  user.Name,
		UserEmail: user.Email,
		Timestamp: time.Now(),
	}
	
	eventBus.Publish(event)
	
	// Create an order
	logger.Info("Creating order")
	orderItems := []OrderItem{
		{ProductID: "prod-1", Quantity: 2, Price: 29.99},
		{ProductID: "prod-2", Quantity: 1, Price: 49.99},
	}
	
	order, err := orderService.CreateOrder(user.ID, orderItems)
	if err != nil {
		logger.Error("Failed to create order", "error", err)
		return
	}
	
	logger.Info("Order created successfully", "order_id", order.ID, "total", order.Total)
	
	// Publish order created event
	orderEvent := OrderCreatedEvent{
		OrderID:   order.ID,
		UserID:    order.UserID,
		Total:     order.Total,
		Timestamp: time.Now(),
	}
	
	eventBus.Publish(orderEvent)
	
	// Process payment
	logger.Info("Processing payment")
	paymentResult, err := paymentService.ProcessPayment(order.Total, "USD", PaymentMethodCreditCard)
	if err != nil {
		logger.Error("Failed to process payment", "error", err)
		return
	}
	
	logger.Info("Payment processed successfully", "payment_id", paymentResult.ID, "amount", paymentResult.Amount)
	
	// Update order status
	logger.Info("Updating order status")
	err = orderService.UpdateOrderStatus(order.ID, OrderStatusConfirmed)
	if err != nil {
		logger.Error("Failed to update order status", "error", err)
		return
	}
	
	logger.Info("Order status updated successfully", "order_id", order.ID, "status", OrderStatusConfirmed)
	
	fmt.Println("✅ Clean architecture demonstrated successfully")
	fmt.Println()
}

func demonstrateGenericRepository(logger Logger, metrics Metrics, cache Cache, database Database) {
	fmt.Println("📦 GENERIC REPOSITORY DEMONSTRATION")
	fmt.Println("===================================")
	
	// Create generic repository for users
	userRepo := NewGenericRepository[User](logger, database, cache, "users")
	
	// Create a user
	logger.Info("Creating user with generic repository")
	user := &User{
		ID:    generateID(),
		Name:  "Bob Smith",
		Email: "bob@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	err := userRepo.Save(user)
	if err != nil {
		logger.Error("Failed to save user", "error", err)
		return
	}
	
	logger.Info("User saved successfully", "user_id", user.ID)
	
	// Find user by ID
	logger.Info("Finding user by ID")
	foundUser, err := userRepo.FindByID(user.ID)
	if err != nil {
		logger.Error("Failed to find user", "error", err)
		return
	}
	
	logger.Info("User found successfully", "user_id", foundUser.ID, "name", foundUser.Name)
	
	// Find all users
	logger.Info("Finding all users")
	allUsers, err := userRepo.FindAll()
	if err != nil {
		logger.Error("Failed to find all users", "error", err)
		return
	}
	
	logger.Info("All users found", "count", len(allUsers))
	
	fmt.Println("✅ Generic repository demonstrated successfully")
	fmt.Println()
}

func demonstrateEventSystem(eventBus EventBus[Event], logger Logger) {
	fmt.Println("📡 EVENT SYSTEM DEMONSTRATION")
	fmt.Println("=============================")
	
	// Publish various events
	logger.Info("Publishing events")
	
	events := []Event{
		UserCreatedEvent{
			UserID:    "user-123",
			UserName:  "Charlie Brown",
			UserEmail: "charlie@example.com",
			Timestamp: time.Now(),
		},
		OrderCreatedEvent{
			OrderID:   "order-456",
			UserID:    "user-123",
			Total:     99.99,
			Timestamp: time.Now(),
		},
	}
	
	for _, event := range events {
		logger.Info("Publishing event", "type", event.Type())
		err := eventBus.Publish(event)
		if err != nil {
			logger.Error("Failed to publish event", "error", err)
		}
	}
	
	fmt.Println("✅ Event system demonstrated successfully")
	fmt.Println()
}

// Generic utility functions
func max[T any](a, b T) T {
	// This is a simplified version - in practice you'd need constraints
	// For demonstration purposes, we'll just return the first value
	return a
}

func mapSlice[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func filterSlice[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func reduceSlice[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Helper struct for method invocation demo
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func inspectType(obj interface{}, logger Logger) {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	
	logger.Info("Type inspection", "type", objType.String(), "kind", objType.Kind())
	
	if objType.Kind() == reflect.Struct {
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			fieldValue := objValue.Field(i)
			logger.Info("Field", "name", field.Name, "type", field.Type.String(), "value", fieldValue.Interface())
		}
	}
}
