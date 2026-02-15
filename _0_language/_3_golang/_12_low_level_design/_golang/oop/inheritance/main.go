package main

import "fmt"

type Animal struct { name string }
func (a Animal) Speak() { fmt.Println(a.name, "makes a sound") }

// Dog "inherits" via embedding

type Dog struct { Animal; breed string }
func (d Dog) Speak() { fmt.Println(d.name, "barks") }

func main() {
	animal := Animal{name: "Generic"}
	dog := Dog{Animal: Animal{name: "Rex"}, breed: "Labrador"}
	animal.Speak() // base
	dog.Speak()    // override
	fmt.Println(dog.breed)
}
