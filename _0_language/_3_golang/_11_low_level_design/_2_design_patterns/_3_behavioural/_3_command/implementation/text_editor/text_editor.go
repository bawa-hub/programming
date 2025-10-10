package texteditor

import "fmt"

type TextEditor struct {
	content string
	cursor  int
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		content: "",
		cursor:  0,
	}
}

func (te *TextEditor) InsertText(text string) {
	te.content = te.content[:te.cursor] + text + te.content[te.cursor:]
	te.cursor += len(text)
	fmt.Printf("TextEditor: Inserted '%s', cursor at %d\n", text, te.cursor)
}

func (te *TextEditor) DeleteText(length int) {
	if te.cursor >= length {
		te.content = te.content[:te.cursor-length] + te.content[te.cursor:]
		te.cursor -= length
		fmt.Printf("TextEditor: Deleted %d characters, cursor at %d\n", length, te.cursor)
	}
}

func (te *TextEditor) MoveCursor(position int) {
	if position >= 0 && position <= len(te.content) {
		te.cursor = position
		fmt.Printf("TextEditor: Moved cursor to %d\n", te.cursor)
	}
}

func (te *TextEditor) GetContent() string {
	return te.content
}

func (te *TextEditor) GetCursor() int {
	return te.cursor
}