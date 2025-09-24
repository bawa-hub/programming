package main

// Interface for polymorphism
type Startable interface {
	Start()
}

// Function that works with any Startable object
func StartVehicle(v Startable) {
	v.Start()
}
