package main

import "fmt"

func rotateLeftBrute(arr []int) {
	n := len(arr)

	temp := make([]int, n)

	for i := 1; i < n; i++ {
		temp[i-1] = arr[i]
	}

	temp[n-1] = arr[0]

	fmt.Println(temp)
}
// TC: O(n)
// SC: O(n)

func rotateLeftOptimized(arr []int) {
	n := len(arr)

	if n == 0 {
		return
	}

	temp := arr[0]

	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[n-1] = temp

	fmt.Println(arr)
}
// TC: O(n)
// SC: O(1)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotateLeftBrute(arr)
}