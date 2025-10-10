package remote

import "fmt"

type TurnOffCommand struct {
	device   Device
	executed bool
}

func NewTurnOffCommand(device Device) *TurnOffCommand {
	return &TurnOffCommand{
		device:   device,
		executed: false,
	}
}

func (tofc *TurnOffCommand) Execute() {
	if !tofc.executed {
		tofc.device.Off()
		tofc.executed = true
	}
}

func (tofc *TurnOffCommand) Undo() {
	if tofc.executed {
		tofc.device.On()
		tofc.executed = false
	}
}

func (tofc *TurnOffCommand) GetName() string {
	return fmt.Sprintf("TurnOff(%s)", tofc.device.GetName())
}