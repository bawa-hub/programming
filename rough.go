package main

import (
	"fmt"
)

func pool(jobs <- chan int, results chan <- int) {
	for job := range jobs {
		results <- job * job
	}
}

func main() {
   
	tasks := 10
	workers := 3


	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	for i:=0;i<workers;i++ {

		go pool(jobs, results)
	}

	for i:=1;i<=tasks;i++ {
        jobs <- i
	}
		close(jobs)

	go func ()  {
		close(results)
	}()

	for i:=0;i<len(results);i++ {
		fmt.Println(<-results)
	}
}