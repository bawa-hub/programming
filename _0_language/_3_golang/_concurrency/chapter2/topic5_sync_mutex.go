// 🧩 What is a Mutex?
//     A Mutex (short for Mutual Exclusion) is used to lock a critical section so only one goroutine can access it at a time.
// It guarantees safety when goroutines share data.

// ✅ Syntax

// var mu sync.Mutex

// mu.Lock()   // acquire lock
// // critical section
// mu.Unlock() // release lock

// 🧠 If you forget Unlock(), the program deadlocks — it’ll just hang.

// 👨‍💻 Challenge: Safe Concurrent Counter with Mutex
// Goal:

// Simulate 1000 users logging in concurrently.
// Each goroutine should increment a shared loginCount safely using a mutex.

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
	loginCount++
	mu.Unlock()
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


// 🧠 Key Concepts:
// Concept | Meaning
// mu.Lock() | Block other goroutines from entering critical section
// mu.Unlock() | Allow next waiting goroutine to enter
// defer wg.Done() | Signal that goroutine is finished
// Critical section | Code modifying shared data (loginCount++)

// ⚠️ Common Interview Bug: Forgetting Unlock()

// Example of a trap:

// mu.Lock()
// loginCount++
// // forgot mu.Unlock()

// 🔒 This causes a deadlock. Always use defer mu.Unlock() if possible:

// mu.Lock()
// defer mu.Unlock()
// loginCount++


// 🧠 Interview Questions
// Question | Best Answer
// What is a mutex? | A lock that prevents multiple goroutines from accessing shared data at the same time.
// What happens if you don’t unlock? | Deadlock — other goroutines wait forever.
// How is a mutex different from a channel? | Mutex protects shared memory; channels enable communication (often to avoid sharing memory).
// When to use a mutex? | When multiple goroutines read/write shared variables directly.