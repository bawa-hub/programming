package main

import "fmt"

var pl = fmt.Println

func main() {

	// Maps are collections of key/value pairs. Keys can be any data type that can be compared
	// syntax: var myMap map [keyType]valueType

	var heroes map[string]string // Declare a map variable
	heroes = make(map[string]string) // create
	villians := make(map[string]string) // You can do it in one step
	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"} 	// Define with map literal

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	fmt.Printf("Batman is %v\n", heroes["Batman"])

	pl("Chip :", superPets[3]) 	// key that doesn't exist you get nil

	_, ok := superPets[3] 	// check if there is a value or nil
	pl("Is there a 3rd pet :", ok)

	// Cycle through map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	delete(heroes, "The Flash") 	// Delete a key value
}
