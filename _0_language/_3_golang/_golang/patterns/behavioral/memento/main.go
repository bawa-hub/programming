package main

import "fmt"

type Memento struct { state string }

type Originator struct { state string }
func (o *Originator) SetState(s string) { o.state = s }
func (o Originator) Save() Memento { return Memento{state: o.state} }
func (o *Originator) Restore(m Memento) { o.state = m.state }

func main() {
	o := &Originator{}
	o.SetState("A")
	snap := o.Save()
	o.SetState("B")

o.Restore(snap)
	fmt.Println(o.state)
}
