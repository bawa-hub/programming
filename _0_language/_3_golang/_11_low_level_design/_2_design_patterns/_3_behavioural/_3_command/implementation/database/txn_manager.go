package database

import "fmt"

// Transaction Manager
type TransactionManager struct {
	commands []DatabaseCommand
	database *Database
}

func NewTransactionManager(database *Database) *TransactionManager {
	return &TransactionManager{
		commands: make([]DatabaseCommand, 0),
		database: database,
	}
}

func (tm *TransactionManager) AddCommand(command DatabaseCommand) {
	tm.commands = append(tm.commands, command)
}

func (tm *TransactionManager) ExecuteTransaction() {
	fmt.Println("Executing transaction...")
	for _, cmd := range tm.commands {
		cmd.Execute()
	}
}

func (tm *TransactionManager) RollbackTransaction() {
	fmt.Println("Rolling back transaction...")
	for i := len(tm.commands) - 1; i >= 0; i-- {
		tm.commands[i].Undo()
	}
}

func (tm *TransactionManager) ClearTransaction() {
	tm.commands = make([]DatabaseCommand, 0)
}