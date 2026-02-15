package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC ABSTRACT FACTORY PATTERN
// =============================================================================

// Abstract Products
type Button interface {
	Render() string
	OnClick() string
}

type Checkbox interface {
	Render() string
	OnCheck() string
}

type TextField interface {
	Render() string
	OnInput() string
}

// Concrete Products for Windows
type WindowsButton struct{}

func (wb *WindowsButton) Render() string {
	return "Windows Button rendered with Windows styling"
}

func (wb *WindowsButton) OnClick() string {
	return "Windows Button clicked - Windows event handled"
}

type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Render() string {
	return "Windows Checkbox rendered with Windows styling"
}

func (wc *WindowsCheckbox) OnCheck() string {
	return "Windows Checkbox checked - Windows event handled"
}

type WindowsTextField struct{}

func (wtf *WindowsTextField) Render() string {
	return "Windows TextField rendered with Windows styling"
}

func (wtf *WindowsTextField) OnInput() string {
	return "Windows TextField input - Windows event handled"
}

// Concrete Products for Mac
type MacButton struct{}

func (mb *MacButton) Render() string {
	return "Mac Button rendered with Mac styling"
}

func (mb *MacButton) OnClick() string {
	return "Mac Button clicked - Mac event handled"
}

type MacCheckbox struct{}

func (mc *MacCheckbox) Render() string {
	return "Mac Checkbox rendered with Mac styling"
}

func (mc *MacCheckbox) OnCheck() string {
	return "Mac Checkbox checked - Mac event handled"
}

type MacTextField struct{}

func (mtf *MacTextField) Render() string {
	return "Mac TextField rendered with Mac styling"
}

func (mtf *MacTextField) OnInput() string {
	return "Mac TextField input - Mac event handled"
}

// Concrete Products for Linux
type LinuxButton struct{}

func (lb *LinuxButton) Render() string {
	return "Linux Button rendered with Linux styling"
}

func (lb *LinuxButton) OnClick() string {
	return "Linux Button clicked - Linux event handled"
}

type LinuxCheckbox struct{}

func (lc *LinuxCheckbox) Render() string {
	return "Linux Checkbox rendered with Linux styling"
}

func (lc *LinuxCheckbox) OnCheck() string {
	return "Linux Checkbox checked - Linux event handled"
}

type LinuxTextField struct{}

func (ltf *LinuxTextField) Render() string {
	return "Linux TextField rendered with Linux styling"
}

func (ltf *LinuxTextField) OnInput() string {
	return "Linux TextField input - Linux event handled"
}

// Abstract Factory
type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
	CreateTextField() TextField
	GetOS() string
}

// Concrete Factories
type WindowsUIFactory struct{}

