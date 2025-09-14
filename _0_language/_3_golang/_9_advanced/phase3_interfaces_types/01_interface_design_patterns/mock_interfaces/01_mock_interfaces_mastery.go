package main

import (
	"fmt"
	"sync"
	"time"
)

// 🎭 MOCK INTERFACES MASTERY
// Understanding mock interfaces for testing in Go

// Define types for mock interfaces
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

func main() {
	fmt.Println("🎭 MOCK INTERFACES MASTERY")
	fmt.Println("=========================")
	fmt.Println()

	// 1. Mock Generation
	mockGeneration()
	fmt.Println()

	// 2. Test Doubles
	testDoubles()
	fmt.Println()

	// 3. Advanced Mocking
	advancedMocking()
	fmt.Println()

	// 4. Testing Patterns
	testingPatterns()
	fmt.Println()

	// 5. Mock Verification
	mockVerification()
	fmt.Println()

	// 6. Mock Expectations
	mockExpectations()
	fmt.Println()

	// 7. Mock Behavior Configuration
	mockBehaviorConfiguration()
	fmt.Println()

	// 8. Best Practices
	mockInterfacesBestPractices()
}

// 1. Mock Generation
func mockGeneration() {
	fmt.Println("1. Mock Generation:")
	fmt.Println("Understanding mock generation techniques...")

	// Demonstrate manual mock creation
	manualMockCreation()
	
	// Show mock generation tools
	mockGenerationTools()
	
	// Demonstrate mock interface patterns
	mockInterfacePatterns()
}

func manualMockCreation() {
	fmt.Println("  📊 Manual mock creation:")
	
	// Create manual mock
	mockLogger := &ManualMockLogger{}
	mockEmailService := &ManualMockEmailService{}
	
	// Use mocks
	userService := NewUserService(mockLogger, mockEmailService)
	user := &User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error creating user: %v\n", err)
	} else {
		fmt.Println("    User created successfully with manual mock")
	}
	
	// Verify mock calls
	fmt.Printf("    Mock logger Info calls: %d\n", mockLogger.InfoCallCount)
	fmt.Printf("    Mock email service calls: %d\n", mockEmailService.SendEmailCallCount)
}

type ManualMockLogger struct {
	InfoCallCount  int
	ErrorCallCount int
	LastInfoMsg    string
	LastErrorMsg   string
}

func (m *ManualMockLogger) Info(format string, args ...interface{}) {
	m.InfoCallCount++
	m.LastInfoMsg = fmt.Sprintf(format, args...)
	fmt.Printf("    [MOCK-INFO] %s\n", m.LastInfoMsg)
}

func (m *ManualMockLogger) Error(format string, args ...interface{}) {
	m.ErrorCallCount++
	m.LastErrorMsg = fmt.Sprintf(format, args...)
	fmt.Printf("    [MOCK-ERROR] %s\n", m.LastErrorMsg)
}

type ManualMockEmailService struct {
	SendEmailCallCount int
	LastTo             string
	LastSubject        string
	LastBody           string
	ShouldFail         bool
}

func (m *ManualMockEmailService) SendEmail(to, subject, body string) error {
	m.SendEmailCallCount++
	m.LastTo = to
	m.LastSubject = subject
	m.LastBody = body
	
	if m.ShouldFail {
		return fmt.Errorf("mock email service failure")
	}
	
	fmt.Printf("    [MOCK-EMAIL] Sending to %s: %s\n", to, subject)
	return nil
}

func mockGenerationTools() {
	fmt.Println("  📊 Mock generation tools:")
	fmt.Println("    Popular Go mock generation tools:")
	fmt.Println("      - mockgen (golang/mock)")
	fmt.Println("      - counterfeiter")
	fmt.Println("      - gomock")
	fmt.Println("      - mockery")
	fmt.Println("    Benefits:")
	fmt.Println("      - Automatic mock generation")
	fmt.Println("      - Type-safe mocks")
	fmt.Println("      - Reduced boilerplate code")
	fmt.Println("      - Consistent mock patterns")
}

