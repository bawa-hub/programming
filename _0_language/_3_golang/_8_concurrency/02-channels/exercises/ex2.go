package exercises

import "fmt"

// Exercise 2: Buffered vs Unbuffered
// Compare the behavior of buffered and unbuffered channels.
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered vs Unbuffered")
	fmt.Println("===================================")
	
	// Unbuffered channel
	unbuffered := make(chan int)
	
	// Buffered channel
	buffered := make(chan int, 3)
	
	// Test unbuffered
	fmt.Println("Unbuffered channel test:")
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Sending %d to unbuffered channel\n", i)
			unbuffered <- i
		}
		close(unbuffered)
	}()
	
	for value := range unbuffered {
		fmt.Printf("Received %d from unbuffered channel\n", value)
	}
	
	// Test buffered
	fmt.Println("\nBuffered channel test:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sending %d to buffered channel\n", i)
		buffered <- i
	}
	close(buffered)
	
	for value := range buffered {
		fmt.Printf("Received %d from buffered channel\n", value)
	}
}