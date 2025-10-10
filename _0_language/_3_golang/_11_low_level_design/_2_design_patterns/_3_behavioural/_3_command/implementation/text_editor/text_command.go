package texteditor

type TextCommand interface {
	Execute()
	Undo()
	GetName() string
}