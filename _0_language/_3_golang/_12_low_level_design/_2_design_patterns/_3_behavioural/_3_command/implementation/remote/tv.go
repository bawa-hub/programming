package remote

import "fmt"

type TV struct {
	name  string
	state string
}

func NewTV(name string) *TV {
	return &TV{
		name:  name,
		state: "off",
	}
}

func (t *TV) On() {
	t.state = "on"
	fmt.Printf("TV %s is now ON\n", t.name)
}

func (t *TV) Off() {
	t.state = "off"
	fmt.Printf("TV %s is now OFF\n", t.name)
}

func (t *TV) GetName() string {
	return t.name
}

func (t *TV) GetState() string {
	return t.state
}
