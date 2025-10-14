package main

import (
	"fmt"
	"io"
	"math"
	"sort"
	"time"
)

// 🎯 INTERFACE TYPES MASTERY
// This file demonstrates comprehensive interface usage and patterns

// Basic interfaces
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Composed interfaces
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// Custom interfaces for our CRUD application
type CRUD interface {
	Create() error
	Read() error
	Update() error
	Delete() error
}

type Identifiable interface {
	GetID() int
	SetID(int)
}

type Timestamped interface {
	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(time.Time)
}

type Validatable interface {
	Validate() error
}

// Shape interface for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Drawable interface for objects that can be drawn
type Drawable interface {
	Draw() string
	GetColor() string
	SetColor(string)
}

// Movable interface for objects that can move
type Movable interface {
	Move(x, y float64)
	GetPosition() (float64, float64)
}

// Animal interface for different types of animals
type Animal interface {
	Speak() string
	Move() string
	Eat() string
	GetName() string
}

// Vehicle interface for different types of vehicles
type Vehicle interface {
	Start() error
	Stop() error
	GetSpeed() float64
	SetSpeed(float64)
	GetFuel() float64
	Refuel(float64)
}

// Database interface for different database implementations
type Database interface {
	Connect() error
	Disconnect() error
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
	BeginTransaction() (Transaction, error)
}

type Transaction interface {
	Commit() error
	Rollback() error
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
}

// Storage interface for different storage backends
type Storage interface {
	Store(key string, value interface{}) error
	Retrieve(key string) (interface{}, error)
	Delete(key string) error
	List() ([]string, error)
}

// Serializer interface for different serialization formats
type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	Deserialize(data []byte, target interface{}) error
	GetContentType() string
}

// Implementations of interfaces

// Shape implementations
type Circle struct {
	Radius float64
	Color  string
	X, Y   float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing a %s circle with radius %.2f", c.Color, c.Radius)
}

func (c Circle) GetColor() string {
	return c.Color
}

func (c Circle) SetColor(color string) {
	c.Color = color
}

func (c Circle) Move(x, y float64) {
	c.X += x
	c.Y += y
}

func (c Circle) GetPosition() (float64, float64) {
	return c.X, c.Y
}

type Rectangle struct {
	Width, Height float64
	Color         string
	X, Y          float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f)", r.Width, r.Height)
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing a %s rectangle %.2fx%.2f", r.Color, r.Width, r.Height)
}

func (r Rectangle) GetColor() string {
	return r.Color
}

func (r Rectangle) SetColor(color string) {
	r.Color = color
}

func (r Rectangle) Move(x, y float64) {
	r.X += x
	r.Y += y
}

func (r Rectangle) GetPosition() (float64, float64) {
	return r.X, r.Y
}

// Animal implementations
type Dog struct {
	Name  string
	Breed string
	Age   int
}

func (d Dog) Speak() string {
	return "Woof! Woof!"
}

func (d Dog) Move() string {
	return "Running on four legs"
}

func (d Dog) Eat() string {
	return "Eating dog food"
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name  string
	Breed string
	Age   int
}

func (c Cat) Speak() string {
	return "Meow! Meow!"
}

func (c Cat) Move() string {
	return "Walking gracefully"
}

func (c Cat) Eat() string {
	return "Eating cat food"
}

func (c Cat) GetName() string {
	return c.Name
}

type Bird struct {
	Name  string
	Species string
	Age   int
}

func (b Bird) Speak() string {
	return "Tweet! Tweet!"
}

func (b Bird) Move() string {
	return "Flying in the sky"
}

func (b Bird) Eat() string {
	return "Eating seeds and insects"
}

func (b Bird) GetName() string {
	return b.Name
}

// Vehicle implementations
type Car struct {
	Make        string
	Model       string
	Year        int
	Speed       float64
	Fuel        float64
	MaxFuel     float64
	IsRunning   bool
}

func (c *Car) Start() error {
	if c.IsRunning {
		return fmt.Errorf("car is already running")
	}
	c.IsRunning = true
	return nil
}

