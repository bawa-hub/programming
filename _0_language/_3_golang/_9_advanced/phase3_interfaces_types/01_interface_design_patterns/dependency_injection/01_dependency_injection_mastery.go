package main

import (
	"fmt"
	"sync"
	"time"
)

// 💉 DEPENDENCY INJECTION MASTERY
// Understanding dependency injection patterns in Go

func main() {
	fmt.Println("💉 DEPENDENCY INJECTION MASTERY")
	fmt.Println("===============================")
	fmt.Println()

	// 1. Constructor Injection
	constructorInjection()
	fmt.Println()

	// 2. Interface Injection
	interfaceInjection()
	fmt.Println()

	// 3. Service Locator Pattern
	serviceLocatorPattern()
	fmt.Println()

	// 4. Advanced DI Patterns
	advancedDIPatterns()
	fmt.Println()

	// 5. Dependency Resolution
	dependencyResolution()
	fmt.Println()

	// 6. Lifecycle Management
	lifecycleManagement()
	fmt.Println()

	// 7. Testing with DI
	testingWithDI()
	fmt.Println()

	// 8. Best Practices
	dependencyInjectionBestPractices()
}

// 1. Constructor Injection
func constructorInjection() {
	fmt.Println("1. Constructor Injection:")
	fmt.Println("Understanding constructor-based dependency injection...")

	// Demonstrate constructor injection
	constructorInjectionExample()
	
	// Show interface-based dependencies
	interfaceBasedDependencies()
	
	// Demonstrate dependency resolution
	dependencyResolutionExample()
}

func constructorInjectionExample() {
	fmt.Println("  📊 Constructor injection example:")
	
	// Create dependencies
	logger := &ConsoleLogger{}
	emailService := &SMTPEmailService{}
	
	// Inject dependencies through constructor
	userService := NewUserService(logger, emailService)
	
	// Use the service
	user := &User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully")
	}
}

type User struct {
	ID    int
	Name  string
	Email string
}

type UserService struct {
	logger       Logger
	emailService EmailService
}

func NewUserService(logger Logger, emailService EmailService) *UserService {
	return &UserService{
		logger:       logger,
		emailService: emailService,
	}
}

func (s *UserService) CreateUser(user *User) error {
	s.logger.Info("Creating user: %s", user.Name)
	
	// Simulate user creation
	time.Sleep(100 * time.Millisecond)
	
	// Send welcome email
	err := s.emailService.SendEmail(user.Email, "Welcome!", "Welcome to our service!")
	if err != nil {
		s.logger.Error("Failed to send welcome email: %v", err)
		return err
	}
	
	s.logger.Info("User created successfully: %s", user.Name)
	return nil
}

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Info(format string, args ...interface{}) {
	fmt.Printf("    [INFO] "+format+"\n", args...)
}

func (l *ConsoleLogger) Error(format string, args ...interface{}) {
	fmt.Printf("    [ERROR] "+format+"\n", args...)
}

type SMTPEmailService struct{}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	fmt.Printf("    Sending email to %s: %s\n", to, subject)
	time.Sleep(50 * time.Millisecond)
	return nil
}

func interfaceBasedDependencies() {
	fmt.Println("  📊 Interface-based dependencies:")
	
	// Create different implementations
	fileLogger := &FileLogger{}
	smsService := &SMSServiceImpl{}
	
	// Inject different implementations
	notificationService := NewNotificationService(fileLogger, smsService)
	
	// Use the service
	err := notificationService.SendNotification("user123", "Your order is ready!")
	if err != nil {
		fmt.Printf("    Error sending notification: %v\n", err)
	} else {
		fmt.Println("    Notification sent successfully")
	}
}

type NotificationService struct {
	logger Logger
	sms    SMSService
}

func NewNotificationService(logger Logger, sms SMSService) *NotificationService {
	return &NotificationService{
		logger: logger,
		sms:    sms,
	}
}

func (s *NotificationService) SendNotification(userID, message string) error {
	s.logger.Info("Sending notification to user %s", userID)
	
	// Simulate notification sending
	time.Sleep(75 * time.Millisecond)
	
	// Send SMS
	err := s.sms.SendSMS(userID, message)
	if err != nil {
		s.logger.Error("Failed to send SMS: %v", err)
		return err
	}
	
	s.logger.Info("Notification sent successfully to user %s", userID)
	return nil
}

type FileLogger struct{}

func (l *FileLogger) Info(format string, args ...interface{}) {
	fmt.Printf("    [FILE-INFO] "+format+"\n", args...)
}

func (l *FileLogger) Error(format string, args ...interface{}) {
	fmt.Printf("    [FILE-ERROR] "+format+"\n", args...)
}

type SMSService interface {
	SendSMS(userID, message string) error
}

