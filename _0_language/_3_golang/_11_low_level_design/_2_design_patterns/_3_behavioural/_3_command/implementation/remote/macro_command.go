package remote

import "fmt"

type MacroCommand struct {
	commands []DeviceCommand
	name     string
}

func NewMacroCommand(name string, commands ...DeviceCommand) *MacroCommand {
	return &MacroCommand{
		commands: commands,
		name:     name,
	}
}

func (mc *MacroCommand) Execute() {
	fmt.Printf("Executing macro: %s\n", mc.name)
	for _, cmd := range mc.commands {
		cmd.Execute()
	}
}

func (mc *MacroCommand) Undo() {
	fmt.Printf("Undoing macro: %s\n", mc.name)
	// Undo in reverse order
	for i := len(mc.commands) - 1; i >= 0; i-- {
		mc.commands[i].Undo()
	}
}

func (mc *MacroCommand) GetName() string {
	return mc.name
}