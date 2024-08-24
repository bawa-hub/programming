package main

import (
	"fmt"
)

var pl = fmt.Println

// passing by reference
func changeValue(ptr *int) {
	*ptr = 10
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	f4 := 5
	var f4Ptr *int = &f4
	pl("f4 Address: ", f4Ptr)
	pl("f4 Value: ", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value: ", *f4Ptr)
	pl("f4 before func: ", f4)
	changeValue(&f4)
	pl("f3 after func: ", f4)

	// passing arrays as pointers
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)
}