func mockInterfacePatterns() {
	fmt.Println("  📊 Mock interface patterns:")
	
	// Demonstrate different mock patterns
	simpleMock := &SimpleMock{}
	advancedMock := NewAdvancedMock()
	
	// Use simple mock
	simpleMock.DoSomething("test")
	fmt.Printf("    Simple mock calls: %d\n", simpleMock.CallCount)
	
	// Use advanced mock
	advancedMock.ConfigureBehavior("test", "mocked result")
	result := advancedMock.DoSomething("test")
	fmt.Printf("    Advanced mock result: %s\n", result)
}

type SimpleMock struct {
	CallCount int
}

func (m *SimpleMock) DoSomething(input string) {
	m.CallCount++
	fmt.Printf("    Simple mock: %s\n", input)
}

type AdvancedMock struct {
	behaviors map[string]string
	callCount int
	mu        sync.Mutex
}

func NewAdvancedMock() *AdvancedMock {
	return &AdvancedMock{
		behaviors: make(map[string]string),
	}
}

func (m *AdvancedMock) ConfigureBehavior(input, output string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.behaviors[input] = output
}

func (m *AdvancedMock) DoSomething(input string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.callCount++
	if output, exists := m.behaviors[input]; exists {
		return output
	}
	return "default output"
}

// 2. Test Doubles
func testDoubles() {
	fmt.Println("2. Test Doubles:")
	fmt.Println("Understanding different types of test doubles...")

	// Demonstrate dummy objects
	dummyObjects()
	
	// Show stub objects
	stubObjects()
	
	// Demonstrate mock objects
	mockObjects()
	
	// Show fake objects
	fakeObjects()
}

func dummyObjects() {
	fmt.Println("  📊 Dummy objects:")
	
	// Create dummy objects
	dummyLogger := &DummyLogger{}
	dummyEmailService := &DummyEmailService{}
	
	// Use dummy objects
	userService := NewUserService(dummyLogger, dummyEmailService)
	user := &User{ID: 2, Name: "Jane Smith", Email: "jane@example.com"}
	
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error with dummy objects: %v\n", err)
	} else {
		fmt.Println("    Dummy objects work correctly")
	}
}

type DummyLogger struct{}

func (d *DummyLogger) Info(format string, args ...interface{}) {
	// Do nothing - dummy implementation
}

func (d *DummyLogger) Error(format string, args ...interface{}) {
	// Do nothing - dummy implementation
}

type DummyEmailService struct{}

func (d *DummyEmailService) SendEmail(to, subject, body string) error {
	// Do nothing - dummy implementation
	return nil
}

func stubObjects() {
	fmt.Println("  📊 Stub objects:")
	
	// Create stub objects
	stubLogger := &StubLogger{}
	stubEmailService := &StubEmailService{}
	
	// Use stub objects
	userService := NewUserService(stubLogger, stubEmailService)
	user := &User{ID: 3, Name: "Bob Johnson", Email: "bob@example.com"}
	
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error with stub objects: %v\n", err)
	} else {
		fmt.Println("    Stub objects work correctly")
	}
	
	// Verify stub behavior
	fmt.Printf("    Stub logger calls: %d\n", stubLogger.CallCount)
	fmt.Printf("    Stub email service calls: %d\n", stubEmailService.CallCount)
}

type StubLogger struct {
	CallCount int
}

func (s *StubLogger) Info(format string, args ...interface{}) {
	s.CallCount++
	// Stub implementation - minimal behavior
}

func (s *StubLogger) Error(format string, args ...interface{}) {
	s.CallCount++
	// Stub implementation - minimal behavior
}

type StubEmailService struct {
	CallCount int
}

func (s *StubEmailService) SendEmail(to, subject, body string) error {
	s.CallCount++
	// Stub implementation - always succeeds
	return nil
}

