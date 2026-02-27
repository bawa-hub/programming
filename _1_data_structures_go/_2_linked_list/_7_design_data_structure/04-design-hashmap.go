// https://leetcode.com/problems/design-hashmap/

package main

import "fmt"

/************ HASH MAP ************/

type Pair struct {
	key   int
	value int
}

type MyHashMap struct {
	bucket int
	table  [][]Pair
}

/************ CONSTRUCTOR ************/

func Constructor() MyHashMap {

	bucket := 1000

	return MyHashMap{
		bucket: bucket,
		table:  make([][]Pair, bucket),
	}
}

/************ HASH FUNCTION ************/

func (h *MyHashMap) hash(key int) int {
	return key % h.bucket
}

/************ SEARCH ************/

func (h *MyHashMap) search(key int) (int, int) {

	i := h.hash(key)

	for idx, p := range h.table[i] {
		if p.key == key {
			return i, idx
		}
	}

	return i, -1
}

/************ PUT ************/

func (h *MyHashMap) Put(key int, value int) {

	i, idx := h.search(key)

	if idx != -1 {
		h.table[i][idx].value = value
		return
	}

	h.table[i] = append(h.table[i], Pair{key, value})
}

/************ GET ************/

func (h *MyHashMap) Get(key int) int {

	i, idx := h.search(key)

	if idx == -1 {
		return -1
	}

	return h.table[i][idx].value
}

/************ REMOVE ************/

func (h *MyHashMap) Remove(key int) {

	i, idx := h.search(key)

	if idx == -1 {
		return
	}

	bucket := h.table[i]
	h.table[i] = append(bucket[:idx], bucket[idx+1:]...)
}

/************ TEST ************/

func main() {

	hm := Constructor()

	hm.Put(1, 10)
	hm.Put(2, 20)

	fmt.Println(hm.Get(1)) // 10
	fmt.Println(hm.Get(3)) // -1

	hm.Put(2, 30)

	fmt.Println(hm.Get(2)) // 30

	hm.Remove(2)

	fmt.Println(hm.Get(2)) // -1
}