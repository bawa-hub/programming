package exercises

import "fmt"

// Exercise 4: Channel Closing
// Implement proper channel closing and cleanup.
func Exercise4() {
	fmt.Println("\nExercise 4: Channel Closing")
	fmt.Println("===========================")
	
	ch := make(chan int)
	
	// Producer
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Producer: Sent %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Producer: Finished sending")
	}()
	
	// Consumer
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Consumer: Channel closed, exiting")
			break
		}
		fmt.Printf("Consumer: Received %d\n", value)
	}
}