package basics

import "fmt"

// Receiver - knows how to perform operations
type Receiver struct {
	state string
}

func NewReceiver() *Receiver {
	return &Receiver{state: "Initial State"}
}

func (r *Receiver) Action1() {
	fmt.Printf("Receiver: Performing Action1, current state: %s\n", r.state)
	r.state = "Action1 Completed"
}

func (r *Receiver) Action2() {
	fmt.Printf("Receiver: Performing Action2, current state: %s\n", r.state)
	r.state = "Action2 Completed"
}

func (r *Receiver) Action3() {
	fmt.Printf("Receiver: Performing Action3, current state: %s\n", r.state)
	r.state = "Action3 Completed"
}

func (r *Receiver) GetState() string {
	return r.state
}