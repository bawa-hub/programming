package remote

import "fmt"

type Light struct {
	name  string
	state string
}

func NewLight(name string) *Light {
	return &Light{
		name:  name,
		state: "off",
	}
}

func (l *Light) On() {
	l.state = "on"
	fmt.Printf("Light %s is now ON\n", l.name)
}

func (l *Light) Off() {
	l.state = "off"
	fmt.Printf("Light %s is now OFF\n", l.name)
}

func (l *Light) GetName() string {
	return l.name
}

func (l *Light) GetState() string {
	return l.state
}