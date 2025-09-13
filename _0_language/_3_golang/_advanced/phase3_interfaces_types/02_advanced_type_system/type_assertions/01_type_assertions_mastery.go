package main

import (
	"fmt"
)

// 🔧 TYPE ASSERTIONS MASTERY
// Understanding type assertions, type switches, and advanced type patterns

func main() {
	fmt.Println("🔧 TYPE ASSERTIONS MASTERY")
	fmt.Println("==========================")

	// 1. Basic Type Assertions
	fmt.Println("\n1. Basic Type Assertions:")
	basicTypeAssertions()

	// 2. Type Switches
	fmt.Println("\n2. Type Switches:")
	typeSwitches()

	// 3. Interface Type Assertions
	fmt.Println("\n3. Interface Type Assertions:")
	interfaceTypeAssertions()

	// 4. Type Assertions with Error Handling
	fmt.Println("\n4. Type Assertions with Error Handling:")
	typeAssertionsWithErrorHandling()

	// 5. Advanced Type Patterns
	fmt.Println("\n5. Advanced Type Patterns:")
	advancedTypePatterns()

	// 6. Type Assertions Best Practices
	fmt.Println("\n6. Type Assertions Best Practices:")
	typeAssertionsBestPractices()
}

// BASIC TYPE ASSERTIONS: Understanding type assertions
func basicTypeAssertions() {
	fmt.Println("Understanding basic type assertions...")
	
	// Create interface value
	var i interface{} = 42
	
	// Type assertion
	if v, ok := i.(int); ok {
		fmt.Printf("  📊 Value is int: %d\n", v)
	} else {
		fmt.Println("  ❌ Value is not int")
	}
	
	// Type assertion without ok check (will panic if wrong type)
	// v := i.(int) // This would panic if i is not int
	
	// Multiple type assertions
	var values []interface{} = []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
	}
	
	for _, val := range values {
		switch v := val.(type) {
		case int:
			fmt.Printf("  📊 Integer: %d\n", v)
		case string:
			fmt.Printf("  📊 String: %s\n", v)
		case float64:
			fmt.Printf("  📊 Float: %.2f\n", v)
		case bool:
			fmt.Printf("  📊 Boolean: %t\n", v)
		case []int:
			fmt.Printf("  📊 Slice: %v\n", v)
		default:
			fmt.Printf("  📊 Unknown type: %T\n", v)
		}
	}
}

// TYPE SWITCHES: Understanding type switches
func typeSwitches() {
	fmt.Println("Understanding type switches...")
	
	// Type switch with interface{}
	var i interface{} = "hello world"
	
	switch v := i.(type) {
	case int:
		fmt.Printf("  📊 Integer: %d\n", v)
	case string:
		fmt.Printf("  📊 String: %s (length: %d)\n", v, len(v))
	case float64:
		fmt.Printf("  📊 Float: %.2f\n", v)
	case bool:
		fmt.Printf("  📊 Boolean: %t\n", v)
	default:
		fmt.Printf("  📊 Unknown type: %T\n", v)
	}
	
	// Type switch with custom types
	var shapes []Shape = []Shape{
		&Circle{Radius: 5.0},
		&Rectangle{Width: 10, Height: 8},
		&Triangle{Base: 6, Height: 4},
	}
	
	for _, shape := range shapes {
		switch s := shape.(type) {
		case *Circle:
			fmt.Printf("  🔵 Circle: radius=%.1f, area=%.2f\n", s.Radius, s.Area())
		case *Rectangle:
			fmt.Printf("  🔲 Rectangle: %dx%d, area=%.2f\n", s.Width, s.Height, s.Area())
		case *Triangle:
			fmt.Printf("  🔺 Triangle: base=%.1f, height=%.1f, area=%.2f\n", s.Base, s.Height, s.Area())
		default:
			fmt.Printf("  ❓ Unknown shape: %T\n", s)
		}
	}
}

// INTERFACE TYPE ASSERTIONS: Understanding interface type assertions
func interfaceTypeAssertions() {
	fmt.Println("Understanding interface type assertions...")
	
	// Create interface values
	var reader Reader = &FileReader{name: "test.txt"}
	var writer Writer = &FileWriter{name: "output.txt"}
	
	// Type assertion to specific interface
	if r, ok := reader.(Reader); ok {
		data, _ := r.Read()
		fmt.Printf("  📖 Reader interface: %s\n", data)
	}
	
	// Type assertion to concrete type
	if fr, ok := reader.(*FileReader); ok {
		fmt.Printf("  📁 FileReader concrete: %s\n", fr.name)
	}
	
	// Check if interface implements another interface
	if rw, ok := reader.(ReadWriter); ok {
		fmt.Printf("  📖✏️  ReadWriter interface: %T\n", rw)
	} else {
		fmt.Printf("  ❌ Not a ReadWriter\n")
	}
	
	// Multiple interface assertions
	interfaces := []interface{}{
		reader,
		writer,
		&FileReadWriter{Reader: reader, Writer: writer},
	}
	
	for _, iface := range interfaces {
		switch v := iface.(type) {
		case Reader:
			data, _ := v.Read()
			fmt.Printf("  📖 Reader: %s\n", data)
		case Writer:
			v.Write("test data")
			fmt.Println("  ✏️  Writer: data written")
		case ReadWriter:
			v.Write("composed data")
			data, _ := v.Read()
			fmt.Printf("  📖✏️  ReadWriter: %s\n", data)
		default:
			fmt.Printf("  ❓ Unknown interface: %T\n", v)
		}
	}
}

