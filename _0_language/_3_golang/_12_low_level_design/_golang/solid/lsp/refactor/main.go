package main

import "fmt"

type Shape interface{ Area() int }

type Rectangle struct { width, height int }
func (r Rectangle) Area() int { return r.width * r.height }

type Square struct { side int }
func (s Square) Area() int { return s.side * s.side }

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	printArea(Rectangle{width: 10, height: 5})
	printArea(Square{side: 5})
}
