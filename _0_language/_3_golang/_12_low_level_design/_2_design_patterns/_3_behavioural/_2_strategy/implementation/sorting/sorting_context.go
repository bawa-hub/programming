package sorting

import (
	"fmt"
	"time"
)

type SortingContext struct {
	strategy SortingStrategy
}

func NewSortingContext(strategy SortingStrategy) *SortingContext {
	return &SortingContext{strategy: strategy}
}

func (sc *SortingContext) SetStrategy(strategy SortingStrategy) {
	sc.strategy = strategy
}

func (sc *SortingContext) Sort(data []int) []int {
	start := time.Now()
	result := sc.strategy.Sort(data)
	duration := time.Since(start)
	
	fmt.Printf("Sorted with %s in %v: %v\n", 
		sc.strategy.GetName(), duration, result)
	return result
}

func (sc *SortingContext) GetCurrentStrategy() string {
	return sc.strategy.GetName()
}