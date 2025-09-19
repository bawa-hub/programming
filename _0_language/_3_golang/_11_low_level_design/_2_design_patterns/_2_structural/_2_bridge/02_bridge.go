package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC BRIDGE PATTERN
// =============================================================================

// Implementor interface - defines the interface for implementation classes
type DrawingAPI interface {
	DrawCircle(x, y, radius float64)
	DrawRectangle(x, y, width, height float64)
	GetAPI() string
}

// Concrete Implementor A - Windows drawing API
type WindowsDrawingAPI struct{}

func (wda *WindowsDrawingAPI) DrawCircle(x, y, radius float64) {
	fmt.Printf("Windows API: Drawing circle at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (wda *WindowsDrawingAPI) DrawRectangle(x, y, width, height float64) {
	fmt.Printf("Windows API: Drawing rectangle at (%.2f, %.2f) with size %.2fx%.2f\n", x, y, width, height)
}

func (wda *WindowsDrawingAPI) GetAPI() string {
	return "Windows Drawing API"
}

// Concrete Implementor B - Mac drawing API
type MacDrawingAPI struct{}

func (mda *MacDrawingAPI) DrawCircle(x, y, radius float64) {
	fmt.Printf("Mac API: Drawing circle at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (mda *MacDrawingAPI) DrawRectangle(x, y, width, height float64) {
	fmt.Printf("Mac API: Drawing rectangle at (%.2f, %.2f) with size %.2fx%.2f\n", x, y, width, height)
}

func (mda *MacDrawingAPI) GetAPI() string {
	return "Mac Drawing API"
}

// Concrete Implementor C - Linux drawing API
type LinuxDrawingAPI struct{}

func (lda *LinuxDrawingAPI) DrawCircle(x, y, radius float64) {
	fmt.Printf("Linux API: Drawing circle at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (lda *LinuxDrawingAPI) DrawRectangle(x, y, width, height float64) {
	fmt.Printf("Linux API: Drawing rectangle at (%.2f, %.2f) with size %.2fx%.2f\n", x, y, width, height)
}

func (lda *LinuxDrawingAPI) GetAPI() string {
	return "Linux Drawing API"
}

// Abstraction - defines the interface for the control part
type Shape interface {
	Draw()
	Resize(factor float64)
	GetShapeType() string
}

// Refined Abstraction - Circle
type Circle struct {
	x, y, radius float64
	drawingAPI   DrawingAPI
}

func NewCircle(x, y, radius float64, drawingAPI DrawingAPI) *Circle {
	return &Circle{
		x:          x,
		y:          y,
		radius:     radius,
		drawingAPI: drawingAPI,
	}
}

func (c *Circle) Draw() {
	c.drawingAPI.DrawCircle(c.x, c.y, c.radius)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
	fmt.Printf("Circle resized by factor %.2f, new radius: %.2f\n", factor, c.radius)
}

func (c *Circle) GetShapeType() string {
	return "Circle"
}

// Refined Abstraction - Rectangle
type Rectangle struct {
	x, y, width, height float64
	drawingAPI          DrawingAPI
}

func NewRectangle(x, y, width, height float64, drawingAPI DrawingAPI) *Rectangle {
	return &Rectangle{
		x:          x,
		y:          y,
		width:      width,
		height:     height,
		drawingAPI: drawingAPI,
	}
}

func (r *Rectangle) Draw() {
	r.drawingAPI.DrawRectangle(r.x, r.y, r.width, r.height)
}

func (r *Rectangle) Resize(factor float64) {
	r.width *= factor
	r.height *= factor
	fmt.Printf("Rectangle resized by factor %.2f, new size: %.2fx%.2f\n", factor, r.width, r.height)
}

func (r *Rectangle) GetShapeType() string {
	return "Rectangle"
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. DATABASE BRIDGE PATTERN
type DatabaseConnection interface {
	Connect() error
	Disconnect() error
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
	GetConnectionType() string
}

// MySQL implementation
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
	fmt.Printf("MySQL executing query: %s\n", sql)
	return []map[string]interface{}{{"id": 1, "name": "John"}}, nil
}

func (mc *MySQLConnection) Execute(sql string) error {
	fmt.Printf("MySQL executing command: %s\n", sql)
	return nil
}

func (mc *MySQLConnection) GetConnectionType() string {
	return "MySQL"
}

// PostgreSQL implementation
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
	fmt.Printf("PostgreSQL executing query: %s\n", sql)
	return []map[string]interface{}{{"id": 1, "name": "John"}}, nil
}

func (pc *PostgreSQLConnection) Execute(sql string) error {
	fmt.Printf("PostgreSQL executing command: %s\n", sql)
	return nil
}

func (pc *PostgreSQLConnection) GetConnectionType() string {
	return "PostgreSQL"
}

// MongoDB implementation
type MongoDBConnection struct {
	host     string
	port     int
	username string
	password string
}

func (mc *MongoDBConnection) Connect() error {
	fmt.Printf("Connecting to MongoDB at %s:%d\n", mc.host, mc.port)
	return nil
}

func (mc *MongoDBConnection) Disconnect() error {
	fmt.Println("Disconnecting from MongoDB")
	return nil
}

func (mc *MongoDBConnection) Query(sql string) ([]map[string]interface{}, error) {
	fmt.Printf("MongoDB executing query: %s\n", sql)
	return []map[string]interface{}{{"_id": "507f1f77bcf86cd799439011", "name": "John"}}, nil
}

func (mc *MongoDBConnection) Execute(sql string) error {
	fmt.Printf("MongoDB executing command: %s\n", sql)
	return nil
}

func (mc *MongoDBConnection) GetConnectionType() string {
	return "MongoDB"
}

// Database abstraction
type Database struct {
	connection DatabaseConnection
}

func NewDatabase(connection DatabaseConnection) *Database {
	return &Database{connection: connection}
}

func (db *Database) Connect() error {
	return db.connection.Connect()
}

func (db *Database) Disconnect() error {
	return db.connection.Disconnect()
}

func (db *Database) Query(sql string) ([]map[string]interface{}, error) {
	return db.connection.Query(sql)
}

func (db *Database) Execute(sql string) error {
	return db.connection.Execute(sql)
}

func (db *Database) GetConnectionType() string {
	return db.connection.GetConnectionType()
}

// 2. NOTIFICATION BRIDGE PATTERN
type NotificationChannel interface {
	Send(message string, recipient string) error
	GetChannelType() string
}

// Email notification
type EmailNotification struct {
	smtpServer string
	port       int
	username   string
	password   string
}

func (en *EmailNotification) Send(message string, recipient string) error {
	fmt.Printf("Sending email to %s: %s\n", recipient, message)
	return nil
}

func (en *EmailNotification) GetChannelType() string {
	return "Email"
}

// SMS notification
type SMSNotification struct {
	apiKey    string
	apiSecret string
}

func (sn *SMSNotification) Send(message string, recipient string) error {
	fmt.Printf("Sending SMS to %s: %s\n", recipient, message)
	return nil
}

func (sn *SMSNotification) GetChannelType() string {
	return "SMS"
}

// Push notification
type PushNotification struct {
	serverKey string
}

func (pn *PushNotification) Send(message string, recipient string) error {
	fmt.Printf("Sending push notification to %s: %s\n", recipient, message)
	return nil
}

func (pn *PushNotification) GetChannelType() string {
	return "Push"
}

// Notification abstraction
type Notification struct {
	channel NotificationChannel
}

func NewNotification(channel NotificationChannel) *Notification {
	return &Notification{channel: channel}
}

func (n *Notification) Send(message string, recipient string) error {
	return n.channel.Send(message, recipient)
}

func (n *Notification) GetChannelType() string {
	return n.channel.GetChannelType()
}

// 3. PAYMENT PROCESSOR BRIDGE PATTERN
type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (string, error)
	RefundPayment(transactionID string) error
	GetProcessorType() string
}

// Stripe processor
type StripeProcessor struct {
	apiKey string
}

func (sp *StripeProcessor) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Stripe processing payment: $%.2f %s\n", amount, currency)
	return fmt.Sprintf("stripe_txn_%d", time.Now().Unix()), nil
}

