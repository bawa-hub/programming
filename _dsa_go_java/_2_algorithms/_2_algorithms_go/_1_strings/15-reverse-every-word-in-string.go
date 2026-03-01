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
		for i >= 0 && s[i] == ' ' {
			i--
		}

		if i < 0 {
			break
		}

		j := i

		// Find word start
		for j >= 0 && s[j] != ' ' {
			j--
		}

		// Append word
		res += s[j+1 : i+1]

		// Add space if not last word
		if j > 0 {
			res += " "
		}

		i = j - 1
	}

	return res
}

func main() {
	fmt.Println(reverseWords("  hello   world  "))
	fmt.Println(reverseWords("a good   example"))
}