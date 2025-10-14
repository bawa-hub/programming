package main

import "fmt"

var pl = fmt.Println

// An interface is a set of method signatures.
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

// Empty Interface (interface{})
	// Old Go version: interface{} meant any type.
	// Modern Go: any is an alias for interface{}.
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
// But values stored in an interface have type + value pair (fat pointer).
// This is why type assertions are needed.

// Type Assertions & Type Switch
func Describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    default:
        fmt.Println("unknown type")
    }
}


func main() {

	// Interfaces allow you to create contractsthat say if anything inherits it that they will implement defined methods
	// If we had animals and wanted to define that they all perform certain actions, but in their specific way we could use an interface
	// With Go you don't have to say a type uses an interface. When your type implements the required methods it is automatic

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	// We can only call methods defined in the interface for Cats because of the contract unless you convert Cat back into a concrete Cat type using a type assertion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}


// Basic interfaces
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Composed interface
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
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
}

// Movable interface for objects that can move
type Movable interface {
	Move(dx, dy float64)
	Position() (float64, float64)
}

// DrawableShape combines Shape and Drawable
type DrawableShape interface {
	Shape
	Drawable
}

// Point represents a 2D point
type Point struct {
	X, Y float64
}

// Circle implements Shape and Drawable
type Circle struct {
	Center Point
	Radius float64
}

// Rectangle implements Shape and Drawable
type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

// Triangle implements Shape and Drawable
type Triangle struct {
	P1, P2, P3 Point
}

// MovableCircle implements Shape, Drawable, and Movable
type MovableCircle struct {
	Circle
}

// Stringer interface (built-in)
type Stringer interface {
	String() string
}

// Custom types implementing Stringer
type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}

// DemonstrateBasicInterfaces shows basic interface concepts
func DemonstrateBasicInterfaces() {
	fmt.Println("=== BASIC INTERFACES ===")
	
	// 1. Interface Declaration
	fmt.Println("\n--- Interface Declaration ---")
	
	var w Writer
	fmt.Printf("Zero value of Writer interface: %v\n", w)
	fmt.Printf("Is Writer nil? %t\n", w == nil)
	
	// 2. Interface Implementation
	fmt.Println("\n--- Interface Implementation ---")
	
	// Create a concrete type that implements Writer
	fileWriter := &FileWriter{filename: "test.txt"}
	w = fileWriter
	
	// Call interface method
	n, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes\n", n)
	}
	
	// 3. Interface Composition
	fmt.Println("\n--- Interface Composition ---")
	
	var rwc ReadWriteCloser = &File{}
	
	// Can call methods from all composed interfaces
	data := make([]byte, 100)
	n, err = rwc.Read(data)
	fmt.Printf("Read %d bytes\n", n)
	
	n, err = rwc.Write([]byte("data"))
	fmt.Printf("Wrote %d bytes\n", n)
	
	err = rwc.Close()
	fmt.Printf("Closed: %v\n", err)
}

// DemonstrateShapeInterfaces shows shape interface implementations
func DemonstrateShapeInterfaces() {
	fmt.Println("\n=== SHAPE INTERFACES ===")
	
	// Create different shapes
	shapes := []Shape{
		Circle{Center: Point{0, 0}, Radius: 5},
		Rectangle{TopLeft: Point{0, 0}, BottomRight: Point{10, 5}},
		Triangle{P1: Point{0, 0}, P2: Point{5, 0}, P3: Point{2.5, 5}},
	}
	
	// 1. Polymorphism - same interface, different implementations
	fmt.Println("\n--- Polymorphism ---")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: %s\n", i+1, shape.String())
		fmt.Printf("  Area: %.2f\n", shape.Area())
		fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
	
	// 2. Type Assertion
	fmt.Println("\n--- Type Assertion ---")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		
		// Type assertion to specific type
		if circle, ok := shape.(Circle); ok {
			fmt.Printf("Circle with radius %.2f\n", circle.Radius)
		} else if rect, ok := shape.(Rectangle); ok {
			width := rect.BottomRight.X - rect.TopLeft.X
			height := rect.BottomRight.Y - rect.TopLeft.Y
			fmt.Printf("Rectangle %.2fx%.2f\n", width, height)
		} else if triangle, ok := shape.(Triangle); ok {
			fmt.Printf("Triangle with points (%.1f,%.1f), (%.1f,%.1f), (%.1f,%.1f)\n",
				triangle.P1.X, triangle.P1.Y,
				triangle.P2.X, triangle.P2.Y,
				triangle.P3.X, triangle.P3.Y)
		}
	}
	
	// 3. Type Switch
	fmt.Println("\n--- Type Switch ---")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", s.Radius)
		case Rectangle:
			width := s.BottomRight.X - s.TopLeft.X
			height := s.BottomRight.Y - s.TopLeft.Y
			fmt.Printf("Rectangle %.2fx%.2f\n", width, height)
		case Triangle:
			fmt.Printf("Triangle\n")
		default:
			fmt.Printf("Unknown shape type\n")
		}
	}
}

