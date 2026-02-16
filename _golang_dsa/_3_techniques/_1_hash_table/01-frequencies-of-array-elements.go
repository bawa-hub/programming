// https://practice.geeksforgeeks.org/problems/frequency-of-array-elements-1587115620/0

package main

import "fmt"

func frequency(arr []int) {
	// create map
	mp := make(map[int]int)

	// go idioms
	for _, value := range arr {
		mp[value]++
	}

	// for i := 0; i < len(arr); i++ {
	// 	mp[arr[i]]++
	// }


	// print result
	for key, value := range mp {
		fmt.Println(key, value)
	}
}

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	frequency(arr)
}
