package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Custom types for demonstration
type MemoryMonitor struct {
	lastGC     time.Time
	gcCount    uint32
	allocCount uint64
}

func (m *MemoryMonitor) Update() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	m.gcCount = stats.NumGC
	m.allocCount = stats.TotalAlloc
}

func (m *MemoryMonitor) String() string {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return fmt.Sprintf("GC Count: %d, Total Alloc: %d KB", 
		stats.NumGC, stats.TotalAlloc/1024)
}

type GoroutinePool struct {
	workers int
	jobs    chan func()
	wg      sync.WaitGroup
}

func NewGoroutinePool(workers int) *GoroutinePool {
	return &GoroutinePool{
		workers: workers,
		jobs:    make(chan func(), workers*2),
	}
}

func (p *GoroutinePool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

func (p *GoroutinePool) worker(id int) {
	defer p.wg.Done()
	for job := range p.jobs {
		job()
	}
}

func (p *GoroutinePool) Submit(job func()) {
	p.jobs <- job
}

func (p *GoroutinePool) Close() {
	close(p.jobs)
	p.wg.Wait()
}

type ResourceMonitor struct {
	startTime time.Time
	initialMem runtime.MemStats
}

func NewResourceMonitor() *ResourceMonitor {
	rm := &ResourceMonitor{
		startTime: time.Now(),
	}
	runtime.ReadMemStats(&rm.initialMem)
	return rm
}

func (rm *ResourceMonitor) Report() {
	var currentMem runtime.MemStats
	runtime.ReadMemStats(&currentMem)
	
	elapsed := time.Since(rm.startTime)
	fmt.Printf("Runtime Report (elapsed: %v):\n", elapsed)
	fmt.Printf("  Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("  Memory Alloc: %d KB (was %d KB)\n", 
		currentMem.Alloc/1024, rm.initialMem.Alloc/1024)
	fmt.Printf("  Memory Sys: %d KB (was %d KB)\n", 
		currentMem.Sys/1024, rm.initialMem.Sys/1024)
	fmt.Printf("  GC Cycles: %d (was %d)\n", 
		currentMem.NumGC, rm.initialMem.NumGC)
}

func main() {
	fmt.Println("🚀 Go runtime Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic Runtime Information
	fmt.Println("\n1. Basic Runtime Information:")
	
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Compiler: %s\n", runtime.Compiler)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Max processors: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of CGO calls: %d\n", runtime.NumCgoCall())

	// 2. Memory Statistics
	fmt.Println("\n2. Memory Statistics:")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys: %d KB\n", m.Sys/1024)
	fmt.Printf("Lookups: %d\n", m.Lookups)
	fmt.Printf("Mallocs: %d\n", m.Mallocs)
	fmt.Printf("Frees: %d\n", m.Frees)
	fmt.Printf("HeapAlloc: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("HeapSys: %d KB\n", m.HeapSys/1024)
	fmt.Printf("HeapIdle: %d KB\n", m.HeapIdle/1024)
	fmt.Printf("HeapInuse: %d KB\n", m.HeapInuse/1024)
	fmt.Printf("HeapReleased: %d KB\n", m.HeapReleased/1024)
	fmt.Printf("HeapObjects: %d\n", m.HeapObjects)
	fmt.Printf("StackInuse: %d KB\n", m.StackInuse/1024)
	fmt.Printf("StackSys: %d KB\n", m.StackSys/1024)
	fmt.Printf("MSpanInuse: %d KB\n", m.MSpanInuse/1024)
	fmt.Printf("MSpanSys: %d KB\n", m.MSpanSys/1024)
	fmt.Printf("MCacheInuse: %d KB\n", m.MCacheInuse/1024)
	fmt.Printf("MCacheSys: %d KB\n", m.MCacheSys/1024)
	fmt.Printf("BuckHashSys: %d KB\n", m.BuckHashSys/1024)
	fmt.Printf("GCSys: %d KB\n", m.GCSys/1024)
	fmt.Printf("OtherSys: %d KB\n", m.OtherSys/1024)
	fmt.Printf("NextGC: %d KB\n", m.NextGC/1024)
	fmt.Printf("LastGC: %v\n", time.Unix(0, int64(m.LastGC)))
	fmt.Printf("PauseTotalNs: %d ns\n", m.PauseTotalNs)
	fmt.Printf("NumGC: %d\n", m.NumGC)
	fmt.Printf("NumForcedGC: %d\n", m.NumForcedGC)
	fmt.Printf("GCCPUFraction: %f\n", m.GCCPUFraction)
	fmt.Printf("EnableGC: %t\n", m.EnableGC)
	fmt.Printf("DebugGC: %t\n", m.DebugGC)

	// 3. Goroutine Management
	fmt.Println("\n3. Goroutine Management:")
	
	fmt.Printf("Initial goroutines: %d\n", runtime.NumGoroutine())
	
	// Create some goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d: Current goroutines: %d\n", 
				id, runtime.NumGoroutine())
		}(i)
	}
	
	fmt.Printf("After starting goroutines: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("After goroutines finished: %d\n", runtime.NumGoroutine())

	// 4. Garbage Collection
	fmt.Println("\n4. Garbage Collection:")
	
	fmt.Printf("Before GC - Alloc: %d KB\n", m.Alloc/1024)
	
	// Allocate some memory
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024) // 1KB each
	}
	
	runtime.ReadMemStats(&m)
	fmt.Printf("After allocation - Alloc: %d KB\n", m.Alloc/1024)
	
	// Force garbage collection
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("After GC - Alloc: %d KB\n", m.Alloc/1024)
	
	// Note: SetGCPercent is not available in all Go versions
	fmt.Printf("GC percent setting not available in this Go version\n")

	// 5. Processor Management
	fmt.Println("\n5. Processor Management:")
	
	oldMaxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Current GOMAXPROCS: %d\n", oldMaxProcs)
	
	// Set to number of CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Set GOMAXPROCS to %d\n", runtime.NumCPU())
	
	// Restore original value
	runtime.GOMAXPROCS(oldMaxProcs)
	fmt.Printf("Restored GOMAXPROCS to %d\n", oldMaxProcs)

	// 6. Goroutine Yielding
	fmt.Println("\n6. Goroutine Yielding:")
	
	fmt.Println("Before Gosched()")
	runtime.Gosched() // Yield processor
	fmt.Println("After Gosched()")

	// 7. Stack Information
	fmt.Println("\n7. Stack Information:")
	
	// Get current stack
	stack := make([]byte, 4096)
	n := runtime.Stack(stack, false)
	fmt.Printf("Current stack (first 500 chars):\n%s\n", 
		string(stack[:min(500, n)]))

	// 8. Caller Information
	fmt.Println("\n8. Caller Information:")
	
	// Get caller information
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Printf("Current function: %s\n", runtime.FuncForPC(pc).Name())
		fmt.Printf("File: %s\n", file)
		fmt.Printf("Line: %d\n", line)
	}

	// 9. Memory Monitoring
	fmt.Println("\n9. Memory Monitoring:")
	
	monitor := &MemoryMonitor{}
	monitor.Update()
	fmt.Printf("Initial monitor: %s\n", monitor.String())
	
	// Allocate some memory
	_ = make([]byte, 1024*1024) // 1MB
	
	monitor.Update()
	fmt.Printf("After allocation: %s\n", monitor.String())

	// 10. Goroutine Pool
	fmt.Println("\n10. Goroutine Pool:")
	
	pool := NewGoroutinePool(3)
	pool.Start()
	
	// Submit some jobs
	for i := 0; i < 10; i++ {
		jobID := i
		pool.Submit(func() {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Job %d completed by goroutine\n", jobID)
		})
	}
	
	pool.Close()
	fmt.Println("All jobs completed")

	// 11. Resource Monitoring
	fmt.Println("\n11. Resource Monitoring:")
	
	rm := NewResourceMonitor()
	
	// Simulate some work
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024)
		if i%100 == 0 {
			runtime.GC()
		}
	}
	
	rm.Report()

	// 12. Finalizers
	fmt.Println("\n12. Finalizers:")
	
	type Resource struct {
		id   int
		data []byte
	}
	
	resource := &Resource{
		id:   1,
		data: make([]byte, 1024),
	}
	
	// Set finalizer
	runtime.SetFinalizer(resource, func(r *Resource) {
		fmt.Printf("Finalizer called for resource %d\n", r.id)
	})
	
	// Clear reference to trigger finalizer
	resource = nil
	runtime.GC()
	runtime.GC() // Second GC to ensure finalizer runs

	// 13. Memory Profile Rate
	fmt.Println("\n13. Memory Profile Rate:")
	
	oldRate := runtime.MemProfileRate
	fmt.Printf("Old memory profile rate: %d\n", oldRate)
	
	// Set memory profile rate
	runtime.MemProfileRate = 1
	fmt.Printf("New memory profile rate: %d\n", runtime.MemProfileRate)
	
	// Restore original rate
	runtime.MemProfileRate = oldRate
	fmt.Printf("Restored memory profile rate: %d\n", oldRate)

	// 14. Block Profile Rate
	fmt.Println("\n14. Block Profile Rate:")
	
	runtime.SetBlockProfileRate(1)
	fmt.Printf("Set block profile rate to 1\n")
	
	// Note: SetBlockProfileRate doesn't return the old value
	fmt.Printf("Block profile rate set\n")

	// 15. Mutex Profile Fraction
	fmt.Println("\n15. Mutex Profile Fraction:")
	
	oldMutexRate := runtime.SetMutexProfileFraction(1)
	fmt.Printf("Old mutex profile fraction: %d\n", oldMutexRate)
	
	// Restore original rate
	runtime.SetMutexProfileFraction(oldMutexRate)
	fmt.Printf("Restored mutex profile fraction: %d\n", oldMutexRate)

	// 16. CPU Profile Rate
	fmt.Println("\n16. CPU Profile Rate:")
	
	runtime.SetCPUProfileRate(100)
	fmt.Printf("Set CPU profile rate to 100\n")
	
	// Note: SetCPUProfileRate doesn't return the old value
	fmt.Printf("CPU profile rate set\n")

	// 17. KeepAlive
	fmt.Println("\n17. KeepAlive:")
	
	// Create a value that might be garbage collected
	value := make([]byte, 1024)
	
	// Use KeepAlive to prevent GC
	runtime.KeepAlive(value)
	fmt.Println("Value kept alive")

	// 18. Goroutine Profile
	fmt.Println("\n18. Goroutine Profile:")
	
	// Get goroutine profile
	goroutineProfile := make([]runtime.StackRecord, 10)
	nGoroutine, okGoroutine := runtime.GoroutineProfile(goroutineProfile)
	if okGoroutine {
		fmt.Printf("Goroutine profile: %d records\n", nGoroutine)
		for i := 0; i < nGoroutine && i < 3; i++ {
			fmt.Printf("Record %d: Stack0 length: %d\n", i, len(goroutineProfile[i].Stack0))
		}
	}

	// 19. Thread Creation Profile
	fmt.Println("\n19. Thread Creation Profile:")
	
	// Get thread creation profile
	threadProfile := make([]runtime.StackRecord, 10)
	n, ok = runtime.ThreadCreateProfile(threadProfile)
	if ok {
		fmt.Printf("Thread creation profile: %d records\n", n)
	}

	// 20. Memory Profile
	fmt.Println("\n20. Memory Profile:")
	
	// Get memory profile
	memProfile := make([]runtime.MemProfileRecord, 10)
	n, ok = runtime.MemProfile(memProfile, true)
	if ok {
		fmt.Printf("Memory profile: %d records\n", n)
		for i := 0; i < n && i < 3; i++ {
			fmt.Printf("Record %d: AllocBytes=%d, FreeBytes=%d, AllocObjects=%d, FreeObjects=%d\n",
				i, memProfile[i].AllocBytes, memProfile[i].FreeBytes,
				memProfile[i].AllocObjects, memProfile[i].FreeObjects)
		}
	}

	// 21. Block Profile
	fmt.Println("\n21. Block Profile:")
	
	// Get block profile
	blockProfile := make([]runtime.BlockProfileRecord, 10)
	n, ok = runtime.BlockProfile(blockProfile)
	if ok {
		fmt.Printf("Block profile: %d records\n", n)
	}

	// 22. Mutex Profile
	fmt.Println("\n22. Mutex Profile:")
	
	// Get mutex profile
	mutexProfile := make([]runtime.BlockProfileRecord, 10)
	n, ok = runtime.MutexProfile(mutexProfile)
	if ok {
		fmt.Printf("Mutex profile: %d records\n", n)
	}

	// 23. CPU Profile
	fmt.Println("\n23. CPU Profile:")
	
	// Note: CPUProfile is not available in all Go versions
	fmt.Printf("CPU profile not available in this Go version\n")

	// 24. Performance Test
	fmt.Println("\n24. Performance Test:")
	
	// Test memory allocation performance
	start := time.Now()
	
	// Allocate memory in chunks
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024)
	}
	
	allocTime := time.Since(start)
	fmt.Printf("Allocated 1000 chunks of 1KB in %v\n", allocTime)
	
	// Test GC performance
	start = time.Now()
	runtime.GC()
	gcTime := time.Since(start)
	fmt.Printf("Garbage collection took %v\n", gcTime)

	// 25. System Information Summary
	fmt.Println("\n25. System Information Summary:")
	
	var finalStats runtime.MemStats
	runtime.ReadMemStats(&finalStats)
	
	fmt.Printf("Final goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Final memory alloc: %d KB\n", finalStats.Alloc/1024)
	fmt.Printf("Total allocations: %d\n", finalStats.Mallocs)
	fmt.Printf("Total frees: %d\n", finalStats.Frees)
	fmt.Printf("GC cycles: %d\n", finalStats.NumGC)
	fmt.Printf("GC CPU fraction: %f\n", finalStats.GCCPUFraction)

	fmt.Println("\n🎉 runtime Package Mastery Complete!")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
