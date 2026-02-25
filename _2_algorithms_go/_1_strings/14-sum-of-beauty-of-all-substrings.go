// https://leetcode.com/problems/sum-of-beauty-of-all-substrings/

package main

import "fmt"

func beautySum(s string) int {
	n := len(s)
	ans := 0

	for i := 0; i < n; i++ {
		var cnt [26]int

		for j := i; j < n; j++ {
			cnt[s[j]-'a']++

			maxF := 0
			minF := 1<<31 - 1

			for k := 0; k < 26; k++ {
				if cnt[k] > 0 {
					if cnt[k] > maxF {
						maxF = cnt[k]
					}
					if cnt[k] < minF {
						minF = cnt[k]
					}
				}
			}

			ans += (maxF - minF)
		}
	}

	return ans
}

func main() {
	fmt.Println(beautySum("aabcb")) // 5
}