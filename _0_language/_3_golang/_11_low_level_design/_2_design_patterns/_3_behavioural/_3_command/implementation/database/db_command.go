package database

type DatabaseCommand interface {
	Execute()
	Undo()
	GetName() string
}