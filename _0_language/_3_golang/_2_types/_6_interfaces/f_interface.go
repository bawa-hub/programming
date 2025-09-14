package main

import "fmt"

var pl = fmt.Println

// An interface is a set of method signatures.
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

// Empty Interface (interface{})
	// Old Go version: interface{} meant any type.
	// Modern Go: any is an alias for interface{}.
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
// But values stored in an interface have type + value pair (fat pointer).
// This is why type assertions are needed.

// Type Assertions & Type Switch
func Describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    default:
        fmt.Println("unknown type")
    }
}


func main() {

	// Interfaces allow you to create contractsthat say if anything inherits it that they will implement defined methods
	// If we had animals and wanted to define that they all perform certain actions, but in their specific way we could use an interface
	// With Go you don't have to say a type uses an interface. When your type implements the required methods it is automatic

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	// We can only call methods defined in the interface for Cats because of the contract unless you convert Cat back into a concrete Cat type using a type assertion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}
