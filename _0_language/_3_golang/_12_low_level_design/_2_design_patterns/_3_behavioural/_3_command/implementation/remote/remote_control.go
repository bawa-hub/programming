package remote

type RemoteControl struct {
	commands []DeviceCommand
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commands: make([]DeviceCommand, 0),
	}
}

func (rc *RemoteControl) PressButton(command DeviceCommand) {
	command.Execute()
	rc.commands = append(rc.commands, command)
}

func (rc *RemoteControl) UndoLastCommand() {
	if len(rc.commands) > 0 {
		lastCommand := rc.commands[len(rc.commands)-1]
		lastCommand.Undo()
		rc.commands = rc.commands[:len(rc.commands)-1]
	}
}