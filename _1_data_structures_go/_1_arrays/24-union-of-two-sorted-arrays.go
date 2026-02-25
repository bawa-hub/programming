// https://takeuforward.org/data-structure/union-of-two-sorted-arrays/

package main

import (
	"fmt"
	"sort"
)

func findUnionMap(arr1, arr2 []int) []int {
	freq := make(map[int]int)

	for _, v := range arr1 {
		freq[v]++
	}
	for _, v := range arr2 {
		freq[v]++
	}

	union := []int{}
	for key := range freq {
		union = append(union, key)
	}

	sort.Ints(union) // because Go map is unordered

	return union
}
// Time Compleixty : O( (m+n)log(m+n) ) .
// Inserting a key in map takes logN times, where N is no of elements in map.
// At max map can store m+n elements {when there are no common elements and elements in arr,arr2 are distntict}. So Inserting m+n th element takes log(m+n) time.
// Upon approximation across insertion of all elements in worst it would take O((m+n)log(m+n) time.
// Using unordered_map also takes the same time, On average insertion in unordered_map takes O(1) time but sorting the union vector takes O((m+n)log(m+n))  time.
// Because at max union vector can have m+n elements.

// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}


// using set
func findUnionSet(arr1, arr2 []int) []int {
	set := make(map[int]struct{})

	for _, v := range arr1 {
		set[v] = struct{}{}
	}
	for _, v := range arr2 {
		set[v] = struct{}{}
	}

	union := []int{}
	for key := range set {
		union = append(union, key)
	}

	sort.Ints(union)
	return union
}
// Time Compleixty : O( (m+n)log(m+n) ) . Inserting a element in set takes logN time, where N is no of elements in set. At max set can store m+n elements {when there are no common elements and elements in arr,arr2 are distntict}. So Inserting m+n th element takes log(m+n) time. Upon approximation across inserting all elements in worst it would take O((m+n)log(m+n) time.
// Using unordered_set also takes the same time, On average insertion in unordered_set takes O(1) time but sorting the union vector takes O((m+n)log(m+n))  time. Because at max union vector can have m+n elements.
// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}

// using pointer (if arrays are sorted)
func findUnionSorted(arr1, arr2 []int) []int {
	i, j := 0, 0
	union := []int{}

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}

	for i < len(arr1) {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}

	return union
}
// Time Complexity: O(m+n), Because at max i runs for n times and j runs for m times. When there are no common elements in arr1 and arr2 and all elements in arr1, arr2 are distinct.
// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}


func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 3, 4, 4, 5, 11, 12}

	union := findUnionSorted(arr1, arr2)

	fmt.Println("Union of arr1 and arr2:")
	for _, val := range union {
		fmt.Print(val, " ")
	}
}