package main

import (
	"fmt"
	"reflect"
	"time"
)

// Custom types for demonstration
type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, City: %s}", p.Name, p.Age, p.City)
}

type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {
	fmt.Println("🚀 Go fmt Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic Print Functions
	fmt.Println("\n1. Basic Print Functions:")
	fmt.Print("Hello ")           // No newline
	fmt.Print("World")            // No newline
	fmt.Println()                 // Just newline
	fmt.Println("Hello World")    // With newline
	fmt.Printf("Formatted: %s\n", "Go") // Formatted

	// 2. String Formatting
	fmt.Println("\n2. String Formatting:")
	name := "Alice"
	age := 30
	height := 5.6
	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
	fmt.Printf("Name: %q, Age: %d, Height: %g\n", name, age, height) // Quoted string

	// 3. Number Formatting
	fmt.Println("\n3. Number Formatting:")
	number := 42
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Octal: %o\n", number)
	fmt.Printf("Hex (lower): %x\n", number)
	fmt.Printf("Hex (upper): %X\n", number)
	fmt.Printf("Width 5: %5d\n", number)
	fmt.Printf("Zero-padded: %05d\n", number)

	// 4. Float Formatting
	fmt.Println("\n4. Float Formatting:")
	pi := 3.14159265359
	fmt.Printf("Default: %f\n", pi)
	fmt.Printf("Precision 2: %.2f\n", pi)
	fmt.Printf("Width 8, precision 2: %8.2f\n", pi)
	fmt.Printf("Scientific: %e\n", pi)
	fmt.Printf("Scientific (upper): %E\n", pi)
	fmt.Printf("Auto format: %g\n", pi)

	// 5. Boolean Formatting
	fmt.Println("\n5. Boolean Formatting:")
	flag := true
	fmt.Printf("Boolean: %t\n", flag)
	fmt.Printf("Boolean (quoted): %q\n", flag)

	// 6. Pointer Formatting
	fmt.Println("\n6. Pointer Formatting:")
	ptr := &number
	fmt.Printf("Pointer: %p\n", ptr)
	fmt.Printf("Pointer value: %v\n", ptr)

	// 7. Type Formatting
	fmt.Println("\n7. Type Formatting:")
	fmt.Printf("Type of number: %T\n", number)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of flag: %T\n", flag)

	// 8. Struct Formatting
	fmt.Println("\n8. Struct Formatting:")
	person := Person{Name: "Bob", Age: 25, City: "New York"}
	fmt.Printf("Default: %v\n", person)
	fmt.Printf("With field names: %+v\n", person)
	fmt.Printf("Go syntax: %#v\n", person)
	fmt.Printf("String method: %s\n", person)

	// 9. Slice and Array Formatting
	fmt.Println("\n9. Slice and Array Formatting:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("Slice (quoted): %q\n", numbers)
	fmt.Printf("Slice (Go syntax): %#v\n", numbers)

	// 10. Map Formatting
	fmt.Println("\n10. Map Formatting:")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Printf("Map: %v\n", colors)
	fmt.Printf("Map (Go syntax): %#v\n", colors)

	// 11. Error Formatting
	fmt.Println("\n11. Error Formatting:")
	err := CustomError{Code: 404, Message: "Not Found"}
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Error (quoted): %q\n", err)

	// 12. Time Formatting
	fmt.Println("\n12. Time Formatting:")
	now := time.Now()
	fmt.Printf("Time: %v\n", now)
	fmt.Printf("Time (RFC3339): %s\n", now.Format(time.RFC3339))
	fmt.Printf("Time (custom): %s\n", now.Format("2006-01-02 15:04:05"))

	// 13. Sprint Functions (String formatting)
	fmt.Println("\n13. Sprint Functions:")
	formatted := fmt.Sprintf("Hello %s, you are %d years old", name, age)
	fmt.Println("Sprintf result:", formatted)
	
	formatted2 := fmt.Sprint("Numbers: ", numbers)
	fmt.Println("Sprint result:", formatted2)

	// 14. Scan Functions (Input)
	fmt.Println("\n14. Scan Functions:")
	fmt.Print("Enter your name: ")
	var inputName string
	fmt.Scanln(&inputName)
	fmt.Printf("Hello, %s!\n", inputName)

	// 15. Advanced Formatting
	fmt.Println("\n15. Advanced Formatting:")
	
	// Width and precision
	fmt.Printf("|%10s|%10d|%10.2f|\n", "Name", 42, 3.14)
	fmt.Printf("|%-10s|%-10d|%-10.2f|\n", "Name", 42, 3.14) // Left-aligned
	
	// Zero padding
	fmt.Printf("ID: %05d\n", 123)
	
	// Space padding
	fmt.Printf("Score: %5d\n", 95)
	
	// Plus sign for positive numbers
	fmt.Printf("Temperature: %+d°C\n", 25)
	fmt.Printf("Temperature: %+d°C\n", -5)

	// 16. Reflection with fmt
	fmt.Println("\n16. Reflection with fmt:")
	value := reflect.ValueOf(person)
	fmt.Printf("Reflection: %v\n", value)
	fmt.Printf("Reflection type: %T\n", value)

	// 17. Custom Formatting
	fmt.Println("\n17. Custom Formatting:")
	fmt.Printf("Custom string: %s\n", person)
	fmt.Printf("Custom error: %v\n", err)

	// 18. Formatting with different bases
	fmt.Println("\n18. Different Number Bases:")
	num := 255
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Octal: %o\n", num)
	fmt.Printf("Hex: %x\n", num)
	fmt.Printf("Hex (upper): %X\n", num)

	// 19. String and Rune Formatting
	fmt.Println("\n19. String and Rune Formatting:")
	text := "Hello, 世界"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("String (quoted): %q\n", text)
	fmt.Printf("String (Go syntax): %#v\n", text)
	
	for i, r := range text {
		fmt.Printf("Rune %d: %c (U+%04X)\n", i, r, r)
	}

	// 20. Practical Examples
	fmt.Println("\n20. Practical Examples:")
	
	// Table formatting
	fmt.Println("| Name     | Age | City      |")
	fmt.Println("|----------|-----|-----------|")
	fmt.Printf("| %-8s | %3d | %-9s |\n", "Alice", 30, "New York")
	fmt.Printf("| %-8s | %3d | %-9s |\n", "Bob", 25, "London")
	
	// Progress bar simulation
	fmt.Print("Progress: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("\rProgress: [%s%s] %d%%", 
			repeatString("=", i), 
			repeatString(" ", 10-i), 
			i*10)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	// File size formatting
	fileSizes := []int64{1024, 1024*1024, 1024*1024*1024}
	for _, size := range fileSizes {
		fmt.Printf("File size: %d bytes (%.2f KB)\n", size, float64(size)/1024)
	}

	fmt.Println("\n🎉 fmt Package Mastery Complete!")
}

// Helper function for string repetition
func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
