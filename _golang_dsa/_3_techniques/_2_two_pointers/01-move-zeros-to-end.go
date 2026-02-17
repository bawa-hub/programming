// https://leetcode.com/problems/move-zeroes/

package main

import "fmt"

// brute force
func zerosToEnd(arr []int) {
	n := len(arr)
	temp := make([]int, n)

	k := 0
	for _, value := range arr {
		if value != 0 {
			temp[k] = value
			k++
		}
	}

	fmt.Println(temp)
}
// Time complexity: o(n)
// Space complexity: o(n)

func zerosToEndSpaceOtimized(arr []int) {
	k := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[k], arr[i] = arr[i], arr[k]
			k++
		}
	}

	fmt.Println(arr)
}
// Time complexity: o(n)
// Space complexity: o(1)


func main() {
	arr := []int{1, 0, 2, 0, 3, 0}
	zerosToEndSpaceOtimized(arr)
}