// TYPE ASSERTIONS WITH ERROR HANDLING: Safe type assertions
func typeAssertionsWithErrorHandling() {
	fmt.Println("Understanding type assertions with error handling...")
	
	// Safe type assertion with error handling
	var i interface{} = "hello"
	
	if v, ok := i.(string); ok {
		fmt.Printf("  ✅ Safe assertion: %s\n", v)
	} else {
		fmt.Printf("  ❌ Type assertion failed\n")
	}
	
	// Function that returns interface and error
	result, err := processValue(42)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		fmt.Printf("  ✅ Result: %v\n", result)
	}
	
	// Type assertion with multiple types
	var values []interface{} = []interface{}{
		"string",
		42,
		3.14,
		[]int{1, 2, 3},
	}
	
	for _, val := range values {
		if result, err := safeTypeAssertion(val); err != nil {
			fmt.Printf("  ❌ Error: %v\n", err)
		} else {
			fmt.Printf("  ✅ Result: %v\n", result)
		}
	}
}

// ADVANCED TYPE PATTERNS: Advanced type patterns
func advancedTypePatterns() {
	fmt.Println("Understanding advanced type patterns...")
	
	// Pattern 1: Type assertion with reflection
	var i interface{} = map[string]int{"a": 1, "b": 2}
	
	if m, ok := i.(map[string]int); ok {
		fmt.Printf("  📊 Map: %v\n", m)
	}
	
	// Pattern 2: Type assertion with function types
	var fn interface{} = func(x int) int { return x * 2 }
	
	if f, ok := fn.(func(int) int); ok {
		result := f(5)
		fmt.Printf("  🔧 Function: %d\n", result)
	}
	
	// Pattern 3: Type assertion with channels
	var ch interface{} = make(chan int, 1)
	
	if c, ok := ch.(chan int); ok {
		c <- 42
		value := <-c
		fmt.Printf("  📡 Channel: %d\n", value)
	}
	
	// Pattern 4: Type assertion with slices
	var slice interface{} = []string{"a", "b", "c"}
	
	if s, ok := slice.([]string); ok {
		fmt.Printf("  📊 Slice: %v\n", s)
	}
	
	// Pattern 5: Type assertion with structs
	var person interface{} = Person{Name: "John", Age: 30}
	
	if p, ok := person.(Person); ok {
		fmt.Printf("  👤 Person: %s, %d\n", p.Name, p.Age)
	}
}

// TYPE ASSERTIONS BEST PRACTICES: Following best practices
func typeAssertionsBestPractices() {
	fmt.Println("Understanding type assertions best practices...")
	
	// 1. Always use the two-value form
	fmt.Println("  📝 Best Practice 1: Always use the two-value form")
	twoValueForm()
	
	// 2. Use type switches for multiple types
	fmt.Println("  📝 Best Practice 2: Use type switches for multiple types")
	typeSwitchBestPractice()
	
	// 3. Avoid type assertions when possible
	fmt.Println("  📝 Best Practice 3: Avoid type assertions when possible")
	avoidTypeAssertions()
	
	// 4. Use interfaces instead of concrete types
	fmt.Println("  📝 Best Practice 4: Use interfaces instead of concrete types")
	useInterfaces()
}

func twoValueForm() {
	var i interface{} = "hello"
	
	// Good: Two-value form
	if v, ok := i.(string); ok {
		fmt.Printf("    ✅ Safe: %s\n", v)
	}
	
	// Bad: Single-value form (can panic)
	// v := i.(string) // This can panic
}

func typeSwitchBestPractice() {
	var i interface{} = 42
	
	// Good: Type switch
	switch v := i.(type) {
	case int:
		fmt.Printf("    ✅ Integer: %d\n", v)
	case string:
		fmt.Printf("    ✅ String: %s\n", v)
	default:
		fmt.Printf("    ✅ Other: %T\n", v)
	}
}

func avoidTypeAssertions() {
	// Good: Use interfaces
	var reader Reader = &FileReader{name: "test.txt"}
	data, _ := reader.Read()
	fmt.Printf("    ✅ Interface: %s\n", data)
	
	// Bad: Type assertion when interface would work
	// if fr, ok := reader.(*FileReader); ok {
	//     data, _ := fr.Read()
	// }
}

func useInterfaces() {
	// Good: Accept interfaces
	processReader(&FileReader{name: "test.txt"})
	processReader(&NetworkReader{url: "https://api.example.com"})
}

func processReader(reader Reader) {
	data, _ := reader.Read()
	fmt.Printf("    ✅ Processed: %s\n", data)
}

// HELPER FUNCTIONS

func processValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case int:
		return v * 2, nil
	case string:
		return len(v), nil
	case float64:
		return v * 3.14, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

func safeTypeAssertion(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("String: %s", v), nil
	case int:
		return fmt.Sprintf("Integer: %d", v), nil
	case float64:
		return fmt.Sprintf("Float: %.2f", v), nil
	case []int:
		return fmt.Sprintf("Slice: %v", v), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

// INTERFACE DEFINITIONS

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(data string) error
}

type ReadWriter interface {
	Reader
	Writer
}

// SHAPE INTERFACES

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

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

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t *Triangle) Perimeter() float64 {
	// Simplified for demo
	return t.Base + t.Height + t.Base
}

// IMPLEMENTATIONS

type FileReader struct {
	name string
}

func (f *FileReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", f.name), nil
}

type FileWriter struct {
	name string
}

func (f *FileWriter) Write(data string) error {
	fmt.Printf("Writing to %s: %s\n", f.name, data)
	return nil
}

type FileReadWriter struct {
	Reader
	Writer
}

type NetworkReader struct {
	url string
}

func (n *NetworkReader) Read() (string, error) {
	return fmt.Sprintf("Data from %s", n.url), nil
}

// DATA STRUCTURES

type Person struct {
	Name string
	Age  int
}
