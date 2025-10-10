package main

import (
	b "command/basics"
	txt "command/text_editor"
	r "command/remote"
	db "command/database"
	"fmt"
)

func main() {
	fmt.Println("=== COMMAND PATTERN DEMONSTRATION ===\n")

	// 1. BASIC COMMAND
	fmt.Println("1. BASIC COMMAND:")
	receiver := b.NewReceiver()
	invoker := b.NewInvoker()
	
	command1 := b.NewConcreteCommand1(receiver)
	command2 := b.NewConcreteCommand2(receiver)
	
	invoker.ExecuteCommand(command1)
	invoker.ExecuteCommand(command2)
	
	fmt.Printf("Receiver state: %s\n", receiver.GetState())
	fmt.Printf("Command history: %v\n", invoker.GetCommandHistory())
	
	invoker.UndoLastCommand()
	invoker.UndoLastCommand()
	invoker.UndoLastCommand() // No more commands to undo
	fmt.Println()

	// // 2. REAL-WORLD EXAMPLES
	// fmt.Println("2. REAL-WORLD EXAMPLES:")

	// // Text Editor Commands
	fmt.Println("Text Editor Commands:")
	editor := txt.NewTextEditor()
	
	insertCmd1 := txt.NewInsertTextCommand(editor, "Hello", 0)
	insertCmd2 := txt.NewInsertTextCommand(editor, " World", 5)
	insertCmd3 := txt.NewInsertTextCommand(editor, "!", 11)
	
	insertCmd1.Execute()
	insertCmd2.Execute()
	insertCmd3.Execute()
	
	fmt.Printf("Content: '%s'\n", editor.GetContent())
	
	insertCmd3.Undo()
	insertCmd2.Undo()
	insertCmd1.Undo()
	
	fmt.Printf("Content after undo: '%s'\n", editor.GetContent())
	fmt.Println()

	// // Remote Control Commands
	fmt.Println("Remote Control Commands:")
	light1 := r.NewLight("Living Room")
	light2 := r.NewLight("Bedroom")
	tv := r.NewTV("Living Room")
	
	remote := r.NewRemoteControl()
	
	// Turn on devices
	remote.PressButton(r.NewTurnOnCommand(light1))
	remote.PressButton(r.NewTurnOnCommand(tv))
	remote.PressButton(r.NewTurnOnCommand(light2))
	
	// Undo last command
	remote.UndoLastCommand()
	
	// // Create macro command
	macroCmd := r.NewMacroCommand("Movie Night", 
		r.NewTurnOnCommand(light1),
		r.NewTurnOnCommand(tv),
		r.NewTurnOffCommand(light2),
	)
	
	remote.PressButton(macroCmd)
	remote.UndoLastCommand()
	fmt.Println()

	// Database Transaction Commands
	fmt.Println("Database Transaction Commands:")
	database := db.NewDatabase()
	transactionManager := db.NewTransactionManager(database)
	
	// Add commands to transaction
	transactionManager.AddCommand(db.NewInsertCommand(database, "user1", "John"))
	transactionManager.AddCommand(db.NewInsertCommand(database, "user2", "Jane"))
	transactionManager.AddCommand(db.NewUpdateCommand(database, "user1", "John Doe"))
	transactionManager.AddCommand(db.NewDeleteCommand(database, "user2"))
	
	// Execute transaction
	transactionManager.ExecuteTransaction()
	
	// Rollback transaction
	transactionManager.RollbackTransaction()
	
	// Clear transaction
	transactionManager.ClearTransaction()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
