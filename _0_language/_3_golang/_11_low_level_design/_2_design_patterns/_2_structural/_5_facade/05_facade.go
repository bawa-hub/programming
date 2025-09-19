package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC FACADE PATTERN
// =============================================================================

// Subsystem classes
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU: Starting...")
}

func (c *CPU) Execute() {
	fmt.Println("CPU: Executing instructions...")
}

func (c *CPU) Stop() {
	fmt.Println("CPU: Stopping...")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory: Loading data...")
}

func (m *Memory) Store() {
	fmt.Println("Memory: Storing data...")
}

func (m *Memory) Clear() {
	fmt.Println("Memory: Clearing data...")
}

type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Hard Drive: Reading data...")
}

func (h *HardDrive) Write() {
	fmt.Println("Hard Drive: Writing data...")
}

func (h *HardDrive) Shutdown() {
	fmt.Println("Hard Drive: Shutting down...")
}

// Facade - provides a simplified interface to the subsystem
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) StartComputer() {
	fmt.Println("=== Starting Computer ===")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.Read()
	fmt.Println("Computer started successfully!")
}

func (cf *ComputerFacade) ShutdownComputer() {
	fmt.Println("=== Shutting Down Computer ===")
	cf.cpu.Stop()
	cf.memory.Clear()
	cf.hardDrive.Shutdown()
	fmt.Println("Computer shut down successfully!")
}

func (cf *ComputerFacade) ExecuteProgram() {
	fmt.Println("=== Executing Program ===")
	cf.cpu.Execute()
	cf.memory.Store()
	cf.hardDrive.Write()
	fmt.Println("Program executed successfully!")
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. HOME AUTOMATION FACADE
type Light struct {
	location string
	isOn     bool
}

func NewLight(location string) *Light {
	return &Light{location: location, isOn: false}
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Printf("Light in %s is turned on\n", l.location)
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Printf("Light in %s is turned off\n", l.location)
}

func (l *Light) IsOn() bool {
	return l.isOn
}

type TV struct {
	location string
	isOn     bool
	channel  int
}

func NewTV(location string) *TV {
	return &TV{location: location, isOn: false, channel: 1}
}

func (t *TV) TurnOn() {
	t.isOn = true
	fmt.Printf("TV in %s is turned on\n", t.location)
}

func (t *TV) TurnOff() {
	t.isOn = false
	fmt.Printf("TV in %s is turned off\n", t.location)
}

func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Printf("TV in %s is set to channel %d\n", t.location, channel)
}

func (t *TV) IsOn() bool {
	return t.isOn
}

type SoundSystem struct {
	location string
	isOn     bool
	volume   int
}

func NewSoundSystem(location string) *SoundSystem {
	return &SoundSystem{location: location, isOn: false, volume: 0}
}

func (s *SoundSystem) TurnOn() {
	s.isOn = true
	fmt.Printf("Sound system in %s is turned on\n", s.location)
}

func (s *SoundSystem) TurnOff() {
	s.isOn = false
	fmt.Printf("Sound system in %s is turned off\n", s.location)
}

func (s *SoundSystem) SetVolume(volume int) {
	s.volume = volume
	fmt.Printf("Sound system in %s volume set to %d\n", s.location, volume)
}

func (s *SoundSystem) IsOn() bool {
	return s.isOn
}

type AirConditioner struct {
	location string
	isOn     bool
	temperature int
}

func NewAirConditioner(location string) *AirConditioner {
	return &AirConditioner{location: location, isOn: false, temperature: 22}
}

func (ac *AirConditioner) TurnOn() {
	ac.isOn = true
	fmt.Printf("Air conditioner in %s is turned on\n", ac.location)
}

func (ac *AirConditioner) TurnOff() {
	ac.isOn = false
	fmt.Printf("Air conditioner in %s is turned off\n", ac.location)
}

func (ac *AirConditioner) SetTemperature(temp int) {
	ac.temperature = temp
	fmt.Printf("Air conditioner in %s temperature set to %d°C\n", ac.location, temp)
}

func (ac *AirConditioner) IsOn() bool {
	return ac.isOn
}

