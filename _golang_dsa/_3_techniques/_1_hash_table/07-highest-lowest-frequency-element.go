// https://takeuforward.org/arrays/find-the-highest-lowest-frequency-element/

package main

import "fmt"

// brute force
func countFreq(arr []int) {
	n := len(arr)

	visited := make([]bool, n)
	maxFreq := 0
	minFreq := n

	var maxEle, minEle int

	for i := 0; i < n; i++ {

		// Skip if already processed
		if visited[i] {
			continue
		}

		count := 1

		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				visited[j] = true
				count++
			}
		}

		if count > maxFreq {
			maxFreq = count
			maxEle = arr[i]
		}

		if count < minFreq {
			minFreq = count
			minEle = arr[i]
		}
	}

	fmt.Println("The highest frequency element is:", maxEle)
	fmt.Println("The lowest frequency element is:", minEle)
}

// Time Complexity: O(N*N), where N = size of the array. We are using the nested loop to find the frequency.
// Space Complexity:  O(N), where N = size of the array. It is for the visited array we are using.

func frequencyOptimized(arr []int) {
	mp := make(map[int]int)

	// Count frequency
	for _, value := range arr {
		mp[value]++
	}

	maxFreq := 0
	minFreq := len(arr)

	var maxEle, minEle int

	// Traverse map
	for element, count := range mp {

		if count > maxFreq {
			maxFreq = count
			maxEle = element
		}

		if count < minFreq {
			minFreq = count
			minEle = element
		}
	}

	fmt.Println("The highest frequency element is:", maxEle)
	fmt.Println("The lowest frequency element is:", minEle)
}
// Time Complexity: O(N), where N = size of the array. The insertion and retrieval operation in the map takes O(1) time.
// Space Complexity:  O(N), where N = size of the array. It is for the map we are using.

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	countFreq(arr)
}
