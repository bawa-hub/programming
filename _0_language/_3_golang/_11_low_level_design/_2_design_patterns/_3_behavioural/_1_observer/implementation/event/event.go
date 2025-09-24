package event

import (
	"fmt"
	"time"
)

type Event struct {
	Type      string
	Data      interface{}
	Timestamp time.Time
	Source    string
}

func (e *Event) String() string {
	return fmt.Sprintf("[%s] %s from %s: %v", 
		e.Timestamp.Format("15:04:05"), e.Type, e.Source, e.Data)
}