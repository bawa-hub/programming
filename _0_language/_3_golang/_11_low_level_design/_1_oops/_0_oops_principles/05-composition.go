package main

import "fmt"

// Using Composition (HAS-A relationship)
type Engine struct {
	Type        string
	Horsepower  int
	IsRunning   bool
}

func (e *Engine) Start() {
	e.IsRunning = true
	fmt.Printf("%s engine started\n", e.Type)
}

func (e *Engine) Stop() {
	e.IsRunning = false
	fmt.Printf("%s engine stopped\n", e.Type)
}

// Car HAS-A Engine (composition)
type ModernCar struct {
	Brand  string
	Model  string
	Engine Engine // Composition
}

func (mc *ModernCar) Start() {
	mc.Engine.Start()
	fmt.Printf("%s %s is ready to drive\n", mc.Brand, mc.Model)
}
