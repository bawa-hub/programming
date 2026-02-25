// https://leetcode.com/problems/largest-odd-number-in-string/

package main

import "fmt"

func largestOddNumber(num string) string {

	for i := len(num) - 1; i >= 0; i-- {

		// convert char to digit
		digit := num[i] - '0'

		if digit%2 != 0 {
			return num[:i+1]
		}
	}

	return ""
}

func main() {
	fmt.Println(largestOddNumber("52"))       // "5"
	fmt.Println(largestOddNumber("4206"))     // ""
	fmt.Println(largestOddNumber("35427"))    // "35427"
}