// DemonstrateEmptyInterface shows empty interface usage
func DemonstrateEmptyInterface() {
	fmt.Println("\n=== EMPTY INTERFACE ===")
	
	// 1. Empty interface can hold any type
	fmt.Println("\n--- Empty Interface ---")
	
	var empty interface{}
	
	empty = 42
	fmt.Printf("empty = %v (type: %T)\n", empty, empty)
	
	empty = "Hello"
	fmt.Printf("empty = %v (type: %T)\n", empty, empty)
	
	empty = []int{1, 2, 3}
	fmt.Printf("empty = %v (type: %T)\n", empty, empty)
	
	// 2. Type assertion with empty interface
	fmt.Println("\n--- Type Assertion with Empty Interface ---")
	
	values := []interface{}{
		42,
		"Hello, World!",
		[]int{1, 2, 3},
		Person{Name: "John", Age: 30},
		true,
		3.14,
	}
	
	for i, value := range values {
		fmt.Printf("Value %d: ", i+1)
		
		switch v := value.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case []int:
			fmt.Printf("Slice of ints: %v\n", v)
		case Person:
			fmt.Printf("Person: %+v\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		case float64:
			fmt.Printf("Float: %.2f\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}
}

// DemonstrateInterfaceComposition shows interface composition
func DemonstrateInterfaceComposition() {
	fmt.Println("\n=== INTERFACE COMPOSITION ===")
	
	// 1. Multiple interface implementation
	fmt.Println("\n--- Multiple Interface Implementation ---")
	
	movableCircle := MovableCircle{
		Circle: Circle{Center: Point{0, 0}, Radius: 5},
	}
	
	// Can be used as different interface types
	var shape Shape = movableCircle
	var drawable Drawable = movableCircle
	var movable Movable = movableCircle
	var drawableShape DrawableShape = movableCircle
	
	fmt.Printf("As Shape: %s\n", shape.String())
	fmt.Printf("As Drawable: %s\n", drawable.Draw())
	x, y := movable.Position()
	fmt.Printf("As Movable: Position (%.1f, %.1f)\n", x, y)
	fmt.Printf("As DrawableShape: %s\n", drawableShape.String())
	
	// 2. Interface embedding
	fmt.Println("\n--- Interface Embedding ---")
	
	// DrawableShape embeds both Shape and Drawable
	fmt.Printf("DrawableShape Area: %.2f\n", drawableShape.Area())
	fmt.Printf("DrawableShape Draw: %s\n", drawableShape.Draw())
}

// DemonstrateStringerInterface shows Stringer interface implementation
func DemonstrateStringerInterface() {
	fmt.Println("\n=== STRINGER INTERFACE ===")
	
	// 1. Custom Stringer implementations
	fmt.Println("\n--- Custom Stringer ---")
	
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	
	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Mouse", Price: 29.99},
		{Name: "Keyboard", Price: 79.99},
	}
	
	fmt.Println("People:")
	for _, person := range people {
		fmt.Printf("  %s\n", person.String())
	}
	
	fmt.Println("Products:")
	for _, product := range products {
		fmt.Printf("  %s\n", product.String())
	}
	
	// 2. Using Stringer in generic functions
	fmt.Println("\n--- Generic Stringer Usage ---")
	
	items := []Stringer{
		Person{Name: "David", Age: 28},
		Product{Name: "Monitor", Price: 299.99},
		Person{Name: "Eve", Age: 32},
		Product{Name: "Headphones", Price: 149.99},
	}
	
	for _, item := range items {
		fmt.Printf("  %s\n", item.String())
	}
}

// DemonstrateInterfaceValues shows interface value behavior
func DemonstrateInterfaceValues() {
	fmt.Println("\n=== INTERFACE VALUES ===")
	
	// 1. Interface value is nil
	fmt.Println("\n--- Nil Interface ---")
	var shape Shape
	fmt.Printf("shape == nil: %t\n", shape == nil)
	fmt.Printf("shape value: %v\n", shape)
	
	// 2. Interface value is not nil, but underlying value is nil
	fmt.Println("\n--- Interface with Nil Underlying Value ---")
	var circle *Circle
	shape = circle
	fmt.Printf("shape == nil: %t\n", shape == nil)
	fmt.Printf("shape value: %v\n", shape)
	
	// 3. Interface value with actual value
	fmt.Println("\n--- Interface with Value ---")
	shape = &Circle{Center: Point{0, 0}, Radius: 5}
	fmt.Printf("shape == nil: %t\n", shape == nil)
	fmt.Printf("shape value: %v\n", shape)
}

// DemonstrateInterfaceSorting shows interface usage in sorting
func DemonstrateInterfaceSorting() {
	fmt.Println("\n=== INTERFACE SORTING ===")
	
	// 1. Sort people by age
	fmt.Println("\n--- Sort People by Age ---")
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "David", Age: 20},
	}
	
	fmt.Println("Before sorting:")
	for _, person := range people {
		fmt.Printf("  %s\n", person.String())
	}
	
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	
	fmt.Println("After sorting by age:")
	for _, person := range people {
		fmt.Printf("  %s\n", person.String())
	}
	
	// 2. Sort products by price
	fmt.Println("\n--- Sort Products by Price ---")
	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Mouse", Price: 29.99},
		{Name: "Keyboard", Price: 79.99},
		{Name: "Monitor", Price: 299.99},
	}
	
	fmt.Println("Before sorting:")
	for _, product := range products {
		fmt.Printf("  %s\n", product.String())
	}
	
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	
	fmt.Println("After sorting by price:")
	for _, product := range products {
		fmt.Printf("  %s\n", product.String())
	}
}

