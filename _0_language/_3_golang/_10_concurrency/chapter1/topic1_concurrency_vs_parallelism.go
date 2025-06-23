// 🧠 1. Concurrency vs Parallelism — Core Differences

// Concept | Concurrency | Parallelism
// Definition | Managing multiple tasks at the same time | Performing multiple tasks at the same time
// Goal | Handle many tasks efficiently | Speed up execution using multiple cores
// Execution | Tasks may overlap, but not necessarily | Tasks must run at the same instant
// Analogy | Chef cooking multiple dishes at once (switching between them) | Multiple chefs cooking dishes simultaneously
// Go Support | Fully supported via goroutines | Supported (if goroutines scheduled on different threads/cores)

// 🛠️ Example Analogy
// ✅ Concurrency (1 worker doing many things):

// Imagine a single cashier switching between 3 customers rapidly — handling all without fully finishing any one until later.
// ✅ Parallelism (3 workers doing 3 things):

// Now imagine 3 cashiers — each handling one customer fully at the same time.

// 💡 Key Insight:

//     Concurrency is about structure: how you design your system to handle multiple things.

//     Parallelism is about execution: how your system uses hardware (like multiple CPUs) to do multiple things at once.

// In Go:

//     goroutine = concurrency primitive (cheap to spawn, managed by Go scheduler)

//     Go may schedule goroutines to run in parallel if multiple cores are available

// 📌 Interview Angle

// Q: Can a concurrent program be non-parallel?

//     Yes. You can have many goroutines running concurrently on a single CPU — they appear simultaneous but are just interleaved.

// Q: Can a parallel program be non-concurrent?

//     Rare, but theoretically yes — if you split a task (e.g., matrix multiplication) across cores but don’t interleave tasks.

// 🧠 Summary

//     Concurrency is about handling multiple tasks well.

//     Parallelism is about doing multiple tasks at once.

//     Go supports both: you write concurrent code, and Go (with GOMAXPROCS > 1) can execute it in parallel.