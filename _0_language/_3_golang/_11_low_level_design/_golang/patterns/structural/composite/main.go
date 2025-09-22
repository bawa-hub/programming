package main

import "fmt"

type Component interface {
	Name() string
	Print(prefix string)
}

type File struct{ name string }
func (f File) Name() string { return f.name }
func (f File) Print(prefix string) { fmt.Println(prefix + f.name) }

type Folder struct{
	name string
	children []Component
}

func (d Folder) Name() string { return d.name }
func (d Folder) Print(prefix string) {
	fmt.Println(prefix + d.name + "/")
	for _, c := range d.children { c.Print(prefix + "  ") }
}

func main() {
	root := Folder{name: "root", children: []Component{
		File{name: "a.txt"},
		Folder{name: "bin", children: []Component{ File{name: "b"} }},
	}}
	root.Print("")
}