func (wf *WindowsUIFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsUIFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

func (wf *WindowsUIFactory) CreateTextField() TextField {
	return &WindowsTextField{}
}

func (wf *WindowsUIFactory) GetOS() string {
	return "Windows"
}

type MacUIFactory struct{}

func (mf *MacUIFactory) CreateButton() Button {
	return &MacButton{}
}

func (mf *MacUIFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

func (mf *MacUIFactory) CreateTextField() TextField {
	return &MacTextField{}
}

func (mf *MacUIFactory) GetOS() string {
	return "Mac"
}

type LinuxUIFactory struct{}

func (lf *LinuxUIFactory) CreateButton() Button {
	return &LinuxButton{}
}

func (lf *LinuxUIFactory) CreateCheckbox() Checkbox {
	return &LinuxCheckbox{}
}

func (lf *LinuxUIFactory) CreateTextField() TextField {
	return &LinuxTextField{}
}

func (lf *LinuxUIFactory) GetOS() string {
	return "Linux"
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. DATABASE ABSTRACT FACTORY
type DatabaseConnection interface {
	Connect() error
	Disconnect() error
	Query(sql string) ([]map[string]interface{}, error)
	GetConnectionType() string
}

type DatabaseTransaction interface {
	Begin() error
	Commit() error
	Rollback() error
	GetTransactionType() string
}

type DatabaseQueryBuilder interface {
	Select(columns ...string) DatabaseQueryBuilder
	From(table string) DatabaseQueryBuilder
	Where(condition string) DatabaseQueryBuilder
	Build() string
	GetBuilderType() string
}

// MySQL Products
type MySQLConnection struct {
	host     string
	port     int
	username string
	password string
}

func (mc *MySQLConnection) Connect() error {
	fmt.Printf("Connecting to MySQL at %s:%d\n", mc.host, mc.port)
	return nil
}

func (mc *MySQLConnection) Disconnect() error {
	fmt.Println("Disconnecting from MySQL")
	return nil
}

func (mc *MySQLConnection) Query(sql string) ([]map[string]interface{}, error) {
	fmt.Printf("Executing MySQL query: %s\n", sql)
	return []map[string]interface{}{{"result": "data"}}, nil
}

func (mc *MySQLConnection) GetConnectionType() string {
	return "MySQL"
}

type MySQLTransaction struct{}

func (mt *MySQLTransaction) Begin() error {
	fmt.Println("Starting MySQL transaction")
	return nil
}

func (mt *MySQLTransaction) Commit() error {
	fmt.Println("Committing MySQL transaction")
	return nil
}

func (mt *MySQLTransaction) Rollback() error {
	fmt.Println("Rolling back MySQL transaction")
	return nil
}

func (mt *MySQLTransaction) GetTransactionType() string {
	return "MySQL"
}

type MySQLQueryBuilder struct {
	query string
}

func (mqb *MySQLQueryBuilder) Select(columns ...string) DatabaseQueryBuilder {
	mqb.query = "SELECT " + fmt.Sprintf("%v", columns)
	return mqb
}

func (mqb *MySQLQueryBuilder) From(table string) DatabaseQueryBuilder {
	mqb.query += " FROM " + table
	return mqb
}

func (mqb *MySQLQueryBuilder) Where(condition string) DatabaseQueryBuilder {
	mqb.query += " WHERE " + condition
	return mqb
}

func (mqb *MySQLQueryBuilder) Build() string {
	return mqb.query
}

func (mqb *MySQLQueryBuilder) GetBuilderType() string {
	return "MySQL"
}

// PostgreSQL Products
type PostgreSQLConnection struct {
	host     string
	port     int
	username string
	password string
}

func (pc *PostgreSQLConnection) Connect() error {
	fmt.Printf("Connecting to PostgreSQL at %s:%d\n", pc.host, pc.port)
	return nil
}

func (pc *PostgreSQLConnection) Disconnect() error {
	fmt.Println("Disconnecting from PostgreSQL")
	return nil
}

func (pc *PostgreSQLConnection) Query(sql string) ([]map[string]interface{}, error) {
	fmt.Printf("Executing PostgreSQL query: %s\n", sql)
	return []map[string]interface{}{{"result": "data"}}, nil
}

func (pc *PostgreSQLConnection) GetConnectionType() string {
	return "PostgreSQL"
}

type PostgreSQLTransaction struct{}

func (pt *PostgreSQLTransaction) Begin() error {
	fmt.Println("Starting PostgreSQL transaction")
	return nil
}

func (pt *PostgreSQLTransaction) Commit() error {
	fmt.Println("Committing PostgreSQL transaction")
	return nil
}

func (pt *PostgreSQLTransaction) Rollback() error {
	fmt.Println("Rolling back PostgreSQL transaction")
	return nil
}

func (pt *PostgreSQLTransaction) GetTransactionType() string {
	return "PostgreSQL"
}

type PostgreSQLQueryBuilder struct {
	query string
}

func (pqb *PostgreSQLQueryBuilder) Select(columns ...string) DatabaseQueryBuilder {
	pqb.query = "SELECT " + fmt.Sprintf("%v", columns)
	return pqb
}

func (pqb *PostgreSQLQueryBuilder) From(table string) DatabaseQueryBuilder {
	pqb.query += " FROM " + table
	return pqb
}

func (pqb *PostgreSQLQueryBuilder) Where(condition string) DatabaseQueryBuilder {
	pqb.query += " WHERE " + condition
	return pqb
}

func (pqb *PostgreSQLQueryBuilder) Build() string {
	return pqb.query
}

func (pqb *PostgreSQLQueryBuilder) GetBuilderType() string {
	return "PostgreSQL"
}

// Database Abstract Factory
type DatabaseFactory interface {
	CreateConnection(host string, port int, username, password string) DatabaseConnection
	CreateTransaction() DatabaseTransaction
	CreateQueryBuilder() DatabaseQueryBuilder
	GetDatabaseType() string
}

// Concrete Database Factories
type MySQLDatabaseFactory struct{}

func (mdf *MySQLDatabaseFactory) CreateConnection(host string, port int, username, password string) DatabaseConnection {
	return &MySQLConnection{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (mdf *MySQLDatabaseFactory) CreateTransaction() DatabaseTransaction {
	return &MySQLTransaction{}
}

func (mdf *MySQLDatabaseFactory) CreateQueryBuilder() DatabaseQueryBuilder {
	return &MySQLQueryBuilder{}
}

func (mdf *MySQLDatabaseFactory) GetDatabaseType() string {
	return "MySQL"
}

type PostgreSQLDatabaseFactory struct{}

func (pdf *PostgreSQLDatabaseFactory) CreateConnection(host string, port int, username, password string) DatabaseConnection {
	return &PostgreSQLConnection{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (pdf *PostgreSQLDatabaseFactory) CreateTransaction() DatabaseTransaction {
	return &PostgreSQLTransaction{}
}

func (pdf *PostgreSQLDatabaseFactory) CreateQueryBuilder() DatabaseQueryBuilder {
	return &PostgreSQLQueryBuilder{}
}

func (pdf *PostgreSQLDatabaseFactory) GetDatabaseType() string {
	return "PostgreSQL"
}

// 2. PAYMENT PROCESSING ABSTRACT FACTORY
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string) error
	GetProcessorType() string
}

type PaymentValidator interface {
	ValidateCard(cardNumber string) bool
	ValidateAmount(amount float64) bool
	GetValidatorType() string
}

type PaymentLogger interface {
	LogTransaction(transactionID string, amount float64) error
	LogRefund(transactionID string, amount float64) error
	GetLoggerType() string
}

// Credit Card Products
type CreditCardProcessor struct {
	apiKey string
}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}

func (ccp *CreditCardProcessor) RefundPayment(transactionID string) error {
	fmt.Printf("Refunding credit card payment: %s\n", transactionID)
	return nil
}

func (ccp *CreditCardProcessor) GetProcessorType() string {
	return "Credit Card"
}

type CreditCardValidator struct{}

func (ccv *CreditCardValidator) ValidateCard(cardNumber string) bool {
	fmt.Printf("Validating credit card: %s\n", cardNumber)
	return len(cardNumber) == 16
}

func (ccv *CreditCardValidator) ValidateAmount(amount float64) bool {
	fmt.Printf("Validating credit card amount: $%.2f\n", amount)
	return amount > 0 && amount <= 10000
}

func (ccv *CreditCardValidator) GetValidatorType() string {
	return "Credit Card"
}

type CreditCardLogger struct{}

func (ccl *CreditCardLogger) LogTransaction(transactionID string, amount float64) error {
	fmt.Printf("Logging credit card transaction: %s, $%.2f\n", transactionID, amount)
	return nil
}

func (ccl *CreditCardLogger) LogRefund(transactionID string, amount float64) error {
	fmt.Printf("Logging credit card refund: %s, $%.2f\n", transactionID, amount)
	return nil
}

func (ccl *CreditCardLogger) GetLoggerType() string {
	return "Credit Card"
}

// PayPal Products
type PayPalProcessor struct {
	clientID string
}

func (ppp *PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}

func (ppp *PayPalProcessor) RefundPayment(transactionID string) error {
	fmt.Printf("Refunding PayPal payment: %s\n", transactionID)
	return nil
}

func (ppp *PayPalProcessor) GetProcessorType() string {
	return "PayPal"
}

type PayPalValidator struct{}

func (ppv *PayPalValidator) ValidateCard(cardNumber string) bool {
	fmt.Printf("Validating PayPal account: %s\n", cardNumber)
	return len(cardNumber) > 0
}

func (ppv *PayPalValidator) ValidateAmount(amount float64) bool {
	fmt.Printf("Validating PayPal amount: $%.2f\n", amount)
	return amount > 0 && amount <= 50000
}

func (ppv *PayPalValidator) GetValidatorType() string {
	return "PayPal"
}

type PayPalLogger struct{}

func (ppl *PayPalLogger) LogTransaction(transactionID string, amount float64) error {
	fmt.Printf("Logging PayPal transaction: %s, $%.2f\n", transactionID, amount)
	return nil
}

func (ppl *PayPalLogger) LogRefund(transactionID string, amount float64) error {
	fmt.Printf("Logging PayPal refund: %s, $%.2f\n", transactionID, amount)
	return nil
}

func (ppl *PayPalLogger) GetLoggerType() string {
	return "PayPal"
}

// Payment Abstract Factory
type PaymentFactory interface {
	CreateProcessor(credentials map[string]string) PaymentProcessor
	CreateValidator() PaymentValidator
	CreateLogger() PaymentLogger
	GetPaymentType() string
}

// Concrete Payment Factories
type CreditCardPaymentFactory struct{}

func (ccpf *CreditCardPaymentFactory) CreateProcessor(credentials map[string]string) PaymentProcessor {
	return &CreditCardProcessor{
		apiKey: credentials["api_key"],
	}
}

func (ccpf *CreditCardPaymentFactory) CreateValidator() PaymentValidator {
	return &CreditCardValidator{}
}

func (ccpf *CreditCardPaymentFactory) CreateLogger() PaymentLogger {
	return &CreditCardLogger{}
}

func (ccpf *CreditCardPaymentFactory) GetPaymentType() string {
	return "Credit Card"
}

type PayPalPaymentFactory struct{}

func (pppf *PayPalPaymentFactory) CreateProcessor(credentials map[string]string) PaymentProcessor {
	return &PayPalProcessor{
		clientID: credentials["client_id"],
	}
}

func (pppf *PayPalPaymentFactory) CreateValidator() PaymentValidator {
	return &PayPalValidator{}
}

func (pppf *PayPalPaymentFactory) CreateLogger() PaymentLogger {
	return &PayPalLogger{}
}

func (pppf *PayPalPaymentFactory) GetPaymentType() string {
	return "PayPal"
}

// =============================================================================
// FACTORY REGISTRY
// =============================================================================

type FactoryRegistry struct {
	factories map[string]interface{}
}

func NewFactoryRegistry() *FactoryRegistry {
	return &FactoryRegistry{
		factories: make(map[string]interface{}),
	}
}

func (fr *FactoryRegistry) Register(name string, factory interface{}) {
	fr.factories[name] = factory
}

func (fr *FactoryRegistry) GetUIFactory(name string) UIFactory {
	if factory, exists := fr.factories[name]; exists {
		return factory.(UIFactory)
	}
	return nil
}

func (fr *FactoryRegistry) GetDatabaseFactory(name string) DatabaseFactory {
	if factory, exists := fr.factories[name]; exists {
		return factory.(DatabaseFactory)
	}
	return nil
}

func (fr *FactoryRegistry) GetPaymentFactory(name string) PaymentFactory {
	if factory, exists := fr.factories[name]; exists {
		return factory.(PaymentFactory)
	}
	return nil
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== ABSTRACT FACTORY PATTERN DEMONSTRATION ===\n")

	// 1. BASIC ABSTRACT FACTORY
	fmt.Println("1. BASIC ABSTRACT FACTORY:")
	
	uiFactories := []UIFactory{
		&WindowsUIFactory{},
		&MacUIFactory{},
		&LinuxUIFactory{},
	}
	
	for _, factory := range uiFactories {
		fmt.Printf("\n%s UI Components:\n", factory.GetOS())
		
		button := factory.CreateButton()
		checkbox := factory.CreateCheckbox()
		textField := factory.CreateTextField()
		
		fmt.Printf("  Button: %s\n", button.Render())
		fmt.Printf("  Button Click: %s\n", button.OnClick())
		fmt.Printf("  Checkbox: %s\n", checkbox.Render())
		fmt.Printf("  Checkbox Check: %s\n", checkbox.OnCheck())
		fmt.Printf("  TextField: %s\n", textField.Render())
		fmt.Printf("  TextField Input: %s\n", textField.OnInput())
	}
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Database Factory
	fmt.Println("Database Factory:")
	dbFactories := []DatabaseFactory{
		&MySQLDatabaseFactory{},
		&PostgreSQLDatabaseFactory{},
	}
	
	for _, factory := range dbFactories {
		fmt.Printf("\n%s Database:\n", factory.GetDatabaseType())
		
		connection := factory.CreateConnection("localhost", 5432, "user", "pass")
		transaction := factory.CreateTransaction()
		queryBuilder := factory.CreateQueryBuilder()
		
		connection.Connect()
		transaction.Begin()
		query := queryBuilder.Select("id", "name").From("users").Where("active = true").Build()
		connection.Query(query)
		transaction.Commit()
		connection.Disconnect()
	}
	fmt.Println()

	// Payment Factory
	fmt.Println("Payment Factory:")
	paymentFactories := []PaymentFactory{
		&CreditCardPaymentFactory{},
		&PayPalPaymentFactory{},
	}
	
	for _, factory := range paymentFactories {
		fmt.Printf("\n%s Payment:\n", factory.GetPaymentType())
		
		processor := factory.CreateProcessor(map[string]string{"api_key": "test_key"})
		validator := factory.CreateValidator()
		logger := factory.CreateLogger()
		
		amount := 100.0
		transactionID := "txn_123"
		
		if validator.ValidateAmount(amount) {
			processor.ProcessPayment(amount)
			logger.LogTransaction(transactionID, amount)
		}
		
		processor.RefundPayment(transactionID)
		logger.LogRefund(transactionID, amount)
	}
	fmt.Println()

	// 3. FACTORY REGISTRY
	fmt.Println("3. FACTORY REGISTRY:")
	
	registry := NewFactoryRegistry()
	registry.Register("windows_ui", &WindowsUIFactory{})
	registry.Register("mac_ui", &MacUIFactory{})
	registry.Register("mysql_db", &MySQLDatabaseFactory{})
	registry.Register("postgresql_db", &PostgreSQLDatabaseFactory{})
	registry.Register("creditcard_payment", &CreditCardPaymentFactory{})
	registry.Register("paypal_payment", &PayPalPaymentFactory{})
	
	// Use registry to get factories
	windowsFactory := registry.GetUIFactory("windows_ui")
	if windowsFactory != nil {
		button := windowsFactory.CreateButton()
		fmt.Printf("Registry Windows Button: %s\n", button.Render())
	}
	
	mysqlFactory := registry.GetDatabaseFactory("mysql_db")
	if mysqlFactory != nil {
		connection := mysqlFactory.CreateConnection("localhost", 3306, "user", "pass")
		fmt.Printf("Registry MySQL Connection: %s\n", connection.GetConnectionType())
	}
	
	creditCardFactory := registry.GetPaymentFactory("creditcard_payment")
	if creditCardFactory != nil {
		processor := creditCardFactory.CreateProcessor(map[string]string{"api_key": "test_key"})
		fmt.Printf("Registry Credit Card Processor: %s\n", processor.GetProcessorType())
	}
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
