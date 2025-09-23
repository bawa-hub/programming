// https://practice.geeksforgeeks.org/problems/frequency-of-array-elements-1587115620/0

package main

import "fmt"

// Frequency function
func Frequency(arr []int) {
	// map[int]int -> key: element, value: frequency
	freqMap := make(map[int]int)

	// Count frequencies
	for _, val := range arr {
		freqMap[val]++
	}

	// Print frequencies
	for key, value := range freqMap {
		fmt.Println(key, value)
	}
}

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	Frequency(arr)
}
