// https://practice.geeksforgeeks.org/problems/subsets-with-xor-value2023/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=subsets-with-xor-value

package main

import "fmt"

func solveBrute(A []int, B int) int {
	count := 0

	for i := 0; i < len(A); i++ {
		currentXor := 0

		for j := i; j < len(A); j++ {
			currentXor ^= A[j]
			if currentXor == B {
				count++
			}
		}
	}

	return count
}
    //     Time Complexity: O(N2)
    // Space Complexity: O(1)

	func solveOptimal(A []int, B int) int {
	visited := make(map[int]int)
	cpx := 0 // cumulative prefix xor
	count := 0

	for _, val := range A {
		cpx ^= val

		// If prefix XOR itself equals B
		if cpx == B {
			count++
		}

		// Check if there exists a prefix with xor = cpx ^ B
		h := cpx ^ B
		if freq, exists := visited[h]; exists {
			count += freq
		}

		visited[cpx]++
	}

	return count
}
    // Time Complexity: O(N)
    // Space Complexity: O(N)

	func main() {
	A := []int{4, 2, 2, 6, 4}
	B := 6

	fmt.Println("Brute:", solveBrute(A, B))
	fmt.Println("Optimal:", solveOptimal(A, B))
}