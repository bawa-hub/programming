package main

import "fmt"

type Visitor interface {
	VisitCircle(c Circle)
	VisitSquare(s Square)
}

type Element interface { Accept(v Visitor) }

type Circle struct{ radius int }
func (c Circle) Accept(v Visitor) { v.VisitCircle(c) }

type Square struct{ side int }
func (s Square) Accept(v Visitor) { v.VisitSquare(s) }

type AreaVisitor struct{}
func (AreaVisitor) VisitCircle(c Circle) { fmt.Println("Circle area:", 3.14*float64(c.radius*c.radius)) }
func (AreaVisitor) VisitSquare(s Square) { fmt.Println("Square area:", s.side*s.side) }

func main() {
	elems := []Element{ Circle{3}, Square{4} }
	v := AreaVisitor{}
	for _, e := range elems { e.Accept(v) }
}
