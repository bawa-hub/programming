package main

import "fmt"

type House struct {
	Walls int
	Doors int
	HasGarage bool
}

type HouseBuilder interface {
	Reset()
	BuildWalls(n int)
	BuildDoors(n int)
	BuildGarage()
	GetResult() House
}

type SimpleHouseBuilder struct{ house House }

func (b *SimpleHouseBuilder) Reset()              { b.house = House{} }
func (b *SimpleHouseBuilder) BuildWalls(n int)    { b.house.Walls = n }
func (b *SimpleHouseBuilder) BuildDoors(n int)    { b.house.Doors = n }
func (b *SimpleHouseBuilder) BuildGarage()        { b.house.HasGarage = true }
func (b *SimpleHouseBuilder) GetResult() House    { return b.house }

type Director struct{ builder HouseBuilder }

func NewDirector(b HouseBuilder) *Director { return &Director{builder: b} }

func (d *Director) constructSimpleHouse() House {
	d.builder.Reset()
	d.builder.BuildWalls(4)
	d.builder.BuildDoors(1)
	return d.builder.GetResult()
}

func (d *Director) constructFamilyHouse() House {
	d.builder.Reset()
	d.builder.BuildWalls(6)
	d.builder.BuildDoors(2)
	d.builder.BuildGarage()
	return d.builder.GetResult()
}

func main() {
	b := &SimpleHouseBuilder{}
	d := NewDirector(b)
	fmt.Printf("%+v\n", d.constructSimpleHouse())
	fmt.Printf("%+v\n", d.constructFamilyHouse())
}
