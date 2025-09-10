package structs

import (
	"encoding/json"
	"fmt"
	"time"
)

// Person represents a basic struct
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

// Address represents an embedded struct
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

// Employee represents a struct with methods
type Employee struct {
	Person
	ID        int
	Position  string
	Salary    float64
	StartDate time.Time
	IsActive  bool
}

// Rectangle represents a struct for demonstrating methods
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents another shape
type Circle struct {
	Radius float64
}

// Shape interface for polymorphism
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// DemonstrateBasicStructs shows basic struct operations
func DemonstrateBasicStructs() {
	fmt.Println("=== BASIC STRUCTS ===")
	
	// 1. Struct Declaration and Initialization
	fmt.Println("\n--- Struct Declaration ---")
	
	// Zero-initialized struct
	var person1 Person
	fmt.Printf("Zero-initialized person: %+v\n", person1)
	
	// Struct literal
	person2 := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
	}
	fmt.Printf("Initialized person: %+v\n", person2)
	
	// Partial initialization
	person3 := Person{
		Name: "Jane Smith",
		Age:  25,
		// Email and Address will be zero values
	}
	fmt.Printf("Partially initialized person: %+v\n", person3)
	
	// 2. Field Access and Modification
	fmt.Println("\n--- Field Access ---")
	
	person := Person{
		Name:  "Alice Johnson",
		Age:   28,
		Email: "alice@example.com",
	}
	
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	
	// Modify fields
	person.Age = 29
	person.Email = "alice.johnson@example.com"
	fmt.Printf("After modification: %+v\n", person)
	
	// 3. Anonymous Structs
	fmt.Println("\n--- Anonymous Structs ---")
	
	// Inline struct definition
	config := struct {
		DatabaseURL string
		Port        int
		Debug       bool
	}{
		DatabaseURL: "localhost:5432",
		Port:        8080,
		Debug:       true,
	}
	
	fmt.Printf("Config: %+v\n", config)
}

// DemonstrateStructMethods shows struct methods
func DemonstrateStructMethods() {
	fmt.Println("\n=== STRUCT METHODS ===")
	
	// 1. Value Receiver Methods
	fmt.Println("\n--- Value Receiver Methods ---")
	
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("String: %s\n", rect.String())
	
	// 2. Pointer Receiver Methods
	fmt.Println("\n--- Pointer Receiver Methods ---")
	
	employee := Employee{
		Person: Person{
			Name:  "Bob Wilson",
			Age:   35,
			Email: "bob@company.com",
		},
		ID:        1001,
		Position:  "Software Engineer",
		Salary:    75000,
		StartDate: time.Now().AddDate(-2, 0, 0), // 2 years ago
		IsActive:  true,
	}
	
	fmt.Printf("Employee before promotion: %+v\n", employee)
	employee.Promote("Senior Software Engineer", 85000)
	fmt.Printf("Employee after promotion: %+v\n", employee)
	
	// 3. Method Chaining
	fmt.Println("\n--- Method Chaining ---")
	
	rect2 := Rectangle{Width: 3, Height: 4}
	rect2.Scale(2).Scale(1.5)
	fmt.Printf("Scaled rectangle: %+v\n", rect2)
}

