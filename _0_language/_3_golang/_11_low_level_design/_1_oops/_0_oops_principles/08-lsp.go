package main

import "fmt"

// 3. LISKOV SUBSTITUTION PRINCIPLE (LSP)
// Objects of superclass should be replaceable with objects of subclass

// Bird interface
type Bird interface {
	Fly() string
	MakeSound() string
}

// Sparrow implements Bird
type Sparrow struct{}

func (s *Sparrow) Fly() string {
	return "Sparrow is flying"
}

func (s *Sparrow) MakeSound() string {
	return "Chirp chirp"
}

// Penguin implements Bird (but can't fly)
type Penguin struct{}

func (p *Penguin) Fly() string {
	return "Penguin cannot fly"
}

func (p *Penguin) MakeSound() string {
	return "Honk honk"
}

// Function that works with any Bird
func MakeBirdFly(b Bird) {
	fmt.Println(b.Fly())
}
