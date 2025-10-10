package sorting

import "fmt"

type QuickSortStrategy struct{}

func (qss *QuickSortStrategy) Sort(data []int) []int {
	fmt.Printf("Quick sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	qss.quickSort(sorted, 0, len(sorted)-1)
	return sorted
}

func (qss *QuickSortStrategy) quickSort(arr []int, low, high int) {
	if low < high {
		pi := qss.partition(arr, low, high)
		qss.quickSort(arr, low, pi-1)
		qss.quickSort(arr, pi+1, high)
	}
}

func (qss *QuickSortStrategy) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (qss *QuickSortStrategy) GetName() string {
	return "Quick Sort"
}

func (qss *QuickSortStrategy) GetTimeComplexity() string {
	return "O(n log n)"
}