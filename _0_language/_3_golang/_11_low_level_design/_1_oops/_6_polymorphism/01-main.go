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

func main() {
		fmt.Println("3. POLYMORPHISM:")
	vehicles := []Startable{car, motorcycle}
	for _, vehicle := range vehicles {
		StartVehicle(vehicle)
	}
	fmt.Println()

}