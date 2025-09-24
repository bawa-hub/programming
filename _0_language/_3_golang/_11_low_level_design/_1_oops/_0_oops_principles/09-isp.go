package main

import "fmt"

// 4. INTERFACE SEGREGATION PRINCIPLE (ISP)
// Clients should not be forced to depend on interfaces they don't use

// Instead of one large interface, create smaller, focused interfaces

// Worker interface
type Worker interface {
	Work()
}

// Eater interface
type Eater interface {
	Eat()
}

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// Human implements all interfaces
type Human struct {
	Name string
}

func (h *Human) Work() {
	fmt.Printf("%s is working\n", h.Name)
}

func (h *Human) Eat() {
	fmt.Printf("%s is eating\n", h.Name)
}

func (h *Human) Sleep() {
	fmt.Printf("%s is sleeping\n", h.Name)
}

// Robot implements only Worker
type Robot struct {
	Model string
}

func (r *Robot) Work() {
	fmt.Printf("Robot %s is working\n", r.Model)
}

func main() {
	human := &Human{Name: "Bob"}
	robot := &Robot{Model: "T-800"}

	workers := []Worker{human, robot}
	for _, worker := range workers {
		worker.Work()
	}
	fmt.Println()
}
