package main

import "fmt"

type Clonable[T any] interface {
	Clone() T
}

type Document struct {
	Title string
	Pages []string
}

func (d Document) Clone() Document {
	pagesCopy := make([]string, len(d.Pages))
	copy(pagesCopy, d.Pages)
	return Document{Title: d.Title, Pages: pagesCopy}
}

func main() {
	d1 := Document{Title: "Spec", Pages: []string{"Intro", "Design"}}
	d2 := d1.Clone()
	d2.Pages[0] = "Overview"
	fmt.Println(d1.Pages[0], "|", d2.Pages[0])
}
