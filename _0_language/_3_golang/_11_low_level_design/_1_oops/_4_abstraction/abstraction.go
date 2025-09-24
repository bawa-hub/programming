package main

import "fmt"

// Abstract interface (Go uses interfaces for abstraction)
type Shape interface {
	Area() float64
	Perimeter() float64
	GetName() string
}

// Concrete implementation of Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) GetName() string {
	return "Rectangle"
}

// Another concrete implementation
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c *Circle) GetName() string {
	return "Circle"
}

// Function that works with any Shape (polymorphism)
func PrintShapeInfo(s Shape) {
	fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", s.GetName(), s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("4. ABSTRACTION:")
	shapes := []Shape{
		&Rectangle{Width: 5, Height: 3},
		&Circle{Radius: 4},
	}
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
	fmt.Println()
}
