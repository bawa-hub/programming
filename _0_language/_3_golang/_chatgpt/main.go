package main

import "fmt"

func main() {
	fmt.Println("Hello, Golang!")

	// 1. Declaring Variables in Go
	// Using var (Explicit Type)
	// var name string = "Alice"
	// var age int = 25
	// var height float64 = 5.9
	// var isStudent bool = false

	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Height:", height)
	// fmt.Println("Is Student:", isStudent)

	// Using := (Type Inference)
	// name := "Alice"  // string
	// age := 25        // int
	// height := 5.9    // float64
	// isStudent := false // bool

	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Height:", height)
	// fmt.Println("Is Student:", isStudent)

	// Rule: := can only be used inside functions, not at the package level.

	// 2. Constants in Go
	// const PI = 3.1415
	// fmt.Println("Value of PI:", PI)
	// Rule: You cannot use := for constants.

	// 3. Basic Data Types in Go
	// 	Data Type	Example
	// int	42
	// float64	3.14
	// string	"Hello"
	// bool	true, false

	// Type Conversion
	// Go does not allow implicit type conversion. You have to do it explicitly:
	// var num int = 10
    // var price float64 = float64(num) // Convert int to float64
    // fmt.Println("Converted price:", price)
}
