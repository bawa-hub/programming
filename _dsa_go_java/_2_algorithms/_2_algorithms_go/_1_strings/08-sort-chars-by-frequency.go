// https://leetcode.com/problems/sort-characters-by-frequency/
// https://practice.geeksforgeeks.org/problems/sorting-elements-of-an-array-by-frequency/0


package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {

	// Step 1: Count frequency
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	// Step 2: Convert map to slice
	type pair struct {
		char  rune
		count int
	}

	var arr []pair
	for k, v := range freq {
		arr = append(arr, pair{k, v})
	}

	// Step 3: Sort by frequency descending
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})

	// Step 4: Build result
	result := ""
	for _, p := range arr {
		for i := 0; i < p.count; i++ {
			result += string(p.char)
		}
	}

	return result
}

func main() {
	fmt.Println(frequencySort("tree"))    // "eert" or "eetr"
	fmt.Println(frequencySort("cccaaa"))  // "cccaaa" or "aaaccc"
	fmt.Println(frequencySort("Aabb"))    // "bbAa"
}