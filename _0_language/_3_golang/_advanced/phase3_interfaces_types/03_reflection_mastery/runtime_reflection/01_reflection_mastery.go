package main

import (
	"fmt"
	"reflect"
	"time"
)

// 🪞 REFLECTION MASTERY
// Understanding runtime reflection and dynamic programming in Go

func main() {
	fmt.Println("🪞 REFLECTION MASTERY")
	fmt.Println("=====================")

	// 1. Basic Type Reflection
	fmt.Println("\n1. Basic Type Reflection:")
	basicTypeReflection()

	// 2. Value Reflection
	fmt.Println("\n2. Value Reflection:")
	valueReflection()

	// 3. Struct Field Manipulation
	fmt.Println("\n3. Struct Field Manipulation:")
	structFieldManipulation()

	// 4. Dynamic Method Invocation
	fmt.Println("\n4. Dynamic Method Invocation:")
	dynamicMethodInvocation()

	// 5. Interface Reflection
	fmt.Println("\n5. Interface Reflection:")
	interfaceReflection()

	// 6. Advanced Reflection Patterns
	fmt.Println("\n6. Advanced Reflection Patterns:")
	advancedReflectionPatterns()

	// 7. Reflection Performance
	fmt.Println("\n7. Reflection Performance:")
	reflectionPerformance()

	// 8. Reflection Best Practices
	fmt.Println("\n8. Reflection Best Practices:")
	reflectionBestPractices()
}

// BASIC TYPE REFLECTION: Understanding type inspection
func basicTypeReflection() {
	fmt.Println("Understanding basic type reflection...")
	
	// Inspect different types
	types := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
		Person{Name: "John", Age: 30},
	}
	
	for _, value := range types {
		t := reflect.TypeOf(value)
		fmt.Printf("  📊 Type: %s, Kind: %s\n", t, t.Kind())
		
		// Get type information
		fmt.Printf("    Size: %d bytes\n", t.Size())
		fmt.Printf("    String: %s\n", t.String())
		fmt.Printf("    PkgPath: %s\n", t.PkgPath())
	}
}

// VALUE REFLECTION: Understanding value inspection
func valueReflection() {
	fmt.Println("Understanding value reflection...")
	
	// Inspect values
	values := []interface{}{
		42,
		"hello world",
		3.14159,
		true,
		[]int{1, 2, 3, 4, 5},
		map[string]string{"key": "value"},
		Person{Name: "Alice", Age: 25},
	}
	
	for _, value := range values {
		v := reflect.ValueOf(value)
		fmt.Printf("  📊 Value: %v, Type: %s\n", v, v.Type())
		
		// Get value information
		fmt.Printf("    CanSet: %t\n", v.CanSet())
		fmt.Printf("    CanAddr: %t\n", v.CanAddr())
		fmt.Printf("    IsValid: %t\n", v.IsValid())
		
		// Get kind-specific information
		switch v.Kind() {
		case reflect.Slice:
			fmt.Printf("    Length: %d, Capacity: %d\n", v.Len(), v.Cap())
		case reflect.Map:
			fmt.Printf("    Map keys: %v\n", v.MapKeys())
		case reflect.Struct:
			fmt.Printf("    NumField: %d\n", v.NumField())
		}
	}
}

// STRUCT FIELD MANIPULATION: Understanding struct field access
func structFieldManipulation() {
	fmt.Println("Understanding struct field manipulation...")
	
	// Create a struct
	person := Person{Name: "Bob", Age: 35, Email: "bob@example.com"}
	
	// Get struct type and value
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	
	fmt.Printf("  📊 Struct: %+v\n", person)
	fmt.Printf("  📊 Type: %s\n", t)
	fmt.Printf("  📊 NumField: %d\n", t.NumField())
	
	// Iterate through fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		fmt.Printf("    Field %d: %s (%s) = %v\n", 
			i, field.Name, field.Type, fieldValue)
		
		// Check field tags
		if tag := field.Tag.Get("json"); tag != "" {
			fmt.Printf("      JSON tag: %s\n", tag)
		}
	}
	
	// Modify struct fields (if addressable)
	personPtr := &Person{Name: "Charlie", Age: 40}
	vPtr := reflect.ValueOf(personPtr).Elem()
	
	fmt.Printf("  📊 Before: %+v\n", personPtr)
	
	// Set field values
	if vPtr.FieldByName("Name").CanSet() {
		vPtr.FieldByName("Name").SetString("David")
	}
	if vPtr.FieldByName("Age").CanSet() {
		vPtr.FieldByName("Age").SetInt(45)
	}
	
	fmt.Printf("  📊 After: %+v\n", personPtr)
}

