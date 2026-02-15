package basic

import "fmt"

type ConcreteObserver struct {
	id string
}

func NewConcreteObserver(id string) *ConcreteObserver {
	return &ConcreteObserver{id: id}
}

func (co *ConcreteObserver) Update(data interface{}) {
	fmt.Printf("Observer %s received update: %v\n", co.id, data)
}

func (co *ConcreteObserver) GetID() string {
	return co.id
}