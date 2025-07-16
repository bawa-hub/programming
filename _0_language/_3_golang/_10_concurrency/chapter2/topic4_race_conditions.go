// 🧠 What’s a Race Condition?

// A race condition occurs when:

//     Two or more goroutines access shared data

//     At least one goroutine writes

//     The accesses happen concurrently

// Result? ❌ Unpredictable behavior — maybe it works, maybe it crashes, maybe it corrupts memory.

package main

import (
	"fmt"
)

var counter = 0

func increment() {
	counter++
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	fmt.Println("Final counter:", counter)
}


// 😱 What Happens Here?

//     1000 goroutines increment counter.

//     But counter++ is not atomic — it’s 3 steps: read → add → write.

//     Goroutines interleave, leading to missed updates.

// You might expect Final counter: 1000, but you’ll likely get something way lower (like 743).


// 🔍 Detecting Race Conditions

// Go has built-in race detection! 🧪
// Run your code with:

// go run -race main.go

// You’ll see detailed info on where data races occur.

// ✅ Solutions: How to Fix Race Conditions in Go
// sync.Mutex — Mutual Exclusion Lock
// sync/atomic — Lock-free Atomic Ops