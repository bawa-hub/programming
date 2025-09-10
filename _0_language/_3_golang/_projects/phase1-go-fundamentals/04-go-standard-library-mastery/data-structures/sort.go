package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Custom types for demonstration
type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d, %s)", p.Name, p.Age, p.City)
}

type Product struct {
	Name     string
	Price    float64
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("%s - $%.2f (%s)", p.Name, p.Price, p.Category)
}

type Student struct {
	Name  string
	Grade float64
	ID    int
}

func (s Student) String() string {
	return fmt.Sprintf("%s (ID: %d, Grade: %.1f)", s.Name, s.ID, s.Grade)
}

// Custom sort interface implementation
type CustomIntSlice []int

func (c CustomIntSlice) Len() int           { return len(c) }
func (c CustomIntSlice) Less(i, j int) bool { return c[i] < c[j] }
func (c CustomIntSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// Custom sort interface for Person
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	fmt.Println("🚀 Go sort Package Mastery Examples")
	fmt.Println("====================================")

	// 1. Basic Sorting - Integers
	fmt.Println("\n1. Basic Sorting - Integers:")
	
	numbers := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Printf("Original: %v\n", numbers)
	
	// Sort in ascending order
	sort.Ints(numbers)
	fmt.Printf("Sorted: %v\n", numbers)
	
	// Check if sorted
	fmt.Printf("Is sorted: %t\n", sort.IntsAreSorted(numbers))
	
	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("Reverse sorted: %v\n", numbers)

	// 2. Basic Sorting - Floats
	fmt.Println("\n2. Basic Sorting - Floats:")
	
	floats := []float64{3.14, 2.71, 1.41, 1.73, 2.24}
	fmt.Printf("Original: %v\n", floats)
	
	sort.Float64s(floats)
	fmt.Printf("Sorted: %v\n", floats)
	
	// Check if sorted
	fmt.Printf("Is sorted: %t\n", sort.Float64sAreSorted(floats))

	// 3. Basic Sorting - Strings
	fmt.Println("\n3. Basic Sorting - Strings:")
	
	names := []string{"Charlie", "Alice", "Bob", "David", "Eve"}
	fmt.Printf("Original: %v\n", names)
	
	sort.Strings(names)
	fmt.Printf("Sorted: %v\n", names)
	
	// Check if sorted
	fmt.Printf("Is sorted: %t\n", sort.StringsAreSorted(names))

	// 4. Custom Sorting with Interface
	fmt.Println("\n4. Custom Sorting with Interface:")
	
	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "Tokyo"},
		{"David", 28, "Paris"},
	}
	
	fmt.Println("Original people:")
	for _, p := range people {
		fmt.Printf("  %s\n", p)
	}
	
	// Sort by age
	sort.Sort(ByAge(people))
	fmt.Println("\nSorted by age:")
	for _, p := range people {
		fmt.Printf("  %s\n", p)
	}
	
	// Sort by name
	sort.Sort(ByName(people))
	fmt.Println("\nSorted by name:")
	for _, p := range people {
		fmt.Printf("  %s\n", p)
	}

	// 5. Function-based Sorting
	fmt.Println("\n5. Function-based Sorting:")
	
	products := []Product{
		{"Laptop", 999.99, "Electronics"},
		{"Book", 19.99, "Education"},
		{"Phone", 699.99, "Electronics"},
		{"Pen", 2.99, "Office"},
		{"Tablet", 499.99, "Electronics"},
	}
	
	fmt.Println("Original products:")
	for _, p := range products {
		fmt.Printf("  %s\n", p)
	}
	
	// Sort by price
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	fmt.Println("\nSorted by price:")
	for _, p := range products {
		fmt.Printf("  %s\n", p)
	}
	
	// Sort by name
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	fmt.Println("\nSorted by name:")
	for _, p := range products {
		fmt.Printf("  %s\n", p)
	}
	
	// Sort by category, then by price
	sort.Slice(products, func(i, j int) bool {
		if products[i].Category != products[j].Category {
			return products[i].Category < products[j].Category
		}
		return products[i].Price < products[j].Price
	})
	fmt.Println("\nSorted by category, then price:")
	for _, p := range products {
		fmt.Printf("  %s\n", p)
	}

	// 6. Stable Sorting
	fmt.Println("\n6. Stable Sorting:")
	
	// Create data with equal elements
	students := []Student{
		{"Alice", 85.5, 1},
		{"Bob", 85.5, 2},
		{"Charlie", 90.0, 3},
		{"David", 85.5, 4},
		{"Eve", 90.0, 5},
	}
	
	fmt.Println("Original students:")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}
	
	// Stable sort by grade
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Grade < students[j].Grade
	})
	fmt.Println("\nStable sorted by grade (preserves original order for equal grades):")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}

	// 7. Binary Search
	fmt.Println("\n7. Binary Search:")
	
	// Create sorted data
	sortedNumbers := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Printf("Sorted numbers: %v\n", sortedNumbers)
	
	// Search for specific values
	searchValues := []int{5, 10, 15, 20}
	for _, val := range searchValues {
		index := sort.SearchInts(sortedNumbers, val)
		if index < len(sortedNumbers) && sortedNumbers[index] == val {
			fmt.Printf("Found %d at index %d\n", val, index)
		} else {
			fmt.Printf("%d not found, would be inserted at index %d\n", val, index)
		}
	}

	// 8. Custom Binary Search
	fmt.Println("\n8. Custom Binary Search:")
	
	// Search in custom data
	sortedPeople := []Person{
		{"Alice", 25, "NYC"},
		{"Bob", 30, "LA"},
		{"Charlie", 35, "Chicago"},
		{"David", 40, "Boston"},
		{"Eve", 45, "Seattle"},
	}
	
	// Search by age
	targetAge := 35
	index := sort.Search(len(sortedPeople), func(i int) bool {
		return sortedPeople[i].Age >= targetAge
	})
	
	if index < len(sortedPeople) && sortedPeople[index].Age == targetAge {
		fmt.Printf("Found person with age %d: %s\n", targetAge, sortedPeople[index])
	} else {
		fmt.Printf("No person found with age %d\n", targetAge)
	}

	// 9. Reverse Sorting
	fmt.Println("\n9. Reverse Sorting:")
	
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", data)
	
	// Reverse sort
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Printf("Reverse sorted: %v\n", data)
	
	// Reverse sort with function
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	fmt.Printf("Function reverse sorted: %v\n", data)

	// 10. Sorting Performance Comparison
	fmt.Println("\n10. Sorting Performance Comparison:")
	
	// Generate random data
	rand.Seed(time.Now().UnixNano())
	size := 10000
	randomData := make([]int, size)
	for i := range randomData {
		randomData[i] = rand.Intn(1000)
	}
	
	// Test different sorting methods
	testData := make([]int, len(randomData))
	
	// Copy and sort with Ints
	copy(testData, randomData)
	start := time.Now()
	sort.Ints(testData)
	intsTime := time.Since(start)
	
	// Copy and sort with Slice
	copy(testData, randomData)
	start = time.Now()
	sort.Slice(testData, func(i, j int) bool {
		return testData[i] < testData[j]
	})
	sliceTime := time.Since(start)
	
	// Copy and sort with custom interface
	copy(testData, randomData)
	customSlice := CustomIntSlice(testData)
	start = time.Now()
	sort.Sort(customSlice)
	customTime := time.Since(start)
	
	fmt.Printf("Sorting %d integers:\n", size)
	fmt.Printf("  sort.Ints(): %v\n", intsTime)
	fmt.Printf("  sort.Slice(): %v\n", sliceTime)
	fmt.Printf("  Custom interface: %v\n", customTime)

	// 11. Sorting with Multiple Criteria
	fmt.Println("\n11. Sorting with Multiple Criteria:")
	
	employees := []struct {
		Name     string
		Department string
		Salary   int
		Age      int
	}{
		{"Alice", "Engineering", 80000, 30},
		{"Bob", "Marketing", 60000, 25},
		{"Charlie", "Engineering", 90000, 35},
		{"David", "Marketing", 70000, 28},
		{"Eve", "Engineering", 75000, 32},
	}
	
	fmt.Println("Original employees:")
	for _, emp := range employees {
		fmt.Printf("  %s (%s, $%d, age %d)\n", emp.Name, emp.Department, emp.Salary, emp.Age)
	}
	
	// Sort by department, then by salary (descending), then by age
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Department != employees[j].Department {
			return employees[i].Department < employees[j].Department
		}
		if employees[i].Salary != employees[j].Salary {
			return employees[i].Salary > employees[j].Salary
		}
		return employees[i].Age < employees[j].Age
	})
	
	fmt.Println("\nSorted by department, salary (desc), age:")
	for _, emp := range employees {
		fmt.Printf("  %s (%s, $%d, age %d)\n", emp.Name, emp.Department, emp.Salary, emp.Age)
	}

	// 12. Sorting Strings with Custom Comparison
	fmt.Println("\n12. Sorting Strings with Custom Comparison:")
	
	words := []string{"apple", "Banana", "cherry", "Date", "elderberry"}
	fmt.Printf("Original: %v\n", words)
	
	// Case-insensitive sorting
	sort.Slice(words, func(i, j int) bool {
		return strings.ToLower(words[i]) < strings.ToLower(words[j])
	})
	fmt.Printf("Case-insensitive sorted: %v\n", words)
	
	// Sort by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Printf("Sorted by length: %v\n", words)

	// 13. Sorting with Validation
	fmt.Println("\n13. Sorting with Validation:")
	
	// Create data with some invalid entries
	mixedData := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Printf("Original: %v\n", mixedData)
	
	// Sort and validate
	sort.Ints(mixedData)
	fmt.Printf("Sorted: %v\n", mixedData)
	fmt.Printf("Is sorted: %t\n", sort.IntsAreSorted(mixedData))
	
	// Test with unsorted data
	unsortedData := []int{5, 2, 8, 1, 9}
	fmt.Printf("Unsorted data: %v\n", unsortedData)
	fmt.Printf("Is sorted: %t\n", sort.IntsAreSorted(unsortedData))

	// 14. Sorting with Error Handling
	fmt.Println("\n14. Sorting with Error Handling:")
	
	// Safe sorting function
	safeSort := func(data []int) error {
		if len(data) == 0 {
			return fmt.Errorf("empty slice")
		}
		if len(data) > 1000000 {
			return fmt.Errorf("slice too large")
		}
		sort.Ints(data)
		return nil
	}
	
	// Test safe sorting
	testCases := [][]int{
		{},
		{1, 2, 3},
		make([]int, 1000001), // Large slice
		{3, 1, 4, 1, 5},
	}
	
	for i, testCase := range testCases {
		err := safeSort(testCase)
		if err != nil {
			fmt.Printf("Test case %d: Error - %v\n", i+1, err)
		} else {
			fmt.Printf("Test case %d: Success - %v\n", i+1, testCase)
		}
	}

	// 15. Advanced Search Patterns
	fmt.Println("\n15. Advanced Search Patterns:")
	
	// Find first occurrence
	sortedWithDuplicates := []int{1, 2, 2, 2, 3, 4, 4, 5}
	fmt.Printf("Data with duplicates: %v\n", sortedWithDuplicates)
	
	target := 2
	firstIndex := sort.SearchInts(sortedWithDuplicates, target)
	fmt.Printf("First occurrence of %d at index: %d\n", target, firstIndex)
	
	// Find last occurrence
	lastIndex := sort.SearchInts(sortedWithDuplicates, target+1) - 1
	fmt.Printf("Last occurrence of %d at index: %d\n", target, lastIndex)
	
	// Count occurrences
	count := lastIndex - firstIndex + 1
	fmt.Printf("Count of %d: %d\n", target, count)

	// 16. Sorting with Custom Data Types
	fmt.Println("\n16. Sorting with Custom Data Types:")
	
	// Sort by multiple fields with different types
	type Record struct {
		ID    int
		Name  string
		Value float64
		Active bool
	}
	
	records := []Record{
		{3, "Charlie", 3.14, true},
		{1, "Alice", 2.71, false},
		{2, "Bob", 1.41, true},
		{4, "David", 1.73, false},
	}
	
	fmt.Println("Original records:")
	for _, r := range records {
		fmt.Printf("  ID: %d, Name: %s, Value: %.2f, Active: %t\n", r.ID, r.Name, r.Value, r.Active)
	}
	
	// Sort by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	fmt.Println("\nSorted by ID:")
	for _, r := range records {
		fmt.Printf("  ID: %d, Name: %s, Value: %.2f, Active: %t\n", r.ID, r.Name, r.Value, r.Active)
	}
	
	// Sort by Active (true first), then by Value (descending)
	sort.Slice(records, func(i, j int) bool {
		if records[i].Active != records[j].Active {
			return records[i].Active && !records[j].Active
		}
		return records[i].Value > records[j].Value
	})
	fmt.Println("\nSorted by Active (true first), then Value (desc):")
	for _, r := range records {
		fmt.Printf("  ID: %d, Name: %s, Value: %.2f, Active: %t\n", r.ID, r.Name, r.Value, r.Active)
	}

	fmt.Println("\n🎉 sort Package Mastery Complete!")
}
