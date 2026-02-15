package main

import (
	"fmt"
	"time"
)

// =============================================================================
// SIMPLE FACTORY PATTERN
// =============================================================================

// Product interface
type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
}

// Concrete Products
type Laptop struct {
	name        string
	price       float64
	description string
}

func (l *Laptop) GetName() string {
	return l.name
}

func (l *Laptop) GetPrice() float64 {
	return l.price
}

func (l *Laptop) GetDescription() string {
	return l.description
}

type Smartphone struct {
	name        string
	price       float64
	description string
}

func (s *Smartphone) GetName() string {
	return s.name
}

func (s *Smartphone) GetPrice() float64 {
	return s.price
}

func (s *Smartphone) GetDescription() string {
	return s.description
}

type Tablet struct {
	name        string
	price       float64
	description string
}

func (t *Tablet) GetName() string {
	return t.name
}

func (t *Tablet) GetPrice() float64 {
	return t.price
}

func (t *Tablet) GetDescription() string {
	return t.description
}

// Simple Factory
type ProductFactory struct{}

func (pf *ProductFactory) CreateProduct(productType string) Product {
	switch productType {
	case "laptop":
		return &Laptop{
			name:        "Gaming Laptop",
			price:       1299.99,
			description: "High-performance gaming laptop with RTX 4080",
		}
	case "smartphone":
		return &Smartphone{
			name:        "Flagship Phone",
			price:       999.99,
			description: "Latest smartphone with advanced camera system",
		}
	case "tablet":
		return &Tablet{
			name:        "Pro Tablet",
			price:       799.99,
			description: "Professional tablet for creative work",
		}
	default:
		return nil
	}
}

// =============================================================================
// FACTORY METHOD PATTERN
// =============================================================================

// Factory interface
type Factory interface {
	CreateProduct() Product
	GetFactoryName() string
}

// Concrete Factories
type LaptopFactory struct{}

func (lf *LaptopFactory) CreateProduct() Product {
	return &Laptop{
		name:        "Business Laptop",
		price:       899.99,
		description: "Reliable business laptop for productivity",
	}
}

func (lf *LaptopFactory) GetFactoryName() string {
	return "Laptop Factory"
}

type SmartphoneFactory struct{}

func (sf *SmartphoneFactory) CreateProduct() Product {
	return &Smartphone{
		name:        "Budget Phone",
		price:       299.99,
		description: "Affordable smartphone with essential features",
	}
}

func (sf *SmartphoneFactory) GetFactoryName() string {
	return "Smartphone Factory"
}

type TabletFactory struct{}

func (tf *TabletFactory) CreateProduct() Product {
	return &Tablet{
		name:        "Student Tablet",
		price:       399.99,
		description: "Educational tablet for students",
	}
}

func (tf *TabletFactory) GetFactoryName() string {
	return "Tablet Factory"
}

// =============================================================================
// ABSTRACT FACTORY PATTERN
// =============================================================================

// Abstract Product A
type Button interface {
	Render() string
	OnClick() string
}

// Abstract Product B
type Checkbox interface {
	Render() string
	OnCheck() string
}

// Concrete Products for Windows
type WindowsButton struct{}

func (wb *WindowsButton) Render() string {
	return "Windows Button rendered"
}

func (wb *WindowsButton) OnClick() string {
	return "Windows Button clicked"
}

type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Render() string {
	return "Windows Checkbox rendered"
}

func (wc *WindowsCheckbox) OnCheck() string {
	return "Windows Checkbox checked"
}

// Concrete Products for Mac
type MacButton struct{}

func (mb *MacButton) Render() string {
	return "Mac Button rendered"
}

func (mb *MacButton) OnClick() string {
	return "Mac Button clicked"
}

type MacCheckbox struct{}

func (mc *MacCheckbox) Render() string {
	return "Mac Checkbox rendered"
}

func (mc *MacCheckbox) OnCheck() string {
	return "Mac Checkbox checked"
}

// Abstract Factory
type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
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

func (mf *MacUIFactory) GetOS() string {
	return "Mac"
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. DATABASE CONNECTION FACTORY
type DatabaseConnection interface {
	Connect() error
	Disconnect() error
	Query(sql string) ([]map[string]interface{}, error)
	GetConnectionType() string
}

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

type DatabaseFactory struct{}

func (df *DatabaseFactory) CreateConnection(dbType string, host string, port int, username, password string) DatabaseConnection {
	switch dbType {
	case "mysql":
		return &MySQLConnection{
			host:     host,
			port:     port,
			username: username,
			password: password,
		}
	case "postgresql":
		return &PostgreSQLConnection{
			host:     host,
			port:     port,
			username: username,
			password: password,
		}
	default:
		return nil
	}
}

// 2. PAYMENT PROCESSOR FACTORY
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string) error
	GetProcessorName() string
}

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

func (ccp *CreditCardProcessor) GetProcessorName() string {
	return "Credit Card"
}

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

func (ppp *PayPalProcessor) GetProcessorName() string {
	return "PayPal"
}

type PaymentFactory struct{}