type SMSServiceImpl struct{}

func (s *SMSServiceImpl) SendSMS(userID, message string) error {
	fmt.Printf("    Sending SMS to %s: %s\n", userID, message)
	time.Sleep(25 * time.Millisecond)
	return nil
}

// Additional types for service interfaces
type UserRepository interface {
	Save(user *User) error
	FindByID(id int) (*User, error)
}

type CacheService interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type InMemoryUserRepository struct {
	users map[int]*User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*User),
	}
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id int) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

type RedisCacheService struct{}

func (c *RedisCacheService) Set(key string, value interface{}) error {
	fmt.Printf("    Caching key %s\n", key)
	return nil
}

func (c *RedisCacheService) Get(key string) (interface{}, error) {
	fmt.Printf("    Getting key %s from cache\n", key)
	return nil, fmt.Errorf("key not found")
}

func dependencyResolutionExample() {
	fmt.Println("  📊 Dependency resolution example:")
	
	// Create a dependency resolver
	resolver := NewDependencyResolver()
	
	// Register dependencies
	resolver.Register("logger", &ConsoleLogger{})
	resolver.Register("email", &SMTPEmailService{})
	
	// Resolve dependencies
	logger, _ := resolver.Resolve("logger").(Logger)
	emailService, _ := resolver.Resolve("email").(EmailService)
	
	// Create service with resolved dependencies
	userService := NewUserService(logger, emailService)
	
	// Use the service
	user := &User{ID: 2, Name: "Jane Smith", Email: "jane@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully with resolved dependencies")
	}
}

type DependencyResolver struct {
	services map[string]interface{}
	mu       sync.RWMutex
}

func NewDependencyResolver() *DependencyResolver {
	return &DependencyResolver{
		services: make(map[string]interface{}),
	}
}

func (r *DependencyResolver) Register(name string, service interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = service
}

func (r *DependencyResolver) Resolve(name string) interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.services[name]
}

// 2. Interface Injection
func interfaceInjection() {
	fmt.Println("2. Interface Injection:")
	fmt.Println("Understanding interface-based injection patterns...")

	// Demonstrate interface injection
	interfaceInjectionExample()
	
	// Show service interfaces
	serviceInterfaces()
	
	// Demonstrate implementation swapping
	implementationSwapping()
}

func interfaceInjectionExample() {
	fmt.Println("  📊 Interface injection example:")
	
	// Create service with interface injection
	service := NewServiceWithInterfaceInjection()
	
	// Inject different implementations
	service.SetLogger(&ConsoleLogger{})
	service.SetEmailService(&SMTPEmailService{})
	
	// Use the service
	user := &User{ID: 3, Name: "Bob Johnson", Email: "bob@example.com"}
	err := service.ProcessUser(user)
	if err != nil {
		fmt.Printf("    Error processing user: %v\n", err)
	} else {
		fmt.Println("    User processed successfully")
	}
}

type ServiceWithInterfaceInjection struct {
	logger       Logger
	emailService EmailService
}

func NewServiceWithInterfaceInjection() *ServiceWithInterfaceInjection {
	return &ServiceWithInterfaceInjection{}
}

func (s *ServiceWithInterfaceInjection) SetLogger(logger Logger) {
	s.logger = logger
}

func (s *ServiceWithInterfaceInjection) SetEmailService(emailService EmailService) {
	s.emailService = emailService
}

func (s *ServiceWithInterfaceInjection) ProcessUser(user *User) error {
	if s.logger == nil || s.emailService == nil {
		return fmt.Errorf("dependencies not injected")
	}
	
	s.logger.Info("Processing user: %s", user.Name)
	
	// Simulate processing
	time.Sleep(100 * time.Millisecond)
	
	// Send email
	err := s.emailService.SendEmail(user.Email, "Processing Complete", "Your account has been processed.")
	if err != nil {
		s.logger.Error("Failed to send email: %v", err)
		return err
	}
	
	s.logger.Info("User processed successfully: %s", user.Name)
	return nil
}

func serviceInterfaces() {
	fmt.Println("  📊 Service interfaces:")
	
	// Use interfaces
	userRepo := NewInMemoryUserRepository()
	cacheService := &RedisCacheService{}
	
	// Save user
	user := &User{ID: 4, Name: "Alice Brown", Email: "alice@example.com"}
	err := userRepo.Save(user)
	if err != nil {
		fmt.Printf("    Error saving user: %v\n", err)
	} else {
		fmt.Println("    User saved successfully")
	}
	
	// Cache user
	cacheService.Set("user:4", user)
	
	// Find user
	foundUser, err := userRepo.FindByID(4)
	if err != nil {
		fmt.Printf("    Error finding user: %v\n", err)
	} else {
		fmt.Printf("    Found user: %s\n", foundUser.Name)
	}
}

