package database

import "fmt"

type DeleteCommand struct {
	database *Database
	key      string
	value    interface{}
	executed bool
}

func NewDeleteCommand(database *Database, key string) *DeleteCommand {
	value := database.Get(key)
	return &DeleteCommand{
		database: database,
		key:      key,
		value:    value,
		executed: false,
	}
}

func (dc *DeleteCommand) Execute() {
	if !dc.executed {
		dc.database.Delete(dc.key)
		dc.executed = true
	}
}

func (dc *DeleteCommand) Undo() {
	if dc.executed {
		dc.database.Insert(dc.key, dc.value)
		dc.executed = false
	}
}

func (dc *DeleteCommand) GetName() string {
	return fmt.Sprintf("Delete(%s)", dc.key)
}