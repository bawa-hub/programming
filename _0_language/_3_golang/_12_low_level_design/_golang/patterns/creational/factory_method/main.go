package main

import "fmt"

type Transport interface { Deliver() }

type Truck struct{}
func (Truck) Deliver() { fmt.Println("Deliver by land in a box") }

type Ship struct{}
func (Ship) Deliver() { fmt.Println("Deliver by sea in a container") }

type Logistics interface { CreateTransport() Transport }

type RoadLogistics struct{}
func (RoadLogistics) CreateTransport() Transport { return Truck{} }

type SeaLogistics struct{}
func (SeaLogistics) CreateTransport() Transport { return Ship{} }

func planDelivery(l Logistics) {
	transport := l.CreateTransport()
	transport.Deliver()
}

func main() {
	planDelivery(RoadLogistics{})
	planDelivery(SeaLogistics{})
}
