// https://takeuforward.org/data-structure/intersection-of-two-sorted-arrays/

package main

import "fmt"

func intersectionBrute(A, B []int) []int {
	ans := []int{}
	visited := make([]int, len(B))

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {

			if A[i] == B[j] && visited[j] == 0 {
				ans = append(ans, B[j])
				visited[j] = 1
				break
			} else if B[j] > A[i] {
				break
			}
		}
	}

	return ans
}

// Time Complexity: O(n2)
// Space Complexity: O(n) for the extra visited vector

func intersectionSorted(A, B []int) []int {
	i, j := 0, 0
	ans := []int{}

	for i < len(A) && j < len(B) {

		if A[i] < B[j] {
			i++
		} else if B[j] < A[i] {
			j++
		} else {
			ans = append(ans, A[i])
			i++
			j++
		}
	}

	return ans
}
// Time Complexity: O(n) n being the min length of the 2 arrays.
// Space Complexity: O(1)

func main() {
	A := []int{1, 2, 3, 3, 4, 5, 6, 7}
	B := []int{3, 3, 4, 4, 5, 8}

	ans := intersectionSorted(A, B)

	fmt.Println("The elements are:")
	for _, val := range ans {
		fmt.Print(val, " ")
	}
}