// Home Automation Facade
type HomeAutomationFacade struct {
	livingRoomLight *Light
	livingRoomTV    *TV
	livingRoomSound *SoundSystem
	livingRoomAC    *AirConditioner
	bedroomLight    *Light
	bedroomTV       *TV
	bedroomAC       *AirConditioner
}

func NewHomeAutomationFacade() *HomeAutomationFacade {
	return &HomeAutomationFacade{
		livingRoomLight: NewLight("Living Room"),
		livingRoomTV:    NewTV("Living Room"),
		livingRoomSound: NewSoundSystem("Living Room"),
		livingRoomAC:    NewAirConditioner("Living Room"),
		bedroomLight:    NewLight("Bedroom"),
		bedroomTV:       NewTV("Bedroom"),
		bedroomAC:       NewAirConditioner("Bedroom"),
	}
}

func (haf *HomeAutomationFacade) MovieNight() {
	fmt.Println("=== Starting Movie Night ===")
	haf.livingRoomLight.TurnOff()
	haf.livingRoomTV.TurnOn()
	haf.livingRoomTV.SetChannel(5)
	haf.livingRoomSound.TurnOn()
	haf.livingRoomSound.SetVolume(70)
	haf.livingRoomAC.TurnOn()
	haf.livingRoomAC.SetTemperature(20)
	fmt.Println("Movie night setup complete!")
}

func (haf *HomeAutomationFacade) Bedtime() {
	fmt.Println("=== Starting Bedtime Routine ===")
	haf.livingRoomLight.TurnOff()
	haf.livingRoomTV.TurnOff()
	haf.livingRoomSound.TurnOff()
	haf.livingRoomAC.TurnOff()
	haf.bedroomLight.TurnOn()
	haf.bedroomTV.TurnOff()
	haf.bedroomAC.TurnOn()
	haf.bedroomAC.SetTemperature(18)
	fmt.Println("Bedtime routine complete!")
}

func (haf *HomeAutomationFacade) WakeUp() {
	fmt.Println("=== Starting Wake Up Routine ===")
	haf.bedroomLight.TurnOn()
	haf.bedroomTV.TurnOn()
	haf.bedroomTV.SetChannel(1)
	haf.bedroomAC.TurnOff()
	haf.livingRoomLight.TurnOn()
	haf.livingRoomAC.TurnOn()
	haf.livingRoomAC.SetTemperature(22)
	fmt.Println("Wake up routine complete!")
}

func (haf *HomeAutomationFacade) TurnOffEverything() {
	fmt.Println("=== Turning Off Everything ===")
	haf.livingRoomLight.TurnOff()
	haf.livingRoomTV.TurnOff()
	haf.livingRoomSound.TurnOff()
	haf.livingRoomAC.TurnOff()
	haf.bedroomLight.TurnOff()
	haf.bedroomTV.TurnOff()
	haf.bedroomAC.TurnOff()
	fmt.Println("Everything turned off!")
}

// 2. DATABASE OPERATIONS FACADE
type DatabaseConnection struct {
	host     string
	port     int
	username string
	password string
	isConnected bool
}

func NewDatabaseConnection(host string, port int, username, password string) *DatabaseConnection {
	return &DatabaseConnection{
		host:     host,
		port:     port,
		username: username,
		password: password,
		isConnected: false,
	}
}

func (dc *DatabaseConnection) Connect() error {
	fmt.Printf("Connecting to database at %s:%d\n", dc.host, dc.port)
	dc.isConnected = true
	return nil
}

func (dc *DatabaseConnection) Disconnect() error {
	fmt.Println("Disconnecting from database")
	dc.isConnected = false
	return nil
}

func (dc *DatabaseConnection) IsConnected() bool {
	return dc.isConnected
}

type QueryExecutor struct {
	connection *DatabaseConnection
}

func NewQueryExecutor(connection *DatabaseConnection) *QueryExecutor {
	return &QueryExecutor{connection: connection}
}

func (qe *QueryExecutor) ExecuteQuery(sql string) ([]map[string]interface{}, error) {
	if !qe.connection.IsConnected() {
		return nil, fmt.Errorf("database not connected")
	}
	fmt.Printf("Executing query: %s\n", sql)
	return []map[string]interface{}{{"id": 1, "name": "John"}}, nil
}

