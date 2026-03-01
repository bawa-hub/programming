// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	minLen := len(first)
	if len(last) < minLen {
		minLen = len(last)
	}

	ans := ""

	for i := 0; i < minLen; i++ {
		if first[i] == last[i] {
			ans += string(first[i])
		} else {
			break
		}
	}

	return ans
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // "fl"
}