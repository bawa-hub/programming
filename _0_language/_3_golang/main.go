package main

import "fmt"

type Worker interface {
	Work()
}

type Engineer struct{}

func (e *Engineer) Work() {
	fmt.Println("Engineering...")
}

func main() {
	var eng *Engineer = nil
	var w Worker = eng

	fmt.Println("Is w nil?", w == nil)
}