func implementationSwapping() {
	fmt.Println("  📊 Implementation swapping:")
	
	// Create service with swappable implementations
	service := NewSwappableService()
	
	// Use first implementation
	service.SetLogger(&ConsoleLogger{})
	service.LogInfo("Using console logger")
	
	// Swap to different implementation
	service.SetLogger(&FileLogger{})
	service.LogInfo("Using file logger")
	
	// Swap to mock implementation
	service.SetLogger(&MockLogger{})
	service.LogInfo("Using mock logger")
}

type SwappableService struct {
	logger Logger
}

func NewSwappableService() *SwappableService {
	return &SwappableService{}
}

func (s *SwappableService) SetLogger(logger Logger) {
	s.logger = logger
}

func (s *SwappableService) LogInfo(message string) {
	if s.logger != nil {
		s.logger.Info(message)
	}
}

type MockLogger struct{
	infoCallCount  int
	errorCallCount int
	mu             sync.Mutex
}

func (l *MockLogger) Info(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.infoCallCount++
	fmt.Printf("    [MOCK-INFO] "+format+"\n", args...)
}

func (l *MockLogger) Error(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.errorCallCount++
	fmt.Printf("    [MOCK-ERROR] "+format+"\n", args...)
}

func (l *MockLogger) GetInfoCallCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.infoCallCount
}

func (l *MockLogger) GetErrorCallCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.errorCallCount
}

// 3. Service Locator Pattern
func serviceLocatorPattern() {
	fmt.Println("3. Service Locator Pattern:")
	fmt.Println("Understanding service locator patterns...")

	// Demonstrate service locator
	serviceLocatorExample()
	
	// Show service containers
	serviceContainers()
	
	// Demonstrate dependency graphs
	dependencyGraphs()
}

func serviceLocatorExample() {
	fmt.Println("  📊 Service locator example:")
	
	// Create service locator
	locator := NewServiceLocator()
	
	// Register services
	locator.Register("logger", &ConsoleLogger{})
	locator.Register("email", &SMTPEmailService{})
	locator.Register("sms", &SMSServiceImpl{})
	
	// Get services
	logger := locator.Get("logger").(Logger)
	emailService := locator.Get("email").(EmailService)
	smsService := locator.Get("sms").(SMSService)
	
	// Use services
	logger.Info("Service locator example")
	emailService.SendEmail("test@example.com", "Test", "Test message")
	smsService.SendSMS("user123", "Test SMS")
}

type ServiceLocator struct {
	services map[string]interface{}
	mu       sync.RWMutex
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]interface{}),
	}
}

func (sl *ServiceLocator) Register(name string, service interface{}) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.services[name] = service
}

func (sl *ServiceLocator) Get(name string) interface{} {
	sl.mu.RLock()
	defer sl.mu.RUnlock()
	return sl.services[name]
}

func (sl *ServiceLocator) IsRegistered(name string) bool {
	sl.mu.RLock()
	defer sl.mu.RUnlock()
	_, exists := sl.services[name]
	return exists
}

func serviceContainers() {
	fmt.Println("  📊 Service containers:")
	
	// Create service container
	container := NewServiceContainer()
	
	// Register services with lifecycle management
	container.RegisterSingleton("logger", &ConsoleLogger{})
	container.RegisterTransient("email", func() interface{} {
		return &SMTPEmailService{}
	})
	container.RegisterScoped("userService", func() interface{} {
		logger := container.Get("logger").(Logger)
		emailService := container.Get("email").(EmailService)
		return NewUserService(logger, emailService)
	})
	
	// Get services
	logger := container.Get("logger").(Logger)
	userService := container.Get("userService").(*UserService)
	
	// Use services
	logger.Info("Service container example")
	user := &User{ID: 5, Name: "Charlie Wilson", Email: "charlie@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully with service container")
	}
}

type ServiceContainer struct {
	services map[string]ServiceRegistration
	mu       sync.RWMutex
}

type ServiceRegistration struct {
	Factory    func() interface{}
	Instance   interface{}
	Lifecycle  Lifecycle
	IsResolved bool
}

type Lifecycle int

const (
	Singleton Lifecycle = iota
	Transient
	Scoped
)

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		services: make(map[string]ServiceRegistration),
	}
}

func (c *ServiceContainer) RegisterSingleton(name string, instance interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = ServiceRegistration{
		Instance:   instance,
		Lifecycle:  Singleton,
		IsResolved: true,
	}
}

func (c *ServiceContainer) RegisterTransient(name string, factory func() interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = ServiceRegistration{
		Factory:   factory,
		Lifecycle: Transient,
	}
}

