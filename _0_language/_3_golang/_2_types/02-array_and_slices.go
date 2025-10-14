package main

import "fmt"
import "sort"
import "strings"

var pl = fmt.Println

func main() {

	// Collection of values with the same data type and the size can't be changed

	var arr1 [5]int // Declare integer array with 5 elements
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5} 	// Declare and initialize
	pl("Index 0 :", arr2[0])

	pl("Arr Length :", len(arr2)) 	// Length

	// Iterate with index
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// Iterate with range
	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}

	// Multidimensional Array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print multidimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	// String into slice of runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)

	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make
	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("Slice Size :", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	sl2 := []int{12, 21, 1974}
	pl(sl2)

	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)

	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}

// DemonstrateArrays shows array operations and characteristics
func DemonstrateArrays() {
	fmt.Println("=== ARRAYS ===")
	
	// 1. Array Declaration and Initialization
	fmt.Println("\n--- Array Declaration ---")
	
	// Different ways to declare arrays
	var arr1 [5]int                    // Zero-initialized
	var arr2 = [5]int{1, 2, 3, 4, 5}  // Initialized with values
	arr3 := [5]int{1, 2, 3}           // Partially initialized (rest are zero)
	arr4 := [...]int{1, 2, 3, 4, 5}   // Compiler determines size
	
	fmt.Printf("arr1 (zero-initialized): %v\n", arr1)
	fmt.Printf("arr2 (fully initialized): %v\n", arr2)
	fmt.Printf("arr3 (partially initialized): %v\n", arr3)
	fmt.Printf("arr4 (compiler-determined size): %v\n", arr4)
	
	// 2. Array Access and Modification
	fmt.Println("\n--- Array Access ---")
	
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Element at index 2: %d\n", arr[2])
	
	// Modify elements
	arr[2] = 35
	fmt.Printf("After modifying index 2: %v\n", arr)
	
	// 3. Array Length and Iteration
	fmt.Println("\n--- Array Length and Iteration ---")
	
	fmt.Printf("Array length: %d\n", len(arr))
	
	fmt.Println("Iterating with for loop:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("  arr[%d] = %d\n", i, arr[i])
	}
	
	fmt.Println("Iterating with range:")
	for index, value := range arr {
		fmt.Printf("  arr[%d] = %d\n", index, value)
	}
	
	// 4. Array Comparison
	fmt.Println("\n--- Array Comparison ---")
	
	arr5 := [3]int{1, 2, 3}
	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 4}
	
	fmt.Printf("arr5 == arr6: %t\n", arr5 == arr6)
	fmt.Printf("arr5 == arr7: %t\n", arr5 == arr7)
	
	// 5. Multi-dimensional Arrays
	fmt.Println("\n--- Multi-dimensional Arrays ---")
	
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	
	fmt.Println("3x3 Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// DemonstrateSlices shows slice operations and characteristics
func DemonstrateSlices() {
	fmt.Println("\n=== SLICES ===")
	
	// 1. Slice Declaration and Initialization
	fmt.Println("\n--- Slice Declaration ---")
	
	// Different ways to create slices
	var slice1 []int                    // nil slice
	slice2 := []int{1, 2, 3, 4, 5}     // slice literal
	slice3 := make([]int, 5)           // make with length
	slice4 := make([]int, 5, 10)       // make with length and capacity
	
	fmt.Printf("slice1 (nil slice): %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 (literal): %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 (make with len): %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 (make with len,cap): %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))
	
	// 2. Slice Operations
	fmt.Println("\n--- Slice Operations ---")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", slice)
	
	// Append elements
	slice = append(slice, 6, 7, 8)
	fmt.Printf("After append(6,7,8): %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))
	
	// Append another slice
	anotherSlice := []int{9, 10, 11}
	slice = append(slice, anotherSlice...)
	fmt.Printf("After append another slice: %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))
	
	// 3. Slice Slicing
	fmt.Println("\n--- Slice Slicing ---")
	
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", original)
	
	// Basic slicing
	fmt.Printf("slice[2:5]: %v\n", original[2:5])     // [2, 3, 4]
	fmt.Printf("slice[:5]: %v\n", original[:5])        // [0, 1, 2, 3, 4]
	fmt.Printf("slice[5:]: %v\n", original[5:])        // [5, 6, 7, 8, 9]
	fmt.Printf("slice[:]: %v\n", original[:])          // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	
	// Slicing with capacity
	fmt.Printf("slice[2:5:8]: %v (len: %d, cap: %d)\n", original[2:5:8], len(original[2:5:8]), cap(original[2:5:8]))
	
	// 4. Slice Copy
	fmt.Println("\n--- Slice Copy ---")
	
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	
	copied := copy(destination, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)
	fmt.Printf("Elements copied: %d\n", copied)
	
	// 5. Slice as Function Parameter
	fmt.Println("\n--- Slice as Function Parameter ---")
	
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before function: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("After function: %v\n", numbers)
	
	// 6. Slice of Strings
	fmt.Println("\n--- Slice of Strings ---")
	
	fruits := []string{"apple", "banana", "cherry", "date"}
	fmt.Printf("Fruits: %v\n", fruits)
	
	// Sort strings
	sort.Strings(fruits)
	fmt.Printf("Sorted fruits: %v\n", fruits)
	
	// Join strings
	joined := strings.Join(fruits, ", ")
	fmt.Printf("Joined: %s\n", joined)
}

// modifySlice demonstrates that slices are passed by reference
func modifySlice(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

// DemonstrateSliceCapacity shows how slice capacity grows
func DemonstrateSliceCapacity() {
	fmt.Println("\n=== SLICE CAPACITY GROWTH ===")
	
	var slice []int
	fmt.Printf("Initial: len=%d, cap=%d\n", len(slice), cap(slice))
	
	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(slice), cap(slice))
	}
}

// DemonstrateSliceMemory shows memory characteristics of slices
func DemonstrateSliceMemory() {
	fmt.Println("\n=== SLICE MEMORY CHARACTERISTICS ===")
	
	// Create a slice
	original := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v (len: %d, cap: %d)\n", original, len(original), cap(original))
	
	// Create a slice that shares the same underlying array
	shared := original[1:4]
	fmt.Printf("Shared slice: %v (len: %d, cap: %d)\n", shared, len(shared), cap(shared))
	
	// Modify the shared slice
	shared[0] = 99
	fmt.Printf("After modifying shared[0]:\n")
	fmt.Printf("  Original: %v\n", original)
	fmt.Printf("  Shared: %v\n", shared)
	
	// Create a copy to avoid sharing
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 100
	fmt.Printf("After modifying copied[0]:\n")
	fmt.Printf("  Original: %v\n", original)
	fmt.Printf("  Copied: %v\n", copied)
}

// DemonstrateSliceOperations shows advanced slice operations
func DemonstrateSliceOperations() {
	fmt.Println("\n=== ADVANCED SLICE OPERATIONS ===")
	
	// 1. Remove element at index
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", slice)
	
	index := 2 // Remove element at index 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("After removing index 2: %v\n", slice)
	
	// 2. Insert element at index
	slice = []int{1, 2, 4, 5}
	fmt.Printf("Before insert: %v\n", slice)
	
	insertIndex := 2
	insertValue := 3
	slice = append(slice[:insertIndex], append([]int{insertValue}, slice[insertIndex:]...)...)
	fmt.Printf("After inserting 3 at index 2: %v\n", slice)
	
	// 3. Reverse slice
	slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("Before reverse: %v\n", slice)
	
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Printf("After reverse: %v\n", slice)
	
	// 4. Filter slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original numbers: %v\n", numbers)
	
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evenNumbers)
}


// ArraySlices demonstrates arrays and slices of different types
type ArraySlices struct {
	// Arrays - Fixed size collections
	IntArray    [5]int           // Fixed size array of 5 integers
	StringArray [3]string        // Fixed size array of 3 strings
	BoolArray   [4]bool          // Fixed size array of 4 booleans
	
	// Slices - Dynamic arrays
	IntSlice    []int            // Dynamic slice of integers
	StringSlice []string         // Dynamic slice of strings
	BoolSlice   []bool           // Dynamic slice of booleans
	
	// 2D Arrays and Slices
	Matrix2D    [3][3]int        // 2D array (3x3 matrix)
	Slice2D     [][]string       // 2D slice (dynamic matrix)
	
	// Slice of slices (jagged array)
	JaggedSlice [][]int          // Each inner slice can have different lengths
}

// NewArraySlices creates a new instance with initialized values
func NewArraySlices() *ArraySlices {
	return &ArraySlices{
		IntArray:    [5]int{1, 2, 3, 4, 5},
		StringArray: [3]string{"Go", "Rust", "Python"},
		BoolArray:   [4]bool{true, false, true, false},
		
		IntSlice:    []int{10, 20, 30, 40, 50},
		StringSlice: []string{"apple", "banana", "cherry"},
		BoolSlice:   []bool{true, true, false},
		
		Matrix2D: [3][3]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		Slice2D: [][]string{
			{"a", "b", "c"},
			{"d", "e"},
			{"f", "g", "h", "i"},
		},
		JaggedSlice: [][]int{
			{1, 2},
			{3, 4, 5, 6},
			{7},
			{8, 9, 10},
		},
	}
}

// CRUD Operations for Arrays

// Create - Initialize array values
func (as *ArraySlices) Create() {
	fmt.Println("🔧 Creating array and slice values...")
	
	// Array initialization methods
	var emptyArray [5]int
	fmt.Printf("Empty array: %v\n", emptyArray)
	
	// Array literal
	literalArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Literal array: %v\n", literalArray)
	
	// Array with specific indices
	sparseArray := [5]int{1: 10, 3: 30}
	fmt.Printf("Sparse array: %v\n", sparseArray)
	
	// Array with ellipsis (let compiler determine size)
	autoSizeArray := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Auto-size array: %v (length: %d)\n", autoSizeArray, len(autoSizeArray))
	
	// Slice initialization methods
	var nilSlice []int
	fmt.Printf("Nil slice: %v (len: %d, cap: %d)\n", nilSlice, len(nilSlice), cap(nilSlice))
	
	// Slice literal
	literalSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Literal slice: %v (len: %d, cap: %d)\n", literalSlice, len(literalSlice), cap(literalSlice))
	
	// Slice from array
	array := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := array[1:4] // elements 1, 2, 3
	fmt.Printf("Slice from array: %v\n", sliceFromArray)
	
	// Slice using make
	madeSlice := make([]int, 5, 10) // length 5, capacity 10
	fmt.Printf("Made slice: %v (len: %d, cap: %d)\n", madeSlice, len(madeSlice), cap(madeSlice))
	
	fmt.Println("✅ Array and slice values created successfully")
}

// Read - Display all array and slice values
func (as *ArraySlices) Read() {
	fmt.Println("\n📖 READING ARRAYS AND SLICES:")
	fmt.Println("=============================")
	
	// Read arrays
	fmt.Printf("Int Array: %v (length: %d)\n", as.IntArray, len(as.IntArray))
	fmt.Printf("String Array: %v (length: %d)\n", as.StringArray, len(as.StringArray))
	fmt.Printf("Bool Array: %v (length: %d)\n", as.BoolArray, len(as.BoolArray))
	
	// Read slices
	fmt.Printf("\nInt Slice: %v (len: %d, cap: %d)\n", as.IntSlice, len(as.IntSlice), cap(as.IntSlice))
	fmt.Printf("String Slice: %v (len: %d, cap: %d)\n", as.StringSlice, len(as.StringSlice), cap(as.StringSlice))
	fmt.Printf("Bool Slice: %v (len: %d, cap: %d)\n", as.BoolSlice, len(as.BoolSlice), cap(as.BoolSlice))
	
	// Read 2D structures
	fmt.Printf("\n2D Matrix:\n")
	for i, row := range as.Matrix2D {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
	
	fmt.Printf("\n2D Slice:\n")
	for i, row := range as.Slice2D {
		fmt.Printf("  Row %d: %v (len: %d)\n", i, row, len(row))
	}
	
	fmt.Printf("\nJagged Slice:\n")
	for i, row := range as.JaggedSlice {
		fmt.Printf("  Row %d: %v (len: %d)\n", i, row, len(row))
	}
	
	// Demonstrate iteration methods
	fmt.Printf("\nIteration Methods:\n")
	fmt.Printf("Index iteration: ")
	for i := 0; i < len(as.IntSlice); i++ {
		fmt.Printf("%d ", as.IntSlice[i])
	}
	fmt.Println()
	
	fmt.Printf("Range iteration: ")
	for index, value := range as.IntSlice {
		fmt.Printf("[%d:%d] ", index, value)
	}
	fmt.Println()
	
	fmt.Printf("Value-only iteration: ")
	for _, value := range as.IntSlice {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

// Update - Modify array and slice values
func (as *ArraySlices) Update() {
	fmt.Println("\n🔄 UPDATING ARRAYS AND SLICES:")
	fmt.Println("==============================")
	
	// Update array elements
	as.IntArray[0] = 100
	as.IntArray[2] = 300
	as.StringArray[1] = "Golang"
	
	// Update slice elements
	as.IntSlice[0] = 1000
	as.StringSlice[1] = "grape"
	
	// Append to slices (dynamic growth)
	as.IntSlice = append(as.IntSlice, 60, 70, 80)
	as.StringSlice = append(as.StringSlice, "date", "elderberry")
	as.BoolSlice = append(as.BoolSlice, true, false, true)
	
	// Insert element at specific position
	as.IntSlice = insertAt(as.IntSlice, 2, 999)
	as.StringSlice = insertAt(as.StringSlice, 1, "blueberry")
	
	// Update 2D structures
	as.Matrix2D[1][1] = 99
	as.Slice2D[0] = append(as.Slice2D[0], "x", "y")
	as.JaggedSlice[1] = append(as.JaggedSlice[1], 11, 12)
	
	// Slice operations
	as.IntSlice = as.IntSlice[1:] // Remove first element
	as.StringSlice = as.StringSlice[:len(as.StringSlice)-1] // Remove last element
	
	fmt.Println("✅ Arrays and slices updated successfully")
}

// Delete - Remove elements from arrays and slices
func (as *ArraySlices) Delete() {
	fmt.Println("\n🗑️  DELETING FROM ARRAYS AND SLICES:")
	fmt.Println("====================================")
	
	// For arrays, we can't actually delete elements, but we can reset them
	for i := range as.IntArray {
		as.IntArray[i] = 0
	}
	
	// For slices, we can actually delete elements
	as.IntSlice = deleteAt(as.IntSlice, 2) // Delete element at index 2
	as.StringSlice = deleteAt(as.StringSlice, 1) // Delete element at index 1
	
	// Delete multiple elements
	as.IntSlice = deleteRange(as.IntSlice, 1, 3) // Delete elements 1-3
	as.StringSlice = deleteByValue(as.StringSlice, "grape")
	
	// Clear slices
	as.BoolSlice = as.BoolSlice[:0] // Clear but keep capacity
	// or
	as.BoolSlice = nil // Clear and reset capacity
	
	// Delete from 2D structures
	if len(as.Slice2D) > 1 {
		as.Slice2D = deleteAt(as.Slice2D, 1) // Delete row 1
	}
	
	fmt.Println("✅ Elements deleted successfully")
}

// Helper functions for slice operations

// insertAt inserts a value at the specified index
func insertAt[T any](slice []T, index int, value T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	
	// Grow slice by 1
	slice = append(slice, value)
	
	// Shift elements to the right
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	
	return slice
}

// deleteAt removes element at the specified index
func deleteAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	
	// Shift elements to the left
	copy(slice[index:], slice[index+1:])
	
	// Truncate slice
	return slice[:len(slice)-1]
}

// deleteRange removes elements in the specified range
func deleteRange[T any](slice []T, start, end int) []T {
	if start < 0 || end > len(slice) || start >= end {
		return slice
	}
	
	// Shift elements to the left
	copy(slice[start:], slice[end:])
	
	// Truncate slice
	return slice[:len(slice)-(end-start)]
}

// deleteByValue removes all occurrences of the specified value
func deleteByValue[T comparable](slice []T, value T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// Advanced slice operations

// DemonstrateSliceCapabilities shows advanced slice features
func (as *ArraySlices) DemonstrateSliceCapabilities() {
	fmt.Println("\n🚀 ADVANCED SLICE CAPABILITIES:")
	fmt.Println("===============================")
	
	// Slice capacity and growth
	original := make([]int, 3, 5)
	fmt.Printf("Original: %v (len: %d, cap: %d)\n", original, len(original), cap(original))
	
	// Append within capacity
	original = append(original, 4, 5)
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", original, len(original), cap(original))
	
	// Append beyond capacity (triggers reallocation)
	original = append(original, 6, 7, 8)
	fmt.Printf("After growth: %v (len: %d, cap: %d)\n", original, len(original), cap(original))
	
	// Slice expressions
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\nSlice expressions on array %v:\n", arr)
	fmt.Printf("arr[2:5] = %v\n", arr[2:5])   // elements 2, 3, 4
	fmt.Printf("arr[:5] = %v\n", arr[:5])     // elements 0, 1, 2, 3, 4
	fmt.Printf("arr[5:] = %v\n", arr[5:])     // elements 5, 6, 7, 8, 9
	fmt.Printf("arr[:] = %v\n", arr[:])       // all elements
	
	// Slice of slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\nSlice expressions on slice %v:\n", slice)
	fmt.Printf("slice[2:5] = %v (len: %d, cap: %d)\n", slice[2:5], len(slice[2:5]), cap(slice[2:5]))
	fmt.Printf("slice[:5] = %v (len: %d, cap: %d)\n", slice[:5], len(slice[:5]), cap(slice[:5]))
	fmt.Printf("slice[5:] = %v (len: %d, cap: %d)\n", slice[5:], len(slice[5:]), cap(slice[5:]))
	
	// Copy function
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, 3)
	copied := copy(destination, source)
	fmt.Printf("\nCopy: source=%v, destination=%v, copied=%d elements\n", source, destination, copied)
	
	// Full copy
	fullCopy := make([]int, len(source))
	copy(fullCopy, source)
	fmt.Printf("Full copy: %v\n", fullCopy)
}

// DemonstrateSortingAndSearching shows sorting and searching operations
func (as *ArraySlices) DemonstrateSortingAndSearching() {
	fmt.Println("\n🔍 SORTING AND SEARCHING:")
	fmt.Println("=========================")
	
	// Sorting slices
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", numbers)
	
	// Sort integers
	sort.Ints(numbers)
	fmt.Printf("Sorted: %v\n", numbers)
	
	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Original words: %v\n", words)
	sort.Strings(words)
	fmt.Printf("Sorted words: %v\n", words)
	
	// Custom sorting
	people := []string{"Alice", "Bob", "Charlie", "David"}
	fmt.Printf("Original people: %v\n", people)
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j]) // Sort by length
	})
	fmt.Printf("Sorted by length: %v\n", people)
	
	// Searching
	searchValue := 25
	index := sort.SearchInts(numbers, searchValue)
	if index < len(numbers) && numbers[index] == searchValue {
		fmt.Printf("Found %d at index %d\n", searchValue, index)
	} else {
		fmt.Printf("%d not found, would be inserted at index %d\n", searchValue, index)
	}
	
	// Binary search
	searchValue = 22
	index = sort.SearchInts(numbers, searchValue)
	if index < len(numbers) && numbers[index] == searchValue {
		fmt.Printf("Binary search found %d at index %d\n", searchValue, index)
	} else {
		fmt.Printf("Binary search: %d not found\n", searchValue)
	}
}

