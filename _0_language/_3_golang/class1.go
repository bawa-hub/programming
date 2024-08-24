package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	// output ways
	// fmt.Println("hello world")
	// pl("hell once again")

	// input output
	// pl("What is your name ?")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	pl("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// variables
	// var vName string = "vikram"
	// var v1, v2 = 1.2, 1.3
	// var v3 = "Hello"
	// v4 := 4.5
	// v5 := 3.5

	// core data types
	// int, float64, bool, string, rune
	// Default values: 0, 0.0, false, ""

	// type of data
	// pl(reflect.TypeOf(10))

	// type casting
	// cV1 := 1.5
	// cV2 := int(cV1)

	// cV3 := "50000000"
	// cV4, err := strconv.Atoi(cV3)
	// pl(cV4, err, reflect.TypeOf(cV4))

	// operators
	// conditionals
	// age := 8
	// if (age >= 1) && (age <= 18) {
	// 	pl("Important Birthday")
	// } else if (age == 21) || (age == 25) {
	// 	pl("Important Birthday")
	// } else {
	// 	pl("Not important ")
	// }

	// strings
	// sV1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// sV2 := replacer.Replace(sV1)
	// pl(sV2)
	// pl("Length: ", len(sV2))
	// pl("Contains Another: ", strings.Contains(sV2, "Another"))
	// pl("o Index: ", strings.Index(sV2, "o"))
	// pl("Replace: ", strings.Replace(sV2, "o", "0", -1))
	// sV3 := "\nSome Words\n"
	// sV3 = strings.TrimSpace(sV3)
	// pl("Split: ", strings.Split("a-b-c-d", "-"))
	// pl("Lower: ", strings.ToLower(sV2))
	// pl("Upper: ", strings.ToUpper(sV2))
	// pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	// pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))

	// runes
	// rStr := "abcdefg"
	// pl("Rune Count: ", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr {
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	// time
	// now := time.Now()
	// pl(now)
	// pl(now.Date())
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	// math

	// for loop
	// for initialization; condition; postStatement {BODY}
	// for x := 1; x <= 5; x++ {
	// 	pl(x)
	// }

	// while loop
	// fx := 0
	// for fx < 5 {
	// 	pl(fx)
	// 	fx++
	// }

	// range
	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }

	// arrays
	// var arr1 [5]int
	// arr1[0] = 1
	// arr2 := [5]int{1, 2, 3, 4, 5}
	// pl("Index 0: ", arr2[0])
	// pl("Arr Length: ", len(arr2))
	// for i := 0; i < len(arr2); i++ {
	// 	pl(arr2[i])
	// }
	// for i, v := range arr2 {
	// 	fmt.Printf("%d : %d", i, v)
	// }

	// arr3 := [2][2]int{
	// 	{1, 2}, {3, 4},
	// }
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		pl(arr3[i][j])
	// 	}
	// }

	// slices
}