func (c *ServiceContainer) RegisterScoped(name string, factory func() interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = ServiceRegistration{
		Factory:   factory,
		Lifecycle: Scoped,
	}
}

func (c *ServiceContainer) Get(name string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	registration, exists := c.services[name]
	if !exists {
		return nil
	}
	
	switch registration.Lifecycle {
	case Singleton:
		return registration.Instance
	case Transient:
		return registration.Factory()
	case Scoped:
		if !registration.IsResolved {
			registration.Instance = registration.Factory()
			registration.IsResolved = true
			c.services[name] = registration
		}
		return registration.Instance
	}
	
	return nil
}

func dependencyGraphs() {
	fmt.Println("  📊 Dependency graphs:")
	
	// Create complex dependency graph
	container := NewServiceContainer()
	
	// Register dependencies in order
	container.RegisterSingleton("logger", &ConsoleLogger{})
	container.RegisterSingleton("email", &SMTPEmailService{})
	container.RegisterSingleton("sms", &SMSServiceImpl{})
	
	// Register services that depend on others
	container.RegisterScoped("userRepo", func() interface{} {
		return NewInMemoryUserRepository()
	})
	
	container.RegisterScoped("userService", func() interface{} {
		logger := container.Get("logger").(Logger)
		emailService := container.Get("email").(EmailService)
		return NewUserService(logger, emailService)
	})
	
	container.RegisterScoped("notificationService", func() interface{} {
		logger := container.Get("logger").(Logger)
		smsService := container.Get("sms").(SMSService)
		return NewNotificationService(logger, smsService)
	})
	
	// Register orchestrator service
	container.RegisterScoped("orchestrator", func() interface{} {
		userService := container.Get("userService").(*UserService)
		notificationService := container.Get("notificationService").(*NotificationService)
		return NewOrchestrator(userService, notificationService)
	})
	
	// Use orchestrator
	orchestrator := container.Get("orchestrator").(*Orchestrator)
	orchestrator.ProcessUserRegistration(&User{ID: 6, Name: "David Lee", Email: "david@example.com"})
}

type Orchestrator struct {
	userService         *UserService
	notificationService *NotificationService
}

func NewOrchestrator(userService *UserService, notificationService *NotificationService) *Orchestrator {
	return &Orchestrator{
		userService:         userService,
		notificationService: notificationService,
	}
}

func (o *Orchestrator) ProcessUserRegistration(user *User) {
	fmt.Printf("    Processing user registration: %s\n", user.Name)
	
	// Create user
	err := o.userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
		return
	}
	
	// Send notification
	err = o.notificationService.SendNotification(fmt.Sprintf("user%d", user.ID), "Welcome to our platform!")
	if err != nil {
		fmt.Printf("    Error sending notification: %v\n", err)
		return
	}
	
	fmt.Println("    User registration processed successfully")
}

// 4. Advanced DI Patterns
func advancedDIPatterns() {
	fmt.Println("4. Advanced DI Patterns:")
	fmt.Println("Understanding advanced dependency injection patterns...")

	// Demonstrate factory patterns
	factoryPatterns()
	
	// Show builder patterns
	builderPatterns()
	
	// Demonstrate singleton patterns
	singletonPatterns()
	
	// Show scoped dependencies
	scopedDependencies()
}

func factoryPatterns() {
	fmt.Println("  📊 Factory patterns:")
	
	// Create service factory
	factory := NewServiceFactory()
	
	// Register factory methods
	factory.RegisterFactory("userService", func() interface{} {
		logger := &ConsoleLogger{}
		emailService := &SMTPEmailService{}
		return NewUserService(logger, emailService)
	})
	
	factory.RegisterFactory("notificationService", func() interface{} {
		logger := &FileLogger{}
		smsService := &SMSServiceImpl{}
		return NewNotificationService(logger, smsService)
	})
	
	// Create services using factory
	userService := factory.Create("userService").(*UserService)
	notificationService := factory.Create("notificationService").(*NotificationService)
	
	// Use services
	user := &User{ID: 7, Name: "Eve Adams", Email: "eve@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully with factory")
	}
	
	err = notificationService.SendNotification("user7", "Factory pattern example")
	if err != nil {
		fmt.Printf("    Error sending notification: %v\n", err)
	} else {
		fmt.Println("    Notification sent successfully with factory")
	}
}

type ServiceFactory struct {
	factories map[string]func() interface{}
	mu        sync.RWMutex
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{
		factories: make(map[string]func() interface{}),
	}
}

func (f *ServiceFactory) RegisterFactory(name string, factory func() interface{}) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.factories[name] = factory
}

func (f *ServiceFactory) Create(name string) interface{} {
	f.mu.RLock()
	defer f.mu.RUnlock()
	
	factory, exists := f.factories[name]
	if !exists {
		return nil
	}
	
	return factory()
}