// DemonstrateSliceTricks shows useful slice tricks and patterns
func (as *ArraySlices) DemonstrateSliceTricks() {
	fmt.Println("\n🎩 SLICE TRICKS AND PATTERNS:")
	fmt.Println("=============================")
	
	// Remove duplicates
	withDuplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := removeDuplicates(withDuplicates)
	fmt.Printf("Original: %v\n", withDuplicates)
	fmt.Printf("Unique: %v\n", unique)
	
	// Reverse slice
	original := []int{1, 2, 3, 4, 5}
	reversed := reverse(original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Reversed: %v\n", reversed)
	
	// Rotate slice
	toRotate := []int{1, 2, 3, 4, 5}
	rotated := rotate(toRotate, 2) // Rotate left by 2 positions
	fmt.Printf("Original: %v\n", toRotate)
	fmt.Printf("Rotated left by 2: %v\n", rotated)
	
	// Chunk slice into smaller pieces
	bigSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := chunk(bigSlice, 3)
	fmt.Printf("Original: %v\n", bigSlice)
	fmt.Printf("Chunks of 3: %v\n", chunks)
	
	// Flatten 2D slice
	matrix := [][]int{{1, 2}, {3, 4, 5}, {6}}
	flattened := flatten(matrix)
	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Printf("Flattened: %v\n", flattened)
}

// Helper functions for slice tricks

func removeDuplicates[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	result := []T{}
	
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	
	return result
}

func reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

func rotate[T any](slice []T, positions int) []T {
	if len(slice) == 0 || positions == 0 {
		return slice
	}
	
	positions = positions % len(slice)
	if positions < 0 {
		positions += len(slice)
	}
	
	result := make([]T, len(slice))
	copy(result, slice[positions:])
	copy(result[len(slice)-positions:], slice[:positions])
	
	return result
}

func chunk[T any](slice []T, size int) [][]T {
	var chunks [][]T
	
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	
	return chunks
}

func flatten[T any](matrix [][]T) []T {
	var result []T
	for _, row := range matrix {
		result = append(result, row...)
	}
	return result
}