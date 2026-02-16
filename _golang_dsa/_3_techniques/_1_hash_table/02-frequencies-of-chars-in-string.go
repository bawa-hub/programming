package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	// precompute frequency
	hash := make([]int, 256)

	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	var q int
	fmt.Scan(&q)

	for q > 0 {
		var c string
		fmt.Scan(&c)

		hashValue := hash[c[0]]
		fmt.Println(hashValue)

		q--
	}
}
