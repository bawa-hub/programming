package main

import "fmt"

var pl = fmt.Println

func changeVal2(myPtr *int) {
	*myPtr = 12
}

// Receives array by reference and doubles values
func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

type User struct {
    ID   int
    Name string
}

func main() {

	// You can pass by reference with the & (Address of Operator)

	f4 := 10
	pl("f4 :", f4)
	pl("f4 Address :", &f4)

	// Store a pointer (Pointer to int)
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)

	// Print value at pointer
	pl("f4 Value :", *f4Ptr)

	// Assign value using pointer
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	// Change value in function
	pl("f4 before function :", f4)
	changeVal2(&f4)
	pl("f4 after function :", f4)

	// Pass an array by reference
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	// Passing a slice to a function works just
	// like when using variadic functions
	// Just add ... after the slice when passing
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))

	// Zero value of a pointer is nil.  
    // Dereferencing a nil pointer → runtime panic.
	var p *int
	pl(p == nil) // true
	// pl(*p)     // panic

	// Why use pointers?
	// Avoid copying large values (structs, arrays).
	// Allow mutation of values inside functions.
	// Efficiently implement methods on large structs.
	// Shared state (e.g., across goroutines).

	// with structs
	u := User{ID: 1, Name: "Alice"}
    ptr := &u
    ptr.Name = "Bob" // no need to deref explicitly
    pl(u) // {1 Bob}
}
