package sorting

import "fmt"

type BubbleSortStrategy struct{}

func (bss *BubbleSortStrategy) Sort(data []int) []int {
	fmt.Printf("Bubble sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	return sorted
}

func (bss *BubbleSortStrategy) GetName() string {
	return "Bubble Sort"
}

func (bss *BubbleSortStrategy) GetTimeComplexity() string {
	return "O(n²)"
}