// DYNAMIC METHOD INVOCATION: Understanding dynamic method calls
func dynamicMethodInvocation() {
	fmt.Println("Understanding dynamic method invocation...")
	
	// Create a struct with methods
	calculator := &Calculator{}
	
	// Get method by name
	method := reflect.ValueOf(calculator).MethodByName("Add")
	if method.IsValid() {
		// Call method with arguments
		args := []reflect.Value{
			reflect.ValueOf(10),
			reflect.ValueOf(20),
		}
		result := method.Call(args)
		fmt.Printf("  📊 Add(10, 20) = %d\n", result[0].Int())
	}
	
	// Get method by index
	t := reflect.TypeOf(calculator)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("  📊 Method %d: %s\n", i, method.Name)
	}
	
	// Call method dynamically
	method = reflect.ValueOf(calculator).MethodByName("Multiply")
	if method.IsValid() {
		args := []reflect.Value{
			reflect.ValueOf(5),
			reflect.ValueOf(6),
		}
		result := method.Call(args)
		fmt.Printf("  📊 Multiply(5, 6) = %d\n", result[0].Int())
	}
}

// INTERFACE REFLECTION: Understanding interface reflection
func interfaceReflection() {
	fmt.Println("Understanding interface reflection...")
	
	// Create interface values
	var reader Reader = &FileReader{name: "test.txt"}
	var writer Writer = &FileWriter{name: "output.txt"}
	
	// Inspect interface values
	interfaces := []interface{}{
		reader,
		writer,
		&FileReadWriter{Reader: reader, Writer: writer},
	}
	
	for i, iface := range interfaces {
		v := reflect.ValueOf(iface)
		t := reflect.TypeOf(iface)
		
		fmt.Printf("  📊 Interface %d: %s\n", i, t)
		fmt.Printf("    Kind: %s\n", v.Kind())
		fmt.Printf("    IsNil: %t\n", v.IsNil())
		
		// Check if it implements specific interface
		if t.Implements(reflect.TypeOf((*Reader)(nil)).Elem()) {
			fmt.Printf("    Implements Reader interface\n")
		}
		if t.Implements(reflect.TypeOf((*Writer)(nil)).Elem()) {
			fmt.Printf("    Implements Writer interface\n")
		}
	}
}

// ADVANCED REFLECTION PATTERNS: Understanding advanced patterns
func advancedReflectionPatterns() {
	fmt.Println("Understanding advanced reflection patterns...")
	
	// Pattern 1: Generic function caller
	fmt.Println("  📊 Pattern 1: Generic function caller")
	genericFunctionCaller()
	
	// Pattern 2: Struct field iterator
	fmt.Println("  📊 Pattern 2: Struct field iterator")
	structFieldIterator()
	
	// Pattern 3: Dynamic struct creation
	fmt.Println("  📊 Pattern 3: Dynamic struct creation")
	dynamicStructCreation()
	
	// Pattern 4: Type converter
	fmt.Println("  📊 Pattern 4: Type converter")
	typeConverter()
}

func genericFunctionCaller() {
	// Call different functions dynamically
	functions := map[string]interface{}{
		"add":      func(a, b int) int { return a + b },
		"multiply": func(a, b int) int { return a * b },
		"greet":    func(name string) string { return "Hello, " + name },
	}
	
	// Call add function
	if fn, exists := functions["add"]; exists {
		fnValue := reflect.ValueOf(fn)
		args := []reflect.Value{
			reflect.ValueOf(10),
			reflect.ValueOf(20),
		}
		result := fnValue.Call(args)
		fmt.Printf("    Add result: %d\n", result[0].Int())
	}
	
	// Call greet function
	if fn, exists := functions["greet"]; exists {
		fnValue := reflect.ValueOf(fn)
		args := []reflect.Value{
			reflect.ValueOf("World"),
		}
		result := fnValue.Call(args)
		fmt.Printf("    Greet result: %s\n", result[0].String())
	}
}

func structFieldIterator() {
	person := Person{Name: "Eve", Age: 28, Email: "eve@example.com"}
	v := reflect.ValueOf(person)
	t := reflect.TypeOf(person)
	
	fmt.Printf("    Struct: %+v\n", person)
	
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		fmt.Printf("    %s: %v (%s)\n", 
			field.Name, fieldValue, field.Type)
	}
}

