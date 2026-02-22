package main

import (
	"fmt"
	"sort"
)

func main() {

	// PART 1 — Basic Sorting in Go

	// These functions sort in ascending order only.
	// Internally Go uses a hybrid quicksort/introsort-like algorithm.
	// O(n log n)

	arr := []int{5, 2, 8, 1, 9}
	sort.Ints(arr)
	fmt.Println(arr)

	floats := []float64{3.2, 1.5, 9.8}
	sort.Float64s(floats)
	fmt.Println(floats)

	strs := []string{"banana", "apple", "mango"}
	sort.Strings(strs)
	fmt.Println(strs)

	// PART 2 — Sorting With Comparator (Custom Sorting)

	// sort.Slice(slice, func(i, j int) bool { ... })
	// ⚠️ Comparator works on indices, not values.

	// Descending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)

	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	fmt.Println(strs)

	// Sorting Structs (Very Important)
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// Sort by Age:
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	// Multi-Key Sorting (Advanced Interview Case)

	// Sort by Age, then Name:
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}
