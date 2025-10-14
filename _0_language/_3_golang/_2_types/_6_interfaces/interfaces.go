package interfaces

import (
	"fmt"
	"math"
	"sort"
)

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
