// https://leetcode.com/problems/online-stock-span/
// https://practice.geeksforgeeks.org/problems/stock-span-problem-1587115621/1
// https://www.geeksforgeeks.org/the-stock-span-problem/

package main

import "fmt"

type Pair struct {
	index int
	price int
}

type StockSpanner struct {
	stack []Pair
	index int
}

/**************** CONSTRUCTOR ****************/

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []Pair{},
		index: -1,
	}
}

/**************** NEXT ****************/

func (s *StockSpanner) Next(price int) int {

	s.index++

	for len(s.stack) > 0 &&
		s.stack[len(s.stack)-1].price <= price {

		s.stack = s.stack[:len(s.stack)-1]
	}

	if len(s.stack) == 0 {
		s.stack = append(s.stack, Pair{s.index, price})
		return s.index + 1
	}

	prevIndex := s.stack[len(s.stack)-1].index
	s.stack = append(s.stack, Pair{s.index, price})

	return s.index - prevIndex
}
//   Time complexity: O(N)
//  Space complexity:O(N)

/**************** TEST ****************/

func main() {

	spanner := Constructor()

	fmt.Println(spanner.Next(100)) // 1
	fmt.Println(spanner.Next(80))  // 1
	fmt.Println(spanner.Next(60))  // 1
	fmt.Println(spanner.Next(70))  // 2
	fmt.Println(spanner.Next(60))  // 1
	fmt.Println(spanner.Next(75))  // 4
	fmt.Println(spanner.Next(85))  // 6
}