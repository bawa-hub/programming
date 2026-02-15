package main

import "fmt"

type TreeType struct { name, color string }
func (t TreeType) Draw(x, y int) { fmt.Printf("Draw %s (%s) at %d,%d\n", t.name, t.color, x, y) }

type TreeFactory struct { types map[string]TreeType }

func NewTreeFactory() *TreeFactory { return &TreeFactory{types: map[string]TreeType{}} }

func (f *TreeFactory) GetTreeType(name, color string) TreeType {
	key := name+":"+color
	if t, ok := f.types[key]; ok { return t }
	t := TreeType{name: name, color: color}
	f.types[key] = t
	return t
}

type Tree struct { x, y int; kind TreeType }

func main() {
	factory := NewTreeFactory()
	pine := factory.GetTreeType("pine", "green")
	Tree{10, 20, pine}.kind.Draw(10,20)
	// Reuses flyweight
	Tree{5, 7, factory.GetTreeType("pine", "green")}.kind.Draw(5,7)
}
