package main

import "fmt"


func main() {
	nums := []int{1,2,3,4}

	i := 0;
	j := len(nums)-1;

	for i<j {
		nums[i] , nums[j] = nums[j] , nums[i]
		i++
		j--
	}

	fmt.Println(nums)
}