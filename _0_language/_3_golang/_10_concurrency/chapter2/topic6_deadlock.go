// 🚨 1. What is a Deadlock?
//     A deadlock happens when two or more goroutines are waiting on each other to release a lock — and no one proceeds.

// 🔒 Classic Deadlock Example:

package main

import (
	"fmt"
	"sync"
)

var mu1, mu2 sync.Mutex

func main() {
	go func() {
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("Goroutine 1 finished")
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()

		mu1.Lock()
		defer mu1.Unlock()

		fmt.Println("Goroutine 2 finished")
	}()

	select {} // block forever
}

// Goroutine 1 locks mu1, waits for mu2
// Goroutine 2 locks mu2, waits for mu1
// → Both wait forever ❌


// ⚠️ 2. What is Starvation?

//     A goroutine never gets CPU time or the lock because other goroutines keep taking it.

// Usually caused by:

//     Holding locks too long

//     Busy goroutines spamming access

// 💡 Best Practices to Avoid Deadlocks & Starvation
// Practice | Why
// Always lock in the same order | Prevents circular wait (classic deadlock cause)
// Use defer mu.Unlock() immediately after mu.Lock() | Guarantees unlock
// Keep critical sections short | Avoid starvation
// Avoid sleeping or I/O in locked sections | Delays release, may starve others
// Minimize lock contention | Structure code to avoid frequent lock clashes

// 🔧 Real-World Example: Lock Order Fix

// BAD: Can deadlock
// func A() {
// 	mu1.Lock()
// 	mu2.Lock()
// 	// ...
// 	mu2.Unlock()
// 	mu1.Unlock()
// }

// func B() {
// 	mu2.Lock()
// 	mu1.Lock()
// 	// ...
// 	mu1.Unlock()
// 	mu2.Unlock()
// }

// ✅ FIX: Always lock in the same order!
// // GOOD: Consistent lock order
// func A() {
// 	mu1.Lock()
// 	mu2.Lock()
// 	// ...
// 	mu2.Unlock()
// 	mu1.Unlock()
// }

// func B() {
// 	mu1.Lock() // lock mu1 first here too
// 	mu2.Lock()
// 	// ...
// 	mu2.Unlock()
// 	mu1.Unlock()
// }


// 🧠 Interview Q&A

// Q: What causes a deadlock in Go?

//     Goroutines waiting on locks held by each other in a cycle.

// Q: How do you avoid it?

//     Lock in consistent order, unlock with defer, keep critical sections short.

// Q: What’s starvation?

//     When a goroutine never gets scheduled because others dominate CPU or locks.

// Q: When is defer mu.Unlock() not safe?

//     Inside loops where the lock is re-acquired — use manually in that case.



// ✅ Recap

//     💀 Deadlocks = everyone waiting on each other

//     🥶 Starvation = someone always last in line

//     🔐 Best practices = defer, short locks, consistent order