// DemonstrateStructEmbedding shows struct embedding (composition)
func DemonstrateStructEmbedding() {
	fmt.Println("\n=== STRUCT EMBEDDING ===")
	
	// 1. Basic Embedding
	fmt.Println("\n--- Basic Embedding ---")
	
	employee := Employee{
		Person: Person{
			Name:  "Charlie Brown",
			Age:   40,
			Email: "charlie@company.com",
			Address: Address{
				Street:  "456 Oak Ave",
				City:    "San Francisco",
				Country: "USA",
				ZipCode: "94102",
			},
		},
		ID:        1002,
		Position:  "Manager",
		Salary:    95000,
		StartDate: time.Now().AddDate(-5, 0, 0),
		IsActive:  true,
	}
	
	// Access embedded struct fields directly
	fmt.Printf("Employee Name: %s\n", employee.Name)        // From Person
	fmt.Printf("Employee Age: %d\n", employee.Age)          // From Person
	fmt.Printf("Employee ID: %d\n", employee.ID)            // From Employee
	fmt.Printf("Employee Position: %s\n", employee.Position) // From Employee
	
	// Access embedded struct fields explicitly
	fmt.Printf("Employee City: %s\n", employee.Person.Address.City)
	
	// 2. Method Promotion
	fmt.Println("\n--- Method Promotion ---")
	
	// Methods from embedded structs are promoted
	fmt.Printf("Employee Info: %s\n", employee.GetInfo())
	
	// 3. Field Shadowing
	fmt.Println("\n--- Field Shadowing ---")
	
	type Manager struct {
		Person
		Name string // This shadows Person.Name
	}
	
	manager := Manager{
		Person: Person{Name: "John Manager"},
		Name:   "Jane Manager",
	}
	
	fmt.Printf("Manager.Person.Name: %s\n", manager.Person.Name)
	fmt.Printf("Manager.Name: %s\n", manager.Name)
}

// DemonstrateStructTags shows struct tags for metadata
func DemonstrateStructTags() {
	fmt.Println("\n=== STRUCT TAGS ===")
	
	// Struct with JSON tags
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email,omitempty"`
		Password string `json:"-"` // Ignore in JSON
		Age      int    `json:"age,string"` // Convert to string in JSON
	}
	
	user := User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "secret123",
		Age:      30,
	}
	
	// Convert to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	
	fmt.Printf("User as JSON:\n%s\n", string(jsonData))
	
	// Convert from JSON
	jsonString := `{"id":2,"username":"janedoe","email":"jane@example.com","age":"25"}`
	var user2 User
	err = json.Unmarshal([]byte(jsonString), &user2)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	
	fmt.Printf("User from JSON: %+v\n", user2)
}

// DemonstrateStructComparison shows struct comparison
func DemonstrateStructComparison() {
	fmt.Println("\n=== STRUCT COMPARISON ===")
	
	// Structs with comparable fields can be compared
	point1 := struct {
		X, Y int
	}{1, 2}
	
	point2 := struct {
		X, Y int
	}{1, 2}
	
	point3 := struct {
		X, Y int
	}{2, 3}
	
	fmt.Printf("point1 == point2: %t\n", point1 == point2)
	fmt.Printf("point1 == point3: %t\n", point1 == point3)
	
	// Structs with non-comparable fields cannot be compared
	// (This would cause a compile error)
	// person1 := Person{Name: "John"}
	// person2 := Person{Name: "John"}
	// fmt.Printf("person1 == person2: %t\n", person1 == person2) // Error!
}

// Rectangle methods (Value receivers)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		Width:  r.Width * factor,
		Height: r.Height * factor,
	}
}

// Circle methods
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Employee methods (Pointer receivers)
func (e *Employee) Promote(newPosition string, newSalary float64) {
	e.Position = newPosition
	e.Salary = newSalary
}

func (e *Employee) GetInfo() string {
	return fmt.Sprintf("%s (%s) - %s", e.Name, e.Position, e.Email)
}

// Person methods
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d years old", p.Name, p.Age)
}

// DemonstratePolymorphism shows interface implementation
func DemonstratePolymorphism() {
	fmt.Println("\n=== POLYMORPHISM WITH INTERFACES ===")
	
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 4},
		Circle{Radius: 2.5},
	}
	
	for i, shape := range shapes {
		fmt.Printf("Shape %d: %s\n", i+1, shape.String())
		fmt.Printf("  Area: %.2f\n", shape.Area())
		fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}

// RunAllStructExamples runs all struct examples
func RunAllStructExamples() {
	DemonstrateBasicStructs()
	DemonstrateStructMethods()
	DemonstrateStructEmbedding()
	DemonstrateStructTags()
	DemonstrateStructComparison()
	DemonstratePolymorphism()
}
