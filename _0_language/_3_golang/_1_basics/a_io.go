package main


// You could use an alias like f "fmt"
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

/*
I'm a block comment
*/

func main() {
	// Prints text and a newline
	pl("Hello Go")

	pl("What is your name?")

	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')
	
	// It is better to handle it
	name, err := reader.ReadString('\n') // get input upto new line
	if err == nil {
		pl("Hello", name)
	} else {
		// Log this error
		log.Fatal(err)
	}
}
