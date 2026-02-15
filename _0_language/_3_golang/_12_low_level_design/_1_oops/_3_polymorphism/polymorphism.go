package main

import "fmt"

// Interface for polymorphism
type Startable interface {
	Start()
}

// Function that works with any Startable object
func StartVehicle(v Startable) {
	v.Start()
}

type Car struct{}

// Method overriding
func (c *Car) Start() {
	fmt.Println("Car is starting..")
}

type Motorcycle struct{}

func (m *Motorcycle) Start() {
	fmt.Println("Motorcycle is starting..")
}

func main() {
	fmt.Println("3. POLYMORPHISM:")

	car := &Car{}
	motorcycle := &Motorcycle{}

	vehicles := []Startable{car, motorcycle}
	for _, vehicle := range vehicles {
		StartVehicle(vehicle)
	}
	fmt.Println()
}
