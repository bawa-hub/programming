package main

import (
	"fmt"
)

// =============================================================================
// ENCAPSULATION EXAMPLE
// =============================================================================

// BankAccount demonstrates encapsulation
type BankAccount struct {
	accountNumber string
	balance      float64
	owner        string
	// private fields - cannot be accessed directly from outside the package
}

// Constructor function (Go doesn't have constructors, but we can create factory functions)
func NewBankAccount(accountNumber, owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		accountNumber: accountNumber,
		balance:      initialBalance,
		owner:        owner,
	}
}

// Public methods to access private data (encapsulation)
func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}

func (ba *BankAccount) GetAccountNumber() string {
	return ba.accountNumber
}

func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	ba.balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > ba.balance {
		return fmt.Errorf("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

// =============================================================================
// INHERITANCE EXAMPLE (Using Composition in Go)
// =============================================================================

// Base class (using struct embedding)
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v *Vehicle) Start() {
	fmt.Printf("%s %s is starting...\n", v.Brand, v.Model)
}

func (v *Vehicle) Stop() {
	fmt.Printf("%s %s is stopping...\n", v.Brand, v.Model)
}

func (v *Vehicle) GetInfo() string {
	return fmt.Sprintf("%s %s (%d)", v.Brand, v.Model, v.Year)
}

// Derived class (inheritance through embedding)
type Car struct {
	Vehicle // Embedded struct - Car "is-a" Vehicle
	Doors   int
	Engine  string
}

// Method overriding
func (c *Car) Start() {
	fmt.Printf("Car %s %s with %s engine is starting...\n", c.Brand, c.Model, c.Engine)
}

// Additional methods specific to Car
func (c *Car) OpenTrunk() {
	fmt.Printf("Opening trunk of %s %s\n", c.Brand, c.Model)
}

// Another derived class
type Motorcycle struct {
	Vehicle
	HasWindshield bool
}

func (m *Motorcycle) Start() {
	fmt.Printf("Motorcycle %s %s is starting...\n", m.Brand, m.Model)
}

func (m *Motorcycle) Wheelie() {
	fmt.Printf("Doing a wheelie on %s %s\n", m.Brand, m.Model)
}

// =============================================================================
// POLYMORPHISM EXAMPLE
// =============================================================================

// Interface for polymorphism
type Startable interface {
	Start()
}

// Function that works with any Startable object
func StartVehicle(v Startable) {
	v.Start()
}

// =============================================================================
// ABSTRACTION EXAMPLE
// =============================================================================

// Abstract interface (Go uses interfaces for abstraction)
type Shape interface {
	Area() float64
	Perimeter() float64
	GetName() string
}

// Concrete implementation of Shape
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

func (r *Rectangle) GetName() string {
	return "Rectangle"
}

// Another concrete implementation
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c *Circle) GetName() string {
	return "Circle"
}

// Function that works with any Shape (polymorphism)
func PrintShapeInfo(s Shape) {
	fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", s.GetName(), s.Area(), s.Perimeter())
}

// =============================================================================
// SOLID PRINCIPLES EXAMPLES
// =============================================================================

// 1. SINGLE RESPONSIBILITY PRINCIPLE (SRP)
// Each class should have only one reason to change

// User - handles only user data
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) GetDisplayName() string {
	return u.Name
}

// EmailService - handles only email operations
type EmailService struct{}

func (es *EmailService) SendEmail(to, subject, body string) error {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}

// 2. OPEN/CLOSED PRINCIPLE (OCP)
// Open for extension, closed for modification

// PaymentProcessor interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

// CreditCardProcessor implements PaymentProcessor
type CreditCardProcessor struct {
	CardNumber string
}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}

// PayPalProcessor implements PaymentProcessor
type PayPalProcessor struct {
	Email string
}

func (ppp *PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}

// PaymentService - closed for modification, open for extension
type PaymentService struct {
	processor PaymentProcessor
}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	return ps.processor.ProcessPayment(amount)
}

// 3. LISKOV SUBSTITUTION PRINCIPLE (LSP)
// Objects of superclass should be replaceable with objects of subclass

// Bird interface
type Bird interface {
	Fly() string
	MakeSound() string
}

// Sparrow implements Bird
type Sparrow struct{}

func (s *Sparrow) Fly() string {
	return "Sparrow is flying"
}

func (s *Sparrow) MakeSound() string {
	return "Chirp chirp"
}

// Penguin implements Bird (but can't fly)
type Penguin struct{}

func (p *Penguin) Fly() string {
	return "Penguin cannot fly"
}

func (p *Penguin) MakeSound() string {
	return "Honk honk"
}

// Function that works with any Bird
func MakeBirdFly(b Bird) {
	fmt.Println(b.Fly())
}

// 4. INTERFACE SEGREGATION PRINCIPLE (ISP)
// Clients should not be forced to depend on interfaces they don't use

// Instead of one large interface, create smaller, focused interfaces

// Worker interface
type Worker interface {
	Work()
}

// Eater interface
type Eater interface {
	Eat()
}

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// Human implements all interfaces
type Human struct {
	Name string
}

