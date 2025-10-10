package basics

import "fmt"

type Invoker struct {
	commands []Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make([]Command, 0),
	}
}

func (i *Invoker) ExecuteCommand(command Command) {
	command.Execute()
	i.commands = append(i.commands, command)
	fmt.Printf("Invoker: Executed %s\n", command.GetName())
}

func (i *Invoker) UndoLastCommand() {
	if len(i.commands) > 0 {
		lastCommand := i.commands[len(i.commands)-1]
		lastCommand.Undo()
		i.commands = i.commands[:len(i.commands)-1]
		fmt.Printf("Invoker: Undone %s\n", lastCommand.GetName())
	} else {
		fmt.Println("Invoker: No commands to undo")
	}
}

func (i *Invoker) GetCommandHistory() []string {
	history := make([]string, len(i.commands))
	for i, cmd := range i.commands {
		history[i] = cmd.GetName()
	}
	return history
}