package main

import (
	"fmt"
	"time"
)


func print() {
    for i := 1;i<=5;i++ {
        fmt.Printf("Number is %d\n", i)
    }
}

func printA() {
    for i := 'a'; i<= 'e';i++ {
        fmt.Println("Char is ", i)
    }
}

func main() {
	fmt.Println("=== Goroutines Exercise ===")
	
	go print()
	go printA()

	time.Sleep(100*time.Millisecond)
	
	fmt.Println("Exercise completed!")
}
