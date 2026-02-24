package main

import (
	"fmt"
	"sync"
)

func doWork(jobs <- chan int, results chan<- int, wg *sync.WaitGroup) {
	wg.Done()
	for job := range jobs {
		result := job * job
		results <- result
	}
}

func main() {
	tasks := 10
	workers := 3

	var wg sync.WaitGroup

	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	for i:=0;i<=workers;i++ {
		wg.Add(1)
		go doWork(jobs, results, &wg)
	}

	for i:=1;i<=tasks;i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
        wg.Wait()
		// close(results)
	}()

	for result := range results {
		fmt.Println("result : ", result)
	}
}