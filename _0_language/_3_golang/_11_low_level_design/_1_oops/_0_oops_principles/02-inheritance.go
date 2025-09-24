package main

import "fmt"

// Base class (using struct embedding)
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v *Vehicle) Start() {
	fmt.Printf("%s %s is starting...\n", v.Brand, v.Model)
}

func (v *Vehicle) Stop() {
	fmt.Printf("%s %s is stopping...\n", v.Brand, v.Model)
}

func (v *Vehicle) GetInfo() string {
	return fmt.Sprintf("%s %s (%d)", v.Brand, v.Model, v.Year)
}

// Derived class (inheritance through embedding)
type Car struct {
	Vehicle // Embedded struct - Car "is-a" Vehicle
	Doors   int
	Engine  string
}

// Method overriding
func (c *Car) Start() {
	fmt.Printf("Car %s %s with %s engine is starting...\n", c.Brand, c.Model, c.Engine)
}

// Additional methods specific to Car
func (c *Car) OpenTrunk() {
	fmt.Printf("Opening trunk of %s %s\n", c.Brand, c.Model)
}

// Another derived class
type Motorcycle struct {
	Vehicle
	HasWindshield bool
}

func (m *Motorcycle) Start() {
	fmt.Printf("Motorcycle %s %s is starting...\n", m.Brand, m.Model)
}

func (m *Motorcycle) Wheelie() {
	fmt.Printf("Doing a wheelie on %s %s\n", m.Brand, m.Model)
}
