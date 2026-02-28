// https://leetcode.com/problems/design-hashset/

package main

import "fmt"

type MyHashSet struct {
	store [][]int
	size  int
}

/************ CONSTRUCTOR ************/

func Constructor() MyHashSet {
	size := 100
	return MyHashSet{
		store: make([][]int, size),
		size:  size,
	}
}

/************ HASH FUNCTION ************/

func (h *MyHashSet) hash(key int) int {
	return key % h.size
}

/************ SEARCH ************/

func (h *MyHashSet) search(key int) (int, int) {
	i := h.hash(key)

	for idx, val := range h.store[i] {
		if val == key {
			return i, idx
		}
	}
	return i, -1
}

/************ ADD ************/

func (h *MyHashSet) Add(key int) {

	if h.Contains(key) {
		return
	}

	i := h.hash(key)
	h.store[i] = append(h.store[i], key)
}

/************ REMOVE ************/

func (h *MyHashSet) Remove(key int) {

	i, idx := h.search(key)
	if idx == -1 {
		return
	}

	bucket := h.store[i]
	h.store[i] = append(bucket[:idx], bucket[idx+1:]...)
}

/************ CONTAINS ************/

func (h *MyHashSet) Contains(key int) bool {

	_, idx := h.search(key)
	return idx != -1
}

/************ TEST ************/

func main() {

	set := Constructor()

	set.Add(10)
	set.Add(20)

	fmt.Println(set.Contains(10)) // true
	fmt.Println(set.Contains(5))  // false

	set.Remove(10)

	fmt.Println(set.Contains(10)) // false
}