func (c *Car) Stop() error {
	if !c.IsRunning {
		return fmt.Errorf("car is not running")
	}
	c.IsRunning = false
	c.Speed = 0
	return nil
}

func (c *Car) GetSpeed() float64 {
	return c.Speed
}

func (c *Car) SetSpeed(speed float64) {
	if c.IsRunning && speed >= 0 {
		c.Speed = speed
	}
}

func (c *Car) GetFuel() float64 {
	return c.Fuel
}

func (c *Car) Refuel(amount float64) {
	if amount > 0 && c.Fuel+amount <= c.MaxFuel {
		c.Fuel += amount
	}
}

type Motorcycle struct {
	Make        string
	Model       string
	Year        int
	Speed       float64
	Fuel        float64
	MaxFuel     float64
	IsRunning   bool
}

func (m *Motorcycle) Start() error {
	if m.IsRunning {
		return fmt.Errorf("motorcycle is already running")
	}
	m.IsRunning = true
	return nil
}

func (m *Motorcycle) Stop() error {
	if !m.IsRunning {
		return fmt.Errorf("motorcycle is not running")
	}
	m.IsRunning = false
	m.Speed = 0
	return nil
}

func (m *Motorcycle) GetSpeed() float64 {
	return m.Speed
}

func (m *Motorcycle) SetSpeed(speed float64) {
	if m.IsRunning && speed >= 0 {
		m.Speed = speed
	}
}

func (m *Motorcycle) GetFuel() float64 {
	return m.Fuel
}

func (m *Motorcycle) Refuel(amount float64) {
	if amount > 0 && m.Fuel+amount <= m.MaxFuel {
		m.Fuel += amount
	}
}

// Database implementations
type MockDatabase struct {
	connected bool
	data      map[string]interface{}
}

func (m *MockDatabase) Connect() error {
	m.connected = true
	m.data = make(map[string]interface{})
	return nil
}

func (m *MockDatabase) Disconnect() error {
	m.connected = false
	return nil
}

func (m *MockDatabase) Query(sql string) ([]map[string]interface{}, error) {
	if !m.connected {
		return nil, fmt.Errorf("not connected to database")
	}
	
	// Mock query result
	result := []map[string]interface{}{
		{"id": 1, "name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"},
	}
	
	return result, nil
}

func (m *MockDatabase) Execute(sql string) error {
	if !m.connected {
		return fmt.Errorf("not connected to database")
	}
	
	// Mock execution
	fmt.Printf("Executing SQL: %s\n", sql)
	return nil
}

func (m *MockDatabase) BeginTransaction() (Transaction, error) {
	if !m.connected {
		return nil, fmt.Errorf("not connected to database")
	}
	
	return &MockTransaction{db: m}, nil
}

type MockTransaction struct {
	db *MockDatabase
}

func (t *MockTransaction) Commit() error {
	fmt.Println("Transaction committed")
	return nil
}

func (t *MockTransaction) Rollback() error {
	fmt.Println("Transaction rolled back")
	return nil
}

func (t *MockTransaction) Query(sql string) ([]map[string]interface{}, error) {
	return t.db.Query(sql)
}

func (t *MockTransaction) Execute(sql string) error {
	return t.db.Execute(sql)
}

// Storage implementations
type MemoryStorage struct {
	data map[string]interface{}
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]interface{}),
	}
}