func builderPatterns() {
	fmt.Println("  📊 Builder patterns:")
	
	// Create service builder
	builder := NewServiceBuilder()
	
	// Build service with builder pattern
	userService := builder.
		WithLogger(&ConsoleLogger{}).
		WithEmailService(&SMTPEmailService{}).
		BuildUserService()
	
	// Use service
	user := &User{ID: 8, Name: "Frank Miller", Email: "frank@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully with builder pattern")
	}
}

type ServiceBuilder struct {
	logger       Logger
	emailService EmailService
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{}
}

func (b *ServiceBuilder) WithLogger(logger Logger) *ServiceBuilder {
	b.logger = logger
	return b
}

func (b *ServiceBuilder) WithEmailService(emailService EmailService) *ServiceBuilder {
	b.emailService = emailService
	return b
}

func (b *ServiceBuilder) BuildUserService() *UserService {
	return NewUserService(b.logger, b.emailService)
}

func singletonPatterns() {
	fmt.Println("  📊 Singleton patterns:")
	
	// Create singleton manager
	singletonManager := NewSingletonManager()
	
	// Get singleton instances
	logger1 := singletonManager.GetLogger()
	logger2 := singletonManager.GetLogger()
	
	// Verify they are the same instance
	if logger1 == logger2 {
		fmt.Println("    Singleton pattern working: same instance returned")
	} else {
		fmt.Println("    Singleton pattern failed: different instances returned")
	}
	
	// Use singleton
	logger1.Info("Singleton pattern example")
}

type SingletonManager struct {
	logger       Logger
	emailService EmailService
	mu           sync.Mutex
}

func NewSingletonManager() *SingletonManager {
	return &SingletonManager{}
}

func (sm *SingletonManager) GetLogger() Logger {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if sm.logger == nil {
		sm.logger = &ConsoleLogger{}
	}
	
	return sm.logger
}

func (sm *SingletonManager) GetEmailService() EmailService {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if sm.emailService == nil {
		sm.emailService = &SMTPEmailService{}
	}
	
	return sm.emailService
}

func scopedDependencies() {
	fmt.Println("  📊 Scoped dependencies:")
	
	// Create scoped service manager
	scopedManager := NewScopedServiceManager()
	
	// Create scope
	scope1 := scopedManager.CreateScope()
	scope2 := scopedManager.CreateScope()
	
	// Get services from different scopes
	userService1 := scope1.GetUserService()
	userService2 := scope2.GetUserService()
	
	// Verify they are different instances
	if userService1 != userService2 {
		fmt.Println("    Scoped dependencies working: different instances for different scopes")
	} else {
		fmt.Println("    Scoped dependencies failed: same instance for different scopes")
	}
	
	// Use services
	user1 := &User{ID: 9, Name: "Grace Wilson", Email: "grace@example.com"}
	user2 := &User{ID: 10, Name: "Henry Brown", Email: "henry@example.com"}
	
	userService1.CreateUser(user1)
	userService2.CreateUser(user2)
}

type ScopedServiceManager struct {
	mu sync.Mutex
}

func NewScopedServiceManager() *ScopedServiceManager {
	return &ScopedServiceManager{}
}

func (sm *ScopedServiceManager) CreateScope() *ServiceScope {
	return &ServiceScope{
		logger:       &ConsoleLogger{},
		emailService: &SMTPEmailService{},
	}
}

type ServiceScope struct {
	logger       Logger
	emailService EmailService
}

func (s *ServiceScope) GetUserService() *UserService {
	return NewUserService(s.logger, s.emailService)
}

// 5. Dependency Resolution
func dependencyResolution() {
	fmt.Println("5. Dependency Resolution:")
	fmt.Println("Understanding dependency resolution strategies...")

	// Demonstrate dependency resolution
	dependencyResolutionStrategies()
	
	// Show circular dependency handling
	circularDependencyHandling()
	
	// Demonstrate dependency validation
	dependencyValidation()
}

func dependencyResolutionStrategies() {
	fmt.Println("  📊 Dependency resolution strategies:")
	
	// Create dependency resolver with different strategies
	resolver := NewAdvancedDependencyResolver()
	
	// Register services
	resolver.Register("logger", &ConsoleLogger{})
	resolver.Register("email", &SMTPEmailService{})
	resolver.Register("userService", func() interface{} {
		logger := resolver.Resolve("logger").(Logger)
		emailService := resolver.Resolve("email").(EmailService)
		return NewUserService(logger, emailService)
	})
	
	// Resolve dependencies
	userServiceFactory := resolver.Resolve("userService")
	if userServiceFactory != nil {
		userService := userServiceFactory.(func() interface{})()
		userServiceInstance := userService.(*UserService)
		
		// Use service
		user := &User{ID: 11, Name: "Ivy Chen", Email: "ivy@example.com"}
		err := userServiceInstance.CreateUser(user)
		if err != nil {
			fmt.Printf("    Error creating user: %v\n", err)
		} else {
			fmt.Println("    User created successfully with dependency resolution")
		}
	}
}

