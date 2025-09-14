package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 🎯 COMPREHENSIVE CRUD APPLICATION
// This file integrates all Go data types into a complete CRUD application

// App represents the main application
type App struct {
	PrimitiveManager  *PrimitiveTypes
	ArraySliceManager *ArraySlices
	StringManager     *StringProcessor
	StructManager     *StructManager
	InterfaceManager  *InterfaceManager
	MapManager        *MapManager
	PointerManager    *PointerManager
	GenericManager    *GenericManager
	Running           bool
}

// NewApp creates a new application instance
func NewApp() *App {
	return &App{
		PrimitiveManager:  NewPrimitiveTypes(),
		ArraySliceManager: NewArraySlices(),
		StringManager:     NewStringProcessor("Hello, Golang CRUD Mastery! 🌍"),
		StructManager:     NewStructManager(),
		InterfaceManager:  NewInterfaceManager(),
		MapManager:        NewMapManager(),
		PointerManager:    NewPointerManager(),
		GenericManager:    NewGenericManager(),
		Running:           true,
	}
}

// Run starts the application
func (app *App) Run() error {
	app.initialize()
	
	for app.Running {
		app.displayMenu()
		choice := app.getUserChoice()
		app.handleChoice(choice)
	}
	
	return nil
}

// initialize sets up the application
func (app *App) initialize() {
	fmt.Println("🚀 Initializing Golang CRUD Mastery Application...")
	
	// Initialize all managers
	app.PrimitiveManager.Create()
	app.ArraySliceManager.Create()
	app.StringManager.Create()
	app.StructManager.Create()
	app.InterfaceManager.Create()
	app.MapManager.Create()
	app.PointerManager.Create()
	app.GenericManager.Create()
	
	fmt.Println("✅ Application initialized successfully!")
	fmt.Println()
}

// displayMenu shows the main menu
func (app *App) displayMenu() {
	fmt.Println("🎯 GOLANG DATA TYPES MASTERY - CRUD APPLICATION")
	fmt.Println("================================================")
	fmt.Println("Choose a data type to explore:")
	fmt.Println()
	fmt.Println("1.  Primitive Types (int, float, bool, string, rune, byte)")
	fmt.Println("2.  Arrays & Slices")
	fmt.Println("3.  Strings & Text Processing")
	fmt.Println("4.  Structs & Data Modeling")
	fmt.Println("5.  Interfaces & Polymorphism")
	fmt.Println("6.  Maps & Key-Value Storage")
	fmt.Println("7.  Pointers & Memory Management")
	fmt.Println("8.  Generics & Type Parameters")
	fmt.Println("9.  Advanced Demonstrations")
	fmt.Println("10. Run All Examples")
	fmt.Println("11. Exit")
	fmt.Println()
	fmt.Print("Enter your choice (1-11): ")
}

// getUserChoice gets user input
func (app *App) getUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > 11 {
		fmt.Println("❌ Invalid choice. Please enter a number between 1 and 11.")
		return app.getUserChoice()
	}
	
	return choice
}

