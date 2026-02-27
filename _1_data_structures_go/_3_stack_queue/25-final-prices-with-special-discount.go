// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/solutions/685406/c-stack-next-smaller-element/


func finalPrices(prices []int) []int {

	n := len(prices)
	res := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 &&
			stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - stack[len(stack)-1]
		}

		stack = append(stack, prices[i])
	}

	return res
}

// | Metric | Complexity |
// | ------ | ---------- |
// | Time   | O(N)       |
// | Space  | O(N)       |
