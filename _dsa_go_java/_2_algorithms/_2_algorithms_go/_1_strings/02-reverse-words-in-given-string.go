// https://leetcode.com/problems/reverse-words-in-a-string/

package main

import (
	"fmt"
)

func reverseWords(s string) string {
	res := ""
	i := len(s) - 1

	for i >= 0 {

		// Skip spaces
		if s[i] == ' ' {
			i--
		} else {

			j := i
			temp := " "

			// Collect word in reverse
			for j >= 0 && s[j] != ' ' {
				temp += string(s[j])
				j--
			}

			// Reverse the collected word
			runes := []rune(temp)
			for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}

			res += string(runes)
			i = j
		}
	}

	// Remove last extra space
	if len(res) > 0 {
		res = res[1:] // remove leading space
	}

	return res
}
// TC:O(N) for traversing the string + O(N) for reversing
// SC:O(N)

func main() {
	s := "  hello   world  "
	fmt.Println(reverseWords(s)) // "world hello"
}