package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
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
