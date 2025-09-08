package main

import "fmt"

var pl = fmt.Println

func main() {

	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not and Important Birthday")
	}

	// ! turns bools into their opposite value
	pl("!true =", !true)

	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

}
