package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	// Precompute character frequencies (ASCII only)
	hash := make([]int, 256)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var c string
		fmt.Scan(&c)
		fmt.Println(hash[c[0]])
	}
}


// using map
// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	// Precompute character frequencies
// 	hash := make(map[rune]int)
// 	for _, ch := range s {
// 		hash[ch]++
// 	}

// 	var q int
// 	fmt.Scan(&q)

// 	for i := 0; i < q; i++ {
// 		var c string
// 		fmt.Scan(&c)

// 		// Take first rune of input (since Go handles UTF-8 properly)
// 		r := []rune(c)[0]
// 		fmt.Println(hash[r])
// 	}
// }
