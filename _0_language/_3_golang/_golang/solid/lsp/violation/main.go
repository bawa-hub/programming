package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) SetWidth(w int)  { r.width = w }
func (r *Rectangle) SetHeight(h int) { r.height = h }
func (r Rectangle) Area() int        { return r.width * r.height }

type Square struct{ Rectangle }

func (s *Square) SetWidth(w int) {
	// Violates LSP expectations of Rectangle
	s.width = w
	s.height = w
}

func (s *Square) SetHeight(h int) {
	s.width = h
	s.height = h
}

func resizeTo10x5(r interface{ SetWidth(int); SetHeight(int); Area() int }) {
	r.SetWidth(10)
	r.SetHeight(5)
	fmt.Println("Expected area 50, got", r.Area())
}

func main() {
	rect := &Rectangle{}
	sq := &Square{}

	resizeTo10x5(rect) // 50
	resizeTo10x5(sq)   // not 50
}
