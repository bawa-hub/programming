package main

import "fmt"

type Shape interface { Draw() }

type Circle struct{}
func (Circle) Draw() { fmt.Println("Drawing Circle") }

type Square struct{}
func (Square) Draw() { fmt.Println("Drawing Square") }

type ShapeType string

const (
	CircleType ShapeType = "circle"
	SquareType ShapeType = "square"
)

type ShapeFactory struct{}

func (ShapeFactory) Create(t ShapeType) Shape {
	switch t {
	case CircleType:
		return Circle{}
	case SquareType:
		return Square{}
	default:
		return nil
	}
}

func main() {
	f := ShapeFactory{}
	f.Create(CircleType).Draw()
	f.Create(SquareType).Draw()
}