// Concrete implementations

// FileWriter implements Writer
type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing to file %s: %s\n", fw.filename, string(data))
	return len(data), nil
}

// File implements ReadWriteCloser
type File struct {
	data []byte
	pos  int
}

func (f *File) Read(data []byte) (int, error) {
	if f.pos >= len(f.data) {
		return 0, fmt.Errorf("EOF")
	}
	
	n := copy(data, f.data[f.pos:])
	f.pos += n
	return n, nil
}

func (f *File) Write(data []byte) (int, error) {
	f.data = append(f.data, data...)
	return len(data), nil
}

func (f *File) Close() error {
	fmt.Println("File closed")
	return nil
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Center: (%.1f, %.1f), Radius: %.1f}", c.Center.X, c.Center.Y, c.Radius)
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing circle at (%.1f, %.1f) with radius %.1f", c.Center.X, c.Center.Y, c.Radius)
}

// Rectangle methods
func (r Rectangle) Area() float64 {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.BottomRight.Y - r.TopLeft.Y
	return width * height
}

func (r Rectangle) Perimeter() float64 {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.BottomRight.Y - r.TopLeft.Y
	return 2 * (width + height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{TopLeft: (%.1f, %.1f), BottomRight: (%.1f, %.1f)}",
		r.TopLeft.X, r.TopLeft.Y, r.BottomRight.X, r.BottomRight.Y)
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing rectangle from (%.1f, %.1f) to (%.1f, %.1f)",
		r.TopLeft.X, r.TopLeft.Y, r.BottomRight.X, r.BottomRight.Y)
}

// Triangle methods
func (t Triangle) Area() float64 {
	// Using the shoelace formula
	return 0.5 * math.Abs((t.P1.X*(t.P2.Y-t.P3.Y) + t.P2.X*(t.P3.Y-t.P1.Y) + t.P3.X*(t.P1.Y-t.P2.Y)))
}

func (t Triangle) Perimeter() float64 {
	// Calculate distances between points
	d1 := math.Sqrt(math.Pow(t.P2.X-t.P1.X, 2) + math.Pow(t.P2.Y-t.P1.Y, 2))
	d2 := math.Sqrt(math.Pow(t.P3.X-t.P2.X, 2) + math.Pow(t.P3.Y-t.P2.Y, 2))
	d3 := math.Sqrt(math.Pow(t.P1.X-t.P3.X, 2) + math.Pow(t.P1.Y-t.P3.Y, 2))
	return d1 + d2 + d3
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle{P1: (%.1f, %.1f), P2: (%.1f, %.1f), P3: (%.1f, %.1f)}",
		t.P1.X, t.P1.Y, t.P2.X, t.P2.Y, t.P3.X, t.P3.Y)
}

func (t Triangle) Draw() string {
	return fmt.Sprintf("Drawing triangle with points (%.1f, %.1f), (%.1f, %.1f), (%.1f, %.1f)",
		t.P1.X, t.P1.Y, t.P2.X, t.P2.Y, t.P3.X, t.P3.Y)
}

// MovableCircle methods
func (mc MovableCircle) Move(dx, dy float64) {
	mc.Center.X += dx
	mc.Center.Y += dy
}

func (mc MovableCircle) Position() (float64, float64) {
	return mc.Center.X, mc.Center.Y
}

// Person methods
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// Product methods
func (p Product) String() string {
	return fmt.Sprintf("%s - $%.2f", p.Name, p.Price)
}

// RunAllInterfaceExamples runs all interface examples
func RunAllInterfaceExamples() {
	DemonstrateBasicInterfaces()
	DemonstrateShapeInterfaces()
	DemonstrateEmptyInterface()
	DemonstrateInterfaceComposition()
	DemonstrateStringerInterface()
	DemonstrateInterfaceValues()
	DemonstrateInterfaceSorting()
}

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
