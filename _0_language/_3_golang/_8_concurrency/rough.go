package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup

	task := 10
	worker := 3

	jobs := make(chan int, task)
	results := make(chan int, task)

	for i:=0;i<worker;i++ {
		wg.Add(1)
       go doWork(jobs, results, &wg)
	}

	for i:=1;i<=task;i++ {
		jobs <- i
	}

	close(jobs)

	go func ()  {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Resuls : ", result)
	}

	

}

func doWork(jobs <- chan int, results chan<- int, wg *sync.WaitGroup) {
	wg.Done()
	for job := range jobs {
		result := job * job
		results <- result
	}
}