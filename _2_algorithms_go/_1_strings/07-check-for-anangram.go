// https://leetcode.com/problems/valid-anagram/

package main

import (
	"fmt"
	"sort"
)

func checkAnagramsSort(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	s1 := []rune(str1)
	s2 := []rune(str2)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
// Time Complexity: O(nlogn) since sorting function requires nlogn iterations.
// Space Complexity: O(1)


func checkAnagramsHash(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	freq := make([]int, 26)

	for i := 0; i < len(str1); i++ {
		freq[str1[i]-'a']++
		freq[str2[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
// Time Complexity: O(n) where n is the length of string
// Space Complexity: O(1)


// case insensitivity
func checkAnagramsGeneral(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	freq := make(map[rune]int)

	for _, c := range str1 {
		freq[c]++
	}
	for _, c := range str2 {
		freq[c]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}