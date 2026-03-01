// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

func threeSumBrute(arr []int) [][]int {
	n := len(arr)
	set := make(map[[3]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					temp := []int{arr[i], arr[j], arr[k]}
					sort.Ints(temp)

					key := [3]int{temp[0], temp[1], temp[2]}
					set[key] = true
				}
			}
		}
	}

	var ans [][]int
	for k := range set {
		ans = append(ans, []int{k[0], k[1], k[2]})
	}
	return ans
}
// Time Complexity: O(N3 * log(no. of unique triplets)), where N = size of the array.
// Reason: Here, we are mainly using 3 nested loops. And inserting triplets into the set takes O(log(no. of unique triplets)) time complexity. But we are not considering the time complexity of sorting as we are just sorting 3 elements every time.
// Space Complexity: O(2 * no. of the unique triplets) as we are using a set data structure and a list to store the triplets.

func threeSumBetter(arr []int) [][]int {
	n := len(arr)
	set := make(map[[3]int]bool)

	for i := 0; i < n; i++ {
		hashset := make(map[int]bool)

		for j := i + 1; j < n; j++ {
			third := -(arr[i] + arr[j])

			if hashset[third] {
				temp := []int{arr[i], arr[j], third}
				sort.Ints(temp)
				key := [3]int{temp[0], temp[1], temp[2]}
				set[key] = true
			}

			hashset[arr[j]] = true
		}
	}

	var ans [][]int
	for k := range set {
		ans = append(ans, []int{k[0], k[1], k[2]})
	}
	return ans
}
// Time Complexity: O(N2 * log(no. of unique triplets)), where N = size of the array.
// Reason: Here, we are mainly using 3 nested loops. And inserting triplets into the set takes O(log(no. of unique triplets)) time complexity. But we are not considering the time complexity of sorting as we are just sorting 3 elements every time.
// Space Complexity: O(2 * no. of the unique triplets) + O(N) as we are using a set data structure and a list to store the triplets and extra O(N) for storing the array elements in another set.

func threeSumOptimal(arr []int) [][]int {
	n := len(arr)
	sort.Ints(arr)

	var ans [][]int

	for i := 0; i < n; i++ {

		// Skip duplicates for i
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			sum := arr[i] + arr[j] + arr[k]

			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				ans = append(ans, []int{arr[i], arr[j], arr[k]})
				j++
				k--

				// Skip duplicates
				for j < k && arr[j] == arr[j-1] {
					j++
				}
				for j < k && arr[k] == arr[k+1] {
					k--
				}
			}
		}
	}
	return ans
}
// Time Complexity: O(NlogN)+O(N2), where N = size of the array.
// Reason: The pointer i, is running for approximately N times. And both the pointers j and k combined can run for approximately N times including the operation of skipping duplicates. So the total time complexity will be O(N2). 
// Space Complexity: O(no. of quadruplets), This space is only used to store the answer. We are not using any extra space to solve this problem. So, from that perspective, space complexity can be written as O(1).


func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSumOptimal(arr)

	for _, triplet := range ans {
		fmt.Println(triplet)
	}
}