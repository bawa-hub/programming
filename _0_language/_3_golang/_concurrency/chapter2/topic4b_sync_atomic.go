// sync/atomic — Lock-free Atomic Ops

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

// ✅ Also race-free — and even faster than using a mutex.

// 🧠 Interview Insights
// Question | Smart Answer
// What’s a race condition? | Concurrent access to shared data where at least one access is a write, leading to undefined behavior.
// How do you fix it in Go? | Use sync.Mutex, sync/atomic, or channels.
// Is counter++ atomic in Go? | No — it's three separate operations.
// How do you detect races? | Use go run -race.

// ✅ Recap

//     Race conditions = subtle and dangerous.

//     Use mutexes or atomic ops to protect shared data.

//     Use -race to find hidden bugs early.