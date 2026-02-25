// https://leetcode.com/problems/max-consecutive-ones/

func findMaxConsecutiveOnes(nums []int) int {
	cnt, maxi := 0, 0

	for _, v := range nums {
		if v == 1 {
			cnt++
			if cnt > maxi {
				maxi = cnt
			}
		} else {
			cnt = 0
		}
	}

	return maxi
}