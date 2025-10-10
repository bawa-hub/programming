package basics

import "fmt"

type ConcreteCommand1 struct {
	receiver *Receiver
}

func NewConcreteCommand1(receiver *Receiver) *ConcreteCommand1 {
	return &ConcreteCommand1{receiver: receiver}
}

func (cc1 *ConcreteCommand1) Execute() {
	cc1.receiver.Action1()
}

func (cc1 *ConcreteCommand1) Undo() {
	fmt.Printf("Command1: Undoing Action1\n")
	cc1.receiver.state = "Action1 Undone"
}

func (cc1 *ConcreteCommand1) GetName() string {
	return "Command1"
}