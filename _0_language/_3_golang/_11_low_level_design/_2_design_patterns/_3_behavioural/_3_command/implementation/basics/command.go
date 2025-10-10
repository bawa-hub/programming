package basics

// Command interface
type Command interface {
	Execute()
	Undo()
	GetName() string
}