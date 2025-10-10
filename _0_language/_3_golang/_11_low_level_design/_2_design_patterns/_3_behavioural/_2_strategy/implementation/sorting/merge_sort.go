package sorting

import "fmt"

type MergeSortStrategy struct{}

func (mss *MergeSortStrategy) Sort(data []int) []int {
	fmt.Printf("Merge sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	mss.mergeSort(sorted, 0, len(sorted)-1)
	return sorted
}

func (mss *MergeSortStrategy) mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mss.mergeSort(arr, left, mid)
		mss.mergeSort(arr, mid+1, right)
		mss.merge(arr, left, mid, right)
	}
}

func (mss *MergeSortStrategy) merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid
	
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)
	
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}
	
	i, j, k := 0, 0, left
	
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func (mss *MergeSortStrategy) GetName() string {
	return "Merge Sort"
}

func (mss *MergeSortStrategy) GetTimeComplexity() string {
	return "O(n log n)"
}