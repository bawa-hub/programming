package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getMultipleSUm(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// pass by value
func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {

	// functions
	// func funcName(params) returnType {BODY}
	sayHello()
	pl(getSum(1, 2))
	pl(getMultipleSUm(3))
	pl(getQuotient(3, 4))
	pl(getQuotient(3, 0))

	// varadic functions
	pl(getSum2(1, 2))

	// passing arrays to functions
	vArr := []int{1, 2, 3}
	pl(getArraySum(vArr))

}