type AdvancedDependencyResolver struct {
	services map[string]interface{}
	factories map[string]func() interface{}
	mu       sync.RWMutex
}

func NewAdvancedDependencyResolver() *AdvancedDependencyResolver {
	return &AdvancedDependencyResolver{
		services:  make(map[string]interface{}),
		factories: make(map[string]func() interface{}),
	}
}

func (r *AdvancedDependencyResolver) Register(name string, service interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = service
}

func (r *AdvancedDependencyResolver) RegisterFactory(name string, factory func() interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.factories[name] = factory
}

func (r *AdvancedDependencyResolver) Resolve(name string) interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	// Try to get from services first
	if service, exists := r.services[name]; exists {
		return service
	}
	
	// Try to create from factory
	if factory, exists := r.factories[name]; exists {
		return factory()
	}
	
	return nil
}

func circularDependencyHandling() {
	fmt.Println("  📊 Circular dependency handling:")
	
	// Create resolver with circular dependency detection
	resolver := NewCircularDependencyResolver()
	
	// Register services that could create circular dependencies
	resolver.Register("serviceA", func() interface{} {
		return &ServiceA{resolver: resolver}
	})
	
	resolver.Register("serviceB", func() interface{} {
		return &ServiceB{resolver: resolver}
	})
	
	// Try to resolve services
	serviceA := resolver.Resolve("serviceA")
	if serviceA != nil {
		fmt.Println("    ServiceA resolved successfully")
	} else {
		fmt.Println("    ServiceA resolution failed")
	}
}

type CircularDependencyResolver struct {
	services map[string]func() interface{}
	resolving map[string]bool
	mu       sync.RWMutex
}

func NewCircularDependencyResolver() *CircularDependencyResolver {
	return &CircularDependencyResolver{
		services:  make(map[string]func() interface{}),
		resolving: make(map[string]bool),
	}
}

func (r *CircularDependencyResolver) Register(name string, factory func() interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = factory
}

func (r *CircularDependencyResolver) Resolve(name string) interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	// Check for circular dependency
	if r.resolving[name] {
		fmt.Printf("    Circular dependency detected for %s\n", name)
		return nil
	}
	
	// Mark as resolving
	r.resolving[name] = true
	defer delete(r.resolving, name)
	
	// Get factory and create instance
	factory, exists := r.services[name]
	if !exists {
		return nil
	}
	
	return factory()
}

type ServiceA struct {
	resolver *CircularDependencyResolver
}

type ServiceB struct {
	resolver *CircularDependencyResolver
}

func dependencyValidation() {
	fmt.Println("  📊 Dependency validation:")
	
	// Create resolver with validation
	resolver := NewValidatingDependencyResolver()
	
	// Register services
	resolver.Register("logger", &ConsoleLogger{})
	resolver.Register("email", &SMTPEmailService{})
	
	// Validate dependencies
	valid := resolver.ValidateDependencies()
	if valid {
		fmt.Println("    All dependencies are valid")
	} else {
		fmt.Println("    Some dependencies are invalid")
	}
}

type ValidatingDependencyResolver struct {
	services map[string]interface{}
	mu       sync.RWMutex
}

func NewValidatingDependencyResolver() *ValidatingDependencyResolver {
	return &ValidatingDependencyResolver{
		services: make(map[string]interface{}),
	}
}

func (r *ValidatingDependencyResolver) Register(name string, service interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = service
}

func (r *ValidatingDependencyResolver) ValidateDependencies() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	// Check if all required services are registered
	requiredServices := []string{"logger", "email"}
	
	for _, serviceName := range requiredServices {
		if _, exists := r.services[serviceName]; !exists {
			fmt.Printf("    Missing required service: %s\n", serviceName)
			return false
		}
	}
	
	return true
}

// 6. Lifecycle Management
func lifecycleManagement() {
	fmt.Println("6. Lifecycle Management:")
	fmt.Println("Understanding dependency lifecycle management...")

	// Demonstrate lifecycle management
	lifecycleManagementExample()
	
	// Show cleanup patterns
	cleanupPatterns()
	
	// Demonstrate resource management
	resourceManagement()
}

