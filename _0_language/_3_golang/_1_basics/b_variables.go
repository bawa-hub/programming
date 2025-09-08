package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var pl = fmt.Println

func main() {

	// var name type
	// Name must begin with letter and then letters or numbers
	// If a variable, function or type starts with a capital letter, it is considered exported and can be accessed outside the package and otherwise is available only in the current package
	// Camal case is the default naming convention

	var vName string = "Derek"
	var v1, v2 = 1.2, 3.4 // type inferred
	pl(vName, v1, v2)

	// Short variable declaration (Type defined by data)
	var v3 = "Hello"
	pl(v3)

	// Variables are mutable by default (Value can change as long as the data type is the same)
	v4 := 2.4
	pl(v4)

	// After declaring variables to assign values to them always use = there after. If you use := you'll create a new variable

	// costants
	const pi = 3.14
	const (
		Monday  = 1
		Tuesday = 2
	)
	pl(pi, Monday, Tuesday)

	// ----- DATA TYPES -----
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	// pointer/map/slice/channel/function → nil
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('🦍'))

	// ----- CASTING -----
	// Doesn't work with bools or strings
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	// Convert string to int (ASCII to Integer)
	// Returns the result with an error if any
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// Convert int to string (Integer to ASCII)
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// Convert string to float
	cV7 := "3.14"
	// Handling potential errors (Prints if err == nil)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)

	// ----- TIME -----
	// Get day, month, year and time data
	// Get current time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// ----- FORMATTED PRINT -----
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A',
		3.14, true, 1, 1)

	// Float formatting
	fmt.Printf("%9f\n", 3.14)      // Width 9
	fmt.Printf("%.2f\n", 3.141592) // Decimal precision 2
	fmt.Printf("%9.f\n", 3.141592) // Width 9 no precision

	// Sprintf returns a formatted string instead of printing
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)

}
