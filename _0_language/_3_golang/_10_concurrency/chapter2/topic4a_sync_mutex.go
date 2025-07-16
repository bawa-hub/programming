// sync.Mutex — Mutual Exclusion Lock

package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

// ✅ Now the output is always 1000, and no race.