func (h *Human) Work() {
	fmt.Printf("%s is working\n", h.Name)
}

func (h *Human) Eat() {
	fmt.Printf("%s is eating\n", h.Name)
}

func (h *Human) Sleep() {
	fmt.Printf("%s is sleeping\n", h.Name)
}

// Robot implements only Worker
type Robot struct {
	Model string
}

func (r *Robot) Work() {
	fmt.Printf("Robot %s is working\n", r.Model)
}

// 5. DEPENDENCY INVERSION PRINCIPLE (DIP)
// High-level modules should not depend on low-level modules

// Database interface (abstraction)
type Database interface {
	Save(data string) error
	Get(id string) (string, error)
}

// MySQLDatabase implements Database
type MySQLDatabase struct{}

func (mdb *MySQLDatabase) Save(data string) error {
	fmt.Printf("Saving to MySQL: %s\n", data)
	return nil
}

func (mdb *MySQLDatabase) Get(id string) (string, error) {
	fmt.Printf("Getting from MySQL: %s\n", id)
	return "data", nil
}

// UserService depends on Database interface, not concrete implementation
type UserService struct {
	db Database
}

func (us *UserService) CreateUser(name string) error {
	return us.db.Save(name)
}

func (us *UserService) GetUser(id string) (string, error) {
	return us.db.Get(id)
}

// =============================================================================
// COMPOSITION VS INHERITANCE EXAMPLE
// =============================================================================

// Using Composition (HAS-A relationship)
type Engine struct {
	Type        string
	Horsepower  int
	IsRunning   bool
}

func (e *Engine) Start() {
	e.IsRunning = true
	fmt.Printf("%s engine started\n", e.Type)
}

func (e *Engine) Stop() {
	e.IsRunning = false
	fmt.Printf("%s engine stopped\n", e.Type)
}

// Car HAS-A Engine (composition)
type ModernCar struct {
	Brand  string
	Model  string
	Engine Engine // Composition
}

func (mc *ModernCar) Start() {
	mc.Engine.Start()
	fmt.Printf("%s %s is ready to drive\n", mc.Brand, mc.Model)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== OOPS FUNDAMENTALS DEMONSTRATION ===\n")

	// 1. ENCAPSULATION
	fmt.Println("1. ENCAPSULATION:")
	account := NewBankAccount("12345", "John Doe", 1000.0)
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
	
	err := account.Deposit(500.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After deposit: $%.2f\n", account.GetBalance())
	}

	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
	}
	fmt.Println()

	// 2. INHERITANCE
	fmt.Println("2. INHERITANCE:")
	car := &Car{
		Vehicle: Vehicle{Brand: "Toyota", Model: "Camry", Year: 2023},
		Doors:   4,
		Engine:  "V6",
	}
	car.Start()
	car.OpenTrunk()

	motorcycle := &Motorcycle{
		Vehicle:       Vehicle{Brand: "Honda", Model: "CBR600", Year: 2023},
		HasWindshield: true,
	}
	motorcycle.Start()
	motorcycle.Wheelie()
	fmt.Println()

	// 3. POLYMORPHISM
	fmt.Println("3. POLYMORPHISM:")
	vehicles := []Startable{car, motorcycle}
	for _, vehicle := range vehicles {
		StartVehicle(vehicle)
	}
	fmt.Println()

	// 4. ABSTRACTION
	fmt.Println("4. ABSTRACTION:")
	shapes := []Shape{
		&Rectangle{Width: 5, Height: 3},
		&Circle{Radius: 4},
	}
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
	fmt.Println()

	// 5. SOLID PRINCIPLES
	fmt.Println("5. SOLID PRINCIPLES:")
	
	// SRP - Separate concerns
	user := &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	emailService := &EmailService{}
	fmt.Printf("User: %s\n", user.GetDisplayName())
	emailService.SendEmail(user.Email, "Welcome", "Welcome to our service!")
	fmt.Println()

	// OCP - Open for extension
	creditCardProcessor := &CreditCardProcessor{CardNumber: "1234-5678-9012-3456"}
	paymentService := &PaymentService{processor: creditCardProcessor}
	paymentService.ProcessPayment(100.0)
	fmt.Println()

	// LSP - Substitutability
	birds := []Bird{&Sparrow{}, &Penguin{}}
	for _, bird := range birds {
		MakeBirdFly(bird)
	}
	fmt.Println()

	// ISP - Interface segregation
	human := &Human{Name: "Bob"}
	robot := &Robot{Model: "T-800"}
	
	workers := []Worker{human, robot}
	for _, worker := range workers {
		worker.Work()
	}
	fmt.Println()

	// DIP - Dependency inversion
	mysqlDB := &MySQLDatabase{}
	userService := &UserService{db: mysqlDB}
	userService.CreateUser("Charlie")
	userService.GetUser("1")
	fmt.Println()

	// 6. COMPOSITION VS INHERITANCE
	fmt.Println("6. COMPOSITION VS INHERITANCE:")
	modernCar := &ModernCar{
		Brand:  "BMW",
		Model:  "X5",
		Engine: Engine{Type: "V8", Horsepower: 400},
	}
	modernCar.Start()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
