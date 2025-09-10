package arrays_slices

import (
	"fmt"
	"sort"
	"strings"
)

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

// RunAllArraySliceExamples runs all array and slice examples
func RunAllArraySliceExamples() {
	DemonstrateArrays()
	DemonstrateSlices()
	DemonstrateSliceCapacity()
	DemonstrateSliceMemory()
	DemonstrateSliceOperations()
}