func (sp *StripeProcessor) RefundPayment(transactionID string) error {
	fmt.Printf("Stripe refunding transaction: %s\n", transactionID)
	return nil
}

func (sp *StripeProcessor) GetProcessorType() string {
	return "Stripe"
}

// PayPal processor
type PayPalProcessor struct {
	clientID string
	secret   string
}

func (pp *PayPalProcessor) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("PayPal processing payment: $%.2f %s\n", amount, currency)
	return fmt.Sprintf("paypal_txn_%d", time.Now().Unix()), nil
}

func (pp *PayPalProcessor) RefundPayment(transactionID string) error {
	fmt.Printf("PayPal refunding transaction: %s\n", transactionID)
	return nil
}

func (pp *PayPalProcessor) GetProcessorType() string {
	return "PayPal"
}

// Square processor
type SquareProcessor struct {
	applicationID string
	accessToken   string
}

func (sp *SquareProcessor) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Square processing payment: $%.2f %s\n", amount, currency)
	return fmt.Sprintf("square_txn_%d", time.Now().Unix()), nil
}

func (sp *SquareProcessor) RefundPayment(transactionID string) error {
	fmt.Printf("Square refunding transaction: %s\n", transactionID)
	return nil
}

func (sp *SquareProcessor) GetProcessorType() string {
	return "Square"
}

// Payment abstraction
type Payment struct {
	processor PaymentProcessor
}

func NewPayment(processor PaymentProcessor) *Payment {
	return &Payment{processor: processor}
}

func (p *Payment) ProcessPayment(amount float64, currency string) (string, error) {
	return p.processor.ProcessPayment(amount, currency)
}

func (p *Payment) RefundPayment(transactionID string) error {
	return p.processor.RefundPayment(transactionID)
}

func (p *Payment) GetProcessorType() string {
	return p.processor.GetProcessorType()
}