func (qe *QueryExecutor) ExecuteCommand(sql string) error {
	if !qe.connection.IsConnected() {
		return fmt.Errorf("database not connected")
	}
	fmt.Printf("Executing command: %s\n", sql)
	return nil
}

type TransactionManager struct {
	connection *DatabaseConnection
}

func NewTransactionManager(connection *DatabaseConnection) *TransactionManager {
	return &TransactionManager{connection: connection}
}

func (tm *TransactionManager) Begin() error {
	if !tm.connection.IsConnected() {
		return fmt.Errorf("database not connected")
	}
	fmt.Println("Beginning transaction")
	return nil
}

func (tm *TransactionManager) Commit() error {
	if !tm.connection.IsConnected() {
		return fmt.Errorf("database not connected")
	}
	fmt.Println("Committing transaction")
	return nil
}

func (tm *TransactionManager) Rollback() error {
	if !tm.connection.IsConnected() {
		return fmt.Errorf("database not connected")
	}
	fmt.Println("Rolling back transaction")
	return nil
}

// Database Operations Facade
type DatabaseFacade struct {
	connection        *DatabaseConnection
	queryExecutor     *QueryExecutor
	transactionManager *TransactionManager
}

func NewDatabaseFacade(host string, port int, username, password string) *DatabaseFacade {
	connection := NewDatabaseConnection(host, port, username, password)
	return &DatabaseFacade{
		connection:        connection,
		queryExecutor:     NewQueryExecutor(connection),
		transactionManager: NewTransactionManager(connection),
	}
}

func (df *DatabaseFacade) Connect() error {
	return df.connection.Connect()
}

func (df *DatabaseFacade) Disconnect() error {
	return df.connection.Disconnect()
}

func (df *DatabaseFacade) ExecuteQuery(sql string) ([]map[string]interface{}, error) {
	return df.queryExecutor.ExecuteQuery(sql)
}

func (df *DatabaseFacade) ExecuteCommand(sql string) error {
	return df.queryExecutor.ExecuteCommand(sql)
}

func (df *DatabaseFacade) BeginTransaction() error {
	return df.transactionManager.Begin()
}

func (df *DatabaseFacade) CommitTransaction() error {
	return df.transactionManager.Commit()
}

func (df *DatabaseFacade) RollbackTransaction() error {
	return df.transactionManager.Rollback()
}

func (df *DatabaseFacade) ExecuteInTransaction(operations func() error) error {
	if err := df.BeginTransaction(); err != nil {
		return err
	}
	
	if err := operations(); err != nil {
		df.RollbackTransaction()
		return err
	}
	
	return df.CommitTransaction()
}

// 3. E-COMMERCE CHECKOUT FACADE
type InventoryService struct{}

func (is *InventoryService) CheckAvailability(productID string, quantity int) bool {
	fmt.Printf("Checking availability for product %s, quantity %d\n", productID, quantity)
	return true
}

func (is *InventoryService) ReserveProduct(productID string, quantity int) error {
	fmt.Printf("Reserving product %s, quantity %d\n", productID, quantity)
	return nil
}

func (is *InventoryService) ReleaseProduct(productID string, quantity int) error {
	fmt.Printf("Releasing product %s, quantity %d\n", productID, quantity)
	return nil
}

type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(amount float64, paymentMethod string) (string, error) {
	fmt.Printf("Processing payment of $%.2f using %s\n", amount, paymentMethod)
	return fmt.Sprintf("payment_%d", time.Now().Unix()), nil
}

func (ps *PaymentService) RefundPayment(transactionID string, amount float64) error {
	fmt.Printf("Refunding payment %s, amount $%.2f\n", transactionID, amount)
	return nil
}

type ShippingService struct{}

func (ss *ShippingService) CalculateShipping(address string, weight float64) float64 {
	fmt.Printf("Calculating shipping to %s, weight %.2f kg\n", address, weight)
	return 10.0
}

func (ss *ShippingService) ScheduleDelivery(orderID string, address string) error {
	fmt.Printf("Scheduling delivery for order %s to %s\n", orderID, address)
	return nil
}

type NotificationService struct{}

func (ns *NotificationService) SendOrderConfirmation(orderID string, email string) error {
	fmt.Printf("Sending order confirmation for order %s to %s\n", orderID, email)
	return nil
}

