package database

import "fmt"


type InsertCommand struct {
	database *Database
	key      string
	value    interface{}
	executed bool
}

func NewInsertCommand(database *Database, key string, value interface{}) *InsertCommand {
	return &InsertCommand{
		database: database,
		key:      key,
		value:    value,
		executed: false,
	}
}

func (ic *InsertCommand) Execute() {
	if !ic.executed {
		ic.database.Insert(ic.key, ic.value)
		ic.executed = true
	}
}

func (ic *InsertCommand) Undo() {
	if ic.executed {
		ic.database.Delete(ic.key)
		ic.executed = false
	}
}

func (ic *InsertCommand) GetName() string {
	return fmt.Sprintf("Insert(%s, %v)", ic.key, ic.value)
}
