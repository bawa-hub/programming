package main

import "fmt"

var pl = fmt.Println

func main() {
	// ----- ARRAYS -----
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	// Declare integer array with 5 elements
	var arr1 [5]int

	// Assign value to index
	arr1[0] = 1

	// Declare and initialize
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Get by index
	pl("Index 0 :", arr2[0])

	// Length
	pl("Arr Length :", len(arr2))

	// Iterate with index
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// Iterate with range
	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}

	// Multidimensional Array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print multidimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	// String into slice of runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)

	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make
	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("Slice Size :", len(sl1))

	// Cycle with for
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}

	// Create a slice literal
	sl2 := []int{12, 21, 1974}
	pl(sl2)

	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)

	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show
	// as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}
