// https://practice.geeksforgeeks.org/problems/max-sum-subarray-of-size-k5313/1

package main

import (
	"fmt"
	"math"
)

func maximumSumSubarray(K int, arr []int) int64 {
	i, j := 0, 0

	var sum int64 = 0
	maxi := int64(math.MinInt64)

	n := len(arr)

	for j < n {
		sum += int64(arr[j])

		if j-i+1 == K {
			if sum > maxi {
				maxi = sum
			}

			sum -= int64(arr[i])
			i++
		}

		j++
	}

	return maxi
}

func main() {
	arr := []int{100, 200, 300, 400}
	K := 2
	fmt.Println(maximumSumSubarray(K, arr)) // Output: 700
}