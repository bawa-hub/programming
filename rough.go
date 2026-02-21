package main

import "fmt"

func main() {
	arr := []int{1,2,3,2,4}

	mp := make(map[int]int)

	for _,el := range arr {
		mp[el]++
	}

	for key, val := range mp {
		fmt.Println(key, val)
	}
}