// 👨‍💻 Challenge: Safe Concurrent Counter with Mutex
// Goal:

// Simulate 1000 users logging in concurrently.
// Each goroutine should increment a shared loginCount safely using a mutex.
// use defer

// Using defer is a best practice when working with sync.Mutex because it guarantees the lock is released — even if the function returns early or panics.

package main

import (
	"fmt"
	"sync"
)

var loginCount int
var mu sync.Mutex

func login(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock() // always unlocks, even if code panics later

	loginCount++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go login(&wg)
	}

	wg.Wait()
	fmt.Println("Total logins:", loginCount)
}

// ✅ Why This Is Better
// Without defer | With defer
// You must remember to unlock | Unlock is guaranteed
// Easy to forget in long funcs | Cleaner & safer
// Bug-prone under conditions | Panic-safe

// 🧠 Interview Angle

// Q: Why use defer with mutex unlocks?

//     Because defer guarantees execution, it prevents deadlocks due to early returns, panics, or logic errors.