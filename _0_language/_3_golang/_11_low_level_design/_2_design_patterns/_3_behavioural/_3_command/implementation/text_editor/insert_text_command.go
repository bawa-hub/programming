package texteditor

import "fmt"

type InsertTextCommand struct {
	editor    *TextEditor
	text      string
	position  int
	executed  bool
}

func NewInsertTextCommand(editor *TextEditor, text string, position int) *InsertTextCommand {
	return &InsertTextCommand{
		editor:   editor,
		text:     text,
		position: position,
		executed: false,
	}
}

func (itc *InsertTextCommand) Execute() {
	if !itc.executed {
		itc.editor.MoveCursor(itc.position)
		itc.editor.InsertText(itc.text)
		itc.executed = true
	}
}

func (itc *InsertTextCommand) Undo() {
	if itc.executed {
		itc.editor.MoveCursor(itc.position)
		itc.editor.DeleteText(len(itc.text))
		itc.executed = false
	}
}

func (itc *InsertTextCommand) GetName() string {
	return fmt.Sprintf("InsertText('%s' at %d)", itc.text, itc.position)
}