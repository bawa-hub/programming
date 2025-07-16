// 🧠 Why This Matters

// You just launched 100 goroutines — ever wondered:

//     How are they executed so fast?

//     Why don’t they crash your system like 100 OS threads?

//     How does Go decide which goroutine runs when?

// Understanding the Go scheduler answers all of that — and is pure gold for interviews.



// 🔩 Go’s M:N Scheduler

// Go uses an M:N scheduler, meaning:

//     M goroutines mapped onto N OS threads.

// This is managed by a smart runtime system that juggles execution efficiently across CPU cores.


// 🔧 Internal Components
// Component | Meaning
// G | A goroutine (your code)
// M | An OS thread (where G runs)
// P | A processor (executes Go code, holds run queue)

// G = actual goroutine
// M = worker thread (bound to CPU thread)
// P = logical processor (schedules G on M)
// Each P maintains a run queue of goroutines and feeds them to an attached M.

// 🖼️ Visual Model
// ┌────────┐       ┌────────┐       ┌────────┐
// │   G1   │       │   G2   │  ...  │   G100 │
// └────────┘       └────────┘       └────────┘
// 	 ↓                ↓                  ↓
// [Run Queue of P1]  [Run Queue of P2]  [Run Queue of P3]
// 	   ↓               ↓                 ↓
//    ┌────────┐     ┌────────┐       ┌────────┐
//    │   M1   │     │   M2   │  ...  │   Mn   │  → OS Threads
//    └────────┘     └────────┘       └────────┘



// ⚙️ How It Works

//     You write go myFunc() → creates a new G

//     Go runtime assigns G to a processor (P)

//     P adds it to its run queue

//     An M attached to that P executes the G

//     If G blocks (e.g., on I/O), M picks another G from the queue

//     Goroutines can migrate across Ps/Ms


// 🔁 Preemptive Scheduling
//     Go uses cooperative preemption, but also adds runtime preemption since Go 1.14.

//     Goroutines yield at safe points (e.g., function calls) to ensure fair scheduling.



// ⚡️ Efficiency

//     Goroutines are super cheap: ~2KB stack (grows/shrinks dynamically)

//     Can scale to millions of goroutines

//     Threads and cores are efficiently shared

// 🔍 Interview Nuggets

// Q: Why are goroutines cheaper than threads?

//     They don’t map 1:1 to threads, use less memory (2KB vs 1MB), and are scheduled by Go, not OS.

// Q: What are G, M, and P in Go’s runtime?

//         G: goroutine

//         M: OS thread

//         P: logical processor (schedules goroutines to M)

// Q: What if I create 10 million goroutines?

//     Go will handle them unless they all block or consume memory. Each one is paused/resumed smartly by the scheduler.