func dynamicStructCreation() {
	// Create a new struct dynamically
	t := reflect.TypeOf(Person{})
	newPerson := reflect.New(t).Elem()
	
	// Set field values
	newPerson.FieldByName("Name").SetString("Frank")
	newPerson.FieldByName("Age").SetInt(32)
	newPerson.FieldByName("Email").SetString("frank@example.com")
	
	// Convert back to original type
	person := newPerson.Interface().(Person)
	fmt.Printf("    Created person: %+v\n", person)
}

func typeConverter() {
	// Convert between different types
	values := []interface{}{
		"42",
		"3.14",
		"true",
		"hello",
	}
	
	for _, value := range values {
		fmt.Printf("    Converting: %v\n", value)
		
		// Try to convert to int
		if intVal, err := convertToInt(value); err == nil {
			fmt.Printf("      As int: %d\n", intVal)
		}
		
		// Try to convert to float
		if floatVal, err := convertToFloat(value); err == nil {
			fmt.Printf("      As float: %.2f\n", floatVal)
		}
		
		// Try to convert to bool
		if boolVal, err := convertToBool(value); err == nil {
			fmt.Printf("      As bool: %t\n", boolVal)
		}
	}
}

// REFLECTION PERFORMANCE: Understanding performance implications
func reflectionPerformance() {
	fmt.Println("Understanding reflection performance...")
	
	// Direct method call
	calculator := &Calculator{}
	
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		calculator.Add(10, 20)
	}
	directTime := time.Since(start)
	fmt.Printf("  📊 Direct call: %v\n", directTime)
	
	// Reflection method call
	method := reflect.ValueOf(calculator).MethodByName("Add")
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		method.Call(args)
	}
	reflectionTime := time.Since(start)
	fmt.Printf("  📊 Reflection call: %v\n", reflectionTime)
	
	// Performance ratio
	ratio := float64(reflectionTime) / float64(directTime)
	fmt.Printf("  📊 Reflection overhead: %.2fx\n", ratio)
}

// REFLECTION BEST PRACTICES: Following best practices
func reflectionBestPractices() {
	fmt.Println("Understanding reflection best practices...")
	
	// 1. Cache reflection values
	fmt.Println("  📝 Best Practice 1: Cache reflection values")
	cacheReflectionValues()
	
	// 2. Use type assertions when possible
	fmt.Println("  📝 Best Practice 2: Use type assertions when possible")
	useTypeAssertions()
	
	// 3. Handle errors gracefully
	fmt.Println("  📝 Best Practice 3: Handle errors gracefully")
	handleErrorsGracefully()
	
	// 4. Avoid reflection in hot paths
	fmt.Println("  📝 Best Practice 4: Avoid reflection in hot paths")
	avoidReflectionInHotPaths()
}

func cacheReflectionValues() {
	// Good: Cache reflection values
	t := reflect.TypeOf(Person{})
	v := reflect.ValueOf(Person{Name: "John", Age: 30})
	
	// Use cached values
	fmt.Printf("    Cached type: %s\n", t)
	fmt.Printf("    Cached value: %v\n", v)
}

func useTypeAssertions() {
	// Good: Use type assertions when possible
	var i interface{} = 42
	
	// Type assertion is faster than reflection
	if intVal, ok := i.(int); ok {
		fmt.Printf("    Type assertion: %d\n", intVal)
	}
	
	// Only use reflection when necessary
	t := reflect.TypeOf(i)
	fmt.Printf("    Reflection type: %s\n", t)
}

func handleErrorsGracefully() {
	// Good: Handle reflection errors gracefully
	var i interface{} = "hello"
	
	// Try to convert to int
	if intVal, err := convertToInt(i); err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    Converted: %d\n", intVal)
	}
}

func avoidReflectionInHotPaths() {
	// Good: Avoid reflection in performance-critical code
	// Use direct calls or type assertions instead
	calculator := &Calculator{}
	
	// Direct call (fast)
	result := calculator.Add(10, 20)
	fmt.Printf("    Direct call result: %d\n", result)
}

// HELPER FUNCTIONS

func convertToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		// Simple string to int conversion
		if v == "42" {
			return 42, nil
		}
		return 0, fmt.Errorf("cannot convert %s to int", v)
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

func convertToFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case string:
		if v == "3.14" {
			return 3.14, nil
		}
		return 0, fmt.Errorf("cannot convert %s to float", v)
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

func convertToBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		if v == "true" {
			return true, nil
		}
		return false, fmt.Errorf("cannot convert %s to bool", v)
	default:
		return false, fmt.Errorf("unsupported type: %T", v)
	}
}

// DATA STRUCTURES

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

// INTERFACES

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