// handleChoice processes user choice
func (app *App) handleChoice(choice int) {
	fmt.Println()
	
	switch choice {
	case 1:
		app.handlePrimitiveTypes()
	case 2:
		app.handleArraysSlices()
	case 3:
		app.handleStrings()
	case 4:
		app.handleStructs()
	case 5:
		app.handleInterfaces()
	case 6:
		app.handleMaps()
	case 7:
		app.handlePointers()
	case 8:
		app.handleGenerics()
	case 9:
		app.handleAdvancedDemonstrations()
	case 10:
		app.handleRunAllExamples()
	case 11:
		app.handleExit()
	default:
		fmt.Println("❌ Invalid choice.")
	}
	
	if app.Running {
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

// handlePrimitiveTypes handles primitive types menu
func (app *App) handlePrimitiveTypes() {
	fmt.Println("🔢 PRIMITIVE TYPES MENU")
	fmt.Println("=======================")
	fmt.Println("1. Create primitive values")
	fmt.Println("2. Read primitive values")
	fmt.Println("3. Update primitive values")
	fmt.Println("4. Delete primitive values")
	fmt.Println("5. Type conversions")
	fmt.Println("6. Constants demonstration")
	fmt.Println("7. Back to main menu")
	fmt.Print("Enter your choice (1-7): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.PrimitiveManager.Create()
	case 2:
		app.PrimitiveManager.Read()
	case 3:
		app.PrimitiveManager.Update()
	case 4:
		app.PrimitiveManager.Delete()
	case 5:
		app.PrimitiveManager.DemonstrateTypeConversions()
	case 6:
		app.PrimitiveManager.DemonstrateConstants()
	case 7:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleArraysSlices handles arrays and slices menu
func (app *App) handleArraysSlices() {
	fmt.Println("📊 ARRAYS & SLICES MENU")
	fmt.Println("=======================")
	fmt.Println("1. Create arrays and slices")
	fmt.Println("2. Read arrays and slices")
	fmt.Println("3. Update arrays and slices")
	fmt.Println("4. Delete from arrays and slices")
	fmt.Println("5. Slice capabilities")
	fmt.Println("6. Sorting and searching")
	fmt.Println("7. Slice tricks and patterns")
	fmt.Println("8. Back to main menu")
	fmt.Print("Enter your choice (1-8): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.ArraySliceManager.Create()
	case 2:
		app.ArraySliceManager.Read()
	case 3:
		app.ArraySliceManager.Update()
	case 4:
		app.ArraySliceManager.Delete()
	case 5:
		app.ArraySliceManager.DemonstrateSliceCapabilities()
	case 6:
		app.ArraySliceManager.DemonstrateSortingAndSearching()
	case 7:
		app.ArraySliceManager.DemonstrateSliceTricks()
	case 8:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleStrings handles strings menu
func (app *App) handleStrings() {
	fmt.Println("📝 STRINGS MENU")
	fmt.Println("===============")
	fmt.Println("1. Create strings")
	fmt.Println("2. Read string information")
	fmt.Println("3. Update string content")
	fmt.Println("4. Delete from strings")
	fmt.Println("5. String searching")
	fmt.Println("6. String transformation")
	fmt.Println("7. String splitting")
	fmt.Println("8. Regex operations")
	fmt.Println("9. String formatting")
	fmt.Println("10. String conversion")
	fmt.Println("11. String builder")
	fmt.Println("12. String comparison")
	fmt.Println("13. Back to main menu")
	fmt.Print("Enter your choice (1-13): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.StringManager.Create()
	case 2:
		app.StringManager.Read()
	case 3:
		app.StringManager.Update()
	case 4:
		app.StringManager.Delete()
	case 5:
		app.StringManager.DemonstrateStringSearching()
	case 6:
		app.StringManager.DemonstrateStringTransformation()
	case 7:
		app.StringManager.DemonstrateStringSplitting()
	case 8:
		app.StringManager.DemonstrateRegexOperations()
	case 9:
		app.StringManager.DemonstrateStringFormatting()
	case 10:
		app.StringManager.DemonstrateStringConversion()
	case 11:
		app.StringManager.DemonstrateStringBuilder()
	case 12:
		app.StringManager.DemonstrateStringComparison()
	case 13:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleStructs handles structs menu
func (app *App) handleStructs() {
	fmt.Println("🏗️  STRUCTS MENU")
	fmt.Println("================")
	fmt.Println("1. Create struct instances")
	fmt.Println("2. Read struct information")
	fmt.Println("3. Update struct data")
	fmt.Println("4. Delete struct instances")
	fmt.Println("5. Struct methods")
	fmt.Println("6. Struct composition")
	fmt.Println("7. Struct tags")
	fmt.Println("8. Struct validation")
	fmt.Println("9. Struct embedding")
	fmt.Println("10. Back to main menu")
	fmt.Print("Enter your choice (1-10): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.StructManager.Create()
	case 2:
		app.StructManager.Read()
	case 3:
		app.StructManager.Update()
	case 4:
		app.StructManager.Delete()
	case 5:
		app.StructManager.DemonstrateStructMethods()
	case 6:
		app.StructManager.DemonstrateStructComposition()
	case 7:
		app.StructManager.DemonstrateStructTags()
	case 8:
		app.StructManager.DemonstrateStructValidation()
	case 9:
		app.StructManager.DemonstrateStructEmbedding()
	case 10:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleInterfaces handles interfaces menu
func (app *App) handleInterfaces() {
	fmt.Println("🔌 INTERFACES MENU")
	fmt.Println("==================")
	fmt.Println("1. Create interface implementations")
	fmt.Println("2. Read interface information")
	fmt.Println("3. Update interface implementations")
	fmt.Println("4. Delete interface implementations")
	fmt.Println("5. Interface polymorphism")
	fmt.Println("6. Interface composition")
	fmt.Println("7. Type assertion")
	fmt.Println("8. Interface sorting")
	fmt.Println("9. Interface validation")
	fmt.Println("10. Interface chaining")
	fmt.Println("11. Interface reflection")
	fmt.Println("12. Back to main menu")
	fmt.Print("Enter your choice (1-12): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.InterfaceManager.Create()
	case 2:
		app.InterfaceManager.Read()
	case 3:
		app.InterfaceManager.Update()
	case 4:
		app.InterfaceManager.Delete()
	case 5:
		app.InterfaceManager.DemonstrateInterfacePolymorphism()
	case 6:
		app.InterfaceManager.DemonstrateInterfaceComposition()
	case 7:
		app.InterfaceManager.DemonstrateInterfaceTypeAssertion()
	case 8:
		app.InterfaceManager.DemonstrateInterfaceSorting()
	case 9:
		app.InterfaceManager.DemonstrateInterfaceValidation()
	case 10:
		app.InterfaceManager.DemonstrateInterfaceChaining()
	case 11:
		app.InterfaceManager.DemonstrateInterfaceReflection()
	case 12:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleMaps handles maps menu
func (app *App) handleMaps() {
	fmt.Println("🗺️  MAPS MENU")
	fmt.Println("=============")
	fmt.Println("1. Create maps")
	fmt.Println("2. Read map contents")
	fmt.Println("3. Update map contents")
	fmt.Println("4. Delete from maps")
	fmt.Println("5. Map iteration")
	fmt.Println("6. Map sorting")
	fmt.Println("7. Map filtering")
	fmt.Println("8. Map transformation")
	fmt.Println("9. Map aggregation")
	fmt.Println("10. Map merging")
	fmt.Println("11. Map concurrency")
	fmt.Println("12. Map validation")
	fmt.Println("13. Back to main menu")
	fmt.Print("Enter your choice (1-13): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.MapManager.Create()
	case 2:
		app.MapManager.Read()
	case 3:
		app.MapManager.Update()
	case 4:
		app.MapManager.Delete()
	case 5:
		app.MapManager.DemonstrateMapIteration()
	case 6:
		app.MapManager.DemonstrateMapSorting()
	case 7:
		app.MapManager.DemonstrateMapFiltering()
	case 8:
		app.MapManager.DemonstrateMapTransformation()
	case 9:
		app.MapManager.DemonstrateMapAggregation()
	case 10:
		app.MapManager.DemonstrateMapMerging()
	case 11:
		app.MapManager.DemonstrateMapConcurrency()
	case 12:
		app.MapManager.DemonstrateMapValidation()
	case 13:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handlePointers handles pointers menu
func (app *App) handlePointers() {
	fmt.Println("📍 POINTERS MENU")
	fmt.Println("================")
	fmt.Println("1. Create pointer instances")
	fmt.Println("2. Read pointer information")
	fmt.Println("3. Update pointer values")
	fmt.Println("4. Delete pointer references")
	fmt.Println("5. Pointer arithmetic")
	fmt.Println("6. Pointer comparison")
	fmt.Println("7. Pointer dereferencing")
	fmt.Println("8. Pointer passing")
	fmt.Println("9. Pointer returning")
	fmt.Println("10. Pointer reflection")
	fmt.Println("11. Memory management")
	fmt.Println("12. Pointer chaining")
	fmt.Println("13. Back to main menu")
	fmt.Print("Enter your choice (1-13): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.PointerManager.Create()
	case 2:
		app.PointerManager.Read()
	case 3:
		app.PointerManager.Update()
	case 4:
		app.PointerManager.Delete()
	case 5:
		app.PointerManager.DemonstratePointerArithmetic()
	case 6:
		app.PointerManager.DemonstratePointerComparison()
	case 7:
		app.PointerManager.DemonstratePointerDereferencing()
	case 8:
		app.PointerManager.DemonstratePointerPassing()
	case 9:
		app.PointerManager.DemonstratePointerReturning()
	case 10:
		app.PointerManager.DemonstratePointerReflection()
	case 11:
		app.PointerManager.DemonstratePointerMemoryManagement()
	case 12:
		app.PointerManager.DemonstratePointerChaining()
	case 13:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleGenerics handles generics menu
func (app *App) handleGenerics() {
	fmt.Println("🔧 GENERICS MENU")
	fmt.Println("================")
	fmt.Println("1. Create generic instances")
	fmt.Println("2. Read generic information")
	fmt.Println("3. Update generic instances")
	fmt.Println("4. Delete generic instances")
	fmt.Println("5. Generic functions")
	fmt.Println("6. Generic constraints")
	fmt.Println("7. Generic interfaces")
	fmt.Println("8. Type inference")
	fmt.Println("9. Generic performance")
	fmt.Println("10. Back to main menu")
	fmt.Print("Enter your choice (1-10): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.GenericManager.Create()
	case 2:
		app.GenericManager.Read()
	case 3:
		app.GenericManager.Update()
	case 4:
		app.GenericManager.Delete()
	case 5:
		app.GenericManager.DemonstrateGenericFunctions()
	case 6:
		app.GenericManager.DemonstrateGenericConstraints()
	case 7:
		app.GenericManager.DemonstrateGenericInterfaces()
	case 8:
		app.GenericManager.DemonstrateGenericTypeInference()
	case 9:
		app.GenericManager.DemonstrateGenericPerformance()
	case 10:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleAdvancedDemonstrations handles advanced demonstrations
func (app *App) handleAdvancedDemonstrations() {
	fmt.Println("🚀 ADVANCED DEMONSTRATIONS")
	fmt.Println("==========================")
	fmt.Println("1. Data type integration")
	fmt.Println("2. Performance comparison")
	fmt.Println("3. Memory usage analysis")
	fmt.Println("4. Type safety demonstration")
	fmt.Println("5. Error handling patterns")
	fmt.Println("6. Concurrency patterns")
	fmt.Println("7. Back to main menu")
	fmt.Print("Enter your choice (1-7): ")
	
	choice := app.getUserChoice()
	
	switch choice {
	case 1:
		app.demonstrateDataTypeIntegration()
	case 2:
		app.demonstratePerformanceComparison()
	case 3:
		app.demonstrateMemoryUsageAnalysis()
	case 4:
		app.demonstrateTypeSafety()
	case 5:
		app.demonstrateErrorHandling()
	case 6:
		app.demonstrateConcurrencyPatterns()
	case 7:
		return
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// handleRunAllExamples runs all examples
func (app *App) handleRunAllExamples() {
	fmt.Println("🎯 RUNNING ALL EXAMPLES")
	fmt.Println("=======================")
	
	// Run all CRUD operations for each data type
	fmt.Println("1. Primitive Types...")
	app.PrimitiveManager.Create()
	app.PrimitiveManager.Read()
	app.PrimitiveManager.Update()
	app.PrimitiveManager.Delete()
	
	fmt.Println("\n2. Arrays & Slices...")
	app.ArraySliceManager.Create()
	app.ArraySliceManager.Read()
	app.ArraySliceManager.Update()
	app.ArraySliceManager.Delete()
	
	fmt.Println("\n3. Strings...")
	app.StringManager.Create()
	app.StringManager.Read()
	app.StringManager.Update()
	app.StringManager.Delete()
	
	fmt.Println("\n4. Structs...")
	app.StructManager.Create()
	app.StructManager.Read()
	app.StructManager.Update()
	app.StructManager.Delete()
	
	fmt.Println("\n5. Interfaces...")
	app.InterfaceManager.Create()
	app.InterfaceManager.Read()
	app.InterfaceManager.Update()
	app.InterfaceManager.Delete()
	
	fmt.Println("\n6. Maps...")
	app.MapManager.Create()
	app.MapManager.Read()
	app.MapManager.Update()
	app.MapManager.Delete()
	
	fmt.Println("\n7. Pointers...")
	app.PointerManager.Create()
	app.PointerManager.Read()
	app.PointerManager.Update()
	app.PointerManager.Delete()
	
	fmt.Println("\n8. Generics...")
	app.GenericManager.Create()
	app.GenericManager.Read()
	app.GenericManager.Update()
	app.GenericManager.Delete()
	
	fmt.Println("\n✅ All examples completed successfully!")
}

// handleExit handles application exit
func (app *App) handleExit() {
	fmt.Println("👋 Thank you for using Golang CRUD Mastery!")
	fmt.Println("You've learned about all the major Go data types:")
	fmt.Println("✅ Primitive types (int, float, bool, string, rune, byte)")
	fmt.Println("✅ Arrays and slices")
	fmt.Println("✅ Strings and text processing")
	fmt.Println("✅ Structs and data modeling")
	fmt.Println("✅ Interfaces and polymorphism")
	fmt.Println("✅ Maps and key-value storage")
	fmt.Println("✅ Pointers and memory management")
	fmt.Println("✅ Generics and type parameters")
	fmt.Println()
	fmt.Println("Keep practicing and building amazing Go applications! 🚀")
	app.Running = false
}

// Advanced demonstration methods

func (app *App) demonstrateDataTypeIntegration() {
	fmt.Println("\n🔗 DATA TYPE INTEGRATION DEMONSTRATION:")
	fmt.Println("======================================")
	
	// Create a complex data structure that uses all types
	type ComplexData struct {
		ID          int                    `json:"id"`
		Name        string                 `json:"name"`
		Values      []float64              `json:"values"`
		Metadata    map[string]interface{} `json:"metadata"`
		IsActive    bool                   `json:"is_active"`
		CreatedAt   int64                  `json:"created_at"`
		Tags        []string               `json:"tags"`
		Config      map[string]string      `json:"config"`
	}
	
	// Create instance
	data := ComplexData{
		ID:    1,
		Name:  "Complex Data Structure",
		Values: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
		Metadata: map[string]interface{}{
			"version": "1.0",
			"author":  "Golang Master",
			"count":   42,
			"active":  true,
		},
		IsActive:  true,
		CreatedAt: 1640995200, // Unix timestamp
		Tags:      []string{"golang", "crud", "mastery", "data-types"},
		Config: map[string]string{
			"database": "postgresql",
			"cache":    "redis",
			"queue":    "rabbitmq",
		},
	}
	
	fmt.Printf("Complex Data Structure: %+v\n", data)
	
	// Demonstrate operations
	fmt.Printf("ID: %d (int)\n", data.ID)
	fmt.Printf("Name: %s (string)\n", data.Name)
	fmt.Printf("Values: %v ([]float64)\n", data.Values)
	fmt.Printf("Metadata: %v (map[string]interface{})\n", data.Metadata)
	fmt.Printf("IsActive: %t (bool)\n", data.IsActive)
	fmt.Printf("CreatedAt: %d (int64)\n", data.CreatedAt)
	fmt.Printf("Tags: %v ([]string)\n", data.Tags)
	fmt.Printf("Config: %v (map[string]string)\n", data.Config)
}

func (app *App) demonstratePerformanceComparison() {
	fmt.Println("\n⚡ PERFORMANCE COMPARISON DEMONSTRATION:")
	fmt.Println("=======================================")
	
	// Compare different data structures for the same operation
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}
	
	// Array vs Slice performance
	fmt.Println("Array vs Slice performance:")
	
	// Slice operations
	slice := make([]int, len(numbers))
	copy(slice, numbers)
	
	// Array operations
	var array [1000000]int
	copy(array[:], numbers)
	
	fmt.Printf("Slice length: %d\n", len(slice))
	fmt.Printf("Array length: %d\n", len(array))
	
	// Map vs Slice for lookup
	fmt.Println("\nMap vs Slice for lookup:")
	
	// Create map
	numberMap := make(map[int]bool)
	for _, num := range numbers {
		numberMap[num] = true
	}
	
	// Test lookup performance
	target := 500000
	
	// Slice lookup (O(n))
	found := false
	for _, num := range slice {
		if num == target {
			found = true
			break
		}
	}
	fmt.Printf("Slice lookup for %d: %t\n", target, found)
	
	// Map lookup (O(1))
	_, found = numberMap[target]
	fmt.Printf("Map lookup for %d: %t\n", target, found)
}

func (app *App) demonstrateMemoryUsageAnalysis() {
	fmt.Println("\n💾 MEMORY USAGE ANALYSIS DEMONSTRATION:")
	fmt.Println("======================================")
	
	// Analyze memory usage of different data types
	fmt.Println("Memory usage analysis:")
	
	// Primitive types
	var intVal int
	var floatVal float64
	var boolVal bool
	var stringVal string
	
	fmt.Printf("int: %d bytes\n", intVal)
	fmt.Printf("float64: %d bytes\n", floatVal)
	fmt.Printf("bool: %d bytes\n", boolVal)
	fmt.Printf("string: %d bytes\n", stringVal)
	
	// Slice vs Array
	slice := make([]int, 1000)
	var array [1000]int
	
	fmt.Printf("Slice of 1000 ints: %d bytes\n", len(slice)*8)
	fmt.Printf("Array of 1000 ints: %d bytes\n", len(array)*8)
	
	// Map memory usage
	numberMap := make(map[int]string)
	for i := 0; i < 1000; i++ {
		numberMap[i] = fmt.Sprintf("value_%d", i)
	}
	
	fmt.Printf("Map with 1000 entries: approximately %d bytes\n", len(numberMap)*32)
}

func (app *App) demonstrateTypeSafety() {
	fmt.Println("\n🛡️  TYPE SAFETY DEMONSTRATION:")
	fmt.Println("=============================")
	
	// Demonstrate type safety
	fmt.Println("Type safety examples:")
	
	// Compile-time type checking
	var intVal int = 42
	var floatVal float64 = 3.14
	var stringVal string = "hello"
	
	// These would cause compile errors:
	// intVal = "hello"        // Cannot assign string to int
	// floatVal = true         // Cannot assign bool to float64
	// stringVal = 42          // Cannot assign int to string
	
	fmt.Printf("intVal: %d (type: %T)\n", intVal, intVal)
	fmt.Printf("floatVal: %.2f (type: %T)\n", floatVal, floatVal)
	fmt.Printf("stringVal: %s (type: %T)\n", stringVal, stringVal)
	
	// Type assertions
	var interfaceVal interface{} = 42
	
	if intVal, ok := interfaceVal.(int); ok {
		fmt.Printf("Interface value as int: %d\n", intVal)
	}
	
	if stringVal, ok := interfaceVal.(string); ok {
		fmt.Printf("Interface value as string: %s\n", stringVal)
	} else {
		fmt.Println("Interface value is not a string")
	}
}

func (app *App) demonstrateErrorHandling() {
	fmt.Println("\n❌ ERROR HANDLING DEMONSTRATION:")
	fmt.Println("================================")
	
	// Demonstrate error handling patterns
	fmt.Println("Error handling examples:")
	
	// Function that returns an error
	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	}
	
	// Test division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}
	
	// Test division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}
	
	// Panic and recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	// This will cause a panic
	// panic("This is a panic!")
}

func (app *App) demonstrateConcurrencyPatterns() {
	fmt.Println("\n🔄 CONCURRENCY PATTERNS DEMONSTRATION:")
	fmt.Println("=====================================")
	
	// Demonstrate basic concurrency patterns
	fmt.Println("Concurrency patterns:")
	
	// Goroutines
	done := make(chan bool)
	
	go func() {
		fmt.Println("Goroutine 1: Hello from goroutine!")
		done <- true
	}()
	
	go func() {
		fmt.Println("Goroutine 2: Another goroutine!")
		done <- true
	}()
	
	// Wait for goroutines to complete
	<-done
	<-done
	
	fmt.Println("Main goroutine: All goroutines completed!")
	
	// Channels
	ch := make(chan int, 2)
	
	// Send values
	ch <- 1
	ch <- 2
	
	// Receive values
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	
	// Close channel
	close(ch)
	
	// Check if channel is closed
	if val, ok := <-ch; ok {
		fmt.Printf("Received: %d\n", val)
	} else {
		fmt.Println("Channel is closed")
	}
}
