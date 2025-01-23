package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	// ----- COMMAND LINE ARGUMENTS -----
	// You can pass values to your program
	// from the command line
	// Create cltest.go
	// go build cltest.go
	// .\cltest 24 43 12 9 10  // ./cltest in mac
	// Returns an array with everything
	// passed with the name of the app
	// in the first index
	// Outputs the max number passed in

	// make build: go build <file_path>
	// Then run ./<file_name> ...args

	pl(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl("Max value: ", max)

}
