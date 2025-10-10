package remote

import "fmt"

type TurnOnCommand struct {
	device   Device
	executed bool
}

func NewTurnOnCommand(device Device) *TurnOnCommand {
	return &TurnOnCommand{
		device:   device,
		executed: false,
	}
}

func (toc *TurnOnCommand) Execute() {
	if !toc.executed {
		toc.device.On()
		toc.executed = true
	}
}

func (toc *TurnOnCommand) Undo() {
	if toc.executed {
		toc.device.Off()
		toc.executed = false
	}
}

func (toc *TurnOnCommand) GetName() string {
	return fmt.Sprintf("TurnOn(%s)", toc.device.GetName())
}