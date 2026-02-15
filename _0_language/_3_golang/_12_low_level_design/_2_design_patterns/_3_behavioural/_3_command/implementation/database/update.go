package database

import "fmt"


type UpdateCommand struct {
	database *Database
	key      string
	value    interface{}
	oldValue interface{}
	executed bool
}

func NewUpdateCommand(database *Database, key string, value interface{}) *UpdateCommand {
	oldValue := database.Get(key)
	return &UpdateCommand{
		database: database,
		key:      key,
		value:    value,
		oldValue: oldValue,
		executed: false,
	}
}

func (uc *UpdateCommand) Execute() {
	if !uc.executed {
		uc.database.Update(uc.key, uc.value)
		uc.executed = true
	}
}

func (uc *UpdateCommand) Undo() {
	if uc.executed {
		uc.database.Update(uc.key, uc.oldValue)
		uc.executed = false
	}
}

func (uc *UpdateCommand) GetName() string {
	return fmt.Sprintf("Update(%s, %v)", uc.key, uc.value)
}