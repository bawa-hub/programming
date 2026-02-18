package patterns

import (
	"fmt"
	"sync/atomic"
)

// Advanced Pattern 1: Channel-based State Machine
type StateMachine struct {
	state    int32
	stateCh  chan int32
	actionCh chan func()
	quitCh   chan bool
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		state:    0,
		stateCh:  make(chan int32, 1),
		actionCh: make(chan func(), 10),
		quitCh:   make(chan bool),
	}
	
	sm.stateCh <- 0 // Initial state
	go sm.run()
	return sm
}

func (sm *StateMachine) run() {
	for {
		select {
		case newState := <-sm.stateCh:
			atomic.StoreInt32(&sm.state, newState)
			fmt.Printf("State changed to: %d\n", newState)
		case action := <-sm.actionCh:
			action()
		case <-sm.quitCh:
			return
		}
	}
}

func (sm *StateMachine) SetState(state int32) {
	sm.stateCh <- state
}

func (sm *StateMachine) GetState() int32 {
	return atomic.LoadInt32(&sm.state)
}

func (sm *StateMachine) DoAction(action func()) {
	sm.actionCh <- action
}

func (sm *StateMachine) Stop() {
	close(sm.quitCh)
}