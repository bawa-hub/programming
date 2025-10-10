package basics

import "fmt"

type ConcreteCommand2 struct {
	receiver *Receiver
}

func NewConcreteCommand2(receiver *Receiver) *ConcreteCommand2 {
	return &ConcreteCommand2{receiver: receiver}
}

func (cc2 *ConcreteCommand2) Execute() {
	cc2.receiver.Action2()
}

func (cc2 *ConcreteCommand2) Undo() {
	fmt.Printf("Command2: Undoing Action2\n")
	cc2.receiver.state = "Action2 Undone"
}

func (cc2 *ConcreteCommand2) GetName() string {
	return "Command2"
}
