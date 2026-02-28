// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// https://practice.geeksforgeeks.org/problems/array-of-alternate-ve-and-ve-nos1401/1
// https://www.geeksforgeeks.org/rearrange-array-alternating-positive-negative-items-o1-extra-space/
// https://takeuforward.org/arrays/rearrange-array-elements-by-sign/

package main

import "fmt"

// brute
func rearrangeBySignBrute(A []int) []int {
	n := len(A)

	var pos []int
	var neg []int

	// Separate positives and negatives
	for _, v := range A {
		if v > 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
	}

	// Place alternately
	for i := 0; i < n/2; i++ {
		A[2*i] = pos[i]
		A[2*i+1] = neg[i]
	}

	return A
}
// Time Complexity: O(N+N/2) { O(N) for traversing the array once for segregating positives and negatives and another O(N/2) for adding those elements alternatively to the array, where N = size of the array A}.
// Space Complexity:  O(N/2 + N/2) = O(N) { N/2 space required for each of the positive and negative element arrays, where N = size of the array A}.


// optimized
func rearrangeBySign(A []int) []int {

	n := len(A)
	ans := make([]int, n)

	posIndex := 0
	negIndex := 1

	for _, v := range A {

		if v < 0 {
			ans[negIndex] = v
			negIndex += 2
		} else {
			ans[posIndex] = v
			posIndex += 2
		}
	}

	return ans
}
// Time Complexity: O(N) { O(N) for traversing the array once and substituting positives and negatives simultaneously using pointers, where N = size of the array A}.
// Space Complexity:  O(N) { Extra Space used to store the rearranged elements separately in an array, where N = size of array A}.

func main() {

	A := []int{1, 2, -4, -5}

	ans := rearrangeBySign(A)

	for _, v := range ans {
		fmt.Print(v, " ")
	}
}