func lifecycleManagementExample() {
	fmt.Println("  📊 Lifecycle management example:")
	
	// Create service with lifecycle management
	service := NewLifecycleManagedService()
	
	// Initialize service
	err := service.Initialize()
	if err != nil {
		fmt.Printf("    Error initializing service: %v\n", err)
		return
	}
	
	// Use service
	service.ProcessRequest("test request")
	
	// Cleanup service
	err = service.Cleanup()
	if err != nil {
		fmt.Printf("    Error cleaning up service: %v\n", err)
	} else {
		fmt.Println("    Service cleaned up successfully")
	}
}

type LifecycleManagedService struct {
	logger       Logger
	emailService EmailService
	initialized  bool
	mu           sync.Mutex
}

func NewLifecycleManagedService() *LifecycleManagedService {
	return &LifecycleManagedService{}
}

func (s *LifecycleManagedService) Initialize() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if s.initialized {
		return fmt.Errorf("service already initialized")
	}
	
	// Initialize dependencies
	s.logger = &ConsoleLogger{}
	s.emailService = &SMTPEmailService{}
	
	s.initialized = true
	fmt.Println("    Service initialized successfully")
	return nil
}

func (s *LifecycleManagedService) ProcessRequest(request string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if !s.initialized {
		fmt.Println("    Service not initialized")
		return
	}
	
	s.logger.Info("Processing request: %s", request)
	time.Sleep(50 * time.Millisecond)
	s.logger.Info("Request processed successfully")
}

func (s *LifecycleManagedService) Cleanup() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if !s.initialized {
		return fmt.Errorf("service not initialized")
	}
	
	// Cleanup resources
	s.logger = nil
	s.emailService = nil
	s.initialized = false
	
	fmt.Println("    Service cleaned up successfully")
	return nil
}

func cleanupPatterns() {
	fmt.Println("  📊 Cleanup patterns:")
	
	// Create service with cleanup
	service := NewCleanupService()
	
	// Use service with defer cleanup
	func() {
		defer service.Cleanup()
		service.DoWork()
	}()
	
	fmt.Println("    Service used with defer cleanup")
}

type CleanupService struct {
	resources []string
	mu        sync.Mutex
}

func NewCleanupService() *CleanupService {
	return &CleanupService{
		resources: make([]string, 0),
	}
}

func (s *CleanupService) DoWork() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Simulate resource allocation
	s.resources = append(s.resources, "resource1", "resource2", "resource3")
	fmt.Printf("    Allocated resources: %v\n", s.resources)
}

func (s *CleanupService) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Simulate resource cleanup
	fmt.Printf("    Cleaning up resources: %v\n", s.resources)
	s.resources = s.resources[:0]
}

func resourceManagement() {
	fmt.Println("  📊 Resource management:")
	
	// Create resource manager
	manager := NewResourceManager()
	
	// Acquire resources
	resource1 := manager.AcquireResource("database")
	resource2 := manager.AcquireResource("cache")
	
	// Use resources
	fmt.Printf("    Using resource: %s\n", resource1)
	fmt.Printf("    Using resource: %s\n", resource2)
	
	// Release resources
	manager.ReleaseResource(resource1)
	manager.ReleaseResource(resource2)
}

type ResourceManager struct {
	resources map[string]bool
	mu        sync.Mutex
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		resources: make(map[string]bool),
	}
}

func (rm *ResourceManager) AcquireResource(name string) string {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	rm.resources[name] = true
	return name
}

func (rm *ResourceManager) ReleaseResource(name string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	delete(rm.resources, name)
	fmt.Printf("    Released resource: %s\n", name)
}

// 7. Testing with DI
func testingWithDI() {
	fmt.Println("7. Testing with DI:")
	fmt.Println("Understanding testing with dependency injection...")

	// Demonstrate testing with mocks
	testingWithMocks()
	
	// Show test doubles
	testDoubles()
	
	// Demonstrate integration testing
	integrationTesting()
}

func testingWithMocks() {
	fmt.Println("  📊 Testing with mocks:")
	
	// Create mock logger
	mockLogger := &MockLogger{}
	
	// Create mock email service
	mockEmailService := &MockEmailService{}
	
	// Create service with mocks
	userService := NewUserService(mockLogger, mockEmailService)
	
	// Test service
	user := &User{ID: 12, Name: "Jack Smith", Email: "jack@example.com"}
	err := userService.CreateUser(user)
	
	if err != nil {
		fmt.Printf("    Test failed: %v\n", err)
	} else {
		fmt.Println("    Test passed: user created successfully")
	}
	
	// Verify mock calls
	if mockLogger.GetInfoCallCount() > 0 {
		fmt.Println("    Mock logger Info method was called")
	}
	
	if mockEmailService.GetSendEmailCallCount() > 0 {
		fmt.Println("    Mock email service SendEmail method was called")
	}
}

type MockEmailService struct {
	SendEmailCalled bool
	LastTo          string
	LastSubject     string
	LastBody        string
	callCount       int
	mu              sync.Mutex
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.SendEmailCalled = true
	m.LastTo = to
	m.LastSubject = subject
	m.LastBody = body
	m.callCount++
	fmt.Printf("    Mock email sent to %s: %s\n", to, subject)
	return nil
}

func (m *MockEmailService) GetSendEmailCallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.callCount
}

func testDoubles() {
	fmt.Println("  📊 Test doubles:")
	
	// Create test double
	testDouble := NewTestDouble()
	
	// Register test implementations
	testDouble.RegisterLogger(&MockLogger{})
	testDouble.RegisterEmailService(&MockEmailService{})
	
	// Create service with test double
	userService := testDouble.CreateUserService()
	
	// Test service
	user := &User{ID: 13, Name: "Kate Johnson", Email: "kate@example.com"}
	err := userService.CreateUser(user)
	
	if err != nil {
		fmt.Printf("    Test failed: %v\n", err)
	} else {
		fmt.Println("    Test passed with test double")
	}
}

type TestDouble struct {
	logger       Logger
	emailService EmailService
}

func NewTestDouble() *TestDouble {
	return &TestDouble{}
}

func (td *TestDouble) RegisterLogger(logger Logger) {
	td.logger = logger
}

func (td *TestDouble) RegisterEmailService(emailService EmailService) {
	td.emailService = emailService
}

func (td *TestDouble) CreateUserService() *UserService {
	return NewUserService(td.logger, td.emailService)
}

func integrationTesting() {
	fmt.Println("  📊 Integration testing:")
	
	// Create integration test setup
	testSetup := NewIntegrationTestSetup()
	
	// Run integration test
	err := testSetup.RunIntegrationTest()
	if err != nil {
		fmt.Printf("    Integration test failed: %v\n", err)
	} else {
		fmt.Println("    Integration test passed")
	}
}

type IntegrationTestSetup struct {
	container *ServiceContainer
}

func NewIntegrationTestSetup() *IntegrationTestSetup {
	container := NewServiceContainer()
	
	// Register real implementations for integration testing
	container.RegisterSingleton("logger", &ConsoleLogger{})
	container.RegisterSingleton("email", &SMTPEmailService{})
	container.RegisterScoped("userService", func() interface{} {
		logger := container.Get("logger").(Logger)
		emailService := container.Get("email").(EmailService)
		return NewUserService(logger, emailService)
	})
	
	return &IntegrationTestSetup{
		container: container,
	}
}

func (its *IntegrationTestSetup) RunIntegrationTest() error {
	// Get service from container
	userService := its.container.Get("userService").(*UserService)
	
	// Test service
	user := &User{ID: 14, Name: "Liam Brown", Email: "liam@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		return err
	}
	
	fmt.Println("    Integration test completed successfully")
	return nil
}

// 8. Best Practices
func dependencyInjectionBestPractices() {
	fmt.Println("8. Dependency Injection Best Practices:")
	fmt.Println("Best practices for dependency injection...")

	fmt.Println("  📝 Best Practice 1: Use interfaces for dependencies")
	fmt.Println("    - Define clear interfaces for all dependencies")
	fmt.Println("    - Make interfaces focused and cohesive")
	fmt.Println("    - Avoid concrete type dependencies")
	
	fmt.Println("  📝 Best Practice 2: Prefer constructor injection")
	fmt.Println("    - Inject dependencies through constructors")
	fmt.Println("    - Make dependencies explicit and required")
	fmt.Println("    - Avoid setter injection when possible")
	
	fmt.Println("  📝 Best Practice 3: Use dependency containers")
	fmt.Println("    - Centralize dependency registration")
	fmt.Println("    - Manage dependency lifecycles")
	fmt.Println("    - Handle complex dependency graphs")
	
	fmt.Println("  📝 Best Practice 4: Implement proper lifecycle management")
	fmt.Println("    - Initialize dependencies in correct order")
	fmt.Println("    - Clean up resources properly")
	fmt.Println("    - Handle errors during initialization")
	
	fmt.Println("  📝 Best Practice 5: Use mocks for testing")
	fmt.Println("    - Create mock implementations for testing")
	fmt.Println("    - Verify mock method calls")
	fmt.Println("    - Test error scenarios with mocks")
	
	fmt.Println("  📝 Best Practice 6: Avoid circular dependencies")
	fmt.Println("    - Design dependency graph to avoid cycles")
	fmt.Println("    - Use events or callbacks for circular references")
	fmt.Println("    - Refactor to break circular dependencies")
	
	fmt.Println("  📝 Best Practice 7: Keep dependencies minimal")
	fmt.Println("    - Only inject what you actually need")
	fmt.Println("    - Avoid god objects with many dependencies")
	fmt.Println("    - Use composition over inheritance")
}
