package main

import "fmt"

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type IntSliceIterator struct {
	items []int
	idx int
}

func (it *IntSliceIterator) HasNext() bool { return it.idx < len(it.items) }
func (it *IntSliceIterator) Next() int { v := it.items[it.idx]; it.idx++; return v }

func main() {
	it := &IntSliceIterator{items: []int{1,2,3}}
	for it.HasNext() { fmt.Println(it.Next()) }
}