func (ns *NotificationService) SendShippingNotification(orderID string, email string) error {
	fmt.Printf("Sending shipping notification for order %s to %s\n", orderID, email)
	return nil
}

// E-commerce Checkout Facade
type CheckoutFacade struct {
	inventoryService    *InventoryService
	paymentService      *PaymentService
	shippingService     *ShippingService
	notificationService *NotificationService
}

func NewCheckoutFacade() *CheckoutFacade {
	return &CheckoutFacade{
		inventoryService:    &InventoryService{},
		paymentService:      &PaymentService{},
		shippingService:     &ShippingService{},
		notificationService: &NotificationService{},
	}
}

func (cf *CheckoutFacade) ProcessOrder(orderID string, productID string, quantity int, amount float64, paymentMethod string, shippingAddress string, customerEmail string) error {
	fmt.Printf("=== Processing Order %s ===\n", orderID)
	
	// Check inventory
	if !cf.inventoryService.CheckAvailability(productID, quantity) {
		return fmt.Errorf("product not available")
	}
	
	// Reserve product
	if err := cf.inventoryService.ReserveProduct(productID, quantity); err != nil {
		return err
	}
	
	// Process payment
	transactionID, err := cf.paymentService.ProcessPayment(amount, paymentMethod)
	if err != nil {
		cf.inventoryService.ReleaseProduct(productID, quantity)
		return err
	}
	
	// Calculate shipping
	shippingCost := cf.shippingService.CalculateShipping(shippingAddress, float64(quantity))
	totalAmount := amount + shippingCost
	
	// Schedule delivery
	if err := cf.shippingService.ScheduleDelivery(orderID, shippingAddress); err != nil {
		cf.paymentService.RefundPayment(transactionID, amount)
		cf.inventoryService.ReleaseProduct(productID, quantity)
		return err
	}
	
	// Send notifications
	cf.notificationService.SendOrderConfirmation(orderID, customerEmail)
	cf.notificationService.SendShippingNotification(orderID, customerEmail)
	
	fmt.Printf("Order %s processed successfully! Total: $%.2f\n", orderID, totalAmount)
	return nil
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== FACADE PATTERN DEMONSTRATION ===\n")

	// 1. BASIC FACADE
	fmt.Println("1. BASIC FACADE:")
	computer := NewComputerFacade()
	computer.StartComputer()
	computer.ExecuteProgram()
	computer.ShutdownComputer()
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Home Automation Facade
	fmt.Println("Home Automation Facade:")
	homeAutomation := NewHomeAutomationFacade()
	homeAutomation.MovieNight()
	fmt.Println()
	homeAutomation.Bedtime()
	fmt.Println()
	homeAutomation.WakeUp()
	fmt.Println()
	homeAutomation.TurnOffEverything()
	fmt.Println()

	// Database Operations Facade
	fmt.Println("Database Operations Facade:")
	dbFacade := NewDatabaseFacade("localhost", 5432, "user", "password")
	dbFacade.Connect()
	
	// Execute simple query
	results, err := dbFacade.ExecuteQuery("SELECT * FROM users")
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
	} else {
		fmt.Printf("Query results: %v\n", results)
	}
	
	// Execute in transaction
	err = dbFacade.ExecuteInTransaction(func() error {
		if err := dbFacade.ExecuteCommand("INSERT INTO users (name) VALUES ('John')"); err != nil {
			return err
		}
		if err := dbFacade.ExecuteCommand("UPDATE users SET name = 'Jane' WHERE id = 1"); err != nil {
			return err
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
	} else {
		fmt.Println("Transaction completed successfully!")
	}
	
	dbFacade.Disconnect()
	fmt.Println()

	// E-commerce Checkout Facade
	fmt.Println("E-commerce Checkout Facade:")
	checkout := NewCheckoutFacade()
	
	err = checkout.ProcessOrder(
		"ORDER_123",
		"PRODUCT_456",
		2,
		100.0,
		"credit_card",
		"123 Main St, City, State",
		"customer@example.com",
	)
	
	if err != nil {
		fmt.Printf("Checkout failed: %v\n", err)
	} else {
		fmt.Println("Checkout completed successfully!")
	}
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