func (pf *PaymentFactory) CreateProcessor(processorType string, credentials map[string]string) PaymentProcessor {
	switch processorType {
	case "creditcard":
		return &CreditCardProcessor{
			apiKey: credentials["api_key"],
		}
	case "paypal":
		return &PayPalProcessor{
			clientID: credentials["client_id"],
		}
	default:
		return nil
	}
}

// 3. LOGGER FACTORY
type Logger interface {
	Log(level string, message string)
	GetLoggerType() string
}

type FileLogger struct {
	filePath string
}

func (fl *FileLogger) Log(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[FILE] %s [%s] %s\n", timestamp, level, message)
}

func (fl *FileLogger) GetLoggerType() string {
	return "File Logger"
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[CONSOLE] %s [%s] %s\n", timestamp, level, message)
}

func (cl *ConsoleLogger) GetLoggerType() string {
	return "Console Logger"
}

type DatabaseLogger struct {
	connectionString string
}

func (dl *DatabaseLogger) Log(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[DATABASE] %s [%s] %s\n", timestamp, level, message)
}

func (dl *DatabaseLogger) GetLoggerType() string {
	return "Database Logger"
}

type LoggerFactory struct{}

func (lf *LoggerFactory) CreateLogger(loggerType string, config map[string]string) Logger {
	switch loggerType {
	case "file":
		return &FileLogger{
			filePath: config["file_path"],
		}
	case "console":
		return &ConsoleLogger{}
	case "database":
		return &DatabaseLogger{
			connectionString: config["connection_string"],
		}
	default:
		return &ConsoleLogger{} // Default to console logger
	}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== FACTORY PATTERN DEMONSTRATION ===\n")

	// 1. SIMPLE FACTORY
	fmt.Println("1. SIMPLE FACTORY:")
	productFactory := &ProductFactory{}
	
	products := []string{"laptop", "smartphone", "tablet"}
	for _, productType := range products {
		product := productFactory.CreateProduct(productType)
		if product != nil {
			fmt.Printf("Created %s: %s - $%.2f\n", productType, product.GetName(), product.GetPrice())
		}
	}
	fmt.Println()

	// 2. FACTORY METHOD
	fmt.Println("2. FACTORY METHOD:")
	factories := []Factory{
		&LaptopFactory{},
		&SmartphoneFactory{},
		&TabletFactory{},
	}
	
	for _, factory := range factories {
		product := factory.CreateProduct()
		fmt.Printf("%s created: %s - $%.2f\n", factory.GetFactoryName(), product.GetName(), product.GetPrice())
	}
	fmt.Println()

	// 3. ABSTRACT FACTORY
	fmt.Println("3. ABSTRACT FACTORY:")
	uiFactories := []UIFactory{
		&WindowsUIFactory{},
		&MacUIFactory{},
	}
	
	for _, factory := range uiFactories {
		fmt.Printf("\n%s UI Components:\n", factory.GetOS())
		button := factory.CreateButton()
		checkbox := factory.CreateCheckbox()
		
		fmt.Printf("  Button: %s\n", button.Render())
		fmt.Printf("  Button Click: %s\n", button.OnClick())
		fmt.Printf("  Checkbox: %s\n", checkbox.Render())
		fmt.Printf("  Checkbox Check: %s\n", checkbox.OnCheck())
	}
	fmt.Println()

	// 4. REAL-WORLD EXAMPLES
	fmt.Println("4. REAL-WORLD EXAMPLES:")

	// Database Factory
	fmt.Println("Database Factory:")
	dbFactory := &DatabaseFactory{}
	mysqlConn := dbFactory.CreateConnection("mysql", "localhost", 3306, "user", "pass")
	postgresConn := dbFactory.CreateConnection("postgresql", "localhost", 5432, "user", "pass")
	
	mysqlConn.Connect()
	mysqlConn.Query("SELECT * FROM users")
	mysqlConn.Disconnect()
	
	postgresConn.Connect()
	postgresConn.Query("SELECT * FROM products")
	postgresConn.Disconnect()
	fmt.Println()

	// Payment Factory
	fmt.Println("Payment Factory:")
	paymentFactory := &PaymentFactory{}
	ccProcessor := paymentFactory.CreateProcessor("creditcard", map[string]string{"api_key": "cc_api_key"})
	paypalProcessor := paymentFactory.CreateProcessor("paypal", map[string]string{"client_id": "pp_client_id"})
	
	ccProcessor.ProcessPayment(100.0)
	ccProcessor.RefundPayment("txn_123")
	
	paypalProcessor.ProcessPayment(50.0)
	paypalProcessor.RefundPayment("txn_456")
	fmt.Println()

	// Logger Factory
	fmt.Println("Logger Factory:")
	loggerFactory := &LoggerFactory{}
	fileLogger := loggerFactory.CreateLogger("file", map[string]string{"file_path": "/var/log/app.log"})
	consoleLogger := loggerFactory.CreateLogger("console", nil)
	dbLogger := loggerFactory.CreateLogger("database", map[string]string{"connection_string": "db://localhost"})
	
	fileLogger.Log("INFO", "Application started")
	consoleLogger.Log("WARN", "Low memory warning")
	dbLogger.Log("ERROR", "Database connection failed")
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
