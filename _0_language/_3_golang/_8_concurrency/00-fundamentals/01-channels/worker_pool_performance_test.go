package main


import "fmt"
import "time"

func serve(waiter int, jobs <-chan int, results chan<- string) {
    for j := range jobs {
        result := fmt.Sprintf("Waiter %d is serving to Customer %d", waiter, j)
        time.Sleep(10*time.Millisecond)
        results <- result
    }
}

func main() {
    start := time.Now()

	// change the waiter (goroutine) and customer (task) to see the performance magic
    waiter := 10000
    cus := 100000
    
    ch_cus := make(chan int, cus)
    results := make(chan string, cus)
    
    // spawn waiters
    for i := 1; i <= waiter; i++ {
        go serve(i, ch_cus, results)
    }
    
    // add customers
    for i := 1; i <= cus; i++ {
        ch_cus <- i
    }
    close(ch_cus) // ✅ important: no more customers
    
    // collect results
    for i := 1; i <= cus; i++ {
        fmt.Println(<-results)
    }
    
    fmt.Println("Time taken: ", time.Since(start))
}
