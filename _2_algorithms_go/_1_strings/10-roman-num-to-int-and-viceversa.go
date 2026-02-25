// https://leetcode.com/problems/roman-to-integer/

package main

import "fmt"

func romanToInt(s string) int {

	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0

	for i := 0; i < len(s); i++ {

		if i+1 < len(s) && mp[s[i]] < mp[s[i+1]] {
			ans -= mp[s[i]]
		} else {
			ans += mp[s[i]]
		}
	}

	return ans
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}