// =============================================================================
// DYNAMIC BRIDGE EXAMPLE
// =============================================================================

// Dynamic bridge that can switch implementations at runtime
type DynamicShape struct {
	x, y, radius float64
	drawingAPI   DrawingAPI
}

func NewDynamicShape(x, y, radius float64, drawingAPI DrawingAPI) *DynamicShape {
	return &DynamicShape{
		x:          x,
		y:          y,
		radius:     radius,
		drawingAPI: drawingAPI,
	}
}

func (ds *DynamicShape) Draw() {
	ds.drawingAPI.DrawCircle(ds.x, ds.y, ds.radius)
}

func (ds *DynamicShape) SetDrawingAPI(drawingAPI DrawingAPI) {
	ds.drawingAPI = drawingAPI
	fmt.Printf("Switched to %s\n", drawingAPI.GetAPI())
}

func (ds *DynamicShape) GetCurrentAPI() string {
	return ds.drawingAPI.GetAPI()
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== BRIDGE PATTERN DEMONSTRATION ===\n")

	// 1. BASIC BRIDGE
	fmt.Println("1. BASIC BRIDGE:")
	
	// Create different drawing APIs
	windowsAPI := &WindowsDrawingAPI{}
	macAPI := &MacDrawingAPI{}
	linuxAPI := &LinuxDrawingAPI{}
	
	// Create shapes with different APIs
	circle1 := NewCircle(10, 10, 5, windowsAPI)
	circle2 := NewCircle(20, 20, 8, macAPI)
	rectangle1 := NewRectangle(5, 5, 10, 15, linuxAPI)
	
	// Draw shapes
	circle1.Draw()
	circle2.Draw()
	rectangle1.Draw()
	
	// Resize shapes
	circle1.Resize(1.5)
	rectangle1.Resize(0.8)
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Database Bridge
	fmt.Println("Database Bridge:")
	mysqlConn := &MySQLConnection{host: "localhost", port: 3306, username: "user", password: "pass"}
	postgresConn := &PostgreSQLConnection{host: "localhost", port: 5432, username: "user", password: "pass"}
	mongoConn := &MongoDBConnection{host: "localhost", port: 27017, username: "user", password: "pass"}
	
	mysqlDB := NewDatabase(mysqlConn)
	postgresDB := NewDatabase(postgresConn)
	mongoDB := NewDatabase(mongoConn)
	
	// Use different databases
	databases := []*Database{mysqlDB, postgresDB, mongoDB}
	for _, db := range databases {
		db.Connect()
		db.Query("SELECT * FROM users")
		db.Execute("INSERT INTO users (name) VALUES ('John')")
		db.Disconnect()
		fmt.Printf("Used %s database\n", db.GetConnectionType())
	}
	fmt.Println()

	// Notification Bridge
	fmt.Println("Notification Bridge:")
	emailChannel := &EmailNotification{smtpServer: "smtp.gmail.com", port: 587, username: "user", password: "pass"}
	smsChannel := &SMSNotification{apiKey: "sms_key", apiSecret: "sms_secret"}
	pushChannel := &PushNotification{serverKey: "push_key"}
	
	emailNotification := NewNotification(emailChannel)
	smsNotification := NewNotification(smsChannel)
	pushNotification := NewNotification(pushChannel)
	
	// Send notifications through different channels
	notifications := []*Notification{emailNotification, smsNotification, pushNotification}
	for _, notification := range notifications {
		notification.Send("Hello from notification system!", "user@example.com")
		fmt.Printf("Sent via %s\n", notification.GetChannelType())
	}
	fmt.Println()

	// Payment Bridge
	fmt.Println("Payment Bridge:")
	stripeProcessor := &StripeProcessor{apiKey: "stripe_key"}
	paypalProcessor := &PayPalProcessor{clientID: "paypal_client", secret: "paypal_secret"}
	squareProcessor := &SquareProcessor{applicationID: "square_app", accessToken: "square_token"}
	
	stripePayment := NewPayment(stripeProcessor)
	paypalPayment := NewPayment(paypalProcessor)
	squarePayment := NewPayment(squareProcessor)
	
	// Process payments through different processors
	payments := []*Payment{stripePayment, paypalPayment, squarePayment}
	for _, payment := range payments {
		transactionID, _ := payment.ProcessPayment(100.0, "USD")
		payment.RefundPayment(transactionID)
		fmt.Printf("Processed via %s\n", payment.GetProcessorType())
	}
	fmt.Println()

	// 3. DYNAMIC BRIDGE
	fmt.Println("3. DYNAMIC BRIDGE:")
	dynamicShape := NewDynamicShape(15, 15, 10, windowsAPI)
	
	fmt.Printf("Initial API: %s\n", dynamicShape.GetCurrentAPI())
	dynamicShape.Draw()
	
	// Switch to Mac API
	dynamicShape.SetDrawingAPI(macAPI)
	dynamicShape.Draw()
	
	// Switch to Linux API
	dynamicShape.SetDrawingAPI(linuxAPI)
	dynamicShape.Draw()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
