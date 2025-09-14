package main

import "fmt"

var pl = fmt.Println

func main() {

	// Collection of values with the same data type and the size can't be changed

	var arr1 [5]int // Declare integer array with 5 elements
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5} 	// Declare and initialize
	pl("Index 0 :", arr2[0])

	pl("Arr Length :", len(arr2)) 	// Length

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
}
