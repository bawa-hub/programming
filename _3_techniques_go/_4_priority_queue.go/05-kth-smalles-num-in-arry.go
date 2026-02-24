// https://practice.geeksforgeeks.org/problems/kth-smallest-element5635/1
// https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/

package main

// ---------- MIN HEAP ----------
type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // MIN heap
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kthSmallest(arr []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, val)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}

// -------- OPTIMIZED FUNCTION --------
func kthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		// keep heap size ≤ k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// root is kth largest
	return (*h)[0]
}
// (O(n log k))

// quick sort algo
func kthSmallest(arr []int, left, right, k int) int {
	if k > 0 && k <= right-left+1 {

		index := partitionSmallest(arr, left, right)

		if index-left == k-1 {
			return arr[index]
		}

		if index-left > k-1 {
			return kthSmallest(arr, left, index-1, k)
		}

		return kthSmallest(arr, index+1, right, k-index+left-1)
	}

	return -1
}

func partitionSmallest(arr []int, left, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j <= right-1; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}
// Time complexity: O(n) , where n = size of the array
// Space complexity: O(1) 

func main() {
	arr := []int{1, 2, 6, 4, 5, 3}
	fmt.Println("3rd Smallest:", kthSmallest(arr, 3))
	fmt.Println("3rd Smallest:", kthLargest(arr, 3))
}