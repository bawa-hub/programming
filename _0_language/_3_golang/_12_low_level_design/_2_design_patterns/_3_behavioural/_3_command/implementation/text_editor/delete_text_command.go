package texteditor

import "fmt"

type DeleteTextCommand struct {
	editor    *TextEditor
	text      string
	position  int
	executed  bool
}

func NewDeleteTextCommand(editor *TextEditor, position int, length int) *DeleteTextCommand {
	text := editor.content[position : position+length]
	return &DeleteTextCommand{
		editor:   editor,
		text:     text,
		position: position,
		executed: false,
	}
}

func (dtc *DeleteTextCommand) Execute() {
	if !dtc.executed {
		dtc.editor.MoveCursor(dtc.position)
		dtc.editor.DeleteText(len(dtc.text))
		dtc.executed = true
	}
}

func (dtc *DeleteTextCommand) Undo() {
	if dtc.executed {
		dtc.editor.MoveCursor(dtc.position)
		dtc.editor.InsertText(dtc.text)
		dtc.executed = false
	}
}

func (dtc *DeleteTextCommand) GetName() string {
	return fmt.Sprintf("DeleteText('%s' at %d)", dtc.text, dtc.position)
}