func mockObjects() {
	fmt.Println("  📊 Mock objects:")
	
	// Create mock objects
	mockLogger := &MockLogger{}
	mockEmailService := &MockEmailService{}
	
	// Configure mock behavior
	mockEmailService.SetShouldFail(false)
	
	// Use mock objects
	userService := NewUserService(mockLogger, mockEmailService)
	user := &User{ID: 4, Name: "Alice Brown", Email: "alice@example.com"}
	
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error with mock objects: %v\n", err)
	} else {
		fmt.Println("    Mock objects work correctly")
	}
	
	// Verify mock calls
	fmt.Printf("    Mock logger Info calls: %d\n", mockLogger.GetInfoCallCount())
	fmt.Printf("    Mock email service calls: %d\n", mockEmailService.GetSendEmailCallCount())
}

type MockLogger struct {
	infoCallCount  int
	errorCallCount int
	mu             sync.Mutex
}

func (m *MockLogger) Info(format string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.infoCallCount++
}

func (m *MockLogger) Error(format string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errorCallCount++
}

func (m *MockLogger) GetInfoCallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.infoCallCount
}

func (m *MockLogger) GetErrorCallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.errorCallCount
}

type MockEmailService struct {
	sendEmailCallCount int
	shouldFail         bool
	mu                 sync.Mutex
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.sendEmailCallCount++
	if m.shouldFail {
		return fmt.Errorf("mock email service failure")
	}
	return nil
}

func (m *MockEmailService) SetShouldFail(shouldFail bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.shouldFail = shouldFail
}

func (m *MockEmailService) GetSendEmailCallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sendEmailCallCount
}

func fakeObjects() {
	fmt.Println("  📊 Fake objects:")
	
	// Create fake objects
	fakeLogger := &FakeLogger{}
	fakeEmailService := &FakeEmailService{}
	
	// Use fake objects
	userService := NewUserService(fakeLogger, fakeEmailService)
	user := &User{ID: 5, Name: "Charlie Wilson", Email: "charlie@example.com"}
	
	err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("    Error with fake objects: %v\n", err)
	} else {
		fmt.Println("    Fake objects work correctly")
	}
	
	// Verify fake behavior
	fmt.Printf("    Fake logger messages: %v\n", fakeLogger.GetMessages())
	fmt.Printf("    Fake email service emails: %v\n", fakeEmailService.GetEmails())
}

type FakeLogger struct {
	messages []string
	mu       sync.Mutex
}

func (f *FakeLogger) Info(format string, args ...interface{}) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.messages = append(f.messages, fmt.Sprintf("[INFO] "+format, args...))
}

func (f *FakeLogger) Error(format string, args ...interface{}) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.messages = append(f.messages, fmt.Sprintf("[ERROR] "+format, args...))
}

func (f *FakeLogger) GetMessages() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.messages
}

type FakeEmailService struct {
	emails []Email
	mu     sync.Mutex
}

type Email struct {
	To      string
	Subject string
	Body    string
}

func (f *FakeEmailService) SendEmail(to, subject, body string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	f.emails = append(f.emails, Email{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	return nil
}

func (f *FakeEmailService) GetEmails() []Email {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.emails
}

// 3. Advanced Mocking
func advancedMocking() {
	fmt.Println("3. Advanced Mocking:")
	fmt.Println("Understanding advanced mocking techniques...")

	// Demonstrate mock verification
	mockVerificationExample()
	
	// Show mock expectations
	mockExpectationsExample()
	
	// Demonstrate mock behavior configuration
	mockBehaviorConfigurationExample()
}

func mockVerificationExample() {
	fmt.Println("  📊 Mock verification example:")
	
	// Create verifiable mock
	verifiableMock := &VerifiableMock{}
	
	// Use mock
	verifiableMock.DoSomething("test1")
	verifiableMock.DoSomething("test2")
	verifiableMock.DoSomething("test3")
	
	// Verify calls
	if verifiableMock.WasCalledWith("test1") {
		fmt.Println("    Mock was called with 'test1'")
	}
	
	if verifiableMock.WasCalledWith("test4") {
		fmt.Println("    Mock was called with 'test4'")
	} else {
		fmt.Println("    Mock was NOT called with 'test4'")
	}
	
	fmt.Printf("    Total calls: %d\n", verifiableMock.GetCallCount())
}

type VerifiableMock struct {
	calls []string
	mu    sync.Mutex
}

func (v *VerifiableMock) DoSomething(input string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.calls = append(v.calls, input)
}

func (v *VerifiableMock) WasCalledWith(input string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	
	for _, call := range v.calls {
		if call == input {
			return true
		}
	}
	return false
}

func (v *VerifiableMock) GetCallCount() int {
	v.mu.Lock()
	defer v.mu.Unlock()
	return len(v.calls)
}

func mockExpectationsExample() {
	fmt.Println("  📊 Mock expectations example:")
	
	// Create mock with expectations
	expectationMock := &ExpectationMock{}
	
	// Set expectations
	expectationMock.ExpectCall("method1", "arg1")
	expectationMock.ExpectCall("method2", "arg2")
	
	// Use mock
	expectationMock.Call("method1", "arg1")
	expectationMock.Call("method2", "arg2")
	
	// Verify expectations
	if expectationMock.VerifyExpectations() {
		fmt.Println("    All expectations met")
	} else {
		fmt.Println("    Some expectations not met")
	}
}

type ExpectationMock struct {
	expectations []Expectation
	calls        []Call
	mu           sync.Mutex
}

type Expectation struct {
	Method string
	Arg    string
}

type Call struct {
	Method string
	Arg    string
}

func (e *ExpectationMock) ExpectCall(method, arg string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.expectations = append(e.expectations, Expectation{Method: method, Arg: arg})
}

func (e *ExpectationMock) Call(method, arg string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.calls = append(e.calls, Call{Method: method, Arg: arg})
}

func (e *ExpectationMock) VerifyExpectations() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	
	if len(e.expectations) != len(e.calls) {
		return false
	}
	
	for i, expectation := range e.expectations {
		if i >= len(e.calls) {
			return false
		}
		call := e.calls[i]
		if expectation.Method != call.Method || expectation.Arg != call.Arg {
			return false
		}
	}
	
	return true
}

func mockBehaviorConfigurationExample() {
	fmt.Println("  📊 Mock behavior configuration example:")
	
	// Create configurable mock
	configurableMock := NewConfigurableMock()
	
	// Configure behavior
	configurableMock.ConfigureBehavior("input1", "output1")
	configurableMock.ConfigureBehavior("input2", "output2")
	configurableMock.ConfigureError("input3", "error3")
	
	// Use mock
	result1 := configurableMock.Process("input1")
	fmt.Printf("    Input 'input1' -> Output: %s\n", result1)
	
	result2 := configurableMock.Process("input2")
	fmt.Printf("    Input 'input2' -> Output: %s\n", result2)
	
	result3 := configurableMock.Process("input3")
	fmt.Printf("    Input 'input3' -> Output: %s\n", result3)
}

type ConfigurableMock struct {
	behaviors map[string]string
	errors    map[string]string
	mu        sync.Mutex
}

func NewConfigurableMock() *ConfigurableMock {
	return &ConfigurableMock{
		behaviors: make(map[string]string),
		errors:    make(map[string]string),
	}
}

func (c *ConfigurableMock) ConfigureBehavior(input, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.behaviors[input] = output
}

func (c *ConfigurableMock) ConfigureError(input, errorMsg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.errors[input] = errorMsg
}

func (c *ConfigurableMock) Process(input string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if errorMsg, exists := c.errors[input]; exists {
		return "Error: " + errorMsg
	}
	
	if output, exists := c.behaviors[input]; exists {
		return output
	}
	
	return "default output"
}

// 4. Testing Patterns
func testingPatterns() {
	fmt.Println("4. Testing Patterns:")
	fmt.Println("Understanding testing patterns with mocks...")

	// Demonstrate unit testing with mocks
	unitTestingWithMocks()
	
	// Show integration testing
	integrationTestingWithMocks()
	
	// Demonstrate test isolation
	testIsolation()
}

func unitTestingWithMocks() {
	fmt.Println("  📊 Unit testing with mocks:")
	
	// Create mocks for unit test
	mockLogger := &MockLogger{}
	mockEmailService := &MockEmailService{}
	
	// Create service under test
	userService := NewUserService(mockLogger, mockEmailService)
	
	// Test case 1: Successful user creation
	user := &User{ID: 6, Name: "David Lee", Email: "david@example.com"}
	err := userService.CreateUser(user)
	
	if err != nil {
		fmt.Printf("    Test case 1 failed: %v\n", err)
	} else {
		fmt.Println("    Test case 1 passed: User created successfully")
	}
	
	// Verify mock calls
	if mockLogger.GetInfoCallCount() > 0 {
		fmt.Println("    Mock logger was called")
	}
	
	if mockEmailService.GetSendEmailCallCount() > 0 {
		fmt.Println("    Mock email service was called")
	}
	
	// Test case 2: Email service failure
	mockEmailService.SetShouldFail(true)
	user2 := &User{ID: 7, Name: "Eve Adams", Email: "eve@example.com"}
	err = userService.CreateUser(user2)
	
	if err != nil {
		fmt.Println("    Test case 2 passed: Email service failure handled correctly")
	} else {
		fmt.Println("    Test case 2 failed: Email service failure not handled")
	}
}

func integrationTestingWithMocks() {
	fmt.Println("  📊 Integration testing with mocks:")
	
	// Create integration test setup
	integrationTest := NewIntegrationTest()
	
	// Run integration test
	err := integrationTest.RunTest()
	if err != nil {
		fmt.Printf("    Integration test failed: %v\n", err)
	} else {
		fmt.Println("    Integration test passed")
	}
}

type IntegrationTest struct {
	mockLogger       *MockLogger
	mockEmailService *MockEmailService
}

func NewIntegrationTest() *IntegrationTest {
	return &IntegrationTest{
		mockLogger:       &MockLogger{},
		mockEmailService: &MockEmailService{},
	}
}

func (it *IntegrationTest) RunTest() error {
	// Create service with mocks
	userService := NewUserService(it.mockLogger, it.mockEmailService)
	
	// Test user creation
	user := &User{ID: 8, Name: "Frank Miller", Email: "frank@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		return err
	}
	
	// Verify integration
	if it.mockLogger.GetInfoCallCount() == 0 {
		return fmt.Errorf("logger not called")
	}
	
	if it.mockEmailService.GetSendEmailCallCount() == 0 {
		return fmt.Errorf("email service not called")
	}
	
	return nil
}

func testIsolation() {
	fmt.Println("  📊 Test isolation:")
	
	// Create isolated test
	isolatedTest := NewIsolatedTest()
	
	// Run isolated test
	err := isolatedTest.RunIsolatedTest()
	if err != nil {
		fmt.Printf("    Isolated test failed: %v\n", err)
	} else {
		fmt.Println("    Isolated test passed")
	}
}

type IsolatedTest struct {
	mockLogger       *MockLogger
	mockEmailService *MockEmailService
}

func NewIsolatedTest() *IsolatedTest {
	return &IsolatedTest{
		mockLogger:       &MockLogger{},
		mockEmailService: &MockEmailService{},
	}
}

func (it *IsolatedTest) RunIsolatedTest() error {
	// Reset mocks for isolation
	it.mockLogger = &MockLogger{}
	it.mockEmailService = &MockEmailService{}
	
	// Create service with fresh mocks
	userService := NewUserService(it.mockLogger, it.mockEmailService)
	
	// Test in isolation
	user := &User{ID: 9, Name: "Grace Wilson", Email: "grace@example.com"}
	err := userService.CreateUser(user)
	if err != nil {
		return err
	}
	
	// Verify isolation
	if it.mockLogger.GetInfoCallCount() != 2 { // 2 expected calls
		return fmt.Errorf("unexpected logger call count: %d", it.mockLogger.GetInfoCallCount())
	}
	
	if it.mockEmailService.GetSendEmailCallCount() != 1 { // 1 expected call
		return fmt.Errorf("unexpected email service call count: %d", it.mockEmailService.GetSendEmailCallCount())
	}
	
	return nil
}

// 5. Mock Verification
func mockVerification() {
	fmt.Println("5. Mock Verification:")
	fmt.Println("Understanding mock verification techniques...")

	// Demonstrate call verification
	callVerification()
	
	// Show argument verification
	argumentVerification()
	
	// Demonstrate call order verification
	callOrderVerification()
}

func callVerification() {
	fmt.Println("  📊 Call verification:")
	
	// Create verifiable mock
	verifiableMock := &VerifiableMock{}
	
	// Use mock
	verifiableMock.DoSomething("test1")
	verifiableMock.DoSomething("test2")
	
	// Verify calls
	fmt.Printf("    Total calls: %d\n", verifiableMock.GetCallCount())
	fmt.Printf("    Called with 'test1': %t\n", verifiableMock.WasCalledWith("test1"))
	fmt.Printf("    Called with 'test2': %t\n", verifiableMock.WasCalledWith("test2"))
	fmt.Printf("    Called with 'test3': %t\n", verifiableMock.WasCalledWith("test3"))
}

func argumentVerification() {
	fmt.Println("  📊 Argument verification:")
	
	// Create argument verifiable mock
	argVerifiableMock := &ArgumentVerifiableMock{}
	
	// Use mock
	argVerifiableMock.Process("input1", 123)
	argVerifiableMock.Process("input2", 456)
	
	// Verify arguments
	fmt.Printf("    Called with ('input1', 123): %t\n", argVerifiableMock.WasCalledWith("input1", 123))
	fmt.Printf("    Called with ('input2', 456): %t\n", argVerifiableMock.WasCalledWith("input2", 456))
	fmt.Printf("    Called with ('input3', 789): %t\n", argVerifiableMock.WasCalledWith("input3", 789))
}

type ArgumentVerifiableMock struct {
	calls []CallWithArgs
	mu    sync.Mutex
}

type CallWithArgs struct {
	StrArg string
	IntArg int
}

func (a *ArgumentVerifiableMock) Process(strArg string, intArg int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.calls = append(a.calls, CallWithArgs{StrArg: strArg, IntArg: intArg})
}

func (a *ArgumentVerifiableMock) WasCalledWith(strArg string, intArg int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	
	for _, call := range a.calls {
		if call.StrArg == strArg && call.IntArg == intArg {
			return true
		}
	}
	return false
}

func callOrderVerification() {
	fmt.Println("  📊 Call order verification:")
	
	// Create order verifiable mock
	orderVerifiableMock := &OrderVerifiableMock{}
	
	// Use mock in specific order
	orderVerifiableMock.Method1("arg1")
	orderVerifiableMock.Method2("arg2")
	orderVerifiableMock.Method1("arg3")
	
	// Verify call order
	if orderVerifiableMock.VerifyOrder([]string{"Method1", "Method2", "Method1"}) {
		fmt.Println("    Call order verified correctly")
	} else {
		fmt.Println("    Call order verification failed")
	}
}

type OrderVerifiableMock struct {
	callOrder []string
	mu        sync.Mutex
}

func (o *OrderVerifiableMock) Method1(arg string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.callOrder = append(o.callOrder, "Method1")
}

func (o *OrderVerifiableMock) Method2(arg string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.callOrder = append(o.callOrder, "Method2")
}

func (o *OrderVerifiableMock) VerifyOrder(expectedOrder []string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	
	if len(o.callOrder) != len(expectedOrder) {
		return false
	}
	
	for i, call := range o.callOrder {
		if call != expectedOrder[i] {
			return false
		}
	}
	
	return true
}

// 6. Mock Expectations
func mockExpectations() {
	fmt.Println("6. Mock Expectations:")
	fmt.Println("Understanding mock expectations...")

	// Demonstrate expectation setting
	expectationSetting()
	
	// Show expectation verification
	expectationVerification()
	
	// Demonstrate expectation failure handling
	expectationFailureHandling()
}

func expectationSetting() {
	fmt.Println("  📊 Expectation setting:")
	
	// Create expectation mock
	expectationMock := &ExpectationMock{}
	
	// Set expectations
	expectationMock.ExpectCall("method1", "arg1")
	expectationMock.ExpectCall("method2", "arg2")
	expectationMock.ExpectCall("method1", "arg3")
	
	fmt.Println("    Expectations set:")
	fmt.Println("      - method1('arg1')")
	fmt.Println("      - method2('arg2')")
	fmt.Println("      - method1('arg3')")
}

func expectationVerification() {
	fmt.Println("  📊 Expectation verification:")
	
	// Create expectation mock
	expectationMock := &ExpectationMock{}
	
	// Set expectations
	expectationMock.ExpectCall("method1", "arg1")
	expectationMock.ExpectCall("method2", "arg2")
	
	// Make calls
	expectationMock.Call("method1", "arg1")
	expectationMock.Call("method2", "arg2")
	
	// Verify expectations
	if expectationMock.VerifyExpectations() {
		fmt.Println("    All expectations met")
	} else {
		fmt.Println("    Some expectations not met")
	}
}

func expectationFailureHandling() {
	fmt.Println("  📊 Expectation failure handling:")
	
	// Create expectation mock
	expectationMock := &ExpectationMock{}
	
	// Set expectations
	expectationMock.ExpectCall("method1", "arg1")
	expectationMock.ExpectCall("method2", "arg2")
	
	// Make calls (missing one expectation)
	expectationMock.Call("method1", "arg1")
	// Missing: expectationMock.Call("method2", "arg2")
	
	// Verify expectations
	if expectationMock.VerifyExpectations() {
		fmt.Println("    All expectations met")
	} else {
		fmt.Println("    Some expectations not met - this is expected")
	}
}

// 7. Mock Behavior Configuration
func mockBehaviorConfiguration() {
	fmt.Println("7. Mock Behavior Configuration:")
	fmt.Println("Understanding mock behavior configuration...")

	// Demonstrate behavior configuration
	behaviorConfiguration()
	
	// Show dynamic behavior
	dynamicBehavior()
	
	// Demonstrate behavior chaining
	behaviorChaining()
}

func behaviorConfiguration() {
	fmt.Println("  📊 Behavior configuration:")
	
	// Create configurable mock
	configurableMock := NewConfigurableMock()
	
	// Configure different behaviors
	configurableMock.ConfigureBehavior("input1", "output1")
	configurableMock.ConfigureBehavior("input2", "output2")
	configurableMock.ConfigureError("input3", "error3")
	
	// Test behaviors
	fmt.Printf("    Input 'input1' -> %s\n", configurableMock.Process("input1"))
	fmt.Printf("    Input 'input2' -> %s\n", configurableMock.Process("input2"))
	fmt.Printf("    Input 'input3' -> %s\n", configurableMock.Process("input3"))
	fmt.Printf("    Input 'unknown' -> %s\n", configurableMock.Process("unknown"))
}

func dynamicBehavior() {
	fmt.Println("  📊 Dynamic behavior:")
	
	// Create dynamic mock
	dynamicMock := &DynamicMock{}
	
	// Configure dynamic behavior
	dynamicMock.ConfigureBehavior(func(input string) string {
		if len(input) > 5 {
			return "long input"
		}
		return "short input"
	})
	
	// Test dynamic behavior
	fmt.Printf("    Input 'test' -> %s\n", dynamicMock.Process("test"))
	fmt.Printf("    Input 'testing' -> %s\n", dynamicMock.Process("testing"))
}

type DynamicMock struct {
	behavior func(string) string
	mu       sync.Mutex
}

func (d *DynamicMock) ConfigureBehavior(behavior func(string) string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.behavior = behavior
}

func (d *DynamicMock) Process(input string) string {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	if d.behavior != nil {
		return d.behavior(input)
	}
	return "default"
}

func behaviorChaining() {
	fmt.Println("  📊 Behavior chaining:")
	
	// Create chainable mock
	chainableMock := NewChainableMock()
	
	// Configure chained behavior
	chainableMock.
		ConfigureBehavior("input1", "output1").
		ConfigureBehavior("input2", "output2").
		ConfigureError("input3", "error3")
	
	// Test chained behavior
	fmt.Printf("    Input 'input1' -> %s\n", chainableMock.Process("input1"))
	fmt.Printf("    Input 'input2' -> %s\n", chainableMock.Process("input2"))
	fmt.Printf("    Input 'input3' -> %s\n", chainableMock.Process("input3"))
}

type ChainableMock struct {
	behaviors map[string]string
	errors    map[string]string
	mu        sync.Mutex
}

func NewChainableMock() *ChainableMock {
	return &ChainableMock{
		behaviors: make(map[string]string),
		errors:    make(map[string]string),
	}
}

func (c *ChainableMock) ConfigureBehavior(input, output string) *ChainableMock {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.behaviors[input] = output
	return c
}

func (c *ChainableMock) ConfigureError(input, errorMsg string) *ChainableMock {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.errors[input] = errorMsg
	return c
}

func (c *ChainableMock) Process(input string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if errorMsg, exists := c.errors[input]; exists {
		return "Error: " + errorMsg
	}
	
	if output, exists := c.behaviors[input]; exists {
		return output
	}
	
	return "default"
}

// 8. Best Practices
func mockInterfacesBestPractices() {
	fmt.Println("8. Mock Interfaces Best Practices:")
	fmt.Println("Best practices for mock interfaces...")

	fmt.Println("  📝 Best Practice 1: Use interfaces for mocking")
	fmt.Println("    - Mock interfaces, not concrete types")
	fmt.Println("    - Define clear, focused interfaces")
	fmt.Println("    - Keep interfaces small and cohesive")
	
	fmt.Println("  📝 Best Practice 2: Create reusable mocks")
	fmt.Println("    - Create mock factories or builders")
	fmt.Println("    - Use configuration for different test scenarios")
	fmt.Println("    - Avoid duplicating mock code")
	
	fmt.Println("  📝 Best Practice 3: Verify mock interactions")
	fmt.Println("    - Verify that mocks were called correctly")
	fmt.Println("    - Check call counts and arguments")
	fmt.Println("    - Verify call order when important")
	
	fmt.Println("  📝 Best Practice 4: Use appropriate test doubles")
	fmt.Println("    - Use dummies for unused dependencies")
	fmt.Println("    - Use stubs for simple responses")
	fmt.Println("    - Use mocks for verification")
	fmt.Println("    - Use fakes for realistic behavior")
	
	fmt.Println("  📝 Best Practice 5: Keep tests focused")
	fmt.Println("    - Test one thing at a time")
	fmt.Println("    - Use minimal mocks necessary")
	fmt.Println("    - Avoid over-mocking")
	
	fmt.Println("  📝 Best Practice 6: Clean up after tests")
	fmt.Println("    - Reset mocks between tests")
	fmt.Println("    - Use test setup and teardown")
	fmt.Println("    - Ensure test isolation")
	
	fmt.Println("  📝 Best Practice 7: Document mock behavior")
	fmt.Println("    - Document expected mock behavior")
	fmt.Println("    - Use descriptive mock names")
	fmt.Println("    - Add comments for complex mock setups")
}
