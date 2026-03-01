// https://practice.geeksforgeeks.org/problems/leaders-in-an-array-1587115620/1

package main

import "fmt"

func printLeadersBrute(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		leader := true

		for j := i + 1; j < n; j++ {
			if arr[j] > arr[i] {
				leader = false
				break
			}
		}

		if leader {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Println(arr[n-1])
}
// TC: O(n^2)
// SC: O(1)

func printLeadersOptimal(arr []int) {
	n := len(arr)

	maxVal := arr[n-1]
	fmt.Print(maxVal, " ")

	for i := n - 2; i >= 0; i-- {
		if arr[i] > maxVal {
			fmt.Print(arr[i], " ")
			maxVal = arr[i]
		}
	}

	fmt.Println()
}
// TC: O(n)
// SC: O(1)

func main() {
	arr1 := []int{4, 7, 1, 0}
	fmt.Println("Leaders of first array:")
	printLeadersOptimal(arr1)

	arr2 := []int{10, 22, 12, 3, 0, 6}
	fmt.Println("Leaders of second array:")
	printLeadersOptimal(arr2)
}