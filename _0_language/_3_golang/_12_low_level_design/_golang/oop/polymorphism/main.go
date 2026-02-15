package main

import "fmt"

type Shape interface { Area() float64 }

type Rect struct { w, h float64 }
func (r Rect) Area() float64 { return r.w * r.h }

type Circle struct { r float64 }
func (c Circle) Area() float64 { return 3.14159 * c.r * c.r }

func totalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes { sum += s.Area() }
	return sum
}

func main() {
	shapes := []Shape{ Rect{2,3}, Circle{1.5} }
	fmt.Println(totalArea(shapes))
}