func (m *MemoryStorage) Store(key string, value interface{}) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Retrieve(key string) (interface{}, error) {
	value, exists := m.data[key]
	if !exists {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return value, nil
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

func (m *MemoryStorage) List() ([]string, error) {
	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys, nil
}

// Serializer implementations
type JSONSerializer struct{}

func (j JSONSerializer) Serialize(data interface{}) ([]byte, error) {
	// Mock JSON serialization
	return []byte(fmt.Sprintf(`{"data": %+v}`, data)), nil
}

func (j JSONSerializer) Deserialize(data []byte, target interface{}) error {
	// Mock JSON deserialization
	fmt.Printf("Deserializing JSON: %s\n", string(data))
	return nil
}

func (j JSONSerializer) GetContentType() string {
	return "application/json"
}

type XMLSerializer struct{}

func (x XMLSerializer) Serialize(data interface{}) ([]byte, error) {
	// Mock XML serialization
	return []byte(fmt.Sprintf(`<data>%+v</data>`, data)), nil
}

func (x XMLSerializer) Deserialize(data []byte, target interface{}) error {
	// Mock XML deserialization
	fmt.Printf("Deserializing XML: %s\n", string(data))
	return nil
}

func (x XMLSerializer) GetContentType() string {
	return "application/xml"
}

// Interface Manager for CRUD operations
type InterfaceManager struct {
	Shapes     []Shape
	Animals    []Animal
	Vehicles   []Vehicle
	Databases  []Database
	Storages   []Storage
	Serializers []Serializer
}

// NewInterfaceManager creates a new interface manager
func NewInterfaceManager() *InterfaceManager {
	return &InterfaceManager{
		Shapes:     make([]Shape, 0),
		Animals:    make([]Animal, 0),
		Vehicles:   make([]Vehicle, 0),
		Databases:  make([]Database, 0),
		Storages:   make([]Storage, 0),
		Serializers: make([]Serializer, 0),
	}
}

// CRUD Operations for Interfaces

// Create - Initialize interface implementations
func (im *InterfaceManager) Create() {
	fmt.Println("🔧 Creating interface implementations...")
	
	// Create shapes
	circle := Circle{Radius: 5.0, Color: "red", X: 0, Y: 0}
	rectangle := Rectangle{Width: 10.0, Height: 8.0, Color: "blue", X: 0, Y: 0}
	
	im.Shapes = append(im.Shapes, circle, rectangle)
	
	// Create animals
	dog := Dog{Name: "Buddy", Breed: "Golden Retriever", Age: 3}
	cat := Cat{Name: "Whiskers", Breed: "Persian", Age: 2}
	bird := Bird{Name: "Tweety", Species: "Canary", Age: 1}
	
	im.Animals = append(im.Animals, dog, cat, bird)
	
	// Create vehicles
	car := &Car{Make: "Toyota", Model: "Camry", Year: 2020, Speed: 0, Fuel: 50.0, MaxFuel: 60.0, IsRunning: false}
	motorcycle := &Motorcycle{Make: "Honda", Model: "CBR", Year: 2021, Speed: 0, Fuel: 20.0, MaxFuel: 25.0, IsRunning: false}
	
	im.Vehicles = append(im.Vehicles, car, motorcycle)
	
	// Create databases
	mockDB := &MockDatabase{}
	im.Databases = append(im.Databases, mockDB)
	
	// Create storages
	memoryStorage := NewMemoryStorage()
	im.Storages = append(im.Storages, memoryStorage)
	
	// Create serializers
	jsonSerializer := JSONSerializer{}
	xmlSerializer := XMLSerializer{}
	
	im.Serializers = append(im.Serializers, jsonSerializer, xmlSerializer)
	
	fmt.Println("✅ Interface implementations created successfully")
}

// Read - Display interface information
func (im *InterfaceManager) Read() {
	fmt.Println("\n📖 READING INTERFACE IMPLEMENTATIONS:")
	fmt.Println("====================================")
	
	// Read shapes
	fmt.Printf("Shapes (%d):\n", len(im.Shapes))
	for i, shape := range im.Shapes {
		fmt.Printf("  %d: %s (Area: %.2f, Perimeter: %.2f)\n", 
			i+1, shape.String(), shape.Area(), shape.Perimeter())
	}
	
	// Read animals
	fmt.Printf("\nAnimals (%d):\n", len(im.Animals))
	for i, animal := range im.Animals {
		fmt.Printf("  %d: %s - %s\n", 
			i+1, animal.GetName(), animal.Speak())
	}
	
	// Read vehicles
	fmt.Printf("\nVehicles (%d):\n", len(im.Vehicles))
	for i, vehicle := range im.Vehicles {
		fmt.Printf("  %d: Speed: %.2f, Fuel: %.2f\n", 
			i+1, vehicle.GetSpeed(), vehicle.GetFuel())
	}
	
	// Read databases
	fmt.Printf("\nDatabases (%d):\n", len(im.Databases))
	for i, db := range im.Databases {
		fmt.Printf("  %d: Database instance (%T)\n", i+1, db)
	}
	
	// Read storages
	fmt.Printf("\nStorages (%d):\n", len(im.Storages))
	for i, storage := range im.Storages {
		keys, _ := storage.List()
		fmt.Printf("  %d: Storage with %d items\n", i+1, len(keys))
	}
	
	// Read serializers
	fmt.Printf("\nSerializers (%d):\n", len(im.Serializers))
	for i, serializer := range im.Serializers {
		fmt.Printf("  %d: %s\n", i+1, serializer.GetContentType())
	}
}

// Update - Modify interface implementations
func (im *InterfaceManager) Update() {
	fmt.Println("\n🔄 UPDATING INTERFACE IMPLEMENTATIONS:")
	fmt.Println("=====================================")
	
	// Update shapes
	for i, shape := range im.Shapes {
		if drawable, ok := shape.(Drawable); ok {
			drawable.SetColor("green")
			fmt.Printf("Updated shape %d color to green\n", i+1)
		}
		
		if movable, ok := shape.(Movable); ok {
			movable.Move(10, 10)
			x, y := movable.GetPosition()
			fmt.Printf("Moved shape %d to position (%.2f, %.2f)\n", i+1, x, y)
		}
	}
	
	// Update vehicles
	for i, vehicle := range im.Vehicles {
		vehicle.Start()
		vehicle.SetSpeed(60.0)
		vehicle.Refuel(10.0)
		fmt.Printf("Updated vehicle %d: started, speed=%.2f, fuel=%.2f\n", 
			i+1, vehicle.GetSpeed(), vehicle.GetFuel())
	}
	
	// Update storages
	for i, storage := range im.Storages {
		storage.Store(fmt.Sprintf("key%d", i+1), fmt.Sprintf("value%d", i+1))
		fmt.Printf("Stored data in storage %d\n", i+1)
	}
	
	fmt.Println("✅ Interface implementations updated successfully")
}

// Delete - Remove interface implementations
func (im *InterfaceManager) Delete() {
	fmt.Println("\n🗑️  DELETING INTERFACE IMPLEMENTATIONS:")
	fmt.Println("======================================")
	
	// Delete shapes
	if len(im.Shapes) > 0 {
		deleted := im.Shapes[len(im.Shapes)-1]
		im.Shapes = im.Shapes[:len(im.Shapes)-1]
		fmt.Printf("Deleted shape: %s\n", deleted.String())
	}
	
	// Delete animals
	if len(im.Animals) > 0 {
		deleted := im.Animals[len(im.Animals)-1]
		im.Animals = im.Animals[:len(im.Animals)-1]
		fmt.Printf("Deleted animal: %s\n", deleted.GetName())
	}
	
	// Delete vehicles
	if len(im.Vehicles) > 0 {
		deleted := im.Vehicles[len(im.Vehicles)-1]
		deleted.Stop()
		im.Vehicles = im.Vehicles[:len(im.Vehicles)-1]
		fmt.Printf("Deleted vehicle\n")
	}
	
	// Clear storages
	for i, storage := range im.Storages {
		keys, _ := storage.List()
		for _, key := range keys {
			storage.Delete(key)
		}
		fmt.Printf("Cleared storage %d\n", i+1)
	}
	
	fmt.Println("✅ Interface implementations deleted successfully")
}

// Advanced Interface Operations

// DemonstrateInterfacePolymorphism shows polymorphism with interfaces
func (im *InterfaceManager) DemonstrateInterfacePolymorphism() {
	fmt.Println("\n🔄 INTERFACE POLYMORPHISM DEMONSTRATION:")
	fmt.Println("=======================================")
	
	// Process shapes polymorphically
	fmt.Println("Processing shapes:")
	for i, shape := range im.Shapes {
		fmt.Printf("  Shape %d: %s\n", i+1, shape.String())
		fmt.Printf("    Area: %.2f\n", shape.Area())
		fmt.Printf("    Perimeter: %.2f\n", shape.Perimeter())
		
		// Type assertion to check for additional interfaces
		if drawable, ok := shape.(Drawable); ok {
			fmt.Printf("    Draw: %s\n", drawable.Draw())
		}
		
		if movable, ok := shape.(Movable); ok {
			x, y := movable.GetPosition()
			fmt.Printf("    Position: (%.2f, %.2f)\n", x, y)
		}
	}
	
	// Process animals polymorphically
	fmt.Println("\nProcessing animals:")
	for i, animal := range im.Animals {
		fmt.Printf("  Animal %d: %s\n", i+1, animal.GetName())
		fmt.Printf("    Speak: %s\n", animal.Speak())
		fmt.Printf("    Move: %s\n", animal.Move())
		fmt.Printf("    Eat: %s\n", animal.Eat())
	}
	
	// Process vehicles polymorphically
	fmt.Println("\nProcessing vehicles:")
	for i, vehicle := range im.Vehicles {
		fmt.Printf("  Vehicle %d:\n", i+1)
		fmt.Printf("    Speed: %.2f\n", vehicle.GetSpeed())
		fmt.Printf("    Fuel: %.2f\n", vehicle.GetFuel())
	}
}

// DemonstrateInterfaceComposition shows interface composition
func (im *InterfaceManager) DemonstrateInterfaceComposition() {
	fmt.Println("\n🧩 INTERFACE COMPOSITION DEMONSTRATION:")
	fmt.Println("======================================")
	
	// Create a composite interface
	type DrawableMovable interface {
		Drawable
		Movable
	}
	
	// Check which shapes implement the composite interface
	fmt.Println("Shapes that are both drawable and movable:")
	for i, shape := range im.Shapes {
		if _, ok := shape.(DrawableMovable); ok {
			fmt.Printf("  Shape %d: %s\n", i+1, shape.String())
		}
	}
	
	// Demonstrate interface embedding
	type ReadWriteCloser interface {
		io.Reader
		io.Writer
		io.Closer
	}
	
	fmt.Println("\nReadWriteCloser interface combines Reader, Writer, and Closer")
}

// DemonstrateInterfaceTypeAssertion shows type assertion
func (im *InterfaceManager) DemonstrateInterfaceTypeAssertion() {
	fmt.Println("\n🔍 INTERFACE TYPE ASSERTION DEMONSTRATION:")
	fmt.Println("=========================================")
	
	// Type assertion with shapes
	for i, shape := range im.Shapes {
		fmt.Printf("Shape %d: %s\n", i+1, shape.String())
		
		// Type assertion to specific type
		if circle, ok := shape.(Circle); ok {
			fmt.Printf("  It's a circle with radius %.2f\n", circle.Radius)
		} else if rectangle, ok := shape.(Rectangle); ok {
			fmt.Printf("  It's a rectangle %.2fx%.2f\n", rectangle.Width, rectangle.Height)
		}
		
		// Type assertion to interface
		if drawable, ok := shape.(Drawable); ok {
			fmt.Printf("  It's drawable: %s\n", drawable.Draw())
		}
	}
	
	// Type assertion with animals
	for i, animal := range im.Animals {
		fmt.Printf("Animal %d: %s\n", i+1, animal.GetName())
		
		if dog, ok := animal.(Dog); ok {
			fmt.Printf("  It's a dog of breed %s\n", dog.Breed)
		} else if cat, ok := animal.(Cat); ok {
			fmt.Printf("  It's a cat of breed %s\n", cat.Breed)
		} else if bird, ok := animal.(Bird); ok {
			fmt.Printf("  It's a bird of species %s\n", bird.Species)
		}
	}
}

// DemonstrateInterfaceSorting shows sorting with interfaces
func (im *InterfaceManager) DemonstrateInterfaceSorting() {
	fmt.Println("\n📊 INTERFACE SORTING DEMONSTRATION:")
	fmt.Println("==================================")
	
	// Sort shapes by area
	fmt.Println("Sorting shapes by area:")
	sort.Slice(im.Shapes, func(i, j int) bool {
		return im.Shapes[i].Area() < im.Shapes[j].Area()
	})
	
	for i, shape := range im.Shapes {
		fmt.Printf("  %d: %s (Area: %.2f)\n", i+1, shape.String(), shape.Area())
	}
	
	// Sort animals by name
	fmt.Println("\nSorting animals by name:")
	sort.Slice(im.Animals, func(i, j int) bool {
		return im.Animals[i].GetName() < im.Animals[j].GetName()
	})
	
	for i, animal := range im.Animals {
		fmt.Printf("  %d: %s\n", i+1, animal.GetName())
	}
}

// DemonstrateInterfaceValidation shows interface validation
func (im *InterfaceManager) DemonstrateInterfaceValidation() {
	fmt.Println("\n✅ INTERFACE VALIDATION DEMONSTRATION:")
	fmt.Println("====================================")
	
	// Validate shapes
	fmt.Println("Validating shapes:")
	for i, shape := range im.Shapes {
		if shape.Area() > 0 {
			fmt.Printf("  Shape %d: Valid (area > 0)\n", i+1)
		} else {
			fmt.Printf("  Shape %d: Invalid (area <= 0)\n", i+1)
		}
	}
	
	// Validate vehicles
	fmt.Println("\nValidating vehicles:")
	for i, vehicle := range im.Vehicles {
		if vehicle.GetFuel() > 0 {
			fmt.Printf("  Vehicle %d: Valid (has fuel)\n", i+1)
		} else {
			fmt.Printf("  Vehicle %d: Invalid (no fuel)\n", i+1)
		}
	}
}

// DemonstrateInterfaceChaining shows interface method chaining
func (im *InterfaceManager) DemonstrateInterfaceChaining() {
	fmt.Println("\n🔗 INTERFACE METHOD CHAINING DEMONSTRATION:")
	fmt.Println("==========================================")
	
	// Chain vehicle operations
	for i, vehicle := range im.Vehicles {
		fmt.Printf("Vehicle %d operations:\n", i+1)
		
		// Start -> Set Speed -> Refuel -> Stop
		if err := vehicle.Start(); err != nil {
			fmt.Printf("  Start failed: %v\n", err)
			continue
		}
		fmt.Printf("  Started successfully\n")
		
		vehicle.SetSpeed(50.0)
		fmt.Printf("  Speed set to %.2f\n", vehicle.GetSpeed())
		
		vehicle.Refuel(5.0)
		fmt.Printf("  Refueled, fuel level: %.2f\n", vehicle.GetFuel())
		
		if err := vehicle.Stop(); err != nil {
			fmt.Printf("  Stop failed: %v\n", err)
		} else {
			fmt.Printf("  Stopped successfully\n")
		}
	}
}

// DemonstrateInterfaceReflection shows interface reflection
func (im *InterfaceManager) DemonstrateInterfaceReflection() {
	fmt.Println("\n🔍 INTERFACE REFLECTION DEMONSTRATION:")
	fmt.Println("=====================================")
	
	// Reflect on interface types
	interfaces := []interface{}{
		Circle{Radius: 5.0},
		Rectangle{Width: 10.0, Height: 8.0},
		Dog{Name: "Buddy", Breed: "Golden Retriever", Age: 3},
		&Car{Make: "Toyota", Model: "Camry", Year: 2020},
	}
	
	for i, item := range interfaces {
		fmt.Printf("Item %d: %T\n", i+1, item)
		
		// Check if item implements specific interfaces
		if _, ok := item.(Shape); ok {
			fmt.Printf("  Implements Shape interface\n")
		}
		if _, ok := item.(Animal); ok {
			fmt.Printf("  Implements Animal interface\n")
		}
		if _, ok := item.(Vehicle); ok {
			fmt.Printf("  Implements Vehicle interface\n")
		}
		if _, ok := item.(Drawable); ok {
			fmt.Printf("  Implements Drawable interface\n")
		}
		if _, ok := item.(Movable); ok {
			fmt.Printf("  Implements Movable interface\n")
		}
	}
}
