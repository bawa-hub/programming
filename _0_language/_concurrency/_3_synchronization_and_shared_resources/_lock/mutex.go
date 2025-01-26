package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    lock    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    lock.Lock()
    for i := 0; i < 1000; i++ {
        counter++
    }
    lock.Unlock()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go increment(&wg)
    go